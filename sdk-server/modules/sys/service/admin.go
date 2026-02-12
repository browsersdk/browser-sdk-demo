package service

import (
	"dilu/common/ckey"
	"dilu/common/codes"
	"dilu/common/utils"
	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"dilu/common/regexps"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AdminService struct {
	*base.BaseService
}

var SerAdmin = AdminService{
	base.NewService("sys"),
}

func (e *AdminService) QueryPage(req *dto.AdminGetPageReq, list *[]models.Admin, total *int64) error {
	db := e.DB().Limit(req.GetSize()).Offset(req.GetOffset())
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	} else {
		db.Order("status desc, updated_at desc")
	}
	if req.DeptId > 0 {
		var dept models.AdminDept
		if err := SerAdminDept.Get(req.DeptId, &dept); err != nil {
			return err
		}
		// 	db.Where("dept_id = ?", req.DeptId)
		// } else if req.DeptPath != "" {
		db.Where("dept_path like ?", dept.DeptPath+"%")
	}
	if req.Name != "" {
		db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Phone != "" {
		db.Where("phone = ?", req.Phone)
	}
	return db.Find(list).Limit(-1).Offset(-1).Count(total).Error
}

// Get 获取admin对象
func (e *AdminService) Get(id int, model *models.Admin) error {
	data, _ := core.Cache.Get(ckey.AdminKey(id))

	if data != "" {
		json.Unmarshal([]byte(data), model)
	}
	if model.Id > 0 {
		return nil
	}

	err := e.DB().First(model, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		return err
	}
	core.Cache.Set(ckey.AdminKey(id), model, time.Duration(core.Cfg.JWT.Expires*2)*time.Minute)
	return err
}

// ResetPwd 修改用户信息
func (e *AdminService) Update(model *models.Admin) error {
	var old models.Admin
	err := core.DB().First(&old, model.Id).Error
	if err != nil {
		return err
	}

	if (model.RoleId == enums.RolePlatformSuper || old.RoleId == enums.RolePlatformSuper) && old.RoleId == model.RoleId {
		return codes.Err403(errors.New("the super administrator is unique."))
	}

	if model.Password == old.Password {
		model.Password = ""
	} else if model.Password != "" {
		en, err := model.GenPwd(model.Password)
		if err != nil {
			return err
		}
		model.Password = en
	}

	if model.DeptId != 0 {
		var dept models.AdminDept
		if err := SerAdminDept.Get(model.DeptId, &dept); err != nil {
			return err
		}
		model.DeptPath = dept.DeptPath
	}

	err = e.BaseService.UpdateById(&model)
	if err != nil {
		return err
	}
	core.Cache.Del(ckey.AdminKey(model.Id))
	return nil
}

// Remove 删除admin
func (e *AdminService) Remove(id int, adminId int) error {
	var err error
	var data models.Admin

	if err := e.Get(id, &data); err != nil {
		return err
	}

	if data.RoleId == enums.RolePlatformSuper {
		return codes.Err403(nil)
	}

	u := map[string]any{
		"id":         id,
		"update_by":  adminId,
		"status":     -1,
		"updated_at": time.Now(),
	}

	db := e.DB().Model(&data).Updates(u)
	if err = db.Error; err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	core.Cache.Del(ckey.AdminKey(id))
	return nil
}

// UpdatePwd 修改admin对象密码
func (e *AdminService) UpdatePwd(id int, oldPassword, newPassword string) errs.IError {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.Admin{}

	err = core.DB().Model(c).Select("Id", "Password").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrNotFound(strconv.Itoa(id), "admin", "", err)
		}
		return codes.ErrSys(err)
	}
	if !c.CompPwd(oldPassword) {
		return errs.ErrWithCode(codes.ErrPwd)
	}
	c.Password = newPassword
	db := core.DB().Model(c).Where("id = ?", id).
		Select("Password", "Salt").
		Updates(c)
	if err = db.Error; err != nil {
		return codes.ErrSys(err)
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		core.Log.Warn("db update error")
		return codes.ErrSys(err)
	}
	core.Cache.Del(ckey.AdminKey(id))
	return nil
}

