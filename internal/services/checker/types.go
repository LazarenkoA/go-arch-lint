package checker

import (
	"context"

	"github.com/LazarenkoA/go-arch-lint/internal/models"
	"github.com/LazarenkoA/go-arch-lint/internal/models/arch"
	"github.com/LazarenkoA/go-arch-lint/internal/models/common"
)

type (
	projectFilesResolver interface {
		ProjectFiles(ctx context.Context, spec arch.Spec) ([]models.FileHold, error)
	}

	checker interface {
		Check(ctx context.Context, spec arch.Spec) (models.CheckResult, error)
	}

	sourceCodeRenderer interface {
		SourceCode(ref common.Reference, highlight bool, showPointer bool) []byte
	}
)
