package pkg

import (
	"log"
	"net/http"

	"my-app-server/server/auth"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct{
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var user UserDB
	log.Println("Registration endpoint hit")

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload!")
	}
	createdUser, err := h.srv.CreateUserdata(user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Failed to create User!")
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *Handler) LoginUser(c echo.Context) error {
    var creds UserDB

    log.Println("Login endpoint hit")
    if err := c.Bind(&creds); err != nil {
        log.Println("Error binding request payload:", err)
        return c.JSON(http.StatusBadRequest, "Invalid request payload!")
    }
	log.Printf("Received credentials: %+v", creds)
    user, err := h.srv.FindUserbyusername(creds.Username)
    if err != nil {
        if err.Error() == "user not found" {
            log.Println("User not found")
            return c.JSON(http.StatusUnauthorized, "Invalid credentials!")
        }
        log.Printf("Error finding user: %v", err)
        return c.JSON(http.StatusInternalServerError, "Failed to authenticate user")
    }

    log.Printf("Retrieved user: %+v", user)
    log.Printf("Provided password: %s", creds.Password)
    log.Printf("Stored hashed password: %s", user.Password)

    if !auth.CheckPassword(user.Password, creds.Password) {
        log.Println("Password mismatch")
        return c.JSON(http.StatusUnauthorized, "Invalid credentials!")
    }

    token, err := auth.CreateToken(user.Username)
    if err != nil {
        log.Printf("Error creating token: %v", err)
        return c.JSON(http.StatusInternalServerError, "Failed to generate token!")
    }

    return c.JSON(http.StatusOK, echo.Map{
        "token": token,
    })
}

func (h *Handler) CreateTasks(c echo.Context) error{
	var task TaskDB
	
	log.Printf("createTask endpoint hit")
	err:= c.Bind(&task)
	if err != nil{
		log.Printf("Error binding request payload: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid request payload!")
	}
	log.Printf("Received task: %+v", task)
	username := c.Get("username").(string)
	user, err := h.srv.FindUserbyusername(username)
	if err != nil{
		log.Printf("user not found")
		return c.JSON(http.StatusNotFound, "User not found!")
	}
	log.Printf("username: %s",username)
	task.UserId = user.ID
	log.Printf("User ID: %s", task.UserId.Hex())
	Createdtask, err := h.srv.CreateTaskdata(task.Title,task.Description,task.Due_date, task.UserId)
	if err != nil {
		log.Printf("Failed to create task")
		return c.JSON(http.StatusUnprocessableEntity, "Failed to create task!")
	}
	log.Printf("Task created")
	return c.JSON(http.StatusCreated, Createdtask)
}

func (h *Handler) FindTask(c echo.Context) error{
	username := c.Get("username").(string)
	user, err := h.srv.FindUser(username)
	log.Printf("Looking up user by username: %s", username)
	if err != nil{
		return c.JSON(http.StatusNotFound, "User not found!")
	}
	log.Printf("User found: %+v", user)
	task, err := h.srv.FindTaskbyUid(user.ID)
	if err != nil{
		log.Printf("Error finding tasks: %v", err)
		return c.JSON(http.StatusNotFound, "Task not found!")
	}
	log.Printf("Task found")
	return c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTask(c echo.Context) error{
	taskIDstr := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(taskIDstr)
	if err!= nil{
		return c.JSON(http.StatusBadRequest, "Invalid task Id format!")
	}
	err = h.srv.Deletetaskdata(taskID)
	if err != nil{
		if err == mongo.ErrNoDocuments{
		return c.JSON(http.StatusNotFound, "Task not found")
		}
		return c.JSON(http.StatusInternalServerError, "Failed to delete task!")
	}
	return nil
}