package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"dilu/common/codes"
	"dilu/common/config"
	"dilu/common/third/sms/aliyun"
	"dilu/common/third/sms/jg"
	"dilu/common/third/tencent"
	"dilu/common/utils"
	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	coreConst "github.com/baowk/dilu-core/common/consts"

	"dilu/common/regexps"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"gorm.io/gorm"
)

type SysUser struct {
	*base.BaseService
}

var SerSysUser = SysUser{
	base.NewService(coreConst.DB_DEF),
}

func (e *SysUser) Page(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {

	db := core.DB().Model(&models.SysUser{})
	if c.Id > 0 {
		db = db.Where("id = ?", c.Id)
	}

	if c.Username != "" {
		db = db.Where("username LIKE =", c.Username)
	}

	if c.Nickname != "" {
		db = db.Where("nickname LIKE ?", c.Nickname+"%")
	}

	if c.Phone != "" {
		db = db.Where("phone = ?", c.Phone)
	}

	if c.Email != "" {
		db = db.Where("email = ?", c.Email)
	}

	if c.Status != 0 {
		db = db.Where("status = ?", c.Status)
	}

	if c.Gender != "" {
		db = db.Where("gender = ?", c.Gender)
	}

	db.Order("updated_at DESC")

	if err := db.Limit(c.GetSize()).Offset(c.GetOffset()).Find(list).
		Limit(-1).Offset(-1).Count(count).Error; err != nil {
		return err
	}
	return nil
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {
	return core.DB().Debug().Preload("Dept").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
}

// Get 获取SysUser对象
func (e *SysUser) Get(id int, model *models.SysUser) error {
	var data models.SysUser

	err := core.DB().Model(&data).Debug().
		First(model, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		return err
	}
	return err
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq) error {
	var err error
	var model models.SysUser
	db := core.DB().First(&model, c.Id)
	if err = db.Error; err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = core.DB().Table(model.TableName()).Where("id =? ", c.Id).Updates(c).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq) error {
	var err error
	var model models.SysUser
	db := core.DB().First(&model, c.Id)
	if err = db.Error; err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = core.DB().Table(model.TableName()).Where("id =? ", c.Id).Updates(c).Error
	if err != nil {
		return err
	}
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq) error {
	var err error
	var model models.SysUser
	db := core.DB().First(&model, c.Id)
	if err = db.Error; err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	c.Generate(&model)
	err = core.DB().Omit("username", "nick_name", "phone", "role_id", "avatar", "sex").Save(&model).Error
	if err != nil {
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *SysUser) Remove(id int) error {
	var err error
	var data models.SysUser

	db := core.DB().Model(&data).Delete(&data, id)
	if err = db.Error; err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string) errs.IError {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = core.DB().Model(c).Select("Id", "Password").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrNotFound(strconv.Itoa(id), "sysuser", "", err)
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
	return nil
}

// 手机注册次数
func (e *SysUser) CountByPhone(phone string, count *int64) error {
	var data models.SysUser
	return core.DB().Model(&data).Where("phone = ?", phone).Count(count).Error

}

// 手机注册次数
func (e *SysUser) CountByEmail(email string, count *int64) error {
	var data models.SysUser
	return core.DB().Model(&data).Where("email = ?", email).Count(count).Error
}

// 注册用户
func (e *SysUser) Register(loginType enums.LoginType, c *dto.RegisterReq, ip string, teamId, appId int) (dto.LoginOK, errs.IError) {
	model := models.SysUser{}
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
	model.SrcId = c.TID
	model.ClientId = c.DevId
	model.TeamId = teamId
	//model.AppId = appId
	if c.InviteCode != "" {
		iu, err := e.GetByInviteCode(c.InviteCode)
		if err == nil {
			model.Inviter = iu.Id
		}
	}
	err := e.create(&model)
	if err != nil {
		core.Log.Error("UserService", "Create", err)
		return lok, codes.ErrSys(err)
	}
	return e.loginOK(&model, 0, loginType, ip, "", c.Ver, c.Os, c.OsVer, teamId, appId)
}

func (e *SysUser) GetByInviteCode(inviteCode string) (*models.SysUser, errs.IError) {
	var model models.SysUser
	err := core.DB().Model(&model).Where("invite_code =?", inviteCode).First(&model).Error
	if err != nil {
		return nil, codes.ErrSys(err)
	}
	return &model, nil
}

func (e *SysUser) loginOK(u *models.SysUser, need int, loginType enums.LoginType, ip, location, ver, os, osVer string, teamId, appId int) (dto.LoginOK, errs.IError) {
	now := time.Now()
	exp := now.Add(time.Duration(core.Cfg.JWT.Expires) * time.Minute)

	token, err := utils.GenerateAppToken(u.Id, core.Cfg.JWT.Issuer, core.Cfg.JWT.Subject, core.Cfg.JWT.SignKey, exp)
	lok := dto.LoginOK{}
	if err != nil {
		return lok, errs.Err(codes.FAILURE, "", err)
	}
	lok.Expire = exp
	lok.AccessToken = token
	lok.Need = need
	// if u.TeamId != 0 {
	// 	lok.Roles = []string{strconv.Itoa(u.TeamId)}
	// }

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
	refT, _ := utils.GenerateAppRefreshToken(u.Id, core.Cfg.JWT.Issuer, core.Cfg.JWT.Subject, core.Cfg.JWT.SignKey, refExp)
	lok.RefreshToken = refT
	lok.RefreshExpire = refExp

	// 记录登录日志
	ll := models.SysUserLoginLog{
		Uid:       uint(u.Id),
		Method:    int8(loginType),
		Ip:        ip,
		Region:    location,
		ClientId:  u.ClientId,
		Ver:       ver,
		Os:        os,
		OsVer:     osVer,
		UpdatedAt: time.Now(),
		//TeamId:    teamId,
		//AppId: appId,
	}
	go SerSysUserLoginLog.Create(&ll)
	return lok, nil
}

// 通过密码登录
func (e *SysUser) LoginPwd(c *dto.LoginReq, ip string, teamId, appId int) (dto.LoginOK, errs.IError) {
	model := models.SysUser{}
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
	return e.loginOK(&model, 0, loginType, ip, "", c.Ver, c.Os, c.OsVer, teamId, appId)
}

func (e *SysUser) RefreshToken(c *dto.RefreshTokenReq, ip string, teamId, appId int) (dto.LoginOK, errs.IError) {
	var cc utils.AppClaims
	if err := utils.ParseAppRefresh(c.RefreshToken, &cc, core.Cfg.JWT.SignKey); err != nil {
		return dto.LoginOK{}, codes.Err401(err)
	}
	var user models.SysUser
	if err := e.Get(cc.Uid, &user); err != nil {
		return dto.LoginOK{}, codes.Err401(err)
	}

	if user.LockTime.After(time.Now()) {
		return dto.LoginOK{}, codes.Err401(errors.New("账号已被锁定，请联系客服"))
	}

	if user.Status > 0 {
		return e.loginOK(&user, 0, enums.LT_UNKNOWN, ip, "", "", "", "", teamId, appId)
	}
	return dto.LoginOK{}, codes.Err401(errors.New("账号已被锁定，请联系客服"))
}

// 通过验证码
func (e *SysUser) LoginCode(c *dto.LoginReq, ip string, teamId, appId int) (dto.LoginOK, errs.IError) {
	var model models.SysUser
	lok := dto.LoginOK{}
	var name string
	loginType := enums.LT_UNKNOWN
	if regexps.CheckMobile(c.Username) {
		if err := e.GetByPhone(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, codes.ErrSys(err)
			} else {
				model.Phone = c.Username
				name = c.Username
			}
		}
		loginType = enums.LT_PHONE
	} else if regexps.CheckEmail(c.Username) { //是否邮箱
		if err := e.GetByEmail(c.Username, &model); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return lok, codes.ErrNotFound(c.Username, "sysuser", "", err)
			} else {
				model.Email = c.Username
				arr := strings.Split(c.Username, "@")
				name = arr[0]
			}
		}
		loginType = enums.LT_EMAIL
	} else {
		return lok, errs.ErrWithCode(codes.ErrMobileOrEmail)
	}
	if model.Id == 0 {
		model.CreatedAt = time.Now()
		model.UpdatedAt = model.CreatedAt
		model.Nickname = name
		model.RegIp = ip
		model.SrcId = c.TID
		model.ClientId = c.DevId
		//model.AppId = appId
		model.TeamId = teamId
		if c.InviteCode != "" {
			iu, err := e.GetByInviteCode(c.InviteCode)
			if err == nil {
				model.Inviter = iu.Id
			}
		}
		err := e.create(&model)
		if err != nil {
			core.Log.Error("sysuser", "create", err)
			return lok, codes.ErrSys(err)
		}
		if c.UUID != "" {
			//e.bindById(c.UUID, model)
		}
		return e.loginOK(&model, 2, loginType, ip, "", c.Ver, c.Os, c.OsVer, teamId, appId)
	}
	if c.UUID != "" {
		//e.bindById(c.UUID, model)
	}
	return e.loginOK(&model, 0, loginType, ip, "", c.Ver, c.Os, c.OsVer, teamId, appId)
}

