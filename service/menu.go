package service

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
)

var menuDao dao.MenuDao = &dao.MenuDaoImpl{}

type IMenuService interface {
	FindAll() []model.Menu
	QueryMenuPageByQo(c *gin.Context, dto dto.MenuQueryDTO)
}

type MenuServiceImpl struct {
}

func (MenuServiceImpl) FindAll() []model.Menu {
	var menuDao dao.MenuDao = &dao.MenuDaoImpl{}
	return menuDao.SelectAllMenu()
}

func (MenuServiceImpl) QueryMenuPageByQo(c *gin.Context, dto dto.MenuQueryDTO) {
	page := menuDao.SelectMenuPageByQo(dto)
	catalogs := page.Records
	if len(catalogs) == 0 {
		response.Success(c, page, "操作成功")
		return
	}

	//补全目录下的菜单
	for k, val := range catalogs {
		var menuQuery = model.Menu{
			Type:     common.TYPE_MENU,
			ParentID: val.ID,
		}
		menuTypeMenus := menuDao.SelectListByQo(menuQuery)
		if len(menuTypeMenus) == 0 {
			continue
		}
		//补全菜单下的按钮
		for j, menuTypeVal := range menuTypeMenus {
			var buttonQuery = model.Menu{
				Type:     common.TYPE_BUTTON,
				ParentID: menuTypeVal.ID,
			}
			buttonMenus := menuDao.SelectListByQo(buttonQuery)
			if len(buttonMenus) == 0 {
				continue
			}
			menuTypeMenus[j].Children = buttonMenus
		}
		catalogs[k].Children = menuTypeMenus
	}
	page.Records = catalogs
	response.Success(c, page, "操作成功")
}
