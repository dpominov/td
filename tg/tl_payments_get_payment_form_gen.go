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
)

// PaymentsGetPaymentFormRequest represents TL type `payments.getPaymentForm#8a333c8d`.
// Get a payment form
//
// See https://core.telegram.org/method/payments.getPaymentForm for reference.
type PaymentsGetPaymentFormRequest struct {
	// Flags field of PaymentsGetPaymentFormRequest.
	Flags bin.Fields
	// Peer field of PaymentsGetPaymentFormRequest.
	Peer InputPeerClass
	// Message ID of payment form
	MsgID int
	// ThemeParams field of PaymentsGetPaymentFormRequest.
	//
	// Use SetThemeParams and GetThemeParams helpers.
	ThemeParams DataJSON
}

// PaymentsGetPaymentFormRequestTypeID is TL type id of PaymentsGetPaymentFormRequest.
const PaymentsGetPaymentFormRequestTypeID = 0x8a333c8d

func (g *PaymentsGetPaymentFormRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}
	if !(g.ThemeParams.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetPaymentFormRequest) String() string {
	if g == nil {
		return "PaymentsGetPaymentFormRequest(nil)"
	}
	type Alias PaymentsGetPaymentFormRequest
	return fmt.Sprintf("PaymentsGetPaymentFormRequest%+v", Alias(*g))
}

// FillFrom fills PaymentsGetPaymentFormRequest from given interface.
func (g *PaymentsGetPaymentFormRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetThemeParams() (value DataJSON, ok bool)
}) {
	g.Peer = from.GetPeer()
	g.MsgID = from.GetMsgID()
	if val, ok := from.GetThemeParams(); ok {
		g.ThemeParams = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetPaymentFormRequest) TypeID() uint32 {
	return PaymentsGetPaymentFormRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetPaymentFormRequest) TypeName() string {
	return "payments.getPaymentForm"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetPaymentFormRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getPaymentForm",
		ID:   PaymentsGetPaymentFormRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "ThemeParams",
			SchemaName: "theme_params",
			Null:       !g.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetPaymentFormRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getPaymentForm#8a333c8d as nil")
	}
	b.PutID(PaymentsGetPaymentFormRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetPaymentFormRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getPaymentForm#8a333c8d as nil")
	}
	if !(g.ThemeParams.Zero()) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getPaymentForm#8a333c8d: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode payments.getPaymentForm#8a333c8d: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getPaymentForm#8a333c8d: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	if g.Flags.Has(0) {
		if err := g.ThemeParams.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.getPaymentForm#8a333c8d: field theme_params: %w", err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *PaymentsGetPaymentFormRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *PaymentsGetPaymentFormRequest) GetMsgID() (value int) {
	return g.MsgID
}

// SetThemeParams sets value of ThemeParams conditional field.
func (g *PaymentsGetPaymentFormRequest) SetThemeParams(value DataJSON) {
	g.Flags.Set(0)
	g.ThemeParams = value
}

// GetThemeParams returns value of ThemeParams conditional field and
// boolean which is true if field was set.
func (g *PaymentsGetPaymentFormRequest) GetThemeParams() (value DataJSON, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.ThemeParams, true
}

// Decode implements bin.Decoder.
func (g *PaymentsGetPaymentFormRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getPaymentForm#8a333c8d to nil")
	}
	if err := b.ConsumeID(PaymentsGetPaymentFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getPaymentForm#8a333c8d: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetPaymentFormRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getPaymentForm#8a333c8d to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.getPaymentForm#8a333c8d: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.getPaymentForm#8a333c8d: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getPaymentForm#8a333c8d: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	if g.Flags.Has(0) {
		if err := g.ThemeParams.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.getPaymentForm#8a333c8d: field theme_params: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsGetPaymentFormRequest.
var (
	_ bin.Encoder     = &PaymentsGetPaymentFormRequest{}
	_ bin.Decoder     = &PaymentsGetPaymentFormRequest{}
	_ bin.BareEncoder = &PaymentsGetPaymentFormRequest{}
	_ bin.BareDecoder = &PaymentsGetPaymentFormRequest{}
)

// PaymentsGetPaymentForm invokes method payments.getPaymentForm#8a333c8d returning error if any.
// Get a payment form
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//
// See https://core.telegram.org/method/payments.getPaymentForm for reference.
func (c *Client) PaymentsGetPaymentForm(ctx context.Context, request *PaymentsGetPaymentFormRequest) (*PaymentsPaymentForm, error) {
	var result PaymentsPaymentForm

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
