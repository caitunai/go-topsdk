package ability417

import (
	"errors"
	"github.com/caitunai/go-topsdk"
	"github.com/caitunai/go-topsdk/ability417/request"
	"github.com/caitunai/go-topsdk/ability417/response"
	"github.com/caitunai/go-topsdk/util"
	"log"
)

type Ability417 struct {
	Client *topsdk.TopClient
}

func NewAbility417(client *topsdk.TopClient) *Ability417 {
	return &Ability417{client}
}

/*
淘宝客-推广者-处罚订单查询
*/
func (ability *Ability417) TaobaoTbkDgPunishOrderGet(req *request.TaobaoTbkDgPunishOrderGetRequest) (*response.TaobaoTbkDgPunishOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability417 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.punish.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgPunishOrderGetResponse{}
	if err != nil {
		log.Println("taobaoTbkDgPunishOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
