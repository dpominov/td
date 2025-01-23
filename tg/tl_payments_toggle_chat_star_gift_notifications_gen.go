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

// PaymentsToggleChatStarGiftNotificationsRequest represents TL type `payments.toggleChatStarGiftNotifications#60eaefa1`.
//
// See https://core.telegram.org/method/payments.toggleChatStarGiftNotifications for reference.
type PaymentsToggleChatStarGiftNotificationsRequest struct {
	// Flags field of PaymentsToggleChatStarGiftNotificationsRequest.
	Flags bin.Fields
	// Enabled field of PaymentsToggleChatStarGiftNotificationsRequest.
	Enabled bool
	// Peer field of PaymentsToggleChatStarGiftNotificationsRequest.
	Peer InputPeerClass
}

// PaymentsToggleChatStarGiftNotificationsRequestTypeID is TL type id of PaymentsToggleChatStarGiftNotificationsRequest.
const PaymentsToggleChatStarGiftNotificationsRequestTypeID = 0x60eaefa1

// Ensuring interfaces in compile-time for PaymentsToggleChatStarGiftNotificationsRequest.
var (
	_ bin.Encoder     = &PaymentsToggleChatStarGiftNotificationsRequest{}
	_ bin.Decoder     = &PaymentsToggleChatStarGiftNotificationsRequest{}
	_ bin.BareEncoder = &PaymentsToggleChatStarGiftNotificationsRequest{}
	_ bin.BareDecoder = &PaymentsToggleChatStarGiftNotificationsRequest{}
)

func (t *PaymentsToggleChatStarGiftNotificationsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}
	if !(t.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) String() string {
	if t == nil {
		return "PaymentsToggleChatStarGiftNotificationsRequest(nil)"
	}
	type Alias PaymentsToggleChatStarGiftNotificationsRequest
	return fmt.Sprintf("PaymentsToggleChatStarGiftNotificationsRequest%+v", Alias(*t))
}

// FillFrom fills PaymentsToggleChatStarGiftNotificationsRequest from given interface.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) FillFrom(from interface {
	GetEnabled() (value bool)
	GetPeer() (value InputPeerClass)
}) {
	t.Enabled = from.GetEnabled()
	t.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsToggleChatStarGiftNotificationsRequest) TypeID() uint32 {
	return PaymentsToggleChatStarGiftNotificationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsToggleChatStarGiftNotificationsRequest) TypeName() string {
	return "payments.toggleChatStarGiftNotifications"
}

// TypeInfo returns info about TL type.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.toggleChatStarGiftNotifications",
		ID:   PaymentsToggleChatStarGiftNotificationsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Enabled",
			SchemaName: "enabled",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) SetFlags() {
	if !(t.Enabled == false) {
		t.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode payments.toggleChatStarGiftNotifications#60eaefa1 as nil")
	}
	b.PutID(PaymentsToggleChatStarGiftNotificationsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode payments.toggleChatStarGiftNotifications#60eaefa1 as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.toggleChatStarGiftNotifications#60eaefa1: field flags: %w", err)
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode payments.toggleChatStarGiftNotifications#60eaefa1: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.toggleChatStarGiftNotifications#60eaefa1: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode payments.toggleChatStarGiftNotifications#60eaefa1 to nil")
	}
	if err := b.ConsumeID(PaymentsToggleChatStarGiftNotificationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.toggleChatStarGiftNotifications#60eaefa1: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode payments.toggleChatStarGiftNotifications#60eaefa1 to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.toggleChatStarGiftNotifications#60eaefa1: field flags: %w", err)
		}
	}
	t.Enabled = t.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.toggleChatStarGiftNotifications#60eaefa1: field peer: %w", err)
		}
		t.Peer = value
	}
	return nil
}

// SetEnabled sets value of Enabled conditional field.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) SetEnabled(value bool) {
	if value {
		t.Flags.Set(0)
		t.Enabled = true
	} else {
		t.Flags.Unset(0)
		t.Enabled = false
	}
}

// GetEnabled returns value of Enabled conditional field.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) GetEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (t *PaymentsToggleChatStarGiftNotificationsRequest) GetPeer() (value InputPeerClass) {
	if t == nil {
		return
	}
	return t.Peer
}

// PaymentsToggleChatStarGiftNotifications invokes method payments.toggleChatStarGiftNotifications#60eaefa1 returning error if any.
//
// See https://core.telegram.org/method/payments.toggleChatStarGiftNotifications for reference.
func (c *Client) PaymentsToggleChatStarGiftNotifications(ctx context.Context, request *PaymentsToggleChatStarGiftNotificationsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
