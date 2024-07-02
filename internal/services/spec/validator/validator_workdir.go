package validator

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/LazarenkoA/go-arch-lint/internal/models/arch"
	"github.com/LazarenkoA/go-arch-lint/internal/services/spec"
)

type validatorWorkDir struct {
	utils *utils
}

func newValidatorWorkDir(utils *utils) *validatorWorkDir {
	return &validatorWorkDir{
		utils: utils,
	}
}

func (v *validatorWorkDir) Validate(doc spec.Document) []arch.Notice {
	notices := make([]arch.Notice, 0)

	rootDir := filepath.Dir(doc.Version().Reference.File)
	absPath := fmt.Sprintf("%s/%s", rootDir, doc.WorkingDirectory().Value)
	absPath = path.Clean(absPath)

	err := v.utils.assertDirectoriesValid(absPath)
	if err != nil {
		notices = append(notices, arch.Notice{
			Notice: fmt.Errorf("invalid workdir '%s' (%s), directory not exist",
				doc.WorkingDirectory().Value,
				absPath,
			),
			Ref: doc.WorkingDirectory().Reference,
		})
	}

	return notices
}
