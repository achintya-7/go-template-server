package v1

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// This file should contain functions to start a Test and
// various functions to be repetited in the testing process

func newTestServer() *RouteHandler {
	testRouteHanlder := NewRouteHandler()

	return testRouteHanlder
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Panic in tests: ", r)
	}
}

func createTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	return context, recorder
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
