package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func helloworld( c *gin.Context){
c.String(http.StatusOK, "Hello World! from root Time : %s", time.Now().Format(time.RFC3339Nano))
}

func jsonHelloWorld(c *gin.Context) {
 c.JSON(http.StatusOK, gin.H{
  "message": "Hello World!  Here we are from now main.go ! Time : " + time.Now().Format(time.RFC3339Nano),
 })
}

func main() {
 engine := gin.New()
 engine.GET("/", helloworld)
 engine.GET("/json", jsonHelloWorld)
 engine.Run("localhost:8080")
}