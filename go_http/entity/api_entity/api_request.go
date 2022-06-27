package api_entity

type HttpBaseData struct {
	Pid string `json:"pid"`
}

type ProductData struct {
	HttpBaseData
	ReqData struct {
		ProductName string `json:"product_name"` //产品名称
		ProductType string `json:"product_type"` //产品类型
		ProductNum  string `json:"product_num"`  //产品批次
	} `json:"req_data"`
}
