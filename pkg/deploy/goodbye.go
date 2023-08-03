package deploy

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (deployService *DeployService) goodbyeHandler(ctx *gin.Context) {
	fmt.Println("Hello from deploy service")

	name := ctx.Param("name")
	message := fmt.Sprintf("Goodbye %s", name)

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
