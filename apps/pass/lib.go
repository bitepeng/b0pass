package app

import (
	"b0go/core/engine"

	"github.com/gin-gonic/gin"
)

/***** HttpRequest *****/

// GET
func GET(url, param, title string, handle ...gin.HandlerFunc) {
	engine.Router(appId, "GET", url, param, title, handle...)
}

// POST
func POST(url, param, title string, handle ...gin.HandlerFunc) {
	engine.Router(appId, "POST", url, param, title, handle...)
}

// GETX
func GETX(url, param, title string, handle gin.HandlerFunc, mode string) {
	engine.Router(appId, "GET", url, param, "(Auth)"+title, JWTMiddleware(mode), handle)
}

// POSTX
func POSTX(url, param, title string, handle gin.HandlerFunc, mode string) {
	engine.Router(appId, "POST", url, param, "(Auth)"+title, JWTMiddleware(mode), handle)
}
