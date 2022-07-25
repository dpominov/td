// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// StatsMegagroupStats represents TL type `stats.megagroupStats#ef7ff916`.
// Supergroup statistics¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/constructor/stats.megagroupStats for reference.
type StatsMegagroupStats struct {
	// Period in consideration
	Period StatsDateRangeDays
	// Member count change for period in consideration
	Members StatsAbsValueAndPrev
	// Message number change for period in consideration
	Messages StatsAbsValueAndPrev
	// Number of users that viewed messages, for range in consideration
	Viewers StatsAbsValueAndPrev
	// Number of users that posted messages, for range in consideration
	Posters StatsAbsValueAndPrev
	// Supergroup growth graph (absolute subscriber count)
	GrowthGraph StatsGraphClass
	// Members growth (relative subscriber count)
	MembersGraph StatsGraphClass
	// New members by source graph
	NewMembersBySourceGraph StatsGraphClass
	// Subscriber language graph (pie chart)
	LanguagesGraph StatsGraphClass
	// Message activity graph (stacked bar graph, message type)
	MessagesGraph StatsGraphClass
	// Group activity graph (deleted, modified messages, blocked users)
	ActionsGraph StatsGraphClass
	// Activity per hour graph (absolute)
	TopHoursGraph StatsGraphClass
	// Activity per day of week graph (absolute)
	WeekdaysGraph StatsGraphClass
	// Info about most active group members
	TopPosters []StatsGroupTopPoster
	// Info about most active group admins
	TopAdmins []StatsGroupTopAdmin
	// Info about most active group inviters
	TopInviters []StatsGroupTopInviter
	// Info about users mentioned in statistics
	Users []UserClass
}

// StatsMegagroupStatsTypeID is TL type id of StatsMegagroupStats.
const StatsMegagroupStatsTypeID = 0xef7ff916

// Ensuring interfaces in compile-time for StatsMegagroupStats.
var (
	_ bin.Encoder     = &StatsMegagroupStats{}
	_ bin.Decoder     = &StatsMegagroupStats{}
	_ bin.BareEncoder = &StatsMegagroupStats{}
	_ bin.BareDecoder = &StatsMegagroupStats{}
)

func (m *StatsMegagroupStats) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Period.Zero()) {
		return false
	}
	if !(m.Members.Zero()) {
		return false
	}
	if !(m.Messages.Zero()) {
		return false
	}
	if !(m.Viewers.Zero()) {
		return false
	}
	if !(m.Posters.Zero()) {
		return false
	}
	if !(m.GrowthGraph == nil) {
		return false
	}
	if !(m.MembersGraph == nil) {
		return false
	}
	if !(m.NewMembersBySourceGraph == nil) {
		return false
	}
	if !(m.LanguagesGraph == nil) {
		return false
	}
	if !(m.MessagesGraph == nil) {
		return false
	}
	if !(m.ActionsGraph == nil) {
		return false
	}
	if !(m.TopHoursGraph == nil) {
		return false
	}
	if !(m.WeekdaysGraph == nil) {
		return false
	}
	if !(m.TopPosters == nil) {
		return false
	}
	if !(m.TopAdmins == nil) {
		return false
	}
	if !(m.TopInviters == nil) {
		return false
	}
	if !(m.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *StatsMegagroupStats) String() string {
	if m == nil {
		return "StatsMegagroupStats(nil)"
	}
	type Alias StatsMegagroupStats
	return fmt.Sprintf("StatsMegagroupStats%+v", Alias(*m))
}

// FillFrom fills StatsMegagroupStats from given interface.
func (m *StatsMegagroupStats) FillFrom(from interface {
	GetPeriod() (value StatsDateRangeDays)
	GetMembers() (value StatsAbsValueAndPrev)
	GetMessages() (value StatsAbsValueAndPrev)
	GetViewers() (value StatsAbsValueAndPrev)
	GetPosters() (value StatsAbsValueAndPrev)
	GetGrowthGraph() (value StatsGraphClass)
	GetMembersGraph() (value StatsGraphClass)
	GetNewMembersBySourceGraph() (value StatsGraphClass)
	GetLanguagesGraph() (value StatsGraphClass)
	GetMessagesGraph() (value StatsGraphClass)
	GetActionsGraph() (value StatsGraphClass)
	GetTopHoursGraph() (value StatsGraphClass)
	GetWeekdaysGraph() (value StatsGraphClass)
	GetTopPosters() (value []StatsGroupTopPoster)
	GetTopAdmins() (value []StatsGroupTopAdmin)
	GetTopInviters() (value []StatsGroupTopInviter)
	GetUsers() (value []UserClass)
}) {
	m.Period = from.GetPeriod()
	m.Members = from.GetMembers()
	m.Messages = from.GetMessages()
	m.Viewers = from.GetViewers()
	m.Posters = from.GetPosters()
	m.GrowthGraph = from.GetGrowthGraph()
	m.MembersGraph = from.GetMembersGraph()
	m.NewMembersBySourceGraph = from.GetNewMembersBySourceGraph()
	m.LanguagesGraph = from.GetLanguagesGraph()
	m.MessagesGraph = from.GetMessagesGraph()
	m.ActionsGraph = from.GetActionsGraph()
	m.TopHoursGraph = from.GetTopHoursGraph()
	m.WeekdaysGraph = from.GetWeekdaysGraph()
	m.TopPosters = from.GetTopPosters()
	m.TopAdmins = from.GetTopAdmins()
	m.TopInviters = from.GetTopInviters()
	m.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsMegagroupStats) TypeID() uint32 {
	return StatsMegagroupStatsTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsMegagroupStats) TypeName() string {
	return "stats.megagroupStats"
}

