package api

import (
	"concierge/permissions/internal/client"
	"concierge/permissions/internal/config"
	"concierge/permissions/internal/database"
	"concierge/permissions/internal/handlers"
	"concierge/permissions/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContextOptions struct {
	DB            database.Database
	AccountClient client.AccountClient
	JWTContext    *middleware.JWT_Context
	ginMode       string
}

type Option func(*ContextOptions)

func DefaultOptions() *ContextOptions {
	return &ContextOptions{
		DB:            config.LoadDB(),
		AccountClient: config.LoadAccountEndpoint(),
		JWTContext:    middleware.NewJWT(),
		ginMode:       gin.ReleaseMode,
	}
}

func WithDB(db database.Database) Option {
	return func(options *ContextOptions) {
		options.DB = db
	}
}

func WithAccountClient(accCli client.AccountClient) Option {
	return func(options *ContextOptions) {
		options.AccountClient = accCli
	}
}

func WithJWTContext(jwtCtx *middleware.JWT_Context) Option {
	return func(options *ContextOptions) {
		options.JWTContext = jwtCtx
	}
}

func WithGinMode(mode string) Option {
	return func(options *ContextOptions) {
		options.ginMode = mode
	}
}

func NewRouter(options ...Option) *gin.Engine {
	ctxOptions := DefaultOptions()
	for _, opt := range options {
		opt(ctxOptions)
	}

	gin.SetMode(ctxOptions.ginMode)
	router := gin.Default()

	router.Use(middleware.SetDb(ctxOptions.DB))
	router.Use(middleware.SetAccountClient(ctxOptions.AccountClient))
	router.Use(middleware.SetJWTContex(ctxOptions.JWTContext))
	router.Use(middleware.EnableCORS())

	router.POST("/sessions", handlers.PostSessionKey)
	router.GET("/sessions/me", handlers.ParseSessionKey)
	router.GET("/sessions/public-key", handlers.GetPublicKey)

	router.GET("/permissions/healthcheck", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	router.GET("/permissions", handlers.GetPermissions)
	router.GET("/permissions/:id", handlers.GetPermissionById)
	router.POST("/permissions", handlers.PostPermission)

	router.GET("/permission-groups", handlers.GetPermissionGroups)
	router.GET("/permission-groups/:id", handlers.GetPermissionGroupById)
	router.POST("/permission-groups", handlers.PostPermissionGroups)
	router.PATCH("/permission-groups/:id", handlers.PatchPermissionGroups)

	return router
}
