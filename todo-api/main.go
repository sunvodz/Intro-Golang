package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"

	"time"

	"github.com/dgrijalva/jwt-go"

)

// Task is ...
type Task struct {
	ID         *int8  `json:"_id"`
	Topic      string `json:"topic"`
	Created    string `json:"created"`
	FinishDate string `json:"finish_date"`
	Status     bool   `json:"status"`
	File       string `json:"file"`
}

// User is ...
type User struct {
	ID       *int8  `json:"_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
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
var users []User

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

	users = []User{
		{
			ID:       intToPointer(1),
			Name:     "Mr.SunvoDz",
			Username: "SunvoDz",
			Password: "SunvoDz",
		},
		{
			ID:       intToPointer(2),
			Name:     "Mr.Johnny",
			Username: "Johnny",
			Password: "Johnny",
		},
		{
			ID:       intToPointer(3),
			Name:     "Mr.Thomas",
			Username: "Thomas",
			Password: "Thomas",
		},
	}

	e := echo.New()
	e.GET("/", Index)

	e.POST("/tasks", AddTask)
	e.GET("/tasks/:task_id", ShowOneTask)
	e.GET("/tasks", ShowTask)
	e.PUT("/tasks/:task_id", UpdateTask)
	e.DELETE("/tasks/:task_id", DeleteTask)

	e.POST("/users", AddUser)
	e.GET("/users/:user_id", ShowOneUser)
	e.GET("/users", ShowUser)
	e.PUT("/users/:user_id", UpdateUser)
	e.DELETE("/users/:user_id", DeleteUser)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	e.GET("/accessible", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)


	e.Logger.Fatal(e.Start(":1323"))
}

// Index ...
func Index(c echo.Context) error {

	return c.JSON(http.StatusOK, "Hello, World!")
}

//Task ...

// AddTask ...
func AddTask(c echo.Context) error {
	var body Task

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if body.ID == nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: "id empty"})
	}

	task := Task{
		ID:         body.ID,
		Topic:      body.Topic,
		Created:    body.Created,
		FinishDate: body.FinishDate,
		File:       body.File,
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
		ID:         body.ID,
		Topic:      body.Topic,
		FinishDate: body.FinishDate,
		Status:     body.Status,
	}

	tasks[index].Topic = task.Topic
	tasks[index].Status = task.Status
	tasks[index].FinishDate = task.FinishDate
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

// USER ...

// AddUser ...
func AddUser(c echo.Context) error {

	var body User

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if body.ID == nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: "id empty"})
	}

	user := User{
		ID:       body.ID,
		Name:     body.Name,
		Username: body.Username,
		Password: body.Password,
	}

	users = append(users, user)

	return c.JSON(http.StatusCreated, user)
}

// ShowOneUser ...
func ShowOneUser(c echo.Context) error {

	uid, _ := strconv.ParseInt(c.Param("user_id"), 10, 32)

	var index int

	for i, v := range users {
		if *v.ID == int8(uid) {
			index = i
			break
		}
	}

	return c.JSON(http.StatusOK, users[index])
}

// ShowUser ...
func ShowUser(c echo.Context) error {

	return c.JSON(http.StatusOK, users)
}

// UpdateUser ...
func UpdateUser(c echo.Context) error {

	uid, _ := strconv.ParseInt(c.Param("user_id"), 10, 32)

	var index int

	for i, v := range tasks {
		if *v.ID == int8(uid) {
			index = i
			break
		}
	}

	var body User

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := User{
		ID:       body.ID,
		Name:     body.Name,
		Username: body.Username,
		Password: body.Password,
	}

	users[index].Name = user.Name
	users[index].Username = user.Username
	users[index].Password = user.Password
	return c.JSON(http.StatusOK, users[index])
}

//DeleteUser ...
func DeleteUser(c echo.Context) error {

	uid, _ := strconv.ParseInt(c.Param("user_id"), 10, 32)
	var index int

	for i, v := range tasks {
		if *v.ID == int8(uid) {
			index = i
			break
		}
	}

	// Remove the element at index i from a.
	copy(users[index:], users[index+1:]) // Shift a[i+1:] left one index.
	users[len(users)-1] = User{}         // Erase last element (write zero value).
	users = users[:len(users)-1]         // Truncate slice.

	fmt.Println(users) // [A B D E]

	return c.NoContent(http.StatusNoContent)
}

func login(c echo.Context) error {

	var body User

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Status: http.StatusBadRequest, Message: err.Error()})
	}

	fmt.Println("Username :",body.Username);
	fmt.Println("Password :",body.Password);
	if body.Username != "SunvoDz" || body.Password != "SunvoDz" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}