package routes

import (
	"fmt"
	"sqap/internal/models"
	"sqap/internal/services"

	"github.com/gofiber/fiber/v2"
)

type todoImpl struct {
	service *services.TodoService
}

func RegisterTodo(app *fiber.App, service *services.TodoService) {
	impl := &todoImpl{
		service: service,
	}

	group := app.Group("todos")
	group.Get("/", impl.Index)
	group.Get("/:uid", impl.Get)
	group.Post("/", impl.Post)

}

func (impl *todoImpl) Index(c *fiber.Ctx) error {
	todo, err := impl.service.Get()

	if err != nil {
		fmt.Println("error: ", err)
		return c.SendStatus(500)
	}

	return c.JSON(todo)
}

func (impl *todoImpl) Get(c *fiber.Ctx) error {
	uid := c.Params("uid")

	todo, err := impl.service.GetTodo(uid)

	if err != nil {
		fmt.Println("error: ", err)
		return c.SendStatus(500)
	}

	return c.JSON(todo)
}

func (impl *todoImpl) Post(c *fiber.Ctx) error {
	todo := models.TodoRequest{}

	if err := c.BodyParser(&todo); err != nil {
		fmt.Println("error: ", err)
		return c.SendStatus(422)
	}

	err := impl.service.PostTodo(&todo)

	if err != nil {
		fmt.Println("error: ", err)
		return c.SendStatus(500)
	}

	return c.JSON(todo)
}
