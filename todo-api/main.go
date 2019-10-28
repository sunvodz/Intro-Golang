package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Task is ...
type Task struct {
	ID     *int8  `json:"_id"`
	Topic  string `json:"topic"`
	Status bool   `json:"status"`
}

// ResponseError is ...
type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func intToPointer(i int8) *int8 {
	return &i
}

var tasks []Task

func main() {
	
	tasks = []Task{
		{
			ID:    intToPointer(1),
			Topic: "Name1",
		},
		{
			ID:    intToPointer(2),
			Topic: "Name2",
		},
		{
			ID:    intToPointer(3),
			Topic: "Name3",
		},
	}

	e := echo.New()
	e.GET("/", Index)

	e.POST("/tasks", AddTask)
	e.GET("/tasks/:task_id", ShowOneTask)
	e.GET("/tasks", ShowTask)
	e.PUT("/tasks/:task_id", UpdateTask)
	e.DELETE("/tasks/:task_id", DeleteTask)


	
	e.Logger.Fatal(e.Start(":1323"))
}

// Index ...
func Index(c echo.Context) error {

	return c.JSON(http.StatusOK, "Hello, World!")
}

// AddTask ...
func AddTask(c echo.Context) error {
	/*
		id := c.FormValue("task_id")
		name := c.FormValue("name")
		return c.String(http.StatusOK, " id:"+id+", name:"+name)
	*/

	var body Task

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if body.ID == nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: "id empty"})
	}

	task := Task{
		ID:    body.ID,
		Topic: body.Topic,
	}

	tasks = append(tasks, task)

	return c.JSON(http.StatusCreated, task)
}

// ShowOneTask ...
func ShowOneTask(c echo.Context) error {

	tid, _ := strconv.ParseInt(c.Param("task_id"), 10, 32)

	var index int

	for i, v := range tasks {
		if *v.ID == int8(tid) {
			index = i
			break
		}
	}

	return c.JSON(http.StatusOK, tasks[index])
}

// ShowTask ...
func ShowTask(c echo.Context) error {

	return c.JSON(http.StatusOK, tasks)
}

// UpdateTask ...
func UpdateTask(c echo.Context) error {

	tid, _ := strconv.ParseInt(c.Param("task_id"), 10, 32)

	var index int

	for i, v := range tasks {
		if *v.ID == int8(tid) {
			index = i
			break
		}
	}

	var body Task

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: err.Error()})
	}

	task := Task{
		ID:     body.ID,
		Topic:  body.Topic,
		Status: body.Status,
	}

	tasks[index].Topic = task.Topic
	tasks[index].Status = task.Status
	return c.JSON(http.StatusOK, tasks[index])
}

//DeleteTask ...
func DeleteTask(c echo.Context) error {

	tid, _ := strconv.ParseInt(c.Param("task_id"), 10, 8)
	var index int

	for i, v := range tasks {
		if *v.ID == int8(tid) {
			index = i
			break
		}
	}

	// Remove the element at index i from a.
	copy(tasks[index:], tasks[index+1:]) // Shift a[i+1:] left one index.
	tasks[len(tasks)-1] = Task{}         // Erase last element (write zero value).
	tasks = tasks[:len(tasks)-1]         // Truncate slice.

	fmt.Println(tasks) // [A B D E]

	return c.NoContent(http.StatusNoContent)
}
