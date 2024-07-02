package decoder

import "github.com/LazarenkoA/go-arch-lint/internal/services/spec"

type doc interface {
	spec.Document

	postSetup()
}
