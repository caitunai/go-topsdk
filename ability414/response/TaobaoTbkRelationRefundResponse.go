package response

import (
	"github.com/caitunai/go-topsdk/ability414/domain"
)

type TaobaoTbkRelationRefundResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果封装对象
	*/
	Result domain.TaobaoTbkRelationRefundRpcResult `json:"result,omitempty" `
}
