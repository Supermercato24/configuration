package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testStorageModel(t *testing.T, storage *storage) {
	// storage
	assert.NotEmpty(t, storage.Token())
}

func TestStorage(t *testing.T) {
	assert.NotPanics(t, func() { StorageConfiguration() })

	storage := StorageConfiguration()

	testStorageModel(t, storage)
}

func BenchmarkStorage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StorageConfiguration()
	}
}
