// Code generated by sysl DO NOT EDIT.
package simple

import (
	"context"
)

// DefaultSimpleImpl  ...
type DefaultSimpleImpl struct {
}

// NewDefaultSimpleImpl for simple
func NewDefaultSimpleImpl() *DefaultSimpleImpl {
	return &DefaultSimpleImpl{}
}

// GetFoobarList Client
type GetFoobarListClient struct {
}

// GetTestList Client
type GetTestListClient struct {
}

// ServiceInterface for simple
type ServiceInterface struct {
	GetFoobarList func(ctx context.Context, req *GetFoobarListRequest, client GetFoobarListClient) (*Foo, error)
	GetTestList   func(ctx context.Context, req *GetTestListRequest, client GetTestListClient) (*Stuff, error)
}
