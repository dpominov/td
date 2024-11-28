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

// ChatInviteExported represents TL type `chatInviteExported#a22cbd96`.
// Exported chat invite
//
// See https://core.telegram.org/constructor/chatInviteExported for reference.
type ChatInviteExported struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this chat invite was revoked
	Revoked bool
	// Whether this chat invite has no expiration
	Permanent bool
	// Whether users importing this invite link will have to be approved to join the channel
	// or group
	RequestNeeded bool
	// Chat invitation link
	Link string
	// ID of the admin that created this chat invite
	AdminID int64
	// When was this chat invite created
	Date int
	// When was this chat invite last modified
	//
	// Use SetStartDate and GetStartDate helpers.
	StartDate int
	// When does this chat invite expire
	//
	// Use SetExpireDate and GetExpireDate helpers.
	ExpireDate int
	// Maximum number of users that can join using this link
	//
	// Use SetUsageLimit and GetUsageLimit helpers.
	UsageLimit int
	// How many users joined using this link
	//
	// Use SetUsage and GetUsage helpers.
	Usage int
	// Number of users that have already used this link to join
	//
	// Use SetRequested and GetRequested helpers.
	Requested int
	// For Telegram Star subscriptions »¹, contains the number of chat members which have
	// already joined the chat using the link, but have already left due to expiration of
	// their subscription.
	//
	// Links:
	//  1) https://core.telegram.org/api/stars#star-subscriptions
	//
	// Use SetSubscriptionExpired and GetSubscriptionExpired helpers.
	SubscriptionExpired int
	// Custom description for the invite link, visible only to admins
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// For Telegram Star subscriptions »¹, contains the pricing of the subscription the
	// user must activate to join the private channel.
	//
	// Links:
	//  1) https://core.telegram.org/api/stars#star-subscriptions
	//
	// Use SetSubscriptionPricing and GetSubscriptionPricing helpers.
	SubscriptionPricing StarsSubscriptionPricing
}

// ChatInviteExportedTypeID is TL type id of ChatInviteExported.
const ChatInviteExportedTypeID = 0xa22cbd96

// construct implements constructor of ExportedChatInviteClass.
func (c ChatInviteExported) construct() ExportedChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInviteExported.
var (
	_ bin.Encoder     = &ChatInviteExported{}
	_ bin.Decoder     = &ChatInviteExported{}
	_ bin.BareEncoder = &ChatInviteExported{}
	_ bin.BareDecoder = &ChatInviteExported{}

	_ ExportedChatInviteClass = &ChatInviteExported{}
)

func (c *ChatInviteExported) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Revoked == false) {
		return false
	}
	if !(c.Permanent == false) {
		return false
	}
	if !(c.RequestNeeded == false) {
		return false
	}
	if !(c.Link == "") {
		return false
	}
	if !(c.AdminID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.StartDate == 0) {
		return false
	}
	if !(c.ExpireDate == 0) {
		return false
	}
	if !(c.UsageLimit == 0) {
		return false
	}
	if !(c.Usage == 0) {
		return false
	}
	if !(c.Requested == 0) {
		return false
	}
	if !(c.SubscriptionExpired == 0) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.SubscriptionPricing.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteExported) String() string {
	if c == nil {
		return "ChatInviteExported(nil)"
	}
	type Alias ChatInviteExported
	return fmt.Sprintf("ChatInviteExported%+v", Alias(*c))
}