func (e *SysUser) GetOneClickToken(req *dto.GetOneClickTokenReq) (*aliyun.GateVerifyResultDTO, error) {
	switch config.Ext.OneClick {
	case "ali":
		return aliyun.VerifyMobile(req.AccessCode, req.Phone, req.OutId)
	case "tencent":
		return nil, nil
	default:
		return nil, errors.New("sms select is not support")
	}
}

func (s *SysUser) OneClickLogin(req *dto.OneClickLoginReq, ip string, teamId, appId int) (*dto.LoginOK, error) {
	var phone string
	var err error
	switch config.Ext.OneClick {
	case "ali":
		phone, err = aliyun.GetMobile(req.VerifyToken, req.OutId)
	case "tencent":
		phone, err = tencent.New().Login(req.Carrier, req.VerifyToken, req.Platform)
	case "jg":
		phone, err = jg.GetMobile(req.VerifyToken, req.OutId)
	default:
		return nil, errors.New("sms select is not support")
	}

	if err != nil {
		return nil, err
	}
	fmt.Println(phone)
	if phone == "" {
		return nil, errors.New("phone is empty")
	}

	var user models.SysUser
	if err := s.GetByPhone(phone, &user); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if user.Id == 0 {
		user = models.SysUser{
			Phone:     phone,
			ClientId:  req.DevId,
			SrcId:     req.Src,
			RegIp:     ip,
			CreatedAt: time.Now(),
			//AppId:     appId,
			TeamId: teamId,
			Status: 1,
		}
		user.UpdatedAt = user.CreatedAt

		if err := s.create(&user); err != nil {
			return nil, err
		}
	} else if user.LockTime.Unix() > time.Now().Unix() {
		return nil, fmt.Errorf("账号已被锁定，请联系客服")
	}

	o, err := s.loginOK(&user, 0, enums.LT_ONE_CLICK, ip, "", req.Ver, req.Os, req.OsVer, teamId, appId)
	return &o, err
}

