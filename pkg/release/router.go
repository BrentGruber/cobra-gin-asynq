package release

import "github.com/gin-gonic/gin"

type ReleaseService struct {
	Router *gin.Engine
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
func NewReleaseService() *ReleaseService {
	releaseService := &ReleaseService{}
	releaseService.setupRouter()
	return releaseService
}
