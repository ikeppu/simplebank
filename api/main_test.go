package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ikeppu/simplebank/db/sqlc"
	"github.com/ikeppu/simplebank/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	server, err := NewServer(util.Config{DBSource: "test", TokenSymmetricKey: util.RandomString(32), AccessTokenDuration: time.Minute}, store)

	if err != nil {
		t.Fatal("cannot create server:", err)
	}

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
