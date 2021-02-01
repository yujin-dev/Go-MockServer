package rest

import (
	"net/http"

	"github.com/YJ-dev/go-server/handlers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./app/build", true)))

	api := router.Group("/api")
	{
		api.GET("/server_status", func(c *gin.Context) {
			result := handlers.GetStatus()
			c.JSON(http.StatusOK, result)
		})

		api.GET("/config", func(c *gin.Context) {
			result := handlers.GetConfig()
			c.JSON(http.StatusOK, result)
		})
	}

	router.Run(":5000")
}

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request) // 1. URL 2. http.Handler
// }

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route := req.URL.Path
	switch route {
	case "/server_status":
		handlers.GetStatus()
	case "/config":
		handlers.GetConfig()
	}
}

func BuiltInRun() {
	http.Handle("/", new(testHandler))
	http.ListenAndServe(":5000", nil)
}
