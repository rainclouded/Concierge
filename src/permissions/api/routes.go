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

// ContextOptions defines the available configuration options for the API context.
// Args:
//     None
// Returns:
//     A new ContextOptions instance with default values
type ContextOptions struct {
	DB            database.Database           // Database configuration
	AccountClient client.AccountClient        // Account client for API interactions
	JWTContext    *middleware.JWT_Context     // JWT context for authentication
	ginMode       string                      // Gin mode (e.g., Debug, Release)
}

// Option is a function that modifies the ContextOptions.
// Args:
//     options - The ContextOptions instance to modify
// Returns:
//     None
type Option func(*ContextOptions)

// DefaultOptions initializes default options for the context.
// Args:
//     None
// Returns:
//     *ContextOptions - Returns a new instance of ContextOptions with default settings
func DefaultOptions() *ContextOptions {
	return &ContextOptions{

		DB:            config.LoadDB(),
		AccountClient: config.LoadAccountEndpoint(),
		// AccountClient: client.NewMockAccountClient(),
		JWTContext: middleware.NewJWT(),
		ginMode:    gin.DebugMode,

	}
}

// WithDB is a function to set a custom database instance in the context options.
// Args:
//     db - The database instance to set
// Returns:
//     Option - A function that modifies the ContextOptions to use the provided database instance
func WithDB(db database.Database) Option {
	return func(options *ContextOptions) {
		options.DB = db
	}
}

// WithAccountClient allows setting a custom account client in the context options.
// Args:
//     accCli - The custom account client to use
// Returns:
//     Option - A function that modifies the ContextOptions to use the provided account client
func WithAccountClient(accCli client.AccountClient) Option {
	return func(options *ContextOptions) {
		options.AccountClient = accCli
	}
}

// WithJWTContext allows setting a custom JWT context in the context options.
// Args:
//     jwtCtx - The custom JWT context to use
// Returns:
//     Option - A function that modifies the ContextOptions to use the provided JWT context
func WithJWTContext(jwtCtx *middleware.JWT_Context) Option {
	return func(options *ContextOptions) {
		options.JWTContext = jwtCtx
	}
}

// WithGinMode sets a custom Gin mode (e.g., Release, Debug) in the context options.
// Args:
//     mode - The Gin mode to use
// Returns:
//     Option - A function that modifies the ContextOptions to use the provided Gin mode
func WithGinMode(mode string) Option {
	return func(options *ContextOptions) {
		options.ginMode = mode
	}
}

// NewRouter creates a new Gin router with the provided options.
// Args:
//     options - A variadic list of functions that modify the ContextOptions
// Returns:
//     *gin.Engine - Returns a new Gin router configured with the provided options
func NewRouter(options ...Option) *gin.Engine {
	// Set default options and apply any overrides from the provided options
	ctxOptions := DefaultOptions()
	for _, opt := range options {
		opt(ctxOptions)
	}

	// Set the Gin mode and create a new router
	gin.SetMode(ctxOptions.ginMode)
	router := gin.Default()

	// Set up middleware for the router
	router.Use(middleware.SetDb(ctxOptions.DB))           // Middleware for setting the DB context
	router.Use(middleware.SetAccountClient(ctxOptions.AccountClient)) // Middleware for setting the account client context
	router.Use(middleware.SetJWTContex(ctxOptions.JWTContext))  // Middleware for setting the JWT context
	router.Use(middleware.EnableCORS())                      // Middleware for enabling CORS

	// Define routes and associate them with handler functions
	router.POST("/sessions", handlers.PostSessionKey)       // Route for creating a session key
	router.GET("/sessions/me", handlers.ParseSessionKey)    // Route for retrieving session details for the current session
	router.GET("/sessions/public-key", handlers.GetPublicKey) // Route for retrieving the public key for session validation

	router.GET("/permissions/healthcheck", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) }) // Health check route for the permissions service
	router.GET("/permissions", handlers.GetPermissions)          // Route for getting all permissions
	router.GET("/permissions/:id", handlers.GetPermissionById)   // Route for getting a specific permission by its ID
	router.POST("/permissions", handlers.PostPermission)         // Route for creating a new permission

	router.GET("/permission-groups", handlers.GetPermissionGroups)   // Route for getting all permission groups
	router.GET("/permission-groups/:id", handlers.GetPermissionGroupById) // Route for getting a specific permission group by ID
	router.POST("/permission-groups", handlers.PostPermissionGroups) // Route for creating a new permission group
	router.PATCH("/permission-groups/:id", handlers.PatchPermissionGroups) // Route for updating a permission group by ID

	// Return the configured router
	return router
}
