package main

//보류 진행중단
import (
	"github.com/Ko-GyeongTae/Gin_Server_Basic/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		api.Ping(c)
	})
	r.Run()
}
