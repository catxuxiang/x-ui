package controller

import (
	"net/http"
	"x-ui/web/session"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func endWith(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func (a *BaseController) checkLogin(c *gin.Context) {
	if !session.IsLogin(c) {
		if endWith(c.Request.RequestURI, "?wIn") {
			c.Request.RequestURI = c.Request.RequestURI[0 : len(c.Request.RequestURI)-4]
			c.Next()
		} else {
			if isAjax(c) {
				pureJsonMsg(c, false, "登录时效已过，请重新登录")
			} else {
				c.Redirect(http.StatusTemporaryRedirect, c.GetString("base_path"))
			}
			c.Abort()
		}
	} else {
		c.Next()
	}
}
