package main

import (
	"encoding/base64"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"quillnote.arunk140.com/m/quilltypes"
)

func GetDB() *gorm.DB {
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func AddUser(username string, password string) {
	db := GetDB()
	var user quilltypes.User
	tx := db.First(&user, "username = ?", strings.TrimSpace(username))
	if tx.RowsAffected != 0 {
		return
	}
	user.Username = strings.TrimSpace(username)
	user.Auth.Username = strings.TrimSpace(username)
	user.Auth.Password = HashPassword(strings.TrimSpace(password))
	db.Create(&user)
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

func ValidateAuthorizationHeader(Authorization string) (*quilltypes.User, bool) {
	decodedString, err := base64.StdEncoding.DecodeString(Authorization[6:])
	if err != nil {
		return nil, false
	}
	usernamePassword := strings.Split(string(decodedString), ":")
	if len(usernamePassword) != 2 {
		return nil, false
	}
	username := usernamePassword[0]
	password := usernamePassword[1]

	db := GetDB()
	var auth quilltypes.Auth
	tx := db.First(&auth, "username = ?", username)
	if tx.RowsAffected == 0 {
		return nil, false
	}

	error := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))

	if error != nil {
		return nil, false
	}

	var user quilltypes.User
	db.First(&user, "auth_id = ?", auth.ID)
	return &user, true
}

func AuthenticateMiddleware(c *fiber.Ctx) error {
	Authorization := c.Get("Authorization")
	if Authorization == "" {
		return c.SendStatus(401)
	}
	user, ok := ValidateAuthorizationHeader(Authorization)
	if !ok {
		return c.SendStatus(401)
	}
	c.Locals("user", user.ID)
	return c.Next()
}

func Migrate() {
	db := GetDB()
	db.AutoMigrate(&quilltypes.Note{})
	db.AutoMigrate(&quilltypes.User{})
	db.AutoMigrate(&quilltypes.Auth{})
}
