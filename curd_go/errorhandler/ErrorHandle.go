package errorhandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type NewError struct {
	Title      string      `json:"title"`
	Details    interface{} `json:"error_details"`
	StatusCode int         `json:"status_code"`
}

func (e *NewError) Error() gin.H {
	return gin.H{"error": fmt.Sprintf("%v", e)}
}

func GetErorr(e interface{}, status_code int, title string) NewError {
	// fmt.Println(e)
	return NewError{
		Title:      title,
		Details:    e,
		StatusCode: status_code,
	}

}
