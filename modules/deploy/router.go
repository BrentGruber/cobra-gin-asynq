package deploy

import "github.com/gin-gonic/gin"

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

// NewdeployService creates a new DeployService struct
// and sets up routing
func NewDeployService(config Config) *DeployService {
	deployService := &DeployService{
		Config: config,
	}
	deployService.setupRouter()
	return deployService
}
