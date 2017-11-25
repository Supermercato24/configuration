package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSoaModel(t *testing.T, soa *soa) {
	// UA
	assert.NotEmpty(t, soa.UA())

	// cp24
	assert.NotEmpty(t, soa.Cp24.Domain())
	assert.NotEmpty(t, soa.Cp24.Token())

	// s24
	assert.NotEmpty(t, soa.S24.Domain())
	assert.NotEmpty(t, soa.S24.Token())

	// sm
	assert.NotEmpty(t, soa.Sm.Domain())
	assert.NotEmpty(t, soa.Sm.Token())

	// Domain
	assert.NotEqual(t, soa.Cp24.Domain(), soa.S24.Domain())
	assert.NotEqual(t, soa.Cp24.Domain(), soa.Sm.Domain())
	assert.Contains(t, soa.Cp24.Domain(), "http")

	assert.NotEqual(t, soa.S24.Domain(), soa.Cp24.Domain())
	assert.NotEqual(t, soa.S24.Domain(), soa.Sm.Domain())
	assert.Contains(t, soa.S24.Domain(), "http")

	assert.NotEqual(t, soa.Sm.Domain(), soa.S24.Domain())
	assert.NotEqual(t, soa.Sm.Domain(), soa.Cp24.Domain())
	assert.Contains(t, soa.Sm.Domain(), "http")

	// Token
	assert.NotEqual(t, soa.Cp24.Token(), soa.S24.Token())
	assert.NotEqual(t, soa.Cp24.Token(), soa.Sm.Token())
	assert.Contains(t, soa.Cp24.Token(), "Basic")

	assert.NotEqual(t, soa.S24.Token(), soa.Cp24.Token())
	assert.NotEqual(t, soa.S24.Token(), soa.Sm.Token())
	assert.Contains(t, soa.Cp24.Token(), "Basic")

	assert.NotEqual(t, soa.Sm.Token(), soa.S24.Token())
	assert.NotEqual(t, soa.Sm.Token(), soa.Cp24.Token())
	assert.Contains(t, soa.Sm.Token(), "Basic")
}

func TestSoa(t *testing.T) {
	assert.NotPanics(t, func() { SoaConfiguration() })

	soa := SoaConfiguration()

	testSoaModel(t, soa)
}

func BenchmarkSoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SoaConfiguration()
	}
}
