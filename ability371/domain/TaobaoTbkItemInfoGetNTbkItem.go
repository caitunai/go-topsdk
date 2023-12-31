package domain

type TaobaoTbkItemInfoGetNTbkItem struct {
	/*
	   一级类目名称     */
	CatName *string `json:"cat_name,omitempty" `

	/*
	   商品ID     */
	NumIid *string `json:"num_iid,omitempty" `

	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   商品主图     */
	PictUrl *string `json:"pict_url,omitempty" `

	/*
	   商品小图列表     */
	SmallImages *[]string `json:"small_images,omitempty" `

	/*
	   商品一口价格     */
	ReservePrice *string `json:"reserve_price,omitempty" `

	/*
	   折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价     */
	ZkFinalPrice *string `json:"zk_final_price,omitempty" `

	/*
	   卖家类型，0表示集市，1表示商城，3表示特价版     */
	UserType *int64 `json:"user_type,omitempty" `

	/*
	   商品所在地     */
	Provcity *string `json:"provcity,omitempty" `

	/*
	   商品链接     */
	ItemUrl *string `json:"item_url,omitempty" `

	/*
	   卖家id     */
	SellerId *int64 `json:"seller_id,omitempty" `

	/*
	   30天销量     */
	Volume *int64 `json:"volume,omitempty" `

	/*
	   店铺名称     */
	Nick *string `json:"nick,omitempty" `

	/*
	   叶子类目名称     */
	CatLeafName *string `json:"cat_leaf_name,omitempty" `

	/*
	   是否加入消费者保障     */
	IsPrepay *bool `json:"is_prepay,omitempty" `

	/*
	   店铺dsr 评分     */
	ShopDsr *int64 `json:"shop_dsr,omitempty" `

	/*
	   卖家等级     */
	Ratesum *int64 `json:"ratesum,omitempty" `

	/*
	   退款率是否低于行业均值     */
	IRfdRate *bool `json:"i_rfd_rate,omitempty" `

	/*
	   好评率是否高于行业均值     */
	HGoodRate *bool `json:"h_good_rate,omitempty" `

	/*
	   成交转化是否高于行业均值     */
	HPayRate30 *bool `json:"h_pay_rate30,omitempty" `

	/*
	   是否包邮     */
	FreeShipment *bool `json:"free_shipment,omitempty" `

	/*
	   商品库类型，支持多库类型输出，以英文逗号分隔“,”分隔，1:营销商品主推库，如果值为空则不属于1这种商品类型     */
	MaterialLibType *string `json:"material_lib_type,omitempty" `

	/*
	   预售商品-商品优惠信息     */
	PresaleDiscountFeeText *string `json:"presale_discount_fee_text,omitempty" `

	/*
	   预售商品-付定金结束时间（毫秒）     */
	PresaleTailEndTime *int64 `json:"presale_tail_end_time,omitempty" `

	/*
	   预售商品-付尾款开始时间（毫秒）     */
	PresaleTailStartTime *int64 `json:"presale_tail_start_time,omitempty" `

	/*
	   预售商品-付定金结束时间（毫秒）     */
	PresaleEndTime *int64 `json:"presale_end_time,omitempty" `

	/*
	   预售商品-付定金开始时间（毫秒）     */
	PresaleStartTime *int64 `json:"presale_start_time,omitempty" `

	/*
	   预售商品-定金（元）     */
	PresaleDeposit *string `json:"presale_deposit,omitempty" `

	/*
	   聚划算满减  -结束时间（毫秒）     */
	JuPlayEndTime *int64 `json:"ju_play_end_time,omitempty" `

	/*
	   聚划算满减  -开始时间（毫秒）     */
	JuPlayStartTime *int64 `json:"ju_play_start_time,omitempty" `

	/*
	   1聚划算满减：满N件减X元，满N件X折，满N件X元）  2天猫限时抢：前N分钟每件X元，前N分钟满N件每件X元，前N件每件X元）     */
	PlayInfo *string `json:"play_info,omitempty" `

	/*
	   天猫限时抢可售  -结束时间（毫秒）     */
	TmallPlayActivityEndTime *int64 `json:"tmall_play_activity_end_time,omitempty" `

	/*
	   天猫限时抢可售  -开始时间（毫秒）     */
	TmallPlayActivityStartTime *int64 `json:"tmall_play_activity_start_time,omitempty" `

	/*
	   聚划算信息-聚淘开始时间（毫秒）     */
	JuOnlineStartTime *string `json:"ju_online_start_time,omitempty" `

	/*
	   聚划算信息-聚淘结束时间（毫秒）     */
	JuOnlineEndTime *string `json:"ju_online_end_time,omitempty" `

	/*
	   聚划算信息-商品预热开始时间（毫秒）     */
	JuPreShowStartTime *string `json:"ju_pre_show_start_time,omitempty" `

	/*
	   聚划算信息-商品预热结束时间（毫秒）     */
	JuPreShowEndTime *string `json:"ju_pre_show_end_time,omitempty" `

	/*
	   活动价     */
	SalePrice *string `json:"sale_price,omitempty" `

	/*
	   跨店满减信息     */
	KuadianPromotionInfo *string `json:"kuadian_promotion_info,omitempty" `

	/*
	   是否品牌精选，0不是，1是     */
	SuperiorBrand *string `json:"superior_brand,omitempty" `

	/*
	   是否是热门商品，0不是，1是     */
	HotFlag *string `json:"hot_flag,omitempty" `

	/*
	   入参的（新）商品ID     */
	InputNumIid *string `json:"input_num_iid,omitempty" `
}

