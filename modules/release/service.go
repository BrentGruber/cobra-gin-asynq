package release

import (
	"database/sql"
	"log"
	db "releasr/modules/release/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
)

type ReleaseService struct {
	Router     *gin.Engine
	TaskClient asynq.Client
	Config     Config

	q *db.Queries
}

// Setup all routes
func (releaseService *ReleaseService) setupRouter() {
	router := gin.Default()
	router.GET("/release/hello/:name", releaseService.helloHandler)
	router.POST("/release/boundedcontext", releaseService.createBoundedContext)

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
	conn, err := sql.Open(config.DatabaseDriver, config.DatabaseUrl)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	queries := db.New(conn)
	releaseService := &ReleaseService{
		Config:     config,
		TaskClient: *client,

		q: queries,
	}
	releaseService.setupRouter()
	return releaseService
}
