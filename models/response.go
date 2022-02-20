package models

type Resp struct {
	// 业务编码,200状态下返回Data，其他状态返回Error
	Code uint32 `json:"code"`
	// 相应数据(出现错误的时候一般是空)
	Data interface{} `json:"data"`
	// 错误数据(出现相应数据的时候一般是空)
	Error string `json:"error"`
}
