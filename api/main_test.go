package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// Run the tests
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
