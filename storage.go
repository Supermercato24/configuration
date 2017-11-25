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
	storageStruct storage
)

type storage struct {
	token string
}

func (s storage) Token() string {
	return s.token
}

func init() {
	buffer, err := cfg(filepath.Join(DirScriptConfig, storageFile), storageRegexp)

	if err != nil {
		return
	}

	storageStruct = storage{
		token: string(buffer["token"]),
	}
}

// Expose Storage configuration.
func StorageConfiguration() *storage {
	if (storage{}) == storageStruct {
		err := errors.New("empty Storage configuration")
		panic(err)
	}

	return &storageStruct
}
