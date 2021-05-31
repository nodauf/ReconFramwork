package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nodauf/ReconFramwork/server/API/controllers"
)

func InitializeRoutes(router *gin.Engine) {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	//router.Use(setUserStatus())

	// Handle the index route
	//router.GET("/", showIndexPage)

	// Group user related routes together
	taskRoutes := router.Group("/task")
	{
		taskRoutes.GET("/list", controllers.ListTask)
		taskRoutes.GET("/view/:taskName", controllers.ViewTask)
	}
	workflowRoutes := router.Group("/workflow")
	{
		workflowRoutes.GET("/list", controllers.ListWorkflow)
		workflowRoutes.GET("/view/:workflowName", controllers.ViewWorkflow)
	}
	runRoutes := router.Group("/run")
	{
		runRoutes.POST("/task", controllers.RunTask)
		runRoutes.POST("/workflow", controllers.RunWorkflow)
	}

	// Group article related routes together
	/*	articleRoutes := router.Group("/article")
		{
			// Handle GET requests at /article/view/some_article_id
			articleRoutes.GET("/view/:article_id", getArticle)

			// Handle the GET requests at /article/create
			// Show the article creation page
			// Ensure that the user is logged in by using the middleware
			articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)

			// Handle POST requests at /article/create
			// Ensure that the user is logged in by using the middleware
			articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
		}*/
}
