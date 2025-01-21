package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/smartwalle/alipay/v3"
	"log"
	"os/exec"
	"strings"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
	"time"
)

var goodDao dao.GoodsDao = &dao.GoodsDaoImpl{}
var goodsCategoryDao dao.GoodsCategoryDao = &dao.GoodsCategoryImpl{}

type IGoodsService interface {
	SelectedGoodsAll(c *gin.Context)
	DelIn(c *gin.Context, cn string)
	QueryOptionsSuppliers(c *gin.Context)
	QueryGoodPageByQo(c *gin.Context, dto dto.QueryGoodsDTO)
	SaveGoods(c *gin.Context, good model.Goods, token string)
	UpOrDownGoods(c *gin.Context, gid string, state string, token string)
	QueryGoodsById(c *gin.Context, id string)
	SelectedStoreAll(c *gin.Context)
	UpdateGoods(c *gin.Context, good model.Goods, token string)
	QueryPageStatisticSaleByQo(c *gin.Context, saleDTO dto.QueryStatisticSaleDTO)
	GoPayById(c *gin.Context, id string)
}

type GoodsServiceImpl struct {
}

func (GoodsServiceImpl) QueryPageStatisticSaleByQo(c *gin.Context, saleDTO dto.QueryStatisticSaleDTO) {
	total := goodDao.QueryPageStatisticSaleByQo(saleDTO.Name)
	var vos vo.SalesStatisticsVo
	vos.Total = total
	var pageQueryDto = dto.QueryGoodsDTO{
		BaseQuery: dto.BaseQuery{
			PageSize:    saleDTO.PageSize,
			CurrentPage: saleDTO.CurrentPage,
		},
		State: common.STATE_UP,
		Name:  saleDTO.Name,
	}
	page := goodDao.SelectPageListByQo(pageQueryDto)
	var saleGoodsVoPage = vo.Page[vo.SaleGoodsVo]{
		Current: page.Current,
		Size:    page.Size,
		Total:   page.Total,
	}
	var saleGoodsVos = make([]vo.SaleGoodsVo, 0)
	records := page.Records
	for _, val := range records {
		var goodsVo vo.SaleGoodsVo
		goodsVo.GoodsId = val.ID
		goodsVo.GoodsName = val.Name
		goodsVo.SalesVolume = val.SalesVolume
		goodsVo.Percentage = total
		goodsVo.CoverUrl = val.CoverUrl
		saleGoodsVos = append(saleGoodsVos, goodsVo)
	}
	saleGoodsVoPage.Records = saleGoodsVos
	vos.VOS = saleGoodsVoPage
	response.Success(c, vos, "操作成功")
}

func (GoodsServiceImpl) GoPayById(c *gin.Context, id string) {
	byId := goodDao.SelectById(utils.ConvertStringToInt64(id))
	WebPageAlipay(byId)
}

// 网站扫码支付
func WebPageAlipay(good model.Goods) {
	client := common.GetPayClient()
	pay := alipay.TradePagePay{}
	// 支付成功之后，支付宝将会重定向到该 URL
	pay.ReturnURL = "http://localhost:9291/goods_management/goods/return"
	//支付标题
	pay.Subject = good.Name + "支付宝支付测试"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()
	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//金额
	pay.TotalAmount = fmt.Sprintf("%0.2f", good.SellPrice)
	url, err := client.TradePagePay(pay)
	if err != nil {
		fmt.Println(err)
	}
	payURL := url.String()
	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(payURL)

	//打开默认浏览器
	payURL = strings.Replace(payURL, "&", "^&", -1)
	exec.Command("cmd", "/c", "start", payURL).Start()
}

func (GoodsServiceImpl) SelectedGoodsAll(c *gin.Context) {
	var goods = model.Goods{
		State: common.STATE_UP,
	}
	list := goodDao.SelectListByQo(goods)
	if len(list) == 0 {
		response.Success(c, nil, "操作成功")
		return
	}
	var listVo = make([]map[string]interface{}, 0)
	for _, val := range list {
		goodMap := make(map[string]interface{})
		goodMap["id"] = val.ID
		goodMap["name"] = val.Name
		listVo = append(listVo, goodMap)
	}
	response.Success(c, listVo, "操作成功")
}

func (GoodsServiceImpl) DelIn(c *gin.Context, cn string) {
	var model = model.DetailStoreGoods{
		Cn:     cn,
		State1: common.STATE_DEL,
	}
	detailStoreGoodsDao.UpdateStoreGoods(model)
}

func (GoodsServiceImpl) QueryOptionsSuppliers(c *gin.Context) {
	var supplierDao dao.SupplierDao = &dao.SupplierDaoImpl{}
	var model = model.Supplier{
		State: common.STATE_NORMAL,
	}
	suppliers := supplierDao.SelectListByQo(model)
	if len(suppliers) == 0 {
		response.Success(c, nil, "操作成功")
		return
	}
	var listVo = make([]map[string]interface{}, len(suppliers))
	for index, val := range suppliers {
		supplierMap := make(map[string]interface{})
		supplierMap["id"] = val.Cn
		supplierMap["name"] = val.Name
		listVo[index] = supplierMap
	}
	response.Success(c, listVo, "操作成功")
}

