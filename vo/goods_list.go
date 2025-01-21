package vo

import (
	"encoding/json"
	"time"
)

type GoodsListVo struct {
	ID              uint64    `json:"id"`                                                           // 商品编码
	CoverUrl        string    `json:"coverUrl"`                                                     // 商品封面
	Name            string    `json:"name"`                                                         // 商品名称
	SellPrice       float64   `json:"sellPrice"`                                                    // 销售价格
	PurchasePrice   float64   `gorm:"type:double(10,2);column:purchash_price" json:"purchashPrice"` // 进货价格
	ResidueNum      int64     `json:"residueNum"`                                                   // 商品数量
	ResidueStoreNum uint64    `json:"residueStoreNum"`                                              // 可用库存
	CategoryID      uint64    `json:"categoryId"`                                                   // 商品类型
	CategoryName    string    `json:"categoryName"`                                                 // 商品分类名称
	State           string    `json:"state"`                                                        // 状态，下架、上架
	UpdateBy        string    `json:"updateby"`                                                     // 操作者
	Info            string    `json:"info"`                                                         // 备注信息
	UpdateTime      time.Time `json:"updateTime"`                                                   // 操作时间
	SalesVolume     int64     `json:"salesVolume"`
}

// UnmarshalJSON 自定义反序列化方法，用于处理时间格式
func (g *GoodsListVo) UnmarshalJSON(data []byte) error {
	type Alias GoodsListVo
	aux := &struct {
		UpdateTime json.RawMessage `json:"updateTime"`
		*Alias
	}{
		Alias: (*Alias)(g),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if string(aux.UpdateTime) != "" {
		layout := "2006-01-02" // 根据@JsonFormat指定的格式来解析时间
		if err := json.Unmarshal(aux.UpdateTime, &g.UpdateTime); err != nil {
			return err
		}
		g.UpdateTime, _ = time.Parse(layout, string(aux.UpdateTime))
	}

	return nil
}
