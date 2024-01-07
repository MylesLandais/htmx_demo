package handlers

import (
	dbconn "example_http/dbConn"
	"net/http"
  "strconv"

	"github.com/gin-gonic/gin"
)

func MainH(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetFormH(c *gin.Context) {
	c.HTML(http.StatusOK, "adduser.html", gin.H{})
}
func AddUserH(c *gin.Context) {
	person := new(dbconn.User)
	person.FirstName = c.Request.FormValue("firstName")
	person.LastName = c.Request.FormValue("lastName")
	person.Email = c.Request.FormValue("email")
	person.PhoneNumber = c.Request.FormValue("phoneNumber")
	dbconn.CreateUser(person)
}

func GetContacts(c *gin.Context) {
	users := dbconn.ListUsers()
	c.HTML(http.StatusOK, "contacts.html", gin.H{
		"Users": users,
	})
}

func DeleteContact(c *gin.Context, id string){
  dbconn.DeleteUser(id) 
  c.Redirect(http.StatusMovedPermanently, "/contacts")
}


func EditUserH(c *gin.Context, id string) {
	person := new(dbconn.User)
	person.FirstName = c.Request.FormValue("firstName")
	person.LastName = c.Request.FormValue("lastName")
	person.Email = c.Request.FormValue("email")
	person.PhoneNumber = c.Request.FormValue("phoneNumber")
  person.Id,_ = strconv.Atoi(id)
  dbconn.EditUser(person)
}

func EditUser(c *gin.Context, id string) {
  user := dbconn.SelectUser(id)
  c.HTML(http.StatusOK,"edit.html", gin.H{
    "User": user,
  })   

}