// FillFrom fills ChatInviteExported from given interface.
func (c *ChatInviteExported) FillFrom(from interface {
	GetRevoked() (value bool)
	GetPermanent() (value bool)
	GetRequestNeeded() (value bool)
	GetLink() (value string)
	GetAdminID() (value int64)
	GetDate() (value int)
	GetStartDate() (value int, ok bool)
	GetExpireDate() (value int, ok bool)
	GetUsageLimit() (value int, ok bool)
	GetUsage() (value int, ok bool)
	GetRequested() (value int, ok bool)
	GetSubscriptionExpired() (value int, ok bool)
	GetTitle() (value string, ok bool)
	GetSubscriptionPricing() (value StarsSubscriptionPricing, ok bool)
}) {
	c.Revoked = from.GetRevoked()
	c.Permanent = from.GetPermanent()
	c.RequestNeeded = from.GetRequestNeeded()
	c.Link = from.GetLink()
	c.AdminID = from.GetAdminID()
	c.Date = from.GetDate()
	if val, ok := from.GetStartDate(); ok {
		c.StartDate = val
	}

	if val, ok := from.GetExpireDate(); ok {
		c.ExpireDate = val
	}

	if val, ok := from.GetUsageLimit(); ok {
		c.UsageLimit = val
	}

	if val, ok := from.GetUsage(); ok {
		c.Usage = val
	}

	if val, ok := from.GetRequested(); ok {
		c.Requested = val
	}

	if val, ok := from.GetSubscriptionExpired(); ok {
		c.SubscriptionExpired = val
	}

	if val, ok := from.GetTitle(); ok {
		c.Title = val
	}

	if val, ok := from.GetSubscriptionPricing(); ok {
		c.SubscriptionPricing = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteExported) TypeID() uint32 {
	return ChatInviteExportedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteExported) TypeName() string {
	return "chatInviteExported"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteExported) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteExported",
		ID:   ChatInviteExportedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Revoked",
			SchemaName: "revoked",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Permanent",
			SchemaName: "permanent",
			Null:       !c.Flags.Has(5),
		},
		{
			Name:       "RequestNeeded",
			SchemaName: "request_needed",
			Null:       !c.Flags.Has(6),
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "AdminID",
			SchemaName: "admin_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "StartDate",
			SchemaName: "start_date",
			Null:       !c.Flags.Has(4),
		},
		{
			Name:       "ExpireDate",
			SchemaName: "expire_date",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "UsageLimit",
			SchemaName: "usage_limit",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Usage",
			SchemaName: "usage",
			Null:       !c.Flags.Has(3),
		},
		{
			Name:       "Requested",
			SchemaName: "requested",
			Null:       !c.Flags.Has(7),
		},
		{
			Name:       "SubscriptionExpired",
			SchemaName: "subscription_expired",
			Null:       !c.Flags.Has(10),
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !c.Flags.Has(8),
		},
		{
			Name:       "SubscriptionPricing",
			SchemaName: "subscription_pricing",
			Null:       !c.Flags.Has(9),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChatInviteExported) SetFlags() {
	if !(c.Revoked == false) {
		c.Flags.Set(0)
	}
	if !(c.Permanent == false) {
		c.Flags.Set(5)
	}
	if !(c.RequestNeeded == false) {
		c.Flags.Set(6)
	}
	if !(c.StartDate == 0) {
		c.Flags.Set(4)
	}
	if !(c.ExpireDate == 0) {
		c.Flags.Set(1)
	}
	if !(c.UsageLimit == 0) {
		c.Flags.Set(2)
	}
	if !(c.Usage == 0) {
		c.Flags.Set(3)
	}
	if !(c.Requested == 0) {
		c.Flags.Set(7)
	}
	if !(c.SubscriptionExpired == 0) {
		c.Flags.Set(10)
	}
	if !(c.Title == "") {
		c.Flags.Set(8)
	}
	if !(c.SubscriptionPricing.Zero()) {
		c.Flags.Set(9)
	}
}

// Encode implements bin.Encoder.
func (c *ChatInviteExported) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteExported#a22cbd96 as nil")
	}
	b.PutID(ChatInviteExportedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteExported) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteExported#a22cbd96 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteExported#a22cbd96: field flags: %w", err)
	}
	b.PutString(c.Link)
	b.PutLong(c.AdminID)
	b.PutInt(c.Date)
	if c.Flags.Has(4) {
		b.PutInt(c.StartDate)
	}
	if c.Flags.Has(1) {
		b.PutInt(c.ExpireDate)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.UsageLimit)
	}
	if c.Flags.Has(3) {
		b.PutInt(c.Usage)
	}
	if c.Flags.Has(7) {
		b.PutInt(c.Requested)
	}
	if c.Flags.Has(10) {
		b.PutInt(c.SubscriptionExpired)
	}
	if c.Flags.Has(8) {
		b.PutString(c.Title)
	}
	if c.Flags.Has(9) {
		if err := c.SubscriptionPricing.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatInviteExported#a22cbd96: field subscription_pricing: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteExported) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteExported#a22cbd96 to nil")
	}
	if err := b.ConsumeID(ChatInviteExportedTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteExported) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteExported#a22cbd96 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field flags: %w", err)
		}
	}
	c.Revoked = c.Flags.Has(0)
	c.Permanent = c.Flags.Has(5)
	c.RequestNeeded = c.Flags.Has(6)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field link: %w", err)
		}
		c.Link = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field admin_id: %w", err)
		}
		c.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field date: %w", err)
		}
		c.Date = value
	}
	if c.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field start_date: %w", err)
		}
		c.StartDate = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field expire_date: %w", err)
		}
		c.ExpireDate = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field usage_limit: %w", err)
		}
		c.UsageLimit = value
	}
	if c.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field usage: %w", err)
		}
		c.Usage = value
	}
	if c.Flags.Has(7) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field requested: %w", err)
		}
		c.Requested = value
	}
	if c.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field subscription_expired: %w", err)
		}
		c.SubscriptionExpired = value
	}
	if c.Flags.Has(8) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(9) {
		if err := c.SubscriptionPricing.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#a22cbd96: field subscription_pricing: %w", err)
		}
	}
	return nil
}

