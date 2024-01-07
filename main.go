package main

import (
	"example_http/handlers"
	"sync"

	"github.com/gin-gonic/gin"
)
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go setGin(&wg)
	wg.Wait()
}

func setGin(wg *sync.WaitGroup) {
	router := gin.Default()
	router.LoadHTMLGlob("html/*.*")
	router.GET("/", func(c *gin.Context) {
		handlers.MainH(c)
	})
	router.GET("/createuser", func(c *gin.Context) {
		handlers.GetFormH(c)
	})
	router.POST("/addUser", func(c *gin.Context) {
		handlers.AddUserH(c)
	})
	router.GET("/contacts", func(c *gin.Context) {
		handlers.GetContacts(c)
	})
  router.DELETE("/rm/:id", func(c *gin.Context){
    userId := c.Param("id")
    handlers.DeleteContact(c, userId)
  })
  router.PATCH("/update/:id", func(c *gin.Context){
    userId := c.Param("id")
    handlers.EditUserH(c, userId)
  })
  router.GET("/update/:id", func(c *gin.Context){
    userId := c.Param("id")
    handlers.EditUser(c, userId)
  })
	router.Run(":8080")

}
