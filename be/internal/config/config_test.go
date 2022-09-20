package config

import (
	"neko-question-box-be/internal/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	logger.InitLogger()

	m.Run()
}

func TestInitConfig(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		assert := assert.New(t)

		InitConfig(true)

		assert.Equal(Conf.Port, 5000)
		assert.Equal(Conf.Database.Port, 15432)
		assert.Equal(Conf.Database.Host, "localhost")
		assert.Equal(Conf.Database.Username, "postgres")
		assert.Equal(Conf.Database.Password, "postgres")
		assert.Equal(Conf.Database.Database, "postgres")
	})
}
