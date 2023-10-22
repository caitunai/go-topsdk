package domain

type TaobaoTbkItemInfoUpgradeGetPublishInfo struct {
	/*
	   商品信息-收入比率(商品佣金比率+补贴比率)。1550表示15.5%     */
	CommissionRate *string `json:"commission_rate,omitempty" `

	/*
	   前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
	TopnInfo *TaobaoTbkItemInfoUpgradeGetTopNInfoDTO `json:"topn_info,omitempty" `
}

func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetCommissionRate(v string) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.CommissionRate = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetTopnInfo(v TaobaoTbkItemInfoUpgradeGetTopNInfoDTO) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.TopnInfo = &v
	return s
}
