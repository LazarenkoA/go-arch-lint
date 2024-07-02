package assembler

import (
	"github.com/LazarenkoA/go-arch-lint/internal/models/arch"
	"github.com/LazarenkoA/go-arch-lint/internal/services/spec"
)

type allowAssembler struct {
}

func newAllowAssembler() *allowAssembler {
	return &allowAssembler{}
}

func (efa *allowAssembler) assemble(spec *arch.Spec, document spec.Document) error {
	spec.Allow = arch.Allow{
		DepOnAnyVendor: document.Options().IsDependOnAnyVendor(),
		DeepScan:       document.Options().DeepScan(),
	}

	return nil
}
