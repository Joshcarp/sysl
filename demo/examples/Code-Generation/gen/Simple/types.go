// Code generated by sysl DO NOT EDIT.
package simple

import (
	"time"

	"github.com/rickb777/date"
	"github.service.anz/sysl/server-lib/validator"
)

// Reference imports to suppress unused errors
var _ = time.Parse

// Reference imports to suppress unused errors
var _ = date.Parse

// Foo ...
type Foo struct {
	Content int64 `json:"Content"`
}

// Stuff ...
type Stuff struct {
	Content string `json:"Content"`
}

// GetFoobarListRequest ...
type GetFoobarListRequest struct {
}

// GetTestListRequest ...
type GetTestListRequest struct {
}

// *Foo validator
func (s *Foo) Validate() error {
	return validator.Validate(s)
}

// *Stuff validator
func (s *Stuff) Validate() error {
	return validator.Validate(s)
}
