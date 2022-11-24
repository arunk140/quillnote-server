package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {

	if len(os.Args) > 1 {
		argsStr := strings.Join(os.Args[1:], " ")
		done := ProcessCommand(argsStr)
		if !done {
			log.Fatal("Invalid Command")
			os.Exit(1)
		} else {
			log.Println("Done")
			os.Exit(0)
		}
	}

	SetupCLIHandler()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(AuthenticateMiddleware)

	app.Get("/", Health)

	app.Get("/ocs/v2.php/cloud/capabilities", CapabilitiesAPI)

	notesAPIs := app.Group("/index.php/apps/notes/api/v1/notes")
	notesAPIs.Get("/", GetNotes)
	notesAPIs.Post("/", CreateNote)
	notesAPIs.Put("/:id", UpdateNote)
	notesAPIs.Delete("/:id", DeleteNote)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Quillnote Server Stats"}))

	app.All("/*", NotFound)
	log.Fatal(app.Listen(":3000"))
}

func _(UserPassListString string) (map[string]string, error) {
	if UserPassListString == "" {
		return nil, errors.New("user list is empty")
	}

	UserPassList := strings.Split(UserPassListString, ",")
	if len(UserPassList) == 0 {
		return nil, errors.New("user list is empty")
	}

	Users := make(map[string]string)
	for _, UserPass := range UserPassList {
		UserPassSplit := strings.Split(UserPass, ":")
		if len(UserPassSplit) != 2 {
			return nil, errors.New("user list is invalid")
		}
		Users[UserPassSplit[0]] = UserPassSplit[1]
	}
	return Users, nil
}
