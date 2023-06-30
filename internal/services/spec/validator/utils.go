package validator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fe3dback/go-arch-lint/internal/models/arch"
)

type (
	utils struct {
		pathResolver  PathResolver
		document      arch.Document
		rootDirectory string
	}
)

func newUtils(
	pathResolver PathResolver,
	document arch.Document,
	rootDirectory string,
) *utils {
	return &utils{
		pathResolver:  pathResolver,
		document:      document,
		rootDirectory: rootDirectory,
	}
}

func (u *utils) assertGlobPathValid(localGlobPath string) error {
	absPath := filepath.Clean(fmt.Sprintf("%s/%s", u.rootDirectory, localGlobPath))
	resolved, err := u.pathResolver.Resolve(absPath)
	if err != nil {
		return fmt.Errorf("failed to resolv path: %w", err)
	}

	if len(resolved) == 0 {
		return fmt.Errorf("not found directories for '%s' in '%s'", localGlobPath, absPath)
	}

	return u.assertDirectoriesValid(resolved...)
}

func (u *utils) assertDirectoriesValid(paths ...string) error {
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("directory '%s' not exist", path)
		}
	}

	return nil
}

func (u *utils) assertKnownComponent(name string) error {
	for knownName := range u.document.Components().Map() {
		if name == knownName {
			return nil
		}
	}

	return fmt.Errorf("unknown component '%s'", name)
}

func (u *utils) assertKnownVendor(name string) error {
	for knownName := range u.document.Vendors().Map() {
		if name == knownName {
			return nil
		}
	}

	return fmt.Errorf("unknown vendor '%s'", name)
}