package models

// JsonResult 通用的 JSON 响应结构
type JsonResult struct {
	// 错误码，0表示成功，非0表示失败
	Code int `json:"code" example:"0"`
	
	// 错误信息
	Message string `json:"message" example:"success"`
	
	// 返回数据
	Data interface{} `json:"data"`
} 