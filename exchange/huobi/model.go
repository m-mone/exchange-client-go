package huobi

type PlaceRequestParams struct {
	AccountId string `json:"account_id"` // 账户ID
	Amount    string `json:"amount"`     // 数量
	Price     string `json:"price"`      // 下单价格
	Source    string `json:"source"`     // 订单来源
	Symbol    string `json:"symbol"`     // 交易对
	Type      string `json:"type"`       // 订单类型
}
type PlaceReturn struct {
	Status  string `json:"status"`
	Data    string `json:"data"`
	ErrCode string `json:"err-code"`
	ErrMsg  string `json:"err-msg"`
}
