// Code generated by ogen, DO NOT EDIT.

package scheme

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode encodes BirthDate as json.
func (s BirthDate) Encode(e *jx.Encoder) {
	unwrapped := time.Time(s)

	json.EncodeDate(e, unwrapped)
}

// Decode decodes BirthDate from json.
func (s *BirthDate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BirthDate to nil")
	}
	var unwrapped time.Time
	if err := func() error {
		v, err := json.DecodeDate(d)
		unwrapped = v
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = BirthDate(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s BirthDate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BirthDate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginPostOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginPostOK) encodeFields(e *jx.Encoder) {
	{
		if s.Token.Set {
			e.FieldStart("token")
			s.Token.Encode(e)
		}
	}
}

var jsonFieldsNameOfLoginPostOK = [1]string{
	0: "token",
}

// Decode decodes LoginPostOK from json.
func (s *LoginPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginPostOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "token":
			if err := func() error {
				s.Token.Reset()
				if err := s.Token.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"token\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginPostOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Password.Set {
			e.FieldStart("password")
			s.Password.Encode(e)
		}
	}
}

var jsonFieldsNameOfLoginPostReq = [2]string{
	0: "id",
	1: "password",
}

// Decode decodes LoginPostReq from json.
func (s *LoginPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginPostReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "password":
			if err := func() error {
				s.Password.Reset()
				if err := s.Password.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes BirthDate as json.
func (o OptBirthDate) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes BirthDate from json.
func (o *OptBirthDate) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBirthDate to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBirthDate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBirthDate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes LoginPostReq as json.
func (o OptLoginPostReq) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes LoginPostReq from json.
func (o *OptLoginPostReq) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptLoginPostReq to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptLoginPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptLoginPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UserId as json.
func (o OptUserId) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes UserId from json.
func (o *OptUserId) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUserId to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUserId) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUserId) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UserRegisterPostReq as json.
func (o OptUserRegisterPostReq) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes UserRegisterPostReq from json.
func (o *OptUserRegisterPostReq) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUserRegisterPostReq to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUserRegisterPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUserRegisterPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *R5xx) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *R5xx) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		if s.RequestID.Set {
			e.FieldStart("request_id")
			s.RequestID.Encode(e)
		}
	}
	{
		if s.Code.Set {
			e.FieldStart("code")
			s.Code.Encode(e)
		}
	}
}

var jsonFieldsNameOfR5xx = [3]string{
	0: "message",
	1: "request_id",
	2: "code",
}

// Decode decodes R5xx from json.
func (s *R5xx) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode R5xx to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "request_id":
			if err := func() error {
				s.RequestID.Reset()
				if err := s.RequestID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"request_id\"")
			}
		case "code":
			if err := func() error {
				s.Code.Reset()
				if err := s.Code.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode R5xx")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfR5xx) {
					name = jsonFieldsNameOfR5xx[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *R5xx) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *R5xx) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *User) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *User) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.FirstName.Set {
			e.FieldStart("first_name")
			s.FirstName.Encode(e)
		}
	}
	{
		if s.SecondName.Set {
			e.FieldStart("second_name")
			s.SecondName.Encode(e)
		}
	}
	{
		if s.Birthdate.Set {
			e.FieldStart("birthdate")
			s.Birthdate.Encode(e)
		}
	}
	{
		if s.Biography.Set {
			e.FieldStart("biography")
			s.Biography.Encode(e)
		}
	}
	{
		if s.City.Set {
			e.FieldStart("city")
			s.City.Encode(e)
		}
	}
}

var jsonFieldsNameOfUser = [6]string{
	0: "id",
	1: "first_name",
	2: "second_name",
	3: "birthdate",
	4: "biography",
	5: "city",
}

// Decode decodes User from json.
func (s *User) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode User to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "first_name":
			if err := func() error {
				s.FirstName.Reset()
				if err := s.FirstName.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"first_name\"")
			}
		case "second_name":
			if err := func() error {
				s.SecondName.Reset()
				if err := s.SecondName.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"second_name\"")
			}
		case "birthdate":
			if err := func() error {
				s.Birthdate.Reset()
				if err := s.Birthdate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birthdate\"")
			}
		case "biography":
			if err := func() error {
				s.Biography.Reset()
				if err := s.Biography.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "city":
			if err := func() error {
				s.City.Reset()
				if err := s.City.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"city\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode User")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *User) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *User) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UserId as json.
func (s UserId) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes UserId from json.
func (s *UserId) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserId to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UserId(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s UserId) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserId) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserRegisterPostOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserRegisterPostOK) encodeFields(e *jx.Encoder) {
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserRegisterPostOK = [1]string{
	0: "user_id",
}

// Decode decodes UserRegisterPostOK from json.
func (s *UserRegisterPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserRegisterPostOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserRegisterPostOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserRegisterPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserRegisterPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserRegisterPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserRegisterPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.FirstName.Set {
			e.FieldStart("first_name")
			s.FirstName.Encode(e)
		}
	}
	{
		if s.SecondName.Set {
			e.FieldStart("second_name")
			s.SecondName.Encode(e)
		}
	}
	{
		if s.Birthdate.Set {
			e.FieldStart("birthdate")
			s.Birthdate.Encode(e)
		}
	}
	{
		if s.Biography.Set {
			e.FieldStart("biography")
			s.Biography.Encode(e)
		}
	}
	{
		if s.City.Set {
			e.FieldStart("city")
			s.City.Encode(e)
		}
	}
	{
		if s.Password.Set {
			e.FieldStart("password")
			s.Password.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserRegisterPostReq = [6]string{
	0: "first_name",
	1: "second_name",
	2: "birthdate",
	3: "biography",
	4: "city",
	5: "password",
}

// Decode decodes UserRegisterPostReq from json.
func (s *UserRegisterPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserRegisterPostReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "first_name":
			if err := func() error {
				s.FirstName.Reset()
				if err := s.FirstName.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"first_name\"")
			}
		case "second_name":
			if err := func() error {
				s.SecondName.Reset()
				if err := s.SecondName.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"second_name\"")
			}
		case "birthdate":
			if err := func() error {
				s.Birthdate.Reset()
				if err := s.Birthdate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birthdate\"")
			}
		case "biography":
			if err := func() error {
				s.Biography.Reset()
				if err := s.Biography.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "city":
			if err := func() error {
				s.City.Reset()
				if err := s.City.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"city\"")
			}
		case "password":
			if err := func() error {
				s.Password.Reset()
				if err := s.Password.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserRegisterPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserRegisterPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserRegisterPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
