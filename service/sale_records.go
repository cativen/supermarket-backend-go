package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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
	"time"
)

var saleRecordDao dao.SaleRecordsDao = &dao.SaleRecordsDaoImpl{}
var detailSaleRecordsServiceDao dao.DetailSaleRecordsServiceDao = &dao.DetailSaleRecordsServiceImpl{}

type ISaleRecordsService interface {
	QueryPageByQoSaleRecords(c *gin.Context, dto dto.SaleRecordPageDTO)
	DelSaleRecords(c *gin.Context, cn string)
	GetOptionSaleRecordsGoods(c *gin.Context)
	SaveSaleRecords(c *gin.Context, cn string, token string)
	PaySaleItems(record model.SaleRecord, token string)
}

type SaleRecordsServiceImpl struct {
}

func (saleRecordsService SaleRecordsServiceImpl) QueryPageByQoSaleRecords(c *gin.Context, dto dto.SaleRecordPageDTO) {
	page := saleRecordDao.QueryPageByQoSaleRecords(dto)
	response.Success(c, page, "操作成功")
}

func (saleRecordsService SaleRecordsServiceImpl) DelSaleRecords(c *gin.Context, cn string) {
	saleRecordDao.DelSaleRecords(cn)
	response.Success(c, true, "操作成功")
}

func (saleRecordsService SaleRecordsServiceImpl) GetOptionSaleRecordsGoods(c *gin.Context) {
	list := goodDao.SelectHaveResidueGood()
	goodsList := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for _, val := range list {
			goodMap := make(map[string]interface{})
			goodMap["id"] = val.ID
			goodMap["name"] = val.Name
			goodMap["residueNum"] = val.ResidueNum
			goodsList = append(goodsList, goodMap)
		}
	}
	response.Success(c, list, "操作成功")
}

func (saleRecordsService SaleRecordsServiceImpl) SaveSaleRecords(c *gin.Context, cn string, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println(err)
	}
	var employee model.Employee
	if err := json.Unmarshal([]byte(val), &employee); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}
	res, err := rdb.Get(ctx, cn).Result()
	if err != nil {
		log.Println(err)
	}

	var saleRecords model.SaleRecord
	if err := json.Unmarshal([]byte(res), &saleRecords); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}

	saleRecords.EID = employee.ID
	saleRecords.SellTime = time.Now()
	saleRecords.SellBy = employee.NickName
	saleRecords.State = common.STATE_NORMAL
	detailSaleRecords := saleRecords.DetailSaleRecords
	for index, detailSaleRecord := range detailSaleRecords {
		detailSaleRecords[index].SellCN = saleRecords.CN
		detailSaleRecords[index].GoodsNum = utils.ConvertStringToInt64(detailSaleRecord.GoodsNumJson)
		goods := goodDao.SelectById(detailSaleRecord.GoodsID)
		queryModel := model.Goods{
			ID: goods.ID,
		}

		updateModel := model.Goods{
			SalesVolume: goods.SalesVolume + detailSaleRecord.GoodsNum,
			ResidueNum:  goods.ResidueNum - detailSaleRecord.GoodsNum,
		}
		goodDao.UpdateDataByQO(queryModel, updateModel)
	}
	detailSaleRecordsServiceDao.SaveDetailSaleRecords(detailSaleRecords)
	saleRecordDao.SaveRecords(saleRecords)
	if saleRecords.Type == "1" {
		s := fmt.Sprintf("%1f", saleRecords.SellTotalMoney*0.05)
		member := memberDao.SelectOneByQo(model.Member{Phone: saleRecords.MemberPhone})
		memberDao.UpdateMemberByQO(model.Member{Integral: member.Integral + utils.ConvertStringToInt64(s)}, model.Member{Phone: saleRecords.MemberPhone})
	}
	response.Success(c, saleRecords, "操作成功")
}

func (saleRecordsService SaleRecordsServiceImpl) PaySaleItems(record model.SaleRecord, token string) {
	b, err := json.Marshal(record)
	if err != nil {
		log.Panic(err)
	}
	value := fmt.Sprintf("%s", b)
	rdb := common.GetRDB()
	ctx := context.Background()
	//存入redis缓存中
	keyStr := utils.ConvertInt64ToString(record.CN)
	rdb.Set(ctx, keyStr, value, 1440*time.Minute).Err()

	client := common.GetPayClient()
	pay := alipay.TradePagePay{}
	// 支付成功之后，支付宝将会重定向到该 URL
	pay.ReturnURL = "http://localhost:9291/sale_management/sale_record/saveSaleRecords?cn=" + keyStr + "&token=" + token
	//支付标题
	pay.Subject = "商品销售"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = keyStr
	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//金额
	pay.TotalAmount = fmt.Sprintf("%0.2f", record.SellTotalMoney)
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