func (GoodsServiceImpl) QueryGoodPageByQo(c *gin.Context, dto dto.QueryGoodsDTO) {
	page := goodDao.SelectPageListByQo(dto)
	records := page.Records
	for index, record := range records {
		residueNum := storeDao.SelectResidueNumByGoodsId(record.ID)
		records[index].ResidueNum = residueNum
	}
	page.Records = records
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	response.Success(c, page, "操作成功")
}

func (GoodsServiceImpl) SaveGoods(c *gin.Context, good model.Goods, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println(err)
	}
	var existEmployee model.Employee
	if err := json.Unmarshal([]byte(val), &existEmployee); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}
	good.State = common.STATE_UP
	good.CreateBy = existEmployee.NickName
	good.UpdateBy = existEmployee.NickName
	good.CreateTime = time.Now()
	good.UpdateTime = time.Now()
	if good.CategoryID != 0 {
		//从缓存中获取分类的信息
		val, err := rdb.Get(ctx, common.GOODS_CATEGORY).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				// 如果返回的错误是key不存在
				log.Println(val)
			} else {
				//查询数据库
				var category = model.GoodsCategory{
					ID: good.CategoryID,
				}
				categoryRes := goodsCategoryDao.SelectByQo(category)
				if categoryRes != (model.GoodsCategory{}) {
					good.CategoryName = categoryRes.Name
				}
			}
		}
	}
	goodsDao.SaveGoods(good)
}

func (GoodsServiceImpl) UpOrDownGoods(c *gin.Context, gid string, state string, token string) {
	var queryGoods = model.Goods{
		ID: uint(utils.ConvertStringToInt64(gid)),
	}
	var UpdateGoods = model.Goods{}
	if state == common.STATE_UP {
		rdb := common.GetRDB()
		ctx := context.Background()
		val, err := rdb.Get(ctx, token).Result()
		if err != nil {
			log.Println(err)
		}
		var existEmployee model.Employee
		if err := json.Unmarshal([]byte(val), &existEmployee); err != nil {
			response.Error(c, "token已过期需要重新登录")
			return
		}
		UpdateGoods.State = common.STATE_DOWN
		goods := goodDao.SelectById(utils.ConvertStringToInt64(gid))
		var goodsStore = model.GoodsStore{
			GoodsID: uint(utils.ConvertStringToInt64(gid)),
		}
		list := goodsStoreDao.SelectListByQo(goodsStore)
		for _, val := range list {
			var detailStoreGoods model.DetailStoreGoods
			detailStoreGoods.CreateID = existEmployee.ID
			detailStoreGoods.CreateBy = existEmployee.NickName
			detailStoreGoods.CreateTime = time.Now()
			detailStoreGoods.GoodsID = utils.ConvertStringToInt64(gid)
			detailStoreGoods.GoodsName = goods.Name
			detailStoreGoods.Type = common.TYPE_IN
			detailStoreGoods.State1 = common.STATE1_UNTREATED
			detailStoreGoods.State = common.STATE_DOWN
			snowflake, _ := utils.NewSnowflake(1)
			generatedId, _ := snowflake.Generate()
			detailStoreGoods.Cn = utils.ConvertInt64ToString(generatedId)
			detailStoreGoods.Info = fmt.Sprintf("%s%s", goods.Name, "下架处理")
			detailStoreGoods.GoodsNum = val.ResidueNum
			detailStoreGoods.UntreatedNum = val.ResidueNum
			detailStoreGoods.StoreID = val.StoreID
			detailStoreGoodsDao.SaveData(detailStoreGoods)
		}
	} else {
		UpdateGoods.ResidueNum = 0
		UpdateGoods.State = common.STATE_UP
		var deleteStoreGoods = model.DetailStoreGoods{
			GoodsID: utils.ConvertStringToInt64(gid),
			State:   common.STATE_DOWN,
			State1:  common.STATE1_UNTREATED,
		}
		detailStoreGoodsDao.DeleteByQo(deleteStoreGoods)
	}
	goodDao.UpdateDataByQO(queryGoods, UpdateGoods)
	response.Success(c, nil, "操作成功")
}

func (GoodsServiceImpl) QueryGoodsById(c *gin.Context, id string) {
	byId := goodDao.SelectById(utils.ConvertStringToInt64(id))
	response.Success(c, byId, "操作成功")
}

func (GoodsServiceImpl) SelectedStoreAll(c *gin.Context) {
	var list = make([]map[string]interface{}, 0)
	queryList := storeDao.SelectListByQueryQo(model.Store{State: common.STATE_NORMAL})
	if len(queryList) > 0 {
		for _, val := range queryList {
			var storeMap = make(map[string]interface{})
			storeMap["id"] = val.ID
			storeMap["name"] = val.Name
			list = append(list, storeMap)
		}
	}
	response.Success(c, list, "操作成功")
}

func (GoodsServiceImpl) UpdateGoods(c *gin.Context, good model.Goods, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println(err)
	}
	var existEmployee model.Employee
	if err := json.Unmarshal([]byte(val), &existEmployee); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}
	good.UpdateBy = existEmployee.NickName
	good.UpdateTime = time.Now()
	if good.CategoryID != 0 {
		category := goodsCategoryDao.SelectByQo(model.GoodsCategory{ID: good.CategoryID})
		if category != (model.GoodsCategory{}) {
			good.CategoryName = category.Name
		}
	}
	var queryGoods = model.Goods{
		ID: good.ID,
	}
	goodDao.UpdateDataByQO(queryGoods, good)
	response.Success(c, nil, "操作成功")
}
