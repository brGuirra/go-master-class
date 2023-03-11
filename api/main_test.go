package api

import (
	db "github.com/brGuirra/simple-bank/db/sqlc"
	"github.com/brGuirra/simple-bank/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   gofakeit.DigitN(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
