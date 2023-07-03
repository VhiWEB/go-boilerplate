package users

import (
	"log"
	"main/internal/pkg/database"
	"main/pkg/hash"
	"time"
)

var db = database.Db
var sqlDb = database.SqlDb

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Get() []User {
	var users []User
	db.Find(&users)
	defer sqlDb.Close()

	return users
}

func (user *User) Create() {
	hashedPassword, err := hash.HashString(user.Password)

	if err == nil {
		log.Fatal(err)
	}

	tx := db.Exec("INSERT INTO Users(Email,Password) VALUES(?,?)", user.Email, hashedPassword)
	defer sqlDb.Close()

	if tx == nil {
		log.Fatal(tx)
	}
}

func (user *User) Authenticate() bool {
	tx := db.Raw("SELECT password FROM users WHERE email = ?", user.Email)

	if tx == nil {
		log.Fatal(tx)
	}

	var hashedPassword string
	tx.Scan(&hashedPassword)
	defer sqlDb.Close()

	return hash.CheckStringHash(user.Password, hashedPassword)
}

// Check if a user exists in database by given email
func GetUserByEmail(email string) (User, error) {
	tx := db.Raw("SELECT * FROM users WHERE email = ?", email)

	if tx == nil {
		log.Fatal(tx)
	}

	var user User
	tx.Scan(&user)
	user.Email = email
	defer sqlDb.Close()

	return user, nil
}
