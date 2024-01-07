package dbconn

import (
	"log"
	"github.com/jmoiron/sqlx"
  "strconv"
	_ "github.com/lib/pq"
  "example_http/config"
)

type User struct {
  Id          int `db:"id"`
	FirstName   string `db:"firstname"`
	LastName    string `db:"lastname"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phonenumber"`
}

func dbConn() *sqlx.DB {
  connString := "user=" + config.GetConfig().Username + " dbname=" + config.GetConfig().Db + " sslmode=disable password=" + config.GetConfig().Password + " host=" + config.GetConfig().Host
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Println(err)
		dbConn()
	}

	return db
}

func CreateUser(user *User) {
	db := dbConn()
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO users(firstName, lastName, email, phoneNumber) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.PhoneNumber)
	tx.Commit()
	defer db.Close()
}

func ListUsers() []User {
	users := []User{}
	db := dbConn()
	db.Select(&users, "SELECT id,firstName,lastName,email,phoneNumber from users")
	defer db.Close()
	return users
}

func DeleteUser(id string) {
  db := dbConn()
  tx := db.MustBegin()
  tx.MustExec("DELETE FROM users where id = $1", id)
  tx.Commit()
  defer db.Close()
}

func EditUser(user *User)  {
  db := dbConn()
  tx := db.MustBegin()
  tx.MustExec("UPDATE users SET firstName = $1,lastName = $2, email = $3, phoneNumber=$4 WHERE id = $5", user.FirstName, user.LastName,user.Email,user.PhoneNumber, user.Id )
  tx.Commit()
  defer db.Close()
}

func SelectUser(id string) User {
  user := User{}
	db := dbConn()
  user.Id,_ = strconv.Atoi(id)
	db.Get(&user, "SELECT * from users WHERE id = $1", user.Id)
	defer db.Close()
	return user
}