// 手机注册次数
func (e *AdminService) CountByPhone(phone string, count *int64) error {
	var data models.Admin
	return core.DB().Model(&data).Where("phone = ?", phone).Count(count).Error

}

// 手机注册次数
func (e *AdminService) CountByEmail(email string, count *int64) error {
	var data models.Admin
	return core.DB().Model(&data).Where("email = ?", email).Count(count).Error
}

// 注册用户
func (e *AdminService) Register(loginType enums.LoginType, c *dto.RegisterReq, ip string) (dto.LoginOK, errs.IError) {
	model := models.Admin{}
	lok := dto.LoginOK{}
	var count int64
	if loginType == 1 {
		if err := e.CountByPhone(c.Username, &count); err != nil {
			return lok, codes.ErrSys(err)
		}
		if count > 0 {
			return lok, errs.ErrWithCode(codes.PhoneExistErr)
		}
		model.Phone = c.Username
		if c.Name == "" {
			c.Name = c.Username
		}
	} else {
		if err := e.CountByEmail(c.Username, &count); err != nil {
			return lok, codes.ErrSys(err)
		}
		if count > 0 {
			return lok, errs.ErrWithCode(codes.EmailExistErr)
		}
		model.Email = c.Username
		if c.Name == "" {
			arr := strings.Split(c.Username, "@")
			c.Name = arr[0]
		}
	}
	model.Password = c.Password
	model.Nickname = c.Name
	model.CreatedAt = time.Now()
	model.UpdatedAt = model.CreatedAt
	model.RegIp = ip
	model.ClientId = c.DevId
	err := core.DB().Create(&model).Error
	if err != nil {
		core.Log.Error("adminService", "Create", err)
		return lok, codes.ErrSys(err)
	}
	return e.loginOK(&model, 0, loginType, ip, "", c.Ver, c.Os, c.Ver)
}

// func (e *AdminService) GetByInviteCode(inviteCode string) (*models.Admin, errs.IError) {
// 	var model models.Admin
// 	err := core.DB().Model(&model).Where("invite_code =?", inviteCode).First(&model).Error
// 	if err != nil {
// 		return nil, codes.ErrSys(err)
// 	}
// 	return &model, nil
// }

func (e *AdminService) loginOK(u *models.Admin, need int, loginType enums.LoginType, ip, location, clientVer, os, osVer string) (dto.LoginOK, errs.IError) {
	now := time.Now()
	exp := now.Add(time.Duration(core.Cfg.JWT.Expires) * time.Minute)

	token, err := utils.GenerateAdminToken(u.Id, u.RoleId, core.Cfg.JWT.Issuer, core.Cfg.JWT.Subject, core.Cfg.JWT.SignKey, exp)
	lok := dto.LoginOK{}
	if err != nil {
		return lok, errs.Err(codes.FAILURE, "", err)
	}
	lok.Expire = exp
	lok.AccessToken = token
	lok.Need = need
	if u.Nickname != "" {
		lok.Username = u.Nickname
	} else if u.Username != "" {
		lok.Username = u.Username
	} else if u.Phone != "" {
		lok.Username = u.Phone
	} else if u.Email != "" {
		lok.Username = u.Email
	}
	refExp := now.Add(time.Duration(core.Cfg.JWT.Refresh) * time.Minute)
	refT, _ := utils.GenerateAdminRefreshToken(u.Id, u.RoleId, core.Cfg.JWT.Issuer, core.Cfg.JWT.Subject, core.Cfg.JWT.SignKey, refExp)
	lok.RefreshToken = refT
	lok.RefreshExpire = refExp

	core.Cache.Set(ckey.AdminKey(u.Id), u, time.Duration(core.Cfg.JWT.Expires*2)*time.Minute)

	return lok, nil
}

