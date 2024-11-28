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

// MessagesGetPaidReactionPrivacyRequest represents TL type `messages.getPaidReactionPrivacy#472455aa`.
// Fetches an updatePaidReactionPrivacy¹ update with the current default paid reaction
// privacy, see here »² for more info.
//
// Links:
//  1. https://core.telegram.org/constructor/updatePaidReactionPrivacy
//  2. https://core.telegram.org/api/reactions#paid-reactions
//
// See https://core.telegram.org/method/messages.getPaidReactionPrivacy for reference.
type MessagesGetPaidReactionPrivacyRequest struct {
}

// MessagesGetPaidReactionPrivacyRequestTypeID is TL type id of MessagesGetPaidReactionPrivacyRequest.
const MessagesGetPaidReactionPrivacyRequestTypeID = 0x472455aa

// Ensuring interfaces in compile-time for MessagesGetPaidReactionPrivacyRequest.
var (
	_ bin.Encoder     = &MessagesGetPaidReactionPrivacyRequest{}
	_ bin.Decoder     = &MessagesGetPaidReactionPrivacyRequest{}
	_ bin.BareEncoder = &MessagesGetPaidReactionPrivacyRequest{}
	_ bin.BareDecoder = &MessagesGetPaidReactionPrivacyRequest{}
)

func (g *MessagesGetPaidReactionPrivacyRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetPaidReactionPrivacyRequest) String() string {
	if g == nil {
		return "MessagesGetPaidReactionPrivacyRequest(nil)"
	}
	type Alias MessagesGetPaidReactionPrivacyRequest
	return fmt.Sprintf("MessagesGetPaidReactionPrivacyRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetPaidReactionPrivacyRequest) TypeID() uint32 {
	return MessagesGetPaidReactionPrivacyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetPaidReactionPrivacyRequest) TypeName() string {
	return "messages.getPaidReactionPrivacy"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetPaidReactionPrivacyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getPaidReactionPrivacy",
		ID:   MessagesGetPaidReactionPrivacyRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetPaidReactionPrivacyRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPaidReactionPrivacy#472455aa as nil")
	}
	b.PutID(MessagesGetPaidReactionPrivacyRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetPaidReactionPrivacyRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPaidReactionPrivacy#472455aa as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPaidReactionPrivacyRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPaidReactionPrivacy#472455aa to nil")
	}
	if err := b.ConsumeID(MessagesGetPaidReactionPrivacyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPaidReactionPrivacy#472455aa: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetPaidReactionPrivacyRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPaidReactionPrivacy#472455aa to nil")
	}
	return nil
}

// MessagesGetPaidReactionPrivacy invokes method messages.getPaidReactionPrivacy#472455aa returning error if any.
// Fetches an updatePaidReactionPrivacy¹ update with the current default paid reaction
// privacy, see here »² for more info.
//
// Links:
//  1. https://core.telegram.org/constructor/updatePaidReactionPrivacy
//  2. https://core.telegram.org/api/reactions#paid-reactions
//
// See https://core.telegram.org/method/messages.getPaidReactionPrivacy for reference.
// Can be used by bots.
func (c *Client) MessagesGetPaidReactionPrivacy(ctx context.Context) (UpdatesClass, error) {
	var result UpdatesBox

	request := &MessagesGetPaidReactionPrivacyRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
