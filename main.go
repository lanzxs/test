package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `binding:"required"`
	Age int
}

func main() {

	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	app := gin.Default()
	app.Use(RateLimiter())
	app.GET("/login", func(c *gin.Context) {
		c.JSON(200,gin.H{"lan" : "lan"})
	})

	/*binding.Validator.Engine().RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})*/
	//token,_ := SetToken("lan")
	//fmt.Println("token,",token)
	//app.Use(JwtToken())
	app.POST("/post", func(c *gin.Context) {
		user := &User{}
		err := c.Bind(user)
		fmt.Println(err)
		fmt.Println(user)
		c.JSON(200,gin.H{
			"lan" : "lan",
		})
	})
	app.Run()
}
