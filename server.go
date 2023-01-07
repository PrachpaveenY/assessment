package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/PrachpaveenY/assessment/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Welcome to Server")

	database.InitDB()
	// h := handler.NewApplication(db)

	e := echo.New()
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "postgres" || password == "2565" {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/expenses", database.CreateExpensesAllHandler)
	e.GET("/expenses", database.GetExpensesHandler)
	e.GET("/expenses/:id", database.GetExpensesIDHandler)
	e.PUT("/expenses/:id", database.UpdateAllExpensesHandler)
	e.PATCH("/expenses/:id", database.UpdateExpensesHandler)
	e.DELETE("/expenses/:id", database.DeleteExpensesHandler)

	// serverPort := os.Getenv("PORT")
	// e.Logger.Fatal(e.Start(serverPort))
	log.Fatal(e.Start(":2565"))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// ====================
// use Thunder Client (PostgresSQL15)
// PORT = 2565
// localhost:2565/expenses
// localhost:2565/expenses/:id

// ====================

// fmt.Println("Please use server.go for main file")
// fmt.Println("start at port:", os.Getenv("PORT"))
//
