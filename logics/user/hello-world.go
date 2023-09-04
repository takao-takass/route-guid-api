package logics

import (
	"github.com/gin-gonic/gin"
)

// 文字列"Hello world!"を返却するメソッド
func HelloWorld(c *gin.Context) {
	c.String(200, "Hello World!")
}
