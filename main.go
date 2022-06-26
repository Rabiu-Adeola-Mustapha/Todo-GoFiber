package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct{
	Id   int `json:"id"`
	Name  string `json:"name"`
	Completed  bool `json:"completed"`
}

var todos = []*Todo{
	{Id: 1, Name: "Walk the dog", Completed: false},
	{Id:2, Name: "Clean the dishes", Completed: false},

}

func main(){

	app := fiber.New();

	app.Get("/", func(c *fiber.Ctx)error{

		return c.SendString("hello world")

	})

	// i can aswell do todosRoutes := app.Group("/todos")
	// todosRoutes.Get("/", ...)
	app.Get("/todos", GetTodos);
	app.Post("/todos", CreateTodo);
	app.Get("/todos/:id", GetSingleTodo);
	app.Delete("/todos/:id", DeleteTodo);
	app.Patch("/todos/:id", UpdateTodo);
	

	app.Listen("3000");
	

}

func GetTodos(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(todos);

};

func CreateTodo(c *fiber.Ctx)error{

	type req struct {
		Name string `json:"name"`
	}

	var body req

	err := c.BodyParser(&body);

	if err != nil {
		
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": "cannot parse json",
		})

		
	
	}

	todo := &Todo {
		Id: len(todos)+1,
		Name: body.Name,
		Completed: false,
	}

	todos = append(todos, todo);

	return c.Status(fiber.StatusOK).JSON(todo);
	
}

func GetSingleTodo(c *fiber.Ctx)error{

	paramsId := c.Params("id");

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": "Cannot parse Id",
		})

	}
	for _,todo := range todos {
		if todo.Id == id {

		  return c.Status(fiber.StatusOK).JSON(todo)
			
		};
	};

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"err": "Request Not Found",
	})
}

func DeleteTodo(c *fiber.Ctx)error{
	paramsId := c.Params("id");

	id, err := strconv.Atoi(paramsId);

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": "Cannot parse Id",
		})
	}

	for i, todo := range todos {

		if todo.Id == id {
			
			 todos = append(todos[0:i], todos[i+1:]...)

			 return c.Status(fiber.StatusOK).JSON(todos);
		}
		
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"err": "Request Not Found",
	})
	
};

func UpdateTodo(c *fiber.Ctx)error{

	type req struct {
    
	   Name       *string `json:"name"`  // the wildcard is doing the checking statement if not empty
	   Completed  *bool `json:"completed"`

	}

	var body req

	err := c.BodyParser(&body);

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": "Cannot parse body",
		});
	}

	paramsId := c.Params("id");

	id, err := strconv.Atoi(paramsId);
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": "Cannot parse id",
		});
	};

	

	var todo *Todo;

    for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	
    }

	if todo == nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"err": "Cannot Find the request",
		})
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}
	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return c.Status(fiber.StatusOK).JSON(todo);

};