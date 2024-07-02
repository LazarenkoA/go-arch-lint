package use_foo

import (
	"github.com/LazarenkoA/go-arch-lint/test/check/project/variadic/foo"
)

var (
	foo1 = foo.FooF("boo")
	foo2 = foo.FooF("boo", 5)
	foo3 = foo.FooF("boo", 5, "test", 15)
)