// TypeInfo returns info about TL type.
func (m *StatsMegagroupStats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.megagroupStats",
		ID:   StatsMegagroupStatsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Period",
			SchemaName: "period",
		},
		{
			Name:       "Members",
			SchemaName: "members",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Viewers",
			SchemaName: "viewers",
		},
		{
			Name:       "Posters",
			SchemaName: "posters",
		},
		{
			Name:       "GrowthGraph",
			SchemaName: "growth_graph",
		},
		{
			Name:       "MembersGraph",
			SchemaName: "members_graph",
		},
		{
			Name:       "NewMembersBySourceGraph",
			SchemaName: "new_members_by_source_graph",
		},
		{
			Name:       "LanguagesGraph",
			SchemaName: "languages_graph",
		},
		{
			Name:       "MessagesGraph",
			SchemaName: "messages_graph",
		},
		{
			Name:       "ActionsGraph",
			SchemaName: "actions_graph",
		},
		{
			Name:       "TopHoursGraph",
			SchemaName: "top_hours_graph",
		},
		{
			Name:       "WeekdaysGraph",
			SchemaName: "weekdays_graph",
		},
		{
			Name:       "TopPosters",
			SchemaName: "top_posters",
		},
		{
			Name:       "TopAdmins",
			SchemaName: "top_admins",
		},
		{
			Name:       "TopInviters",
			SchemaName: "top_inviters",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *StatsMegagroupStats) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.megagroupStats#ef7ff916 as nil")
	}
	b.PutID(StatsMegagroupStatsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *StatsMegagroupStats) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.megagroupStats#ef7ff916 as nil")
	}
	if err := m.Period.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field period: %w", err)
	}
	if err := m.Members.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members: %w", err)
	}
	if err := m.Messages.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages: %w", err)
	}
	if err := m.Viewers.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field viewers: %w", err)
	}
	if err := m.Posters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field posters: %w", err)
	}
	if m.GrowthGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field growth_graph is nil")
	}
	if err := m.GrowthGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field growth_graph: %w", err)
	}
	if m.MembersGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members_graph is nil")
	}
	if err := m.MembersGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members_graph: %w", err)
	}
	if m.NewMembersBySourceGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field new_members_by_source_graph is nil")
	}
	if err := m.NewMembersBySourceGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field new_members_by_source_graph: %w", err)
	}
	if m.LanguagesGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field languages_graph is nil")
	}
	if err := m.LanguagesGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field languages_graph: %w", err)
	}
	if m.MessagesGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages_graph is nil")
	}
	if err := m.MessagesGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages_graph: %w", err)
	}
	if m.ActionsGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field actions_graph is nil")
	}
	if err := m.ActionsGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field actions_graph: %w", err)
	}
	if m.TopHoursGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_hours_graph is nil")
	}
	if err := m.TopHoursGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_hours_graph: %w", err)
	}
	if m.WeekdaysGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field weekdays_graph is nil")
	}
	if err := m.WeekdaysGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field weekdays_graph: %w", err)
	}
	b.PutVectorHeader(len(m.TopPosters))
	for idx, v := range m.TopPosters {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_posters element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.TopAdmins))
	for idx, v := range m.TopAdmins {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_admins element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.TopInviters))
	for idx, v := range m.TopInviters {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_inviters element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *StatsMegagroupStats) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.megagroupStats#ef7ff916 to nil")
	}
	if err := b.ConsumeID(StatsMegagroupStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *StatsMegagroupStats) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.megagroupStats#ef7ff916 to nil")
	}
	{
		if err := m.Period.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field period: %w", err)
		}
	}
	{
		if err := m.Members.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field members: %w", err)
		}
	}
	{
		if err := m.Messages.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field messages: %w", err)
		}
	}
	{
		if err := m.Viewers.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field viewers: %w", err)
		}
	}
	{
		if err := m.Posters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field posters: %w", err)
		}
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field growth_graph: %w", err)
		}
		m.GrowthGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field members_graph: %w", err)
		}
		m.MembersGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field new_members_by_source_graph: %w", err)
		}
		m.NewMembersBySourceGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field languages_graph: %w", err)
		}
		m.LanguagesGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field messages_graph: %w", err)
		}
		m.MessagesGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field actions_graph: %w", err)
		}
		m.ActionsGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_hours_graph: %w", err)
		}
		m.TopHoursGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field weekdays_graph: %w", err)
		}
		m.WeekdaysGraph = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_posters: %w", err)
		}

		if headerLen > 0 {
			m.TopPosters = make([]StatsGroupTopPoster, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopPoster
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_posters: %w", err)
			}
			m.TopPosters = append(m.TopPosters, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_admins: %w", err)
		}

		if headerLen > 0 {
			m.TopAdmins = make([]StatsGroupTopAdmin, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopAdmin
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_admins: %w", err)
			}
			m.TopAdmins = append(m.TopAdmins, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_inviters: %w", err)
		}

		if headerLen > 0 {
			m.TopInviters = make([]StatsGroupTopInviter, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopInviter
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_inviters: %w", err)
			}
			m.TopInviters = append(m.TopInviters, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field users: %w", err)
		}

		if headerLen > 0 {
			m.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	return nil
}

// GetPeriod returns value of Period field.
func (m *StatsMegagroupStats) GetPeriod() (value StatsDateRangeDays) {
	if m == nil {
		return
	}
	return m.Period
}

// GetMembers returns value of Members field.
func (m *StatsMegagroupStats) GetMembers() (value StatsAbsValueAndPrev) {
	if m == nil {
		return
	}
	return m.Members
}

// GetMessages returns value of Messages field.
func (m *StatsMegagroupStats) GetMessages() (value StatsAbsValueAndPrev) {
	if m == nil {
		return
	}
	return m.Messages
}

// GetViewers returns value of Viewers field.
func (m *StatsMegagroupStats) GetViewers() (value StatsAbsValueAndPrev) {
	if m == nil {
		return
	}
	return m.Viewers
}

// GetPosters returns value of Posters field.
func (m *StatsMegagroupStats) GetPosters() (value StatsAbsValueAndPrev) {
	if m == nil {
		return
	}
	return m.Posters
}

// GetGrowthGraph returns value of GrowthGraph field.
func (m *StatsMegagroupStats) GetGrowthGraph() (value StatsGraphClass) {
	if m == nil {
		return
	}
	return m.GrowthGraph
}

// GetMembersGraph returns value of MembersGraph field.
func (m *StatsMegagroupStats) GetMembersGraph() (value StatsGraphClass) {
	if m == nil {
		return
	}
	return m.MembersGraph
}

// GetNewMembersBySourceGraph returns value of NewMembersBySourceGraph field.
func (m *StatsMegagroupStats) GetNewMembersBySourceGraph() (value StatsGraphClass) {
	if m == nil {
		return
	}
	return m.NewMembersBySourceGraph
}

// GetLanguagesGraph returns value of LanguagesGraph field.
func (m *StatsMegagroupStats) GetLanguagesGraph() (value StatsGraphClass) {
	if m == nil {
		return
	}
	return m.LanguagesGraph
}

// GetMessagesGraph returns value of MessagesGraph field.
func (m *StatsMegagroupStats) GetMessagesGraph() (value StatsGraphClass) {
	if m == nil {
		return
	}
	return m.MessagesGraph
}

// GetActionsGraph returns value of ActionsGraph field.
func (m *StatsMegagroupStats) GetActionsGraph() (value StatsGraphClass) {
	if m == nil {
		return
	}
	return m.ActionsGraph
}

// GetTopHoursGraph returns value of TopHoursGraph field.
func (m *StatsMegagroupStats) GetTopHoursGraph() (value StatsGraphClass) {
	if m == nil {
		return
	}
	return m.TopHoursGraph
}

// GetWeekdaysGraph returns value of WeekdaysGraph field.
func (m *StatsMegagroupStats) GetWeekdaysGraph() (value StatsGraphClass) {
	if m == nil {
		return
	}
	return m.WeekdaysGraph
}

// GetTopPosters returns value of TopPosters field.
func (m *StatsMegagroupStats) GetTopPosters() (value []StatsGroupTopPoster) {
	if m == nil {
		return
	}
	return m.TopPosters
}

// GetTopAdmins returns value of TopAdmins field.
func (m *StatsMegagroupStats) GetTopAdmins() (value []StatsGroupTopAdmin) {
	if m == nil {
		return
	}
	return m.TopAdmins
}

// GetTopInviters returns value of TopInviters field.
func (m *StatsMegagroupStats) GetTopInviters() (value []StatsGroupTopInviter) {
	if m == nil {
		return
	}
	return m.TopInviters
}

// GetUsers returns value of Users field.
func (m *StatsMegagroupStats) GetUsers() (value []UserClass) {
	if m == nil {
		return
	}
	return m.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (m *StatsMegagroupStats) MapUsers() (value UserClassArray) {
	return UserClassArray(m.Users)
}
