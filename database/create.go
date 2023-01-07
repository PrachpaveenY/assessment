package database

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// e.POST All
func CreateExpensesAllHandler(c echo.Context) error {
	start := time.Now()
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't Bind :" + err.Error()})
	}

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id", &u.Title, &u.Amount, &u.Note, (pq.Array(&u.Tags)))
	err = row.Scan(&u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan :" + err.Error()})
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Run Time : ", time.Since(start), "sec")
	return c.JSON(http.StatusCreated, u)

}

// ====================

// use Thunder Client (PostgresSQL15)

// ***e.POST ID
// localhost:2565/expenses/:id (TC)
// Request Body
// ```json
