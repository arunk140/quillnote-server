package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"quillnote.arunk140.com/m/quilltypes"
)

func GetNotes(c *fiber.Ctx) error {
	UserID := c.Locals("user")
	gormDb := GetDB()
	var notes []quilltypes.Note
	gormDb.Find(&notes, "user_id = ?", int(UserID.(uint)))
	return c.JSON(notes)
}

func CapabilitiesAPI(c *fiber.Ctx) error {
	capabilities := quilltypes.Capabilities{
		OCS: quilltypes.OCS{
			Data: quilltypes.OCSData{
				Capabilities: quilltypes.OCSCapabilities{
					Notes: quilltypes.OCSNotes{
						APIVersion: []string{"0.2", "1.0"},
						Version:    "3.6.0",
					},
				},
			},
		},
	}
	return c.JSON(capabilities)
}

func CreateNote(c *fiber.Ctx) error {
	var note quilltypes.Note

	if err := c.BodyParser(&note); err != nil {
		return err
	}

	gormDb := GetDB()
	ts := gormDb.Create(&note)

	if ts.RowsAffected == 0 {
		return c.SendStatus(400)
	}

	// Update the note with the user id
	UserID := c.Locals("user")
	note.UserID = UserID.(uint)

	ts = gormDb.Model(&note).Updates(map[string]interface{}{"user_id": UserID.(uint)})
	if ts.RowsAffected == 0 {
		return c.SendStatus(400)
	}

	return c.JSON(note)
}

func UpdateNote(c *fiber.Ctx) error {
	var noteId = c.Params("id")

	noteIdInt, err := strconv.Atoi(noteId)
	if err != nil {
		return c.SendStatus(400)
	}

	UserID := c.Locals("user")
	var note quilltypes.Note
	note.UserID = UserID.(uint)
	if err := c.BodyParser(&note); err != nil {
		return c.SendStatus(400)
	}
	gormDb := GetDB()

	var noteToUpdate quilltypes.Note
	ts := gormDb.First(&noteToUpdate, quilltypes.Note{ID: noteIdInt, UserID: UserID.(uint)})
	if ts.RowsAffected == 0 {
		return CreateNote(c)
	}
	noteToUpdate = note
	gormDb.Save(&noteToUpdate)
	return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	var noteId = c.Params("id")
	noteIdInt, err := strconv.Atoi(noteId)
	if err != nil {
		return c.SendStatus(400)
	}

	gormDb := GetDB()
	var noteToDelete quilltypes.Note
	UserID := c.Locals("user")
	ts := gormDb.First(&noteToDelete, quilltypes.Note{ID: noteIdInt, UserID: UserID.(uint)})
	if ts.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	gormDb.Delete(&noteToDelete)
	return c.SendStatus(200)
}

func NotFound(c *fiber.Ctx) error {
	return c.SendStatus(404)
}

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"version": "0.0.1",
	})
}