// 通过密码登录
func (e *AdminService) LoginPwd(c *dto.LoginReq, ip string) (dto.LoginOK, errs.IError) {
	model := models.Admin{}
	lok := dto.LoginOK{}
	loginType := enums.LT_USER_PWD
	if regexps.CheckMobile(c.Username) {
		if err := e.GetByPhone(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
			}
			return lok, errs.ErrWithCode(codes.FAILURE)
		}
		loginType = enums.LT_PHONE
	} else if regexps.CheckEmail(c.Username) { //是否邮箱
		if err := e.GetByEmail(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
			}
			return lok, errs.ErrWithCode(codes.FAILURE)
		}
		loginType = enums.LT_EMAIL
	} else { //用户名密码登录
		if err := e.GetByUsername(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
			}
			return lok, errs.ErrWithCode(codes.FAILURE)
		}
	}
	if model.Password == "" {
		return lok, errs.ErrWithCode(codes.PwdNotExist)
	}
	if !model.CompPwd(c.Password) {
		return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
	}
	if c.UUID != "" {
		//e.bindById(c.UUID, model)
	}
	return e.loginOK(&model, 0, loginType, ip, "", c.Ver, c.Os, c.Ver)
}

func (e *AdminService) RefreshToken(c *dto.RefreshTokenReq, ip string) (dto.LoginOK, errs.IError) {
	var cc utils.AdminClaims
	if err := utils.ParseAdminRefresh(c.RefreshToken, &cc, core.Cfg.JWT.SignKey); err != nil {
		return dto.LoginOK{}, codes.Err401(err)
	}
	var user models.Admin
	if err := e.Get(cc.Uid, &user); err != nil {
		return dto.LoginOK{}, codes.Err401(err)
	}

	if user.Status > 0 {
		return e.loginOK(&user, 0, enums.LT_UNKNOWN, ip, "", "", "", "")
	}
	return dto.LoginOK{}, codes.Err401(errors.New("账号已被锁定，请联系客服"))
}

// 通过验证码
func (e *AdminService) LoginCode(c *dto.LoginReq, ip string) (dto.LoginOK, errs.IError) {
	var model models.Admin
	lok := dto.LoginOK{}
	//var name string
	loginType := enums.LT_UNKNOWN
	if regexps.CheckMobile(c.Username) {
		if err := e.GetByPhone(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, codes.ErrSys(err)
			} else {
				model.Phone = c.Username
				//name = c.Username
			}
		}
		loginType = enums.LT_PHONE
	} else if regexps.CheckEmail(c.Username) { //是否邮箱
		if err := e.GetByEmail(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, codes.ErrNotFound(c.Username, "admin", "", err)
			} else {
				model.Email = c.Username
				//arr := strings.Split(c.Username, "@")
				//name = arr[0]
			}
		}
		loginType = enums.LT_EMAIL
	} else {
		return lok, errs.ErrWithCode(codes.ErrMobileOrEmail)
	}
	if model.Id == 0 {
		// model.CreatedAt = time.Now()
		// model.UpdatedAt = model.CreatedAt
		// model.Nickname = name
		// model.RegIp = ip
		// model.ClientId = c.ClientId
		// err := core.DB().Create(&model).Error
		// if err != nil {
		// 	core.Log.Error("admin", "create", err)
		// 	return lok, codes.ErrSys(err)
		// }
		// if c.UUID != "" {
		// 	//e.bindById(c.UUID, model)
		// }
		// return e.loginOK(&model, 2, loginType, ip, "", c.ClientVer, c.Os, c.OsVer)
		return lok, errs.ErrWithCode(codes.ErrUsernameOrPwd)
	}
	if c.UUID != "" {
		//e.bindById(c.UUID, model)
	}
	return e.loginOK(&model, 0, loginType, ip, "", c.Ver, c.Os, c.Ver)
}

