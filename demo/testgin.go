package demo

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func fake_main() {
    r := gin.Default()
    r.GET("/test/user/:user/psw/:psw", TestHandler)
    r.Run(":8080")
}

// This Handler is used to test parameter binding
func TestHandler(c *gin.Context) {
    user := c.Params.ByName("user")
    psw := c.Params.ByName("psw")
    result := fmt.Sprintf("user is %v and psw is %v\n", user, psw)
    c.JSON(200, gin.H{"status": "ok", "Result": result})
}