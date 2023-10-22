package request

type TaobaoAilabAicloudTopSubscribeRequest struct {
	/*
	   开放用户id     */
	OpenUserId *string `json:"open_user_id,omitempty" required:"false" `
}

func (s *TaobaoAilabAicloudTopSubscribeRequest) SetOpenUserId(v string) *TaobaoAilabAicloudTopSubscribeRequest {
	s.OpenUserId = &v
	return s
}

func (req *TaobaoAilabAicloudTopSubscribeRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.OpenUserId != nil {
		paramMap["open_user_id"] = *req.OpenUserId
	}
	return paramMap
}

func (req *TaobaoAilabAicloudTopSubscribeRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
