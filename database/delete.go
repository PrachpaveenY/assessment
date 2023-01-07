package database

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

// e.DELETE
func DeleteExpensesHandler(c echo.Context) error {
	start := time.Now()
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't Bind :" + err.Error()})
	}

	row := db.QueryRow("DELETE FROM expenses WHERE id = $1", &u.ID)
	err = row.Scan(&u.ID)

	time.Sleep(1 * time.Second)
	fmt.Println("Run Time : ", time.Since(start), "sec")
	return c.JSON(http.StatusOK, u)
}

// ====================

// ***e.DELETE
// DELETE /expenses/:id
// Request Body
// ```json
// {
// 	"id": 1
// }
// ```
