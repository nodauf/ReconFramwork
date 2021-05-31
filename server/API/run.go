package webAPI

import (
	"github.com/gin-gonic/gin"
	"github.com/nodauf/ReconFramwork/server/API/controllers"
	"github.com/nodauf/ReconFramwork/server/API/routers"
	"github.com/nodauf/ReconFramwork/server/server/orchestrator"
)

func Run(options orchestrator.Options) {
	defer options.Wg.Done()
	controllers.OptionsOrchestrator = options

	r := gin.Default()
	routers.InitializeRoutes(r)
	r.Run(":1234")
}
