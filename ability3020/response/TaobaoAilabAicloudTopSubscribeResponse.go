package response

type TaobaoAilabAicloudTopSubscribeResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   0-成功
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   订阅结果
	*/
	ResultData bool `json:"result_data,omitempty" `
	/*
	   错误信息
	*/
	ErrorMsg string `json:"error_msg,omitempty" `
	/*
	   traceId
	*/
	TraceId string `json:"trace_id,omitempty" `
}