func (e *SysUser) GetFusionToken(req *dto.GetFusionTokenReq) (string, error) {
	switch config.Ext.Sms.Select {
	case "ali":
		return aliyun.GetFusionAuthToken(req.Platform)
	case "tencent":
		return "", nil
	default:
		return "", errors.New("sms select is not support")
	}
}

func (s *SysUser) VerifyFusion(req *dto.OneClickLoginReq, ip string, teamId, appId int) (*dto.LoginOK, error) {
	var phone string
	var err error
	switch config.Ext.Sms.Select {
	case "ali":
		phone, err = aliyun.VerifyWithFusionAuthToken(req.VerifyToken)
	case "tencent":
		phone, err = tencent.New().Login(req.Carrier, req.VerifyToken, req.Platform)
	default:
		return nil, errors.New("sms select is not support")
	}

	if err != nil {
		return nil, err
	}
	fmt.Println(phone)
	if phone == "" {
		return nil, errors.New("phone is empty")
	}

	var user models.SysUser
	if err := s.GetByPhone(phone, &user); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if user.Id == 0 {
		user = models.SysUser{
			Phone:     phone,
			ClientId:  req.DevId,
			SrcId:     req.Src,
			RegIp:     ip,
			CreatedAt: time.Now(),
			//AppId:     appId,
			TeamId: teamId,
			Status: 1,
		}
		user.UpdatedAt = user.CreatedAt

		if err := s.create(&user); err != nil {
			return nil, err
		}
	} else if user.LockTime.Unix() > time.Now().Unix() {
		return nil, fmt.Errorf("账号已被锁定，请联系客服")
	}

	o, err := s.loginOK(&user, 0, enums.LT_ONE_CLICK, ip, "", req.Ver, req.Os, req.OsVer, teamId, appId)
	return &o, err
}

func (s *SysUser) create(user *models.SysUser) error {
	if user.Status == 0 {
		user.Status = 1
	}
	if err := s.Create(user); err != nil {
		return err
	}
	return nil
}

