// Code generated by ogen, DO NOT EDIT.

package scheme

import (
	"github.com/go-faster/errors"
)

func (s UserSearchGetOKApplicationJSON) Validate() error {
	alias := ([]User)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}