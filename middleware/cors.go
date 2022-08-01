package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware cors
// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
// 		fmt.Println("origin is ", c.Request.Header.Get("Origin"))
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

// 		if c.Request.Method == "OPTIONS" {
// 			// c.AbortWithStatus(204)
// 			c.JSON(200, "ok")
// 			return
// 		}

// 		c.Next()
// 	}
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowList := map[string]bool{
			"http://localhost:3000": true,
		}

		if origin := c.Request.Header.Get("Origin"); allowList[origin] {
			fmt.Println("HERE", c.Request.Header.Get("Origin"))
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			`Content-Type, 
			 Content-Length, 
			 Accept-Encoding, 
			 X-CSRF-Token, 
			 Authorization, 
			 accept, 
			 origin, 
			 Cache-Control, 
			 X-Requested-With, 
			 sessionID
			`,
		)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
			c.AbortWithStatus(204)
			fmt.Println(c.Writer.Header().Get(("Access-Control-Allow-Origin")))
			return
		}

		c.Next()
	}
}
