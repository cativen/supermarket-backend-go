package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/service"
	"superMarket-backend/utils"
)

var goodService service.IGoodsService = &service.GoodsServiceImpl{}

func SelectedGoodsAll(c *gin.Context) {
	goodService.SelectedGoodsAll(c)
}

func QueryGoodPageByQo(c *gin.Context) {
	var dto dto.QueryGoodsDTO
	c.ShouldBind(&dto)
	goodService.QueryGoodPageByQo(c, dto)
}

func SaveGoods(c *gin.Context) {
	var good model.Goods
	token := c.GetHeader("token")
	c.ShouldBind(&good)
	goodService.SaveGoods(c, good, token)
}

func UpOrDownGoods(c *gin.Context) {
	gid, _ := c.GetPostForm("gid")
	state, _ := c.GetPostForm("state")
	token := c.GetHeader("token")
	goodService.UpOrDownGoods(c, gid, state, token)
}

func QueryGoodsById(c *gin.Context) {
	id := c.Query("id")
	goodService.QueryGoodsById(c, id)
}

func SelectedStoreAll(c *gin.Context) {
	goodService.SelectedStoreAll(c)
}

func UploadGoodImg(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		log.Println(file.Filename)
		dst, newName := utils.UploadUrl(file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		var imgMap = make(map[string]interface{})
		imgMap["uploaded"] = 1 //成功
		imgMap["url"] = "/static/img" + newName
		c.JSON(200, imgMap)
	}
}

func UpdateGoods(c *gin.Context) {
	var good model.Goods
	token := c.GetHeader("token")
	c.ShouldBind(&good)
	goodService.UpdateGoods(c, good, token)
}

func GoPayById(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	goodService.GoPayById(c, id)
}

func Return(c *gin.Context) {
	response.Success(c, "pay success", "操作成功")
}
