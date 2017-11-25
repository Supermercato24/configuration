package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testEnvModel(t *testing.T, env *env) {
	// env
	if env.IsProduction() {
		t.Error("env must be staging or local")
	}

	// db
	assert.NotEmpty(t, env.Db.Host)
	assert.NotEmpty(t, env.Db.Database)
	assert.NotEmpty(t, env.Db.Credentials.Username())
	assert.NotEmpty(t, env.Db.Credentials.Password())

	// mail
	assert.NotEmpty(t, env.Mail.Credentials.Username())
	assert.NotEmpty(t, env.Mail.Credentials.Password())

	// Username
	assert.NotEqual(t, env.Db.Credentials.Username(), env.Mail.Credentials.Username())

	// Password
	assert.NotEqual(t, env.Db.Credentials.Password(), env.Mail.Credentials.Password())
}

func TestEnv(t *testing.T) {
	assert.NotPanics(t, func() { EnvConfiguration() })

	env := EnvConfiguration()

	testEnvModel(t, env)
}

func BenchmarkEnv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EnvConfiguration()
	}
}