// Get 获取User对象
func (e *AdminService) GetByEmail(email string, model *models.Admin) error {
	return core.DB().Where("email = ?", email).First(model).Error
}

// Get 获取User对象
func (e *AdminService) GetByPhone(mobile string, model *models.Admin) error {
	return core.DB().Where("phone = ?", mobile).First(model).Error
}

// func (e *AdminService) bindById(enCode string, user models.Admin) error {
// 	dstr, err := cryptos.RSA_Decrypt(enCode, consts.PriKey)
// 	if err != nil {
// 		return err
// 	}
// 	arr := strings.Split(string(dstr), "-")
// 	if len(arr) != 2 {
// 		return errors.New("参数错误")
// 	}

// 	var tlm models.ThirdLogin
// 	id, err := strconv.Atoi(arr[1])
// 	if err != nil {
// 		return errors.New("参数错误")
// 	}
// 	err = SerThirdLogin.GetById(id, &tlm)
// 	if err != nil {
// 		return err
// 	}
// 	if tlm.Id < 1 {
// 		return errors.New("参数错误")
// 	}

// 	if err := SerThirdLogin.UpdateUserId(user.Id, tlm); err != nil {
// 		return err
// 	}
// 	return nil

// }

// 通过验证码
func (e *AdminService) ChangePwd(mobile, email, password string) errs.IError {
	var user models.Admin
	if mobile != "" {
		if err := e.GetByPhone(mobile, &user); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return codes.ErrSys(err)
			} else {
				return errs.ErrWithCode(codes.UserNotExist)
			}
		}
	}
	if email != "" {
		if err := e.GetByEmail(email, &user); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return codes.ErrSys(err)
			} else {
				return errs.ErrWithCode(codes.UserNotExist)
			}
		}
	}
	user.Password = password
	err := e.Update(&user)
	if err != nil {
		errs.AsBizError(err)
	}

	return nil
}

// 微信登录
// func (e *AdminService) LoginWechatMp(req dto.MpSceneReq, openId, ip string) (dto.LoginOK, errs.IError) {

// 	lok := dto.LoginOK{}

// 	var tl models.ThirdLogin
// 	if err := SerThirdLogin.GetTL(3, openId, "", &tl); err != nil {
// 		if !errors.Is(err, gorm.ErrRecordNotFound) {
// 			return lok, codes.ErrSys(err)
// 		}
// 	}
// 	var user models.Admin
// 	if tl.Id < 1 {
// 		tl.OpenId = openId
// 		tl.Platform = 3
// 		tl.CreatedAt = time.Now().Unix()
// 		tl.UpdatedAt = tl.CreatedAt
// 		if err := SerThirdLogin.Create(&tl); err != nil {
// 			return lok, codes.ErrSys(err)
// 		}
// 		needMobile(tl.Platform, tl.Id, &lok)
// 		return lok, nil
// 	} else {
// 		if tl.Id == 0 {
// 			needMobile(tl.Platform, tl.Id, &lok)
// 			return lok, nil
// 		}
// 		if err := e.Get(tl.Id, &user); err != nil {
// 			return lok, codes.ErrSys(err)
// 		}
// 	}
// 	return e.loginOK(&user, 0, enums.LT_WECHAT_MP, ip, "", req.ClientVer, req.Os, req.OsVer)
// }

// 钉钉登录
// func (e *AdminService) LoginDing(c *dto.LoginDingReq, id string) (dto.LoginOK, errs.IError) {
// 	lok := dto.LoginOK{}

// 	if id == "" {
// 		return lok, errs.ErrWithCode(codes.ThirdNotScan)
// 	}
// 	var user models.Admin
// 	var tlModel models.ThirdLogin
// 	if err := SerThirdLogin.GetTL(2, id, "", &tlModel); err != nil {
// 		if !errors.Is(err, gorm.ErrRecordNotFound) {
// 			return lok, codes.ErrSys(err)
// 		}
// 	}

