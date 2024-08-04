package app

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gowesmart/api-gowesmart/controllers"
	"github.com/gowesmart/api-gowesmart/middlewares"
	"github.com/joho/godotenv"

	"github.com/gowesmart/api-gowesmart/docs"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewRouter() *gin.Engine {

	swaggerSchemes := []string{"https"}
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		utils.PanicIfError(err)
		swaggerSchemes = []string{"http"}
	}

	docs.SwaggerInfo.Title = "Car Review REST API"
	docs.SwaggerInfo.Description = "This is a Car Review REST API Docs."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.MustGetEnv("SERVER_HOST")
	docs.SwaggerInfo.Schemes = swaggerSchemes

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("no_space", func(fl validator.FieldLevel) bool {
			return !strings.Contains(fl.Field().String(), " ")
		})
		v.RegisterValidation("lowercase", func(fl validator.FieldLevel) bool {
			return fl.Field().String() == strings.ToLower(fl.Field().String())
		})
		v.RegisterValidation("uppercase", func(fl validator.FieldLevel) bool {
			return fl.Field().String() == strings.ToUpper(fl.Field().String())
		})
		v.RegisterValidation("url", func(fl validator.FieldLevel) bool {
			_, err := url.ParseRequestURI(fl.Field().String())
			return err == nil
		})
	}

	cfg := zap.Config{
		OutputPaths: []string{"stdout", "./log/log.log"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
			TimeKey:    "time_stamp",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
	}

	logger, err := cfg.Build()
	utils.PanicIfError(err)
	defer logger.Sync()

	db := NewConnection()

	userService := services.NewUserService()
	transactionService := services.NewTransactionService()

	// ======================== USER =======================

	userController := controllers.NewUserController(userService)
	transactionController := controllers.NewTransactionController(*transactionService)

	r := gin.Default()

	r.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
			AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization", "Pragma", "Cache-Control", "Expires"},
			MaxAge:           12 * time.Hour,
		},
	))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("logger", logger)
	})

	r.Use(exceptions.GlobalErrorHandler)

	r.NoRoute(func(c *gin.Context) {
		panic(exceptions.NewCustomError(http.StatusNotFound, fmt.Sprintf("path not found, use https://%s for API docs", utils.MustGetEnv("SERVER_HOST")+"/docs/index.html")))
	})

	apiRouter := r.Group("/api")

	// ======================== AUTH ROUTE =======================

	authRouter := apiRouter.Group("/auth")

	authRouter.POST("/register", userController.Register)
	authRouter.POST("/login", userController.Login)
	authRouter.POST("/forgot-password", userController.ForgotPassword)
	authRouter.POST("/reset-password", userController.ResetPassword)

	// ======================== USERS ROUTE =======================

	userRouter := apiRouter.Group("/users")

	userRouter.Use(middlewares.JwtAuthMiddleware)

	userRouter.GET("/current", userController.GetCurrentUser)

	// ======================== TRANSACTION ROUTE ======================
	transactionRouter := apiRouter.Group("/transactions")

	transactionRouter.GET("", transactionController.GetAll)
	transactionRouter.GET("/:id", transactionController.GetById)
	transactionRouter.POST("/:userId", transactionController.Create)
	transactionRouter.PATCH("/:id", transactionController.Update)
	transactionRouter.DELETE("/:id", transactionController.Delete)
	transactionRouter.PATCH("/payment/:id", transactionController.Pay)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	return r
}
