package release

import (
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type ReleaseService struct {
	Router     *gin.Engine
	TaskClient asynq.Client
	Config     Config
}

// Setup all routes
func (releaseService *ReleaseService) setupRouter() {
	router := gin.Default()
	router.GET("/release/hello/:name", releaseService.helloHandler)

	releaseService.Router = router
}

// Run the service
func (releaseService *ReleaseService) Run() {
	releaseService.Router.Run(":8001")
}

// NewreleaseService creates a new ReleaseService struct
// and sets up routing
func NewReleaseService(config Config, redisAddr string) *ReleaseService {
	// TODO: is there a better way to pass global configs
	// possibly give the service the whole config object
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	releaseService := &ReleaseService{
		Config:     config,
		TaskClient: *client,
	}
	releaseService.setupRouter()
	return releaseService
}
