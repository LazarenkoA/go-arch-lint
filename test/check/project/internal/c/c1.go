package c

import "github.com/LazarenkoA/go-arch-lint/test/check/project/internal/a"

func C1() {
	a.A1() // not allowed
}
