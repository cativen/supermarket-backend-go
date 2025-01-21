package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
	"time"
)

var exchangePointProductsDao dao.ExchangePointProductsDao = &dao.ExchangePointProductsDaoImpl{}
var memberDao dao.MemberDao = &dao.MemberDaoImpl{}
var goodsDao dao.GoodsDao = &dao.GoodsDaoImpl{}

type IExchangePointProductsService interface {
	QueryPageByQoExchangePointProducts(c *gin.Context, dto dto.QueryExchangePointProductsRecordsDTO)
	QueryOptionsMemberPhone(c *gin.Context)
	DelExchangePointProducts(c *gin.Context, cn string)
	QueryOptionsMember(c *gin.Context)
	QueryOptionsPointProducts(c *gin.Context)
	QueryPointProductByMemberId(c *gin.Context, memberId string)
	QueryMemberByGoodsId(c *gin.Context, goodsId string)
	QueryPointProductByGoodsId(c *gin.Context, id string)
	SaveExchangePointProductRecords(c *gin.Context, products model.ExchangePointProductsRecord, token string)
}

type ExchangePointProductsServiceImpl struct {
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) QueryPageByQoExchangePointProducts(c *gin.Context, dto dto.QueryExchangePointProductsRecordsDTO) {
	page := exchangePointProductsDao.QueryPageByQoExchangePointProducts(dto)
	records := page.Records
	for index, record := range records {
		records[index].UpdateTimeRes = utils.FormatTimeToString(record.UpdateTime)
		member := memberDao.SelectById(record.MemberID)
		records[index].MemberPhone = member.Phone
		goods := goodsDao.SelectById(record.GoodsID)
		records[index].GoodsCoverUrl = goods.CoverURL
		records[index].GoodsName = goods.Name
	}
	page.Records = records
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	response.Success(c, page, "操作成功")
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) QueryOptionsMemberPhone(c *gin.Context) {
	memberIds := exchangePointProductsDao.QueryOptionsMemberPhone()
	memberList := memberDao.SelectMemberListByIds(memberIds)
	var vos = make([]map[string]interface{}, len(memberList))
	for index, val := range memberList {
		var memberMap = make(map[string]interface{})
		memberMap["id"] = val.ID
		memberMap["name"] = val.Phone
		vos[index] = memberMap
	}
	response.Success(c, vos, "操作成功")
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) DelExchangePointProducts(c *gin.Context, cn string) {
	exchangePointProductsDao.DelByCn(cn)
	response.Success(c, "success", "操作成功")
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) QueryOptionsMember(c *gin.Context) {
	list := memberDao.SelectListByStateGroupById()
	vos := make([]map[string]interface{}, len(list))
	for index, val := range list {
		memberMap := make(map[string]interface{})
		memberMap["id"] = val.ID
		memberMap["name"] = val.Phone
		vos[index] = memberMap
	}
	response.Success(c, vos, "操作成功")
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) QueryOptionsPointProducts(c *gin.Context) {
	list := pointProductsDao.SelectListByQo(model.PointProduct{State: common.STATE_NORMAL})
	vos := make([]map[string]interface{}, len(list))
	for index, val := range list {
		pointMap := make(map[string]interface{})
		pointMap["id"] = val.GoodsID
		pointMap["name"] = val.GoodsName
		vos[index] = pointMap
	}
	response.Success(c, vos, "操作成功")
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) QueryPointProductByMemberId(c *gin.Context, memberId string) {
	member := memberDao.SelectById(utils.ConvertStringToInt64(memberId))

	var list []model.PointProduct
	if len(memberId) == 0 {
		list = pointProductsDao.SelectListByGreater(0)
	} else {
		list = pointProductsDao.SelectListByGreater(member.Integral)
	}

	mapArrayList := make([]map[string]interface{}, len(list))
	for index, val := range list {
		pointMap := make(map[string]interface{})
		pointMap["id"] = val.GoodsID
		pointMap["name"] = val.GoodsName
		mapArrayList[index] = pointMap
	}
	response.Success(c, mapArrayList, "操作成功")
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) QueryMemberByGoodsId(c *gin.Context, goodsId string) {
	var members []model.Member
	if len(goodsId) != 0 {
		pointProducts := pointProductsDao.SelectByQo(model.PointProduct{GoodsID: uint(utils.ConvertStringToInt64(goodsId)), State: common.STATE_NORMAL})
		members = memberDao.SelectListByGreaterEqual(pointProducts.Integral)
	} else {
		members = memberDao.SelectListByGreater(0)
	}

	vos := make([]map[string]interface{}, len(members))
	for index, val := range members {
		pointMap := make(map[string]interface{})
		pointMap["id"] = val.ID
		pointMap["name"] = val.Phone
		vos[index] = pointMap
	}
	response.Success(c, vos, "操作成功")
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) QueryPointProductByGoodsId(c *gin.Context, id string) {
	qo := pointProductsDao.SelectByQo(model.PointProduct{GoodsID: uint(utils.ConvertStringToInt64(id))})
	response.Success(c, qo, "操作成功")
}

func (exchangePointProductsService ExchangePointProductsServiceImpl) SaveExchangePointProductRecords(c *gin.Context, exchangePointProducts model.ExchangePointProductsRecord, token string) {
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
	snowflake, _ := utils.NewSnowflake(1)
	generateId, _ := snowflake.Generate()
	exchangePointProducts.CN = utils.ConvertInt64ToString(generateId)
	exchangePointProducts.UpdateBy = existEmployee.NickName
	exchangePointProducts.UpdateID = existEmployee.ID
	exchangePointProducts.UpdateTime = time.Now()
	exchangePointProducts.State = common.STATE_NORMAL
	//修改会员的积分
	member := memberDao.SelectById(exchangePointProducts.MemberID)
	member.Integral = member.Integral - exchangePointProducts.Integral
	memberDao.UpdateMemberByQO(member, model.Member{ID: member.ID})
	exchangePointProductsDao.SaveData(exchangePointProducts)
	response.Success(c, nil, "操作成功")
}
