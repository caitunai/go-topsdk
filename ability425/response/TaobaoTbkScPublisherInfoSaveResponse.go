package response

import (
	"github.com/caitunai/go-topsdk/ability425/domain"
)

type TaobaoTbkScPublisherInfoSaveResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   data
	*/
	Data domain.TaobaoTbkScPublisherInfoSaveData `json:"data,omitempty" `
}
