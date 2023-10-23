package domain

type TaobaoTbkDgMaterialRecommendPublishInfo struct {
	/*
	   商品信息-收入比率(%)；商品佣金比率+补贴比率     */
	IncomeRate *string `json:"income_rate,omitempty" `

	/*
	   链接-宝贝推广链接     */
	ClickUrl *string `json:"click_url,omitempty" `

	/*
	   链接-宝贝+券二合一页面链接     */
	CouponShareUrl *string `json:"coupon_share_url,omitempty" `

	/*
	   额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=预售单单奖励，1=618超级U选单单补     */
	CpaRewardType *string `json:"cpa_reward_type,omitempty" `

	/*
	   额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割     */
	CpaRewardAmount *string `json:"cpa_reward_amount,omitempty" `

	/*
	   定向计划集合     */
	SpCampaignList *[]TaobaoTbkDgMaterialRecommendSpCampaign `json:"sp_campaign_list,omitempty" `

	/*
	   预热活动到手价对应的佣金比率     */
	FutureActivityCommissionRate *string `json:"future_activity_commission_rate,omitempty" `

	/*
	   预热价活动开始时间     */
	FutureActivityTime *string `json:"future_activity_time,omitempty" `

	/*
	   商品佣金信息     */
	IncomeInfo *TaobaoTbkDgMaterialRecommendIncomeInfo `json:"income_info,omitempty" `
}

func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetIncomeRate(v string) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.IncomeRate = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetClickUrl(v string) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.ClickUrl = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetCouponShareUrl(v string) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.CouponShareUrl = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetCpaRewardType(v string) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.CpaRewardType = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetCpaRewardAmount(v string) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.CpaRewardAmount = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetSpCampaignList(v []TaobaoTbkDgMaterialRecommendSpCampaign) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.SpCampaignList = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetFutureActivityCommissionRate(v string) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.FutureActivityCommissionRate = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetFutureActivityTime(v string) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.FutureActivityTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendPublishInfo) SetIncomeInfo(v TaobaoTbkDgMaterialRecommendIncomeInfo) *TaobaoTbkDgMaterialRecommendPublishInfo {
	s.IncomeInfo = &v
	return s
}
