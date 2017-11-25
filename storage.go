// Package config implements methods for get storage variables.
//
// script/app/config/storage file.
package configuration

import (
	"errors"
	"path/filepath"
)

const (
	storageFile   = "storage.py"
	storageRegexp = `(?:[\s]{2}'token':[\s]{1}'(?P<token>[\S]+)')`
)

var (
	storageStruct   storage
	storageErrEmpty = errors.New("empty Storage configuration")
	storageLoaded   = false
)

type storage struct {
	token string
}

func (s storage) Token() string {
	return s.token
}

func storageLoad() {
	storageLoaded = true

	buffer, err := cfg(filepath.Join(DirScriptConfig, storageFile), storageRegexp)

	if err != nil {
		panic(storageErrEmpty)
	}

	storageStruct = storage{
		token: string(buffer["token"]),
	}

	if (storage{}) == storageStruct {
		panic(storageErrEmpty)
	}
}

// Expose Storage configuration.
func StorageConfiguration() *storage {
	if !storageLoaded {
		storageLoad()
	}

	return &storageStruct
}
