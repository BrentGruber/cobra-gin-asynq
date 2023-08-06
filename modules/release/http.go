package release

import (
	"fmt"
	"log"
	"net/http"
	"releasr/internal/tasks"
	db "releasr/modules/release/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createBCRequest struct {
	name string
}

func (releaseService *ReleaseService) createBoundedContext(ctx *gin.Context) {
	fmt.Println("Creating bounded context")
	var req createBCRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	arg := db.Createbounded_contextsParams{
		Name:    req.name,
		Manager: "John Doe",
	}

	bc, err := releaseService.q.Createbounded_contexts(ctx, arg)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	ctx.JSON(http.StatusOK, bc)
}

func (releaseService *ReleaseService) helloHandler(ctx *gin.Context) {
	fmt.Println("Hello from release service")

	name := ctx.Param("name")
	message := fmt.Sprintf("Hello %s", name)

	task, err := tasks.NewEmailDeliveryTask(name, "some:template:id")
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
