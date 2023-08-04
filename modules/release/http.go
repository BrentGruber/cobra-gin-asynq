package release

import (
	"fmt"
	"log"
	"net/http"
	"releasr/internal/tasks"

	"github.com/gin-gonic/gin"
)

func (releaseService *ReleaseService) helloHandler(ctx *gin.Context) {
	fmt.Println("Hello from release service")

	name := ctx.Param("name")
	message := fmt.Sprintf("Hello %s", name)

	task, err := tasks.NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := releaseService.TaskClient.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"enabled": releaseService.Config.Enabled,
	})
}
