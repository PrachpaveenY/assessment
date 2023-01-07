package database

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// e.PUT
func UpdateAllExpensesHandler(c echo.Context) error {
	start := time.Now()
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't Bind :" + err.Error()})
	}

	stmt, err := db.Prepare("UPDATE expenses SET title = $2, amount = $3, note = $4, tags = $5 WHERE id = $1;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't Prepare :" + err.Error()})
	}

	users, err := stmt.Exec(&u.Title, &u.Amount, &u.Note, pq.Array(&u.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan :" + err.Error()})
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Run Time : ", time.Since(start), "sec")
	return c.JSON(http.StatusCreated, users)
}

// e.PATCH
func UpdateExpensesHandler(c echo.Context) error {
	start := time.Now()
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't Bind :" + err.Error()})
	}

	stmt, err := db.Prepare("UPDATE expenses SET amount=$2 WHERE id=$1;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't Prepare :" + err.Error()})
	}

	users, err := stmt.Exec(&u.ID, &u.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't Query :" + err.Error()})
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Run Time : ", time.Since(start), "sec")
	return c.JSON(http.StatusCreated, users)
}

// ====================

// ***e.PUT
// EXP03
// PUT /expenses/:id
// :id = 1
// Request Body
// ```json
// {
//     "id": = 1,
//     "title": "apple smoothie",
//     "amount": 89,
//     "note": "no discount",
//     "tags": ["beverage"]
// }
// ```

// ***e.PATCH
// PATCH /expenses/:id
// Request Body
// ```json
// {
//     "id": = 1,
//     "amount": 11,
// }
// ```
