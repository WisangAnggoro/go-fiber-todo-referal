package routes

import (
	"github.com/WisangAnggoro/go-fiber-todo-referal/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
