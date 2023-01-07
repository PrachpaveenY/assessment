package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/lib/pq"
)

// e.GET All
func GetExpensesHandler(c echo.Context) error {
	start := time.Now()
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't Prepare :" + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't Query :" + err.Error()})
	}

	users := []User{}
	for rows.Next() {
		u := User{}
		err := rows.Scan(
			&u.ID,
			&u.Title,
			&u.Amount,
			&u.Note,
			(pq.Array(&u.Tags)),
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan :" + err.Error()})
		}
		users = append(users, u)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Run Time : ", time.Since(start), "sec")
	return c.JSON(http.StatusOK, users)
}

// e.GET ID edit
func GetExpensesIDHandler(c echo.Context) error {
	start := time.Now()
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare :" + err.Error()})
	}

	row := stmt.QueryRow(id)
	u := User{}
	err = row.Scan(
		&u.ID,
		&u.Title,
		&u.Amount,
		&u.Note,
		(pq.Array(&u.Tags)),
	)

	time.Sleep(1 * time.Second)
	fmt.Println("Run Time : ", time.Since(start), "sec")
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, u)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan expenses :" + err.Error()})
	}
}

// ====================

// use Thunder Client (PostgresSQL15)

// ***e.GET All
// EXP04
// GET /expenses
// Request Body
// ```json
// nil
// ```

// ***e.GET ID
// use = localhost:2565/expenses/1 (TC) (ใส่เลขที่x "/x" เพื่อกำหนด id row ที่ต้องการ SELECT)
// EXP02
// GET /expenses/:id
// :id = 1
// Response Body
// ```json
