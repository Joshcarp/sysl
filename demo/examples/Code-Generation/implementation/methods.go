package implementation

import (
	"context"

	"github.com/anz-bank/sysl/demo/examples/Code-Generation/simple"
)

func GetStuffList(ctx context.Context, req *simple.GetStuffListRequest, client simple.GetStuffListClient) (*simple.Stuff, error) {
	return &simple.Stuff{Content: "response"}, nil
}

func GetFoobarList(ctx context.Context, req *simple.GetFoobarListRequest, client simple.GetFoobarListClient) (*simple.Foo, error) {
	return &simple.Foo{Content: 123123}, nil
}
