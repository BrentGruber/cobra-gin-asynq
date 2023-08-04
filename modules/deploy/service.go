package deploy

import (
	"log"
	"releasr/internal/tasks"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type DeployService struct {
	Router *gin.Engine
	Config Config
}

// Setup all routes
func (deployService *DeployService) setupRouter() {
	router := gin.Default()
	router.GET("/deploy/goodbye/:name", deployService.goodbyeHandler)

	deployService.Router = router
}

// Run the service
func (deployService *DeployService) Run() {
	deployService.Router.Run(":8002")
}

// Start the asynq handler
func (deployService *DeployService) startHandler(redisAddr string) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, HandleEmailDeliveryTask)
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

// NewdeployService creates a new DeployService struct
// and sets up routing
func NewDeployService(config Config, redisAddr string) *DeployService {

	deployService := &DeployService{
		Config: config,
	}
	deployService.setupRouter()
	go deployService.startHandler(redisAddr)
	return deployService
}