// SetRevoked sets value of Revoked conditional field.
func (c *ChatInviteExported) SetRevoked(value bool) {
	if value {
		c.Flags.Set(0)
		c.Revoked = true
	} else {
		c.Flags.Unset(0)
		c.Revoked = false
	}
}

// GetRevoked returns value of Revoked conditional field.
func (c *ChatInviteExported) GetRevoked() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// SetPermanent sets value of Permanent conditional field.
func (c *ChatInviteExported) SetPermanent(value bool) {
	if value {
		c.Flags.Set(5)
		c.Permanent = true
	} else {
		c.Flags.Unset(5)
		c.Permanent = false
	}
}

// GetPermanent returns value of Permanent conditional field.
func (c *ChatInviteExported) GetPermanent() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(5)
}

// SetRequestNeeded sets value of RequestNeeded conditional field.
func (c *ChatInviteExported) SetRequestNeeded(value bool) {
	if value {
		c.Flags.Set(6)
		c.RequestNeeded = true
	} else {
		c.Flags.Unset(6)
		c.RequestNeeded = false
	}
}

// GetRequestNeeded returns value of RequestNeeded conditional field.
func (c *ChatInviteExported) GetRequestNeeded() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(6)
}

// GetLink returns value of Link field.
func (c *ChatInviteExported) GetLink() (value string) {
	if c == nil {
		return
	}
	return c.Link
}

// GetAdminID returns value of AdminID field.
func (c *ChatInviteExported) GetAdminID() (value int64) {
	if c == nil {
		return
	}
	return c.AdminID
}

// GetDate returns value of Date field.
func (c *ChatInviteExported) GetDate() (value int) {
	if c == nil {
		return
	}
	return c.Date
}

// SetStartDate sets value of StartDate conditional field.
func (c *ChatInviteExported) SetStartDate(value int) {
	c.Flags.Set(4)
	c.StartDate = value
}

// GetStartDate returns value of StartDate conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetStartDate() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(4) {
		return value, false
	}
	return c.StartDate, true
}

// SetExpireDate sets value of ExpireDate conditional field.
func (c *ChatInviteExported) SetExpireDate(value int) {
	c.Flags.Set(1)
	c.ExpireDate = value
}

// GetExpireDate returns value of ExpireDate conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetExpireDate() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.ExpireDate, true
}

// SetUsageLimit sets value of UsageLimit conditional field.
func (c *ChatInviteExported) SetUsageLimit(value int) {
	c.Flags.Set(2)
	c.UsageLimit = value
}

// GetUsageLimit returns value of UsageLimit conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetUsageLimit() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.UsageLimit, true
}

// SetUsage sets value of Usage conditional field.
func (c *ChatInviteExported) SetUsage(value int) {
	c.Flags.Set(3)
	c.Usage = value
}

// GetUsage returns value of Usage conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetUsage() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(3) {
		return value, false
	}
	return c.Usage, true
}

// SetRequested sets value of Requested conditional field.
func (c *ChatInviteExported) SetRequested(value int) {
	c.Flags.Set(7)
	c.Requested = value
}

// GetRequested returns value of Requested conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetRequested() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(7) {
		return value, false
	}
	return c.Requested, true
}

// SetSubscriptionExpired sets value of SubscriptionExpired conditional field.
func (c *ChatInviteExported) SetSubscriptionExpired(value int) {
	c.Flags.Set(10)
	c.SubscriptionExpired = value
}

// GetSubscriptionExpired returns value of SubscriptionExpired conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetSubscriptionExpired() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(10) {
		return value, false
	}
	return c.SubscriptionExpired, true
}

// SetTitle sets value of Title conditional field.
func (c *ChatInviteExported) SetTitle(value string) {
	c.Flags.Set(8)
	c.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetTitle() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(8) {
		return value, false
	}
	return c.Title, true
}

// SetSubscriptionPricing sets value of SubscriptionPricing conditional field.
func (c *ChatInviteExported) SetSubscriptionPricing(value StarsSubscriptionPricing) {
	c.Flags.Set(9)
	c.SubscriptionPricing = value
}

// GetSubscriptionPricing returns value of SubscriptionPricing conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetSubscriptionPricing() (value StarsSubscriptionPricing, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(9) {
		return value, false
	}
	return c.SubscriptionPricing, true
}

