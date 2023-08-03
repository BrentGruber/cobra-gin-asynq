package release

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (releaseService *ReleaseService) helloHandler(ctx *gin.Context) {
	fmt.Println("Hello from release service")

	name := ctx.Param("name")
	message := fmt.Sprintf("Hello %s", name)

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
