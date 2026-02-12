package service

import (
	"dilu/common/ckey"
	"dilu/common/codes"
	"dilu/common/utils"
	"encoding/json"
	"fmt"
	"time"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
)

type SysMenu struct {
	*base.BaseService
}

var SerSysMenu = SysMenu{
	base.NewService(consts.DB_DEF),
}

// Get 获取SysMenu对象
func (e *SysMenu) GetAndApi(id int, model *models.SysMenu) (*SysMenu, errs.IError) {
	var err error
	var data models.SysMenu

	db := e.DB().Model(&data).Preload("SysApi").
		First(model, id)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	if err != nil {
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	return e, nil
}

func (e *SysMenu) Get(id int, model *models.SysMenu) error {
	return e.DB().First(model, id).Error
}

// Insert 创建SysMenu对象
func (e *SysMenu) Insert(data *models.SysMenu) errs.IError {
	var err error
	tx := core.DB().Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		berr := errs.Err(codes.FAILURE, "", err)
		return berr
	}
	tx.Commit()
	return nil
}

// Remove 删除SysMenu
func (e *SysMenu) Remove(d *dto.SysMenuDeleteReq) (*SysMenu, errs.IError) {
	var err error
	var data models.SysMenu

	db := core.DB().Model(&data).Delete(&data, d.Ids)
	if err = db.Error; err != nil {
		err = db.Error
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	return e, nil
}

func (e *SysMenu) GetMenus(req *dto.SysMenuReq, mvs *[]models.SysMenu) errs.IError {
	db := e.DB()
	if req.MenuType != 0 {
		db = db.Where("menu_type = ?", req.MenuType)
	}
	if req.Title != "" {
		db = db.Where("title like ?", "%"+req.Title+"%")
	}

	if err := db.Where("platform_type = ?", req.PlatformType).Order("sort asc").Find(mvs).Error; err != nil {
		return codes.ErrSys(err)
	}
	return nil
}

func (e *SysMenu) GetUserMenus(teamId, uid int, mvs *[]dto.MenuVo) error {
	var member models.SysMember
	if err := SerSysMember.GetMember(teamId, uid, &member); err != nil {
		return err
	}

	if member.RoleId < 1 {
		return codes.Err403(nil)
	}
	data, _ := core.Cache.Get(ckey.RoleUserMenu(int(enums.MenuPriTeam), member.RoleId))
	if data != "" {
		err := json.Unmarshal([]byte(data), mvs)
		if err == nil {
			return nil
		}
	}

	db := e.DB().Where("platform_type = ?", enums.MenuPriTeam).Order("sort asc").Order("parent_id asc")
	if member.RoleId > enums.RoleTeamOwner {
		db.Joins(" left join team_role_menu on team_role_menu.menu_id = sys_menu.id").
			Where("team_role_menu.role_id = ?", member.RoleId)
	}

	var ms []models.SysMenu
	if err := db.Find(&ms).Error; err != nil {
		return codes.ErrSys(err)
	}
	*mvs = e.treeMenuVo(ms)
	core.Cache.Set(ckey.RoleUserMenu(int(enums.MenuPriTeam), member.RoleId), mvs, time.Hour)
	return nil
}

func (e *SysMenu) TeamUserCanAccess(teamId, uid int, method, url string) error {
	if teamId < 1 || uid < 1 {
		return codes.Err403(nil)
	}
	var member models.SysMember
	if err := SerSysMember.GetMember(teamId, uid, &member); err != nil {
		return err
	}
	if member.RoleId == enums.RoleTeamOwner {
		return nil
	} else if member.RoleId < 1 {
		return codes.Err403(nil)
	}

	var aid int
	var apis []models.SysApi

	SerSysApi.GetByType(3, &apis)
	for _, api := range apis {
		if api.Method == method && utils.KeyMatch2(url, api.Path) {
			aid = api.Id
			break
		}
	}
	if aid < 1 {
		return codes.Err403(nil)
	}
	var ids []int

	if err := SerSysRole.GetApis(member.RoleId, &ids); err != nil {
		return err
	}

	for _, id := range ids {
		if id == aid {
			return nil
		}
	}
	return codes.Err403(nil)
}

func (e *SysMenu) GetAdminMenus(uid int, mvs *[]dto.MenuVo) error {
	var admin models.Admin
	if err := SerAdmin.Get(uid, &admin); err != nil {
		return err
	}
	if admin.RoleId < 1 {
		return errors.New("角色不存在")
	}

	data, _ := core.Cache.Get(ckey.RoleUserMenu(int(enums.MenuPriPlatform), admin.RoleId))
	if data != "" {
		err := json.Unmarshal([]byte(data), mvs)
		if err == nil {
			return nil
		}
	}

	db := e.DB().Where("platform_type = ?", enums.MenuPriPlatform).Order("sort asc").Order("parent_id asc")
	if admin.RoleId > enums.RolePlatformSuper {
		var ar models.AdminRole
		if err := SerAdminRole.Get(admin.RoleId, &ar); err != nil {
			return err
		}
		db.Where("id in ?", []int(ar.MenuIds))
	}

	var ms []models.SysMenu
	if err := db.Find(&ms).Error; err != nil {
		return codes.ErrSys(err)
	}
	*mvs = e.treeMenuVo(ms)

	core.Cache.Set(ckey.RoleUserMenu(int(enums.MenuPriPlatform), admin.RoleId), mvs, time.Hour)

	return nil
}

func (e *SysMenu) AdminCanAccess(adminId int, method, url string) error {
	if adminId < 1 {
		return codes.Err403(nil)
	}

	var admin models.Admin
	if err := SerAdmin.Get(adminId, &admin); err != nil {
		return codes.Err403(nil)
	}

	if admin.RoleId == enums.RolePlatformSuper {
		return nil
	} else if admin.RoleId < 1 {
		return codes.Err403(nil)
	}

	if admin.Id < 1 {
		return codes.Err403(nil)
	}

	var aid int
	var apis []models.SysApi

	SerSysApi.GetByType(3, &apis)
	for _, api := range apis {
		if api.Method == method && utils.KeyMatch2(url, api.Path) {
			aid = api.Id
			break
		}
	}
	//fmt.Println(aid)
	if aid < 1 {
		return codes.Err403(nil)
	}

	var ids []int

	var ar models.AdminRole
	if err := SerAdminRole.Get(admin.RoleId, &ar); err != nil {
		return err
	}
	if err := e.DB().Raw("select sys_api_id from sys_menu_api_rule r where  r.sys_menu_id in ?", []int(ar.MenuIds)).
		Find(&ids).Error; err != nil {
		return err
	}

	for _, id := range ids {
		if id == aid {
			return nil
		}
	}
	return codes.Err403(nil)
}

func (e *SysMenu) treeMenuVo(ms []models.SysMenu) []dto.MenuVo {
	mvs := make([]dto.MenuVo, 0)
	topM := make(map[int]struct{}, 0)
	for _, menu := range ms {
		if menu.ParentId == 0 {
			topM[menu.Id] = struct{}{}
		}
	}

	for _, menu := range ms {
		if menu.MenuType == 2 {
			if _, ok := topM[menu.ParentId]; !ok {
				var pm models.SysMenu
				if err := e.Get(menu.ParentId, &pm); err != nil {
					continue
				}
				vo := menuToVo(pm)
				e.menuCallVo(ms, &vo)
				mvs = append(mvs, vo)
				topM[menu.ParentId] = struct{}{}
			}
		} else if menu.ParentId == 0 {
			vo := menuToVo(menu)
			e.menuCallVo(ms, &vo)
			mvs = append(mvs, vo)
		}
	}
	return mvs
}

// menuCall 构建菜单树
func (e *SysMenu) menuCallVo(ms []models.SysMenu, menu *dto.MenuVo) {
	children := make([]dto.MenuVo, 0)
	for _, m := range ms {
		if menu.Id != m.ParentId {
			continue
		}
		if m.MenuType < 3 {
			vo := menuToVo(m)
			e.menuCallVo(ms, &vo)
			children = append(children, vo)
		} else {
			menu.Meta.Auths = append(menu.Meta.Auths, m.Permission)
		}
	}
	menu.Children = children
}

func menuToVo(menu models.SysMenu) dto.MenuVo {
	meta := dto.RouteMeta{
		Title: menu.Title,
		Icon:  menu.Icon,
	}
	if !menu.Hidden {
		meta.ShowLink = true
	}
	if !menu.NoCache {
		meta.KeepAlive = true
	}
	if menu.Sort > 0 {
		meta.Rank = menu.Sort
	}
	vo := dto.MenuVo{
		Name:      menu.MenuName,
		Meta:      meta,
		Path:      menu.Path,
		Component: menu.Component,
		Id:        menu.Id,
	}
	return vo
}

func (e *SysMenu) GetUserPerms(roleId int, mvs *[]string) errs.IError {
	var sql string
	if roleId == 1 {
		sql = "Select permission from sys_menu  where menu_type > 1 "
	} else {
		sql = fmt.Sprintf("Select permission from sys_menu m,sys_role_menu r where role_id = %d and menu_type > 1 and m.id = r.menu_id", roleId)
	}
	var ms []string
	if err := core.DB().Raw(sql).Find(&ms).Error; err != nil {
		return codes.ErrSys(err)
	}
	*mvs = ms
	return nil
}

func (e *SysMenu) GetApisByPage(menuId int, path string, total *int64, size, offset int) ([]dto.SysApiExt, error) {
	var orderBy = func(menuId int) string {
		if menuId > 0 {
			return "sys_menu_api_rule.sys_menu_id DESC"
		}
		return "sys_api.id DESC"
	}
	var apis []dto.SysApiExt
	db := e.DB().Model(&models.SysApi{}).Select("*")
	if menuId > 0 {
		db.Joins("LEFT JOIN sys_menu_api_rule ON sys_api.id = sys_menu_api_rule.sys_api_id AND sys_menu_api_rule.sys_menu_id = ?", menuId)
	}
	if path != "" {
		db.Where("sys_api.path LIKE ?", "%"+path+"%")
	}
	if err := db.Limit(size).Offset(offset).Order(orderBy(menuId)).Find(&apis).Limit(-1).Offset(-1).Count(total).Error; err != nil {
		return nil, err
	}
	return apis, nil
}

// SetApis 给 menu 设置 api
func (e *SysMenu) SetApis(menuId int, apiIds []int) error {
	for _, id := range apiIds {
		var cnt int64
		if err := e.DB().Model(&models.SysMenuApiRule{}).Where("sys_menu_id = ? and sys_api_id = ?", menuId, id).Count(&cnt).Error; err != nil {
			return err
		}
		if cnt > 0 {
			continue
		} else {
			if err := e.DB().Model(&models.SysMenuApiRule{}).Create(&models.SysMenuApiRule{SysMenuId: uint(menuId), SysApiId: uint(id)}).Error; err != nil {
				return err
			}
		}
	}
	var err error
	if len(apiIds) > 0 {
		err = e.DB().Model(&models.SysMenuApiRule{}).Where("sys_menu_id = ? AND sys_api_id NOT IN (?)", menuId, apiIds).Delete(&models.SysMenuApiRule{}).Error
	} else {
		err = e.DB().Model(&models.SysMenuApiRule{}).Where("sys_menu_id = ?", menuId).Delete(&models.SysMenuApiRule{}).Error
	}
	return err
}
