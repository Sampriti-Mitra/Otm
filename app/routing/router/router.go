package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tylerb/graceful"
	"os"
	"os/signal"
	"otm/app/config"
	"otm/app/controllers"
	"otm/app/routing/middleware"
	"syscall"
	"time"
)

const (
	GracefulTimeoutDuration = 120
)

func Initialize() {
	router := gin.New()
	//prometheusRouter := gin.New()
	router.Use(middleware.Recovery)
	//prometheusRouter.Use(middleware.Recovery)

	registerDefaultMiddlewares(router)

	//prometheusRouterGroup := prometheusRouter.Group("/")
	rootGroup := router.Group("/")
	v1Group := router.Group("/v1")

	//	registerMetricRoutes(prometheusRouterGroup)
	registerPublicRoutes(rootGroup)
	registerPrivateRoutes(v1Group)

	listenAddress := getListenAddress()
	//prometheusListenAddress := getPrometheusListenAddress()

	application := config.GetConfig().Application

	traceData := map[string]interface{}{
		"name":             "otm",
		"application_ip":   application.ListenIP,
		"application_port": application.ListenPort,
	}
	fmt.Println(traceData)
	initializeSignals()
	//go graceful.Run(prometheusListenAddress, GracefulTimeoutDuration*time.Second, prometheusRouter)
	graceful.Run(listenAddress, GracefulTimeoutDuration*time.Second, router)
}

func registerMetricRoutes(router *gin.RouterGroup) {
	//router.GET("metrics", prometheusHandler())
}

//func prometheusHandler() gin.HandlerFunc {
//	h := promhttp.Handler()
//
//	return func(c *gin.Context) {
//		h.ServeHTTP(c.Writer, c.Request)
//	}
//}

func registerDefaultMiddlewares(router *gin.Engine) {
	router.Use(middleware.SetRequestPayload())
}

func getListenAddress() string {
	application := config.GetConfig().Application
	return fmt.Sprintf("%s:%d", application.ListenIP, application.ListenPort)
}

//func getPrometheusListenAddress() string {
//	application := config.GetConfig().Application
//	prometheus := config.GetConfig().Prometheus
//	return fmt.Sprintf("%s:%d", application.ListenIP, prometheus.ListenPort)
//}

func initializeSignals() {
	controllers.AppRunningStatus = true
	c := make(chan os.Signal)

	signal.Notify(c, syscall.SIGUSR1)

	go func() {
		for range c {
			controllers.AppRunningStatus = false
		}
	}()
}

func registerPrivateRoutes(router *gin.RouterGroup) {

}

func registerPublicRoutes(router *gin.RouterGroup) {
	router.GET("", controllers.App.Welcome)
	router.GET("status", controllers.Status.Status)
	router.GET("ping", controllers.Status.Ping)

	router.POST("register", controllers.LoginController.CreateRegister)
	router.GET("register/:register_id", controllers.LoginController.GetRegister)
	router.PUT("register/:register_id", controllers.LoginController.UpdateRegister)
	router.DELETE("register/:register_id", controllers.LoginController.DeleteRegister)
	router.GET("login", controllers.LoginController.GetLogin)

	router.POST("post", controllers.ProfileController.CreatePost)
	router.GET("profile/:profile_id", controllers.ProfileController.ListPost)
	router.GET("post/:video_id", controllers.ProfileController.GetPost)
	router.PUT("post/:video_id", controllers.ProfileController.UpdatePost)
	router.DELETE("post/:video_id", controllers.ProfileController.DeletePost)
	router.PUT("post/:video_id/applause", controllers.ProfileController.ApplausePost)

	router.POST("about", controllers.AboutController.CreateAbout)
	router.GET("about/:user_id", controllers.AboutController.GetAbout)
	router.PUT("about/:user_id", controllers.AboutController.UpdateAbout)

	router.POST("profile/:profile_id/follow", controllers.FollowerController.Follow)
	router.GET("profile/:profile_id/followers", controllers.FollowerController.ListFollowers)
	router.GET("profile/:profile_id/follower_request", controllers.FollowerController.ListFollowerRequest)
	router.GET("profile/:profile_id/following", controllers.FollowerController.ListFollowing)
	router.GET("profile/:profile_id/followers/:request_by", controllers.FollowerController.GetFollower)
	router.PUT("profile/:profile_id/followers/:request_by/accept", controllers.FollowerController.AcceptFollowers)
	router.PUT("profile/:profile_id/followers/:request_by/reject", controllers.FollowerController.RejectFollowers)

	router.GET("profile/:profile_id/feed", controllers.FeedController.Feed)
	router.GET("profile/:profile_id/trending", controllers.FeedController.FeedTrending)
}
