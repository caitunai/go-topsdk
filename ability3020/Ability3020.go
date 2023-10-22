package ability3020

import (
	"errors"
	"github.com/caitunai/go-topsdk"
	"github.com/caitunai/go-topsdk/ability3020/request"
	"github.com/caitunai/go-topsdk/ability3020/response"
	"github.com/caitunai/go-topsdk/util"
	"log"
)

type Ability3020 struct {
	Client *topsdk.TopClient
}

func NewAbility3020(client *topsdk.TopClient) *Ability3020 {
	return &Ability3020{client}
}

/*
消息号订阅接口
*/
func (ability *Ability3020) TaobaoAilabAicloudTopSubscribe(req *request.TaobaoAilabAicloudTopSubscribeRequest, session string) (*response.TaobaoAilabAicloudTopSubscribeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability3020 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.ailab.aicloud.top.subscribe", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoAilabAicloudTopSubscribeResponse{}
	if err != nil {
		log.Println("taobaoAilabAicloudTopSubscribe error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
