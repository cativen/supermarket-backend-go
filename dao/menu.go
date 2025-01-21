package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type MenuDao interface {
	SelectAllMenu() []model.Menu
	SelectMenuPageByQo(dto dto.MenuQueryDTO) vo.Page[model.Menu]
	SelectListByQo(menu model.Menu) []model.Menu
}

type MenuDaoImpl struct {
}

func (menuDao *MenuDaoImpl) SelectAllMenu() []model.Menu {
	db := common.GetDB()
	var catalogMenus []model.Menu
	db.Where("type = ? and state= ?", common.TYPE_CATALOGUE, common.STATE_NORMAL).Find(&catalogMenus)

	//对父级菜单进行遍历
	if catalogMenus == nil {
		return catalogMenus
	}

	for j, val := range catalogMenus {
		var menuTypeMenus []model.Menu
		db.Where("type = ? and state= ? and parent_id = ?", common.TYPE_MENU, common.STATE_NORMAL, val.ID).Find(&menuTypeMenus)
		for k, menuTypeVal := range menuTypeMenus {
			var buttonMenus []model.Menu
			db.Where("type = ? and state= ? and parent_id = ?", common.TYPE_BUTTON, common.STATE_NORMAL, menuTypeVal.ID).Find(&buttonMenus)
			if buttonMenus != nil {
				menuTypeMenus[k].Children = buttonMenus
			}
		}
		if menuTypeMenus != nil {
			catalogMenus[j].Children = menuTypeMenus
		}
	}
	return catalogMenus
}

func (menuDao *MenuDaoImpl) SelectMenuPageByQo(dto dto.MenuQueryDTO) vo.Page[model.Menu] {
	var page vo.Page[model.Menu]
	db := common.GetDB()
	//分页查询出所有的菜单记录
	db = db.Table("t_menu").Select("*").Where("type = ?", common.TYPE_CATALOGUE)

	name := dto.Name
	if name != "" && len(name) > 0 {
		db = db.Where("label like ?", "%"+name+"%")
	}

	// 查询菜单列表
	var records []model.Menu
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (menuDao *MenuDaoImpl) SelectListByQo(menu model.Menu) []model.Menu {
	db := common.GetDB()
	var catalogMenus []model.Menu
	db.Where(menu).Find(&catalogMenus)
	return catalogMenus
}
