package users

import (
	"log"
	"main/internal/pkg/database"

	"main/pkg/hash"

	"gorm.io/gorm"
)

var db = database.Db

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (user *User) create() {
	hashedPassword, err := hash.hashString(user.Password)

	tx := database.Db.Exec("INSERT INTO Users(Email,Password) VALUES(?,?)", user.Email, hashedPassword)

	if tx == nil {
		log.Fatal(tx)
	}
}

func (user *User) authenticate() bool {
	tx := database.Db.Exec("SELECT password FROM users WHERE email = ?", user.Email)

	if tx == nil {
		log.Fatal(tx)
	}

	var hashedPassword string
	err = tx.Scan(&hashedPassword)

	if err != nil {
		log.Fatal(err)
	}

	return hash.checkStringHash(user.Password, hashedPassword)
}

// getUserIdByEmail check if a user exists in database by given email
func getUserIdByEmail(email string) (uint, error) {
	statement := database.Db.Raw("SELECT id FROM users WHERE email = ?")
	// if  != nil {
	// 	log.Fatal(err)
	// }

	var userId uint
	err = sr.Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

// GetUserByID check if a user exists in database and return the user object.
func getEmailById(userId uint) (User, error) {
	statement := database.Db.Raw("SELECT email FROM users WHERE id = ?", userId)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	var email string
	err = row.Scan(&email)

	if err != nil {
		return User{}, err
	}

	return User{ID: userId, Email: email}, nil
}
