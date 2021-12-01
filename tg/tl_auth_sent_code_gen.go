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

// AuthSentCode represents TL type `auth.sentCode#5e002502`.
// Contains info about a sent verification code.
//
// See https://core.telegram.org/constructor/auth.sentCode for reference.
type AuthSentCode struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Phone code type
	Type AuthSentCodeTypeClass
	// Phone code hash, to be stored and later re-used with auth.signIn¹
	//
	// Links:
	//  1) https://core.telegram.org/method/auth.signIn
	PhoneCodeHash string
	// Phone code type that will be sent next, if the phone code is not received within
	// timeout seconds: to send it use auth.resendCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/auth.resendCode
	//
	// Use SetNextType and GetNextType helpers.
	NextType AuthCodeTypeClass
	// Timeout for reception of the phone code
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int
}

// AuthSentCodeTypeID is TL type id of AuthSentCode.
const AuthSentCodeTypeID = 0x5e002502

// Ensuring interfaces in compile-time for AuthSentCode.
var (
	_ bin.Encoder     = &AuthSentCode{}
	_ bin.Decoder     = &AuthSentCode{}
	_ bin.BareEncoder = &AuthSentCode{}
	_ bin.BareDecoder = &AuthSentCode{}
)

func (s *AuthSentCode) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Type == nil) {
		return false
	}
	if !(s.PhoneCodeHash == "") {
		return false
	}
	if !(s.NextType == nil) {
		return false
	}
	if !(s.Timeout == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCode) String() string {
	if s == nil {
		return "AuthSentCode(nil)"
	}
	type Alias AuthSentCode
	return fmt.Sprintf("AuthSentCode%+v", Alias(*s))
}

// FillFrom fills AuthSentCode from given interface.
func (s *AuthSentCode) FillFrom(from interface {
	GetType() (value AuthSentCodeTypeClass)
	GetPhoneCodeHash() (value string)
	GetNextType() (value AuthCodeTypeClass, ok bool)
	GetTimeout() (value int, ok bool)
}) {
	s.Type = from.GetType()
	s.PhoneCodeHash = from.GetPhoneCodeHash()
	if val, ok := from.GetNextType(); ok {
		s.NextType = val
	}

	if val, ok := from.GetTimeout(); ok {
		s.Timeout = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthSentCode) TypeID() uint32 {
	return AuthSentCodeTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthSentCode) TypeName() string {
	return "auth.sentCode"
}

// TypeInfo returns info about TL type.
func (s *AuthSentCode) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.sentCode",
		ID:   AuthSentCodeTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "PhoneCodeHash",
			SchemaName: "phone_code_hash",
		},
		{
			Name:       "NextType",
			SchemaName: "next_type",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Timeout",
			SchemaName: "timeout",
			Null:       !s.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *AuthSentCode) SetFlags() {
	if !(s.NextType == nil) {
		s.Flags.Set(1)
	}
	if !(s.Timeout == 0) {
		s.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (s *AuthSentCode) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCode#5e002502 as nil")
	}
	b.PutID(AuthSentCodeTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AuthSentCode) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCode#5e002502 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.sentCode#5e002502: field flags: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode auth.sentCode#5e002502: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.sentCode#5e002502: field type: %w", err)
	}
	b.PutString(s.PhoneCodeHash)
	if s.Flags.Has(1) {
		if s.NextType == nil {
			return fmt.Errorf("unable to encode auth.sentCode#5e002502: field next_type is nil")
		}
		if err := s.NextType.Encode(b); err != nil {
			return fmt.Errorf("unable to encode auth.sentCode#5e002502: field next_type: %w", err)
		}
	}
	if s.Flags.Has(2) {
		b.PutInt(s.Timeout)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSentCode) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCode#5e002502 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCode#5e002502: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AuthSentCode) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCode#5e002502 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.sentCode#5e002502: field flags: %w", err)
		}
	}
	{
		value, err := DecodeAuthSentCodeType(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCode#5e002502: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCode#5e002502: field phone_code_hash: %w", err)
		}
		s.PhoneCodeHash = value
	}
	if s.Flags.Has(1) {
		value, err := DecodeAuthCodeType(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCode#5e002502: field next_type: %w", err)
		}
		s.NextType = value
	}
	if s.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCode#5e002502: field timeout: %w", err)
		}
		s.Timeout = value
	}
	return nil
}

// GetType returns value of Type field.
func (s *AuthSentCode) GetType() (value AuthSentCodeTypeClass) {
	return s.Type
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (s *AuthSentCode) GetPhoneCodeHash() (value string) {
	return s.PhoneCodeHash
}

// SetNextType sets value of NextType conditional field.
func (s *AuthSentCode) SetNextType(value AuthCodeTypeClass) {
	s.Flags.Set(1)
	s.NextType = value
}

// GetNextType returns value of NextType conditional field and
// boolean which is true if field was set.
func (s *AuthSentCode) GetNextType() (value AuthCodeTypeClass, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.NextType, true
}

// SetTimeout sets value of Timeout conditional field.
func (s *AuthSentCode) SetTimeout(value int) {
	s.Flags.Set(2)
	s.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (s *AuthSentCode) GetTimeout() (value int, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.Timeout, true
}
