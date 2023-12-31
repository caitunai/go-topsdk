package domain

import (
	"github.com/caitunai/go-topsdk/util"
)

type TaobaoTbkDgPunishOrderGetTopApiAfOrderOption struct {
	/*
	   渠道关系id     */
	RelationId *int64 `json:"relation_id,omitempty" `

	/*
	   子订单号     */
	TbTradeId *int64 `json:"tb_trade_id,omitempty" `

	/*
	   此参数不再使用，请勿入参     */
	TbTradeParentId *int64 `json:"tb_trade_parent_id,omitempty" `

	/*
	   pagesize     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   pageNo     */
	PageNo *int64 `json:"page_no,omitempty" `

	/*
	   查询时间跨度，不超过30天，单位是天     */
	Span *int64 `json:"span,omitempty" `

	/*
	   查询开始时间，以taoke订单创建时间开始     */
	StartTime *util.LocalTime `json:"start_time,omitempty" `

	/*
	   此参数不再使用，请勿入参     */
	SpecialId *int64 `json:"special_id,omitempty" `

	/*
	   pid中的第三段，adzoneId     */
	AdzoneId *int64 `json:"adzone_id,omitempty" `

	/*
	   pid中的第二段，siteId     */
	SiteId *int64 `json:"site_id,omitempty" `

	/*
	   此参数不再使用，请勿入参     */
	ViolationType *int64 `json:"violation_type,omitempty" `

	/*
	   此参数不再使用，请勿入参     */
	PunishStatus *int64 `json:"punish_status,omitempty" `
}

func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetRelationId(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetTbTradeId(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.TbTradeId = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetTbTradeParentId(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.TbTradeParentId = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetPageSize(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetPageNo(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetSpan(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.Span = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetStartTime(v util.LocalTime) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.StartTime = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetSpecialId(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetAdzoneId(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetSiteId(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.SiteId = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetViolationType(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.ViolationType = &v
	return s
}
func (s *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption) SetPunishStatus(v int64) *TaobaoTbkDgPunishOrderGetTopApiAfOrderOption {
	s.PunishStatus = &v
	return s
}