// ChatInvitePublicJoinRequests represents TL type `chatInvitePublicJoinRequests#ed107ab7`.
// Used in updates and in the channel log to indicate when a user is requesting to join
// or has joined a discussion group¹
//
// Links:
//  1. https://core.telegram.org/api/discussion#requiring-users-to-join-the-group
//
// See https://core.telegram.org/constructor/chatInvitePublicJoinRequests for reference.
type ChatInvitePublicJoinRequests struct {
}

// ChatInvitePublicJoinRequestsTypeID is TL type id of ChatInvitePublicJoinRequests.
const ChatInvitePublicJoinRequestsTypeID = 0xed107ab7

// construct implements constructor of ExportedChatInviteClass.
func (c ChatInvitePublicJoinRequests) construct() ExportedChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInvitePublicJoinRequests.
var (
	_ bin.Encoder     = &ChatInvitePublicJoinRequests{}
	_ bin.Decoder     = &ChatInvitePublicJoinRequests{}
	_ bin.BareEncoder = &ChatInvitePublicJoinRequests{}
	_ bin.BareDecoder = &ChatInvitePublicJoinRequests{}

	_ ExportedChatInviteClass = &ChatInvitePublicJoinRequests{}
)

func (c *ChatInvitePublicJoinRequests) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInvitePublicJoinRequests) String() string {
	if c == nil {
		return "ChatInvitePublicJoinRequests(nil)"
	}
	type Alias ChatInvitePublicJoinRequests
	return fmt.Sprintf("ChatInvitePublicJoinRequests%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInvitePublicJoinRequests) TypeID() uint32 {
	return ChatInvitePublicJoinRequestsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInvitePublicJoinRequests) TypeName() string {
	return "chatInvitePublicJoinRequests"
}

// TypeInfo returns info about TL type.
func (c *ChatInvitePublicJoinRequests) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInvitePublicJoinRequests",
		ID:   ChatInvitePublicJoinRequestsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInvitePublicJoinRequests) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInvitePublicJoinRequests#ed107ab7 as nil")
	}
	b.PutID(ChatInvitePublicJoinRequestsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInvitePublicJoinRequests) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInvitePublicJoinRequests#ed107ab7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInvitePublicJoinRequests) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInvitePublicJoinRequests#ed107ab7 to nil")
	}
	if err := b.ConsumeID(ChatInvitePublicJoinRequestsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInvitePublicJoinRequests#ed107ab7: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInvitePublicJoinRequests) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInvitePublicJoinRequests#ed107ab7 to nil")
	}
	return nil
}

// ExportedChatInviteClassName is schema name of ExportedChatInviteClass.
const ExportedChatInviteClassName = "ExportedChatInvite"

// ExportedChatInviteClass represents ExportedChatInvite generic type.
//
// See https://core.telegram.org/type/ExportedChatInvite for reference.
//
// Example:
//
//	g, err := tg.DecodeExportedChatInvite(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ChatInviteExported: // chatInviteExported#a22cbd96
//	case *tg.ChatInvitePublicJoinRequests: // chatInvitePublicJoinRequests#ed107ab7
//	default: panic(v)
//	}
type ExportedChatInviteClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ExportedChatInviteClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeExportedChatInvite implements binary de-serialization for ExportedChatInviteClass.
func DecodeExportedChatInvite(buf *bin.Buffer) (ExportedChatInviteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatInviteExportedTypeID:
		// Decoding chatInviteExported#a22cbd96.
		v := ChatInviteExported{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ExportedChatInviteClass: %w", err)
		}
		return &v, nil
	case ChatInvitePublicJoinRequestsTypeID:
		// Decoding chatInvitePublicJoinRequests#ed107ab7.
		v := ChatInvitePublicJoinRequests{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ExportedChatInviteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ExportedChatInviteClass: %w", bin.NewUnexpectedID(id))
	}
}

// ExportedChatInvite boxes the ExportedChatInviteClass providing a helper.
type ExportedChatInviteBox struct {
	ExportedChatInvite ExportedChatInviteClass
}

// Decode implements bin.Decoder for ExportedChatInviteBox.
func (b *ExportedChatInviteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ExportedChatInviteBox to nil")
	}
	v, err := DecodeExportedChatInvite(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ExportedChatInvite = v
	return nil
}

// Encode implements bin.Encode for ExportedChatInviteBox.
func (b *ExportedChatInviteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ExportedChatInvite == nil {
		return fmt.Errorf("unable to encode ExportedChatInviteClass as nil")
	}
	return b.ExportedChatInvite.Encode(buf)
}