// Get 获取User对象
func (e *SysUser) GetByEmail(email string, model *models.SysUser) error {
	return core.DB().Where("email = ?", email).First(model).Error
}

// Get 获取User对象
func (e *SysUser) GetByPhone(mobile string, model *models.SysUser) error {
	return core.DB().Where("phone = ?", mobile).First(model).Error
}

// func (e *SysUser) bindById(enCode string, user models.SysUser) error {
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
func (e *SysUser) ChangePwd(mobile, email, password string) errs.IError {

	var user models.SysUser
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
	enPwd, err := user.GenPwd(password)
	if err != nil {
		return codes.ErrSys(err)
	}
	updates := models.SysUser{
		Password: string(enPwd),
	}
	updates.UpdatedAt = time.Now()
	db := core.DB().Model(&user).Updates(updates)
	if err = db.Error; err != nil {
		core.Log.Error("sysuser", "update", err)
		return codes.ErrSys(err)
	}
	return nil
}

// 微信登录
// func (e *SysUser) LoginWechatMp(req dto.MpSceneReq, openId, ip string) (dto.LoginOK, errs.IError) {

// 	lok := dto.LoginOK{}

// 	var tl models.ThirdLogin
// 	if err := SerThirdLogin.GetTL(3, openId, "", &tl); err != nil {
// 		if !errors.Is(err, gorm.ErrRecordNotFound) {
// 			return lok, codes.ErrSys(err)
// 		}
// 	}
// 	var user models.SysUser
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
// func (e *SysUser) LoginDing(c *dto.LoginDingReq, id string) (dto.LoginOK, errs.IError) {
// 	lok := dto.LoginOK{}

// 	if id == "" {
// 		return lok, errs.ErrWithCode(codes.ThirdNotScan)
// 	}
// 	var user models.SysUser
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
func (e *SysUser) GetByUsername(username string, model *models.SysUser) errs.IError {
	// str, err := core.Cache.Get("username:" + username)
	// fmt.Println("get:" + str)
	// if err == nil && str != "" {
	// 	if err := json.Unmarshal([]byte(str), model); err == nil {
	// 		return nil
	// 	} else {
	// 		fmt.Println("err:" + err.Error())
	// 	}
	// }
	var data models.SysUser
	err := core.DB().Model(&data).Where("username = ?", username).First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		berr := errs.Err(codes.ErrUserExist, "", err)
		core.Log.Error("sysuser", "get", berr)
		return berr
	}
	if err != nil {
		core.Log.Error("sysuser", "err", err)
		return codes.ErrSys(err)
	}
	// if err := core.Cache.Set("username:"+username, model, time.Hour); err == nil {
	// 	fmt.Println("set================")

	// } else {
	// 	fmt.Println(err)
	// }
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
func (e *SysUser) ChangePwdByOld(id int, oldPwd, newPwd, inviteCode string) errs.IError {
	var user models.SysUser
	if id != 0 {
		if err := e.Get(id, &user); err != nil {
			return codes.ErrNotFound(strconv.Itoa(id), "sysuser", "", err)
		}
	}
	if user.Password != "" { //已设置密码
		if !user.CompPwd(oldPwd) {
			return errs.ErrWithCode(codes.ErrPwd)
		}
	}
	enPwd, err := user.GenPwd(newPwd)
	if err != nil {
		return codes.ErrSys(err)
	}
	updates := models.SysUser{
		Password: string(enPwd),
	}
	db := core.DB().Model(user).Updates(updates)
	if err = db.Error; err != nil {
		core.Log.Error("UserService", "update", err)
		return codes.ErrSys(err)
	}
	return nil
}

// 绑定
func (e *SysUser) Bind(id int, c *dto.BindReq) error {
	var user models.SysUser
	if id != 0 {
		if err := e.Get(id, &user); err != nil {
			return errors.New("用户不存在")
		}
	}

	updates := models.SysUser{}

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
		core.Log.Error("UserService", "update", err)
		return err
	}
	return nil
}

// 修改用户信息
func (e *SysUser) ChangeUserinfo(userId int, user models.SysUser) error {
	if user.Password != "" {
		enPwd, err := user.GenPwd(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(enPwd)
	}
	user.UpdateBy = userId
	err := e.UpdateById(user)
	if err != nil {
		core.Log.Error("UserService", "update", err)
		return err
	}
	return nil
}