func (s *TaobaoTbkItemInfoGetNTbkItem) SetCatName(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.CatName = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetNumIid(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.NumIid = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetTitle(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.Title = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetPictUrl(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.PictUrl = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetSmallImages(v []string) *TaobaoTbkItemInfoGetNTbkItem {
	s.SmallImages = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetReservePrice(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.ReservePrice = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetZkFinalPrice(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.ZkFinalPrice = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetUserType(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.UserType = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetProvcity(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.Provcity = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetItemUrl(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.ItemUrl = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetSellerId(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.SellerId = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetVolume(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.Volume = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetNick(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.Nick = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetCatLeafName(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.CatLeafName = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetIsPrepay(v bool) *TaobaoTbkItemInfoGetNTbkItem {
	s.IsPrepay = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetShopDsr(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.ShopDsr = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetRatesum(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.Ratesum = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetIRfdRate(v bool) *TaobaoTbkItemInfoGetNTbkItem {
	s.IRfdRate = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetHGoodRate(v bool) *TaobaoTbkItemInfoGetNTbkItem {
	s.HGoodRate = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetHPayRate30(v bool) *TaobaoTbkItemInfoGetNTbkItem {
	s.HPayRate30 = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetFreeShipment(v bool) *TaobaoTbkItemInfoGetNTbkItem {
	s.FreeShipment = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetMaterialLibType(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.MaterialLibType = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetPresaleDiscountFeeText(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.PresaleDiscountFeeText = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetPresaleTailEndTime(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.PresaleTailEndTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetPresaleTailStartTime(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.PresaleTailStartTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetPresaleEndTime(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.PresaleEndTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetPresaleStartTime(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.PresaleStartTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetPresaleDeposit(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.PresaleDeposit = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetJuPlayEndTime(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.JuPlayEndTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetJuPlayStartTime(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.JuPlayStartTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetPlayInfo(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.PlayInfo = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetTmallPlayActivityEndTime(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.TmallPlayActivityEndTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetTmallPlayActivityStartTime(v int64) *TaobaoTbkItemInfoGetNTbkItem {
	s.TmallPlayActivityStartTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetJuOnlineStartTime(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.JuOnlineStartTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetJuOnlineEndTime(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.JuOnlineEndTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetJuPreShowStartTime(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.JuPreShowStartTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetJuPreShowEndTime(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.JuPreShowEndTime = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetSalePrice(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.SalePrice = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetKuadianPromotionInfo(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.KuadianPromotionInfo = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetSuperiorBrand(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.SuperiorBrand = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetHotFlag(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.HotFlag = &v
	return s
}
func (s *TaobaoTbkItemInfoGetNTbkItem) SetInputNumIid(v string) *TaobaoTbkItemInfoGetNTbkItem {
	s.InputNumIid = &v
	return s
}