// 	if tlModel.Id < 1 {
// 		tlModel.OpenId = id
// 		tlModel.Platform = 2
// 		tlModel.CreatedAt = time.Now().Unix()
// 		tlModel.UpdatedAt = tlModel.CreatedAt
// 		if err := SerThirdLogin.Create(&tlModel); err != nil {
// 			return lok, codes.ErrSys(err)
// 		}
// 		needMobile(tlModel.Platform, tlModel.Id, &lok)
// 		return lok, nil
// 	} else {
// 		if tlModel.Id == 0 {
// 			needMobile(tlModel.Platform, tlModel.Id, &lok)
// 			return lok, nil
// 		}
// 		if err := e.Get(tlModel.Id, &user); err != nil {
// 			return lok, codes.ErrSys(err)
// 		}
// 	}
// 	return e.loginOK(&user, 0, enums.LT_WECHAT_MP, "", "", c.ClientVer, c.Os, c.OsVer)
// }

// Get 获取User对象
func (e *AdminService) GetByUsername(username string, model *models.Admin) errs.IError {
	var data models.Admin
	err := core.DB().Model(&data).Where("username = ?", username).First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		berr := errs.Err(codes.ErrUserExist, "", err)
		core.Log.Error("admin", "get", berr)
		return berr
	}
	if err != nil {
		core.Log.Error("admin", "err", err)
		return codes.ErrSys(err)
	}
	return nil
}

// func needMobile(platform, id int, lod *dto.LoginOK) error {
// 	enS, err := cryptos.RSA_Encrypt([]byte(fmt.Sprintf("%d-%d", platform, id)), consts.PubKey)
// 	if err != nil {
// 		return err
// 	}
// 	lod.Need = 1
// 	lod.AccessToken = enS
// 	return nil
// }

// 通过老密码修改
func (e *AdminService) ChangePwdByOld(id int, oldPwd, newPwd, inviteCode string) errs.IError {
	var user models.Admin
	if id != 0 {
		if err := e.Get(id, &user); err != nil {
			return codes.ErrNotFound(strconv.Itoa(id), "admin", "", err)
		}
	}
	if user.Password != "" { //已设置密码
		if !user.CompPwd(oldPwd) {
			return errs.ErrWithCode(codes.ErrPwd)
		}
	}

	user.Password = newPwd
	err := e.Update(&user)
	if err != nil {
		errs.AsBizError(err)
	}
	return nil
}

// 绑定
func (e *AdminService) Bind(id int, c *dto.BindReq) error {
	var user models.Admin
	if id != 0 {
		if err := e.Get(id, &user); err != nil {
			return errors.New("用户不存在")
		}
	}

	updates := models.Admin{}

	if regexps.CheckMobile(c.Username) {
		updates.Phone = c.Username
		var count int64
		if err := e.CountByPhone(c.Username, &count); err != nil {
			return err
		}
		if count > 0 {
			return errors.New("该手机号已存在")
		}
	} else if regexps.CheckEmail(c.Username) {
		updates.Email = c.Username
		var count int64
		if err := e.CountByEmail(c.Username, &count); err != nil {
			return err
		}
		if count > 0 {
			return errors.New("该邮箱已存在")
		}
	} else {
		return errors.New("请输入正确的手机号或者邮箱")
	}

	db := core.DB().Model(user).Updates(updates)
	if err := db.Error; err != nil {
		core.Log.Error("adminService", "update", err)
		return err
	}
	return nil
}

func (e *AdminService) Create(model *models.Admin) error {
	if model.RoleId == enums.RolePlatformSuper {
		return errors.New("超级管理员角色不允许创建")
	}
	if model.DeptId != 0 {
		var dept models.AdminDept
		if err := SerAdminDept.Get(model.DeptId, &dept); err != nil {
			return err
		}
		model.DeptPath = dept.DeptPath
	}
	return e.BaseService.Create(model)
}
