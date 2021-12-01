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

// PhoneCallProtocol represents TL type `phoneCallProtocol#fc878fc8`.
// Protocol info for libtgvoip
//
// See https://core.telegram.org/constructor/phoneCallProtocol for reference.
type PhoneCallProtocol struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to allow P2P connection to the other participant
	UDPP2P bool
	// Whether to allow connection to the other participants through the reflector servers
	UDPReflector bool
	// Minimum layer for remote libtgvoip
	MinLayer int
	// Maximum layer for remote libtgvoip
	MaxLayer int
	// When using phone.requestCall¹ and phone.acceptCall², specify all library versions
	// supported by the client. The server will merge and choose the best library version
	// supported by both peers, returning only the best value in the result of the callee's
	// phone.acceptCall³ and in the phoneCallAccepted⁴ update received by the caller.
	//
	// Links:
	//  1) https://core.telegram.org/method/phone.requestCall
	//  2) https://core.telegram.org/method/phone.acceptCall
	//  3) https://core.telegram.org/method/phone.acceptCall
	//  4) https://core.telegram.org/constructor/phoneCallAccepted
	LibraryVersions []string
}

// PhoneCallProtocolTypeID is TL type id of PhoneCallProtocol.
const PhoneCallProtocolTypeID = 0xfc878fc8

// Ensuring interfaces in compile-time for PhoneCallProtocol.
var (
	_ bin.Encoder     = &PhoneCallProtocol{}
	_ bin.Decoder     = &PhoneCallProtocol{}
	_ bin.BareEncoder = &PhoneCallProtocol{}
	_ bin.BareDecoder = &PhoneCallProtocol{}
)

func (p *PhoneCallProtocol) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.UDPP2P == false) {
		return false
	}
	if !(p.UDPReflector == false) {
		return false
	}
	if !(p.MinLayer == 0) {
		return false
	}
	if !(p.MaxLayer == 0) {
		return false
	}
	if !(p.LibraryVersions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneCallProtocol) String() string {
	if p == nil {
		return "PhoneCallProtocol(nil)"
	}
	type Alias PhoneCallProtocol
	return fmt.Sprintf("PhoneCallProtocol%+v", Alias(*p))
}

// FillFrom fills PhoneCallProtocol from given interface.
func (p *PhoneCallProtocol) FillFrom(from interface {
	GetUDPP2P() (value bool)
	GetUDPReflector() (value bool)
	GetMinLayer() (value int)
	GetMaxLayer() (value int)
	GetLibraryVersions() (value []string)
}) {
	p.UDPP2P = from.GetUDPP2P()
	p.UDPReflector = from.GetUDPReflector()
	p.MinLayer = from.GetMinLayer()
	p.MaxLayer = from.GetMaxLayer()
	p.LibraryVersions = from.GetLibraryVersions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneCallProtocol) TypeID() uint32 {
	return PhoneCallProtocolTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneCallProtocol) TypeName() string {
	return "phoneCallProtocol"
}

// TypeInfo returns info about TL type.
func (p *PhoneCallProtocol) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phoneCallProtocol",
		ID:   PhoneCallProtocolTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UDPP2P",
			SchemaName: "udp_p2p",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "UDPReflector",
			SchemaName: "udp_reflector",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "MinLayer",
			SchemaName: "min_layer",
		},
		{
			Name:       "MaxLayer",
			SchemaName: "max_layer",
		},
		{
			Name:       "LibraryVersions",
			SchemaName: "library_versions",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *PhoneCallProtocol) SetFlags() {
	if !(p.UDPP2P == false) {
		p.Flags.Set(0)
	}
	if !(p.UDPReflector == false) {
		p.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (p *PhoneCallProtocol) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallProtocol#fc878fc8 as nil")
	}
	b.PutID(PhoneCallProtocolTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhoneCallProtocol) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallProtocol#fc878fc8 as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallProtocol#fc878fc8: field flags: %w", err)
	}
	b.PutInt(p.MinLayer)
	b.PutInt(p.MaxLayer)
	b.PutVectorHeader(len(p.LibraryVersions))
	for _, v := range p.LibraryVersions {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneCallProtocol) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallProtocol#fc878fc8 to nil")
	}
	if err := b.ConsumeID(PhoneCallProtocolTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhoneCallProtocol) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallProtocol#fc878fc8 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field flags: %w", err)
		}
	}
	p.UDPP2P = p.Flags.Has(0)
	p.UDPReflector = p.Flags.Has(1)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field min_layer: %w", err)
		}
		p.MinLayer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field max_layer: %w", err)
		}
		p.MaxLayer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field library_versions: %w", err)
		}

		if headerLen > 0 {
			p.LibraryVersions = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field library_versions: %w", err)
			}
			p.LibraryVersions = append(p.LibraryVersions, value)
		}
	}
	return nil
}

// SetUDPP2P sets value of UDPP2P conditional field.
func (p *PhoneCallProtocol) SetUDPP2P(value bool) {
	if value {
		p.Flags.Set(0)
		p.UDPP2P = true
	} else {
		p.Flags.Unset(0)
		p.UDPP2P = false
	}
}

// GetUDPP2P returns value of UDPP2P conditional field.
func (p *PhoneCallProtocol) GetUDPP2P() (value bool) {
	return p.Flags.Has(0)
}

// SetUDPReflector sets value of UDPReflector conditional field.
func (p *PhoneCallProtocol) SetUDPReflector(value bool) {
	if value {
		p.Flags.Set(1)
		p.UDPReflector = true
	} else {
		p.Flags.Unset(1)
		p.UDPReflector = false
	}
}

// GetUDPReflector returns value of UDPReflector conditional field.
func (p *PhoneCallProtocol) GetUDPReflector() (value bool) {
	return p.Flags.Has(1)
}

// GetMinLayer returns value of MinLayer field.
func (p *PhoneCallProtocol) GetMinLayer() (value int) {
	return p.MinLayer
}

// GetMaxLayer returns value of MaxLayer field.
func (p *PhoneCallProtocol) GetMaxLayer() (value int) {
	return p.MaxLayer
}

// GetLibraryVersions returns value of LibraryVersions field.
func (p *PhoneCallProtocol) GetLibraryVersions() (value []string) {
	return p.LibraryVersions
}
