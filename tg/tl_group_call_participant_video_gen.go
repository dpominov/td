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

// GroupCallParticipantVideo represents TL type `groupCallParticipantVideo#67753ac8`.
// Info about a video stream
//
// See https://core.telegram.org/constructor/groupCallParticipantVideo for reference.
type GroupCallParticipantVideo struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the stream is currently paused
	Paused bool
	// Endpoint
	Endpoint string
	// Source groups
	SourceGroups []GroupCallParticipantVideoSourceGroup
	// Audio source ID
	//
	// Use SetAudioSource and GetAudioSource helpers.
	AudioSource int
}

// GroupCallParticipantVideoTypeID is TL type id of GroupCallParticipantVideo.
const GroupCallParticipantVideoTypeID = 0x67753ac8

// Ensuring interfaces in compile-time for GroupCallParticipantVideo.
var (
	_ bin.Encoder     = &GroupCallParticipantVideo{}
	_ bin.Decoder     = &GroupCallParticipantVideo{}
	_ bin.BareEncoder = &GroupCallParticipantVideo{}
	_ bin.BareDecoder = &GroupCallParticipantVideo{}
)

func (g *GroupCallParticipantVideo) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Paused == false) {
		return false
	}
	if !(g.Endpoint == "") {
		return false
	}
	if !(g.SourceGroups == nil) {
		return false
	}
	if !(g.AudioSource == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCallParticipantVideo) String() string {
	if g == nil {
		return "GroupCallParticipantVideo(nil)"
	}
	type Alias GroupCallParticipantVideo
	return fmt.Sprintf("GroupCallParticipantVideo%+v", Alias(*g))
}

// FillFrom fills GroupCallParticipantVideo from given interface.
func (g *GroupCallParticipantVideo) FillFrom(from interface {
	GetPaused() (value bool)
	GetEndpoint() (value string)
	GetSourceGroups() (value []GroupCallParticipantVideoSourceGroup)
	GetAudioSource() (value int, ok bool)
}) {
	g.Paused = from.GetPaused()
	g.Endpoint = from.GetEndpoint()
	g.SourceGroups = from.GetSourceGroups()
	if val, ok := from.GetAudioSource(); ok {
		g.AudioSource = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GroupCallParticipantVideo) TypeID() uint32 {
	return GroupCallParticipantVideoTypeID
}

// TypeName returns name of type in TL schema.
func (*GroupCallParticipantVideo) TypeName() string {
	return "groupCallParticipantVideo"
}

// TypeInfo returns info about TL type.
func (g *GroupCallParticipantVideo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "groupCallParticipantVideo",
		ID:   GroupCallParticipantVideoTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Paused",
			SchemaName: "paused",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Endpoint",
			SchemaName: "endpoint",
		},
		{
			Name:       "SourceGroups",
			SchemaName: "source_groups",
		},
		{
			Name:       "AudioSource",
			SchemaName: "audio_source",
			Null:       !g.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *GroupCallParticipantVideo) SetFlags() {
	if !(g.Paused == false) {
		g.Flags.Set(0)
	}
	if !(g.AudioSource == 0) {
		g.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (g *GroupCallParticipantVideo) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallParticipantVideo#67753ac8 as nil")
	}
	b.PutID(GroupCallParticipantVideoTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GroupCallParticipantVideo) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallParticipantVideo#67753ac8 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode groupCallParticipantVideo#67753ac8: field flags: %w", err)
	}
	b.PutString(g.Endpoint)
	b.PutVectorHeader(len(g.SourceGroups))
	for idx, v := range g.SourceGroups {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode groupCallParticipantVideo#67753ac8: field source_groups element with index %d: %w", idx, err)
		}
	}
	if g.Flags.Has(1) {
		b.PutInt(g.AudioSource)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GroupCallParticipantVideo) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallParticipantVideo#67753ac8 to nil")
	}
	if err := b.ConsumeID(GroupCallParticipantVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCallParticipantVideo#67753ac8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GroupCallParticipantVideo) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallParticipantVideo#67753ac8 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode groupCallParticipantVideo#67753ac8: field flags: %w", err)
		}
	}
	g.Paused = g.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipantVideo#67753ac8: field endpoint: %w", err)
		}
		g.Endpoint = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipantVideo#67753ac8: field source_groups: %w", err)
		}

		if headerLen > 0 {
			g.SourceGroups = make([]GroupCallParticipantVideoSourceGroup, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value GroupCallParticipantVideoSourceGroup
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode groupCallParticipantVideo#67753ac8: field source_groups: %w", err)
			}
			g.SourceGroups = append(g.SourceGroups, value)
		}
	}
	if g.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipantVideo#67753ac8: field audio_source: %w", err)
		}
		g.AudioSource = value
	}
	return nil
}

// SetPaused sets value of Paused conditional field.
func (g *GroupCallParticipantVideo) SetPaused(value bool) {
	if value {
		g.Flags.Set(0)
		g.Paused = true
	} else {
		g.Flags.Unset(0)
		g.Paused = false
	}
}

// GetPaused returns value of Paused conditional field.
func (g *GroupCallParticipantVideo) GetPaused() (value bool) {
	return g.Flags.Has(0)
}

// GetEndpoint returns value of Endpoint field.
func (g *GroupCallParticipantVideo) GetEndpoint() (value string) {
	return g.Endpoint
}

// GetSourceGroups returns value of SourceGroups field.
func (g *GroupCallParticipantVideo) GetSourceGroups() (value []GroupCallParticipantVideoSourceGroup) {
	return g.SourceGroups
}

// SetAudioSource sets value of AudioSource conditional field.
func (g *GroupCallParticipantVideo) SetAudioSource(value int) {
	g.Flags.Set(1)
	g.AudioSource = value
}

// GetAudioSource returns value of AudioSource conditional field and
// boolean which is true if field was set.
func (g *GroupCallParticipantVideo) GetAudioSource() (value int, ok bool) {
	if !g.Flags.Has(1) {
		return value, false
	}
	return g.AudioSource, true
}
