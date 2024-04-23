package app_api

import (
	"api/app_api/components"
	"api/app_api/pages"
	"api/lib"
	"github.com/gin-gonic/gin"
	g "github.com/maragudk/gomponents"
	"net/http"
)

func InitAppAPI(r *gin.Engine, app *lib.App) {
	InitDevReloadWebsocket(r)

	r.GET("/", func(c *gin.Context) {
		body := pages.IndexPage()
		createHandler(&body)(c.Writer, c.Request)
	})
}

func createHandler(body *g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = components.Page(body).Render(w)
	}
}
