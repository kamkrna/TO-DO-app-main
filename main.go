package main

import (
	"context"
	"log"
	"my-app-server/server/middleware"
	"my-app-server/server/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
  
  func main() {
	mongoUri := "mongodb+srv://turabali57972:3zkPgdR8zGVMQddA@cluster0.sbl2ptv.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
	  log.Fatal("Error connecting to MongoDB!")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}
	srv := pkg.NewService(client)
	hdl := pkg.NewHandler(srv)
	e := echo.New()

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOriginFunc: func(origin string) (bool, error) {return true, nil},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	e.POST("/registration", hdl.CreateUser)
	e.POST("/", hdl.LoginUser)

	e.POST("/todo", hdl.CreateTasks, middleware.ValidateToken)
	e.GET("/todo", hdl.FindTask, middleware.ValidateToken)
	e.DELETE("/todo/:id",hdl.DeleteTask,middleware.ValidateToken)


	e.Logger.Fatal(e.Start(":8000"))
  
  }
  