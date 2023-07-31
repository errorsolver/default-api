package main

import (
	"api/delivery"
)

func main() {
	delivery.NewServer().Run()
}

// var ApplicationName = "API Default"
// var JwtSigningMethod = jwt.SigningMethodHS256
// var JwtSignatureKey = []byte("P@ssword")

// type MyClaims struct {
// 	jwt.RegisteredClaims
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// }

// type Credential struct {
// 	Username string
// 	Password string
// }

// func main() {
// 	router := gin.Default()
// 	router.Use(AuthTokenMiddleware())

// 	router.POST("/login", func(c *gin.Context) {
// 		var user Credential
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(401, gin.H{
// 				"message": "can't bind struct",
// 			})
// 			return
// 		}

// 		if user.Username != "admin" && user.Password != "123" {
// 			c.AbortWithStatus(401)
// 		}
// 		c.JSON(200, gin.H{
// 			"message": "Success Login",
// 		})
// 	})

// 	router.GET("/customer", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "customer",
// 		})
// 	})

// 	if err := router.Run(":8080"); err != nil {
// 		panic(err)
// 	}
// }

// func GenerateToken(username string, email string) (string, error) {

// }

// func AuthTokenMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if c.Request.URL.Path != "/login" {
// 			h := authHeader{}
// 			if err := c.ShouldBindHeader(&h); err != nil {
// 				c.AbortWithStatusJSON(401, gin.H{
// 					"message": "Unauthorized",
// 				})
// 			}
// 			// if h.AuthorizationHeader != "123" {
// 			// 	c.AbortWithStatusJSON(401, gin.H{
// 			// 		"message": "Unauthorized",
// 			// 	})
// 			// }
// 			// c.Next()

// 			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer", "", -1)
// 			fmt.Println(tokenString)
// 			if tokenString == "" {
// 				c.AbortWithStatusJSON(401, gin.H{
// 					"message": "Unauthorized",
// 				})
// 				return
// 			}

// 			token, err := ParseToken(tokenString)
// 			if err != nil {
// 				c.AbortWithStatusJSON(401, gin.H{
// 					"message": "Unauthorized",
// 				})
// 				return
// 			}
// 			fmt.Println(token)

// 			if token["iss"] != ApplicationName {
// 				c.AbortWithStatusJSON(401, gin.H{
// 					"message": "Unauthorized",
// 				})
// 				return
// 			}
// 		}
// 		c.Next()
// 	}
// }
