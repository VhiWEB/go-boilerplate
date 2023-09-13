package users

import (
	"log"
	"time"

	"main/internal/pkg/database"
	"main/pkg/hash"

	"gorm.io/gorm"
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
	// defer sqlDb.Close()

	return users
}

func (user *User) Create() (bool, error) {
	hashedPassword, err := hash.HashString(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	tx := db.Exec("INSERT INTO users(email,password) VALUES(?,?)", user.Email, hashedPassword)
	// defer sqlDb.Close()

	if tx == nil {
		log.Fatal(tx.Error)
		return false, tx.Error
	}

	return true, nil
}

func Authenticate(email string, password string) (bool, *User, *gorm.DB) {
	tx := db.Raw("SELECT * FROM users WHERE email = ?", email)

	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	var user *User
	tx.Scan(&user)
	// defer sqlDb.Close()

	if hash.CheckStringHash(password, user.Password) {
		return true, user, tx
	}

	return false, nil, tx
}

// Check if a user exists in database by given ID
func GetById(id float64) (User, *gorm.DB) {
	tx := db.Raw("SELECT * FROM users WHERE id = ?", id)

	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	var user User
	tx.Scan(&user)
	// defer sqlDb.Close()

	return user, tx
}
