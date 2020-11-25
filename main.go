package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
	}))

	unAuthorized := router.Group("/")

	authorized := router.Group("/", func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" || token != os.Getenv("AUTH_TOKEN") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "please provide valid token"})
		}
	})

	ok := InitApp(unAuthorized, authorized)
	if !ok {
		println("Unable to Initialize Application")
		return
	}
	router.Run()
}

// func mainX() {
// 	router := gin.Default()
// 	router.GET("/ping", ping)
// 	// This handler will match /user/john but will not match /user/ or /user
// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	// However, this one will match /user/john/ and also /user/john/send
// 	// If no other routers match /user/john, it will redirect to /user/john/
// 	router.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")
// 		message := name + " is " + action
// 		c.String(http.StatusOK, message)
// 	})

// 	// For each matched request Context will hold the route definition
// 	router.POST("/user/:name/*action", func(c *gin.Context) {
// 		//c.FullPath() == "/user/:name/*action" // true
// 	})
// 	// Query string parameters are parsed using the existing underlying request object.
// 	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
// 	router.GET("/welcome", func(c *gin.Context) {
// 		firstname := c.DefaultQuery("firstname", "Guest")
// 		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
// 		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
// 	})
// 	router.Run()
// }
// func ping(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// }
