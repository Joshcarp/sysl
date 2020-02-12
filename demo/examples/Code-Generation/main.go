package main

import (
	"context"

	"github.com/anz-bank/sysl/demo/examples/Code-Generation/implementation"
)

func main() {
	implementation.LoadServices(context.Background())
}
