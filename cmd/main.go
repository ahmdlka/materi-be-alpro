package main

import (
	"log"

	"github.com/Mobilizes/materi-be-alpro/config"
	"github.com/Mobilizes/materi-be-alpro/middlewares"
	"github.com/Mobilizes/materi-be-alpro/modules/user"
	userController "github.com/Mobilizes/materi-be-alpro/modules/user/controller"
	userRepository "github.com/Mobilizes/materi-be-alpro/modules/user/repository"
	userService "github.com/Mobilizes/materi-be-alpro/modules/user/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Mobilizes/materi-be-alpro/modules/auth"
	authController "github.com/Mobilizes/materi-be-alpro/modules/auth/controller"
	authService "github.com/Mobilizes/materi-be-alpro/modules/auth/service"

	"github.com/Mobilizes/materi-be-alpro/modules/meal_plan"
	mpController "github.com/Mobilizes/materi-be-alpro/modules/meal_plan/controller"
	mpRepository "github.com/Mobilizes/materi-be-alpro/modules/meal_plan/repository"
	mpService "github.com/Mobilizes/materi-be-alpro/modules/meal_plan/service"

	"github.com/Mobilizes/materi-be-alpro/modules/chat"
	chatController "github.com/Mobilizes/materi-be-alpro/modules/chat/controller"
	chatRepository "github.com/Mobilizes/materi-be-alpro/modules/chat/repository"
	chatService "github.com/Mobilizes/materi-be-alpro/modules/chat/service"

	"github.com/Mobilizes/materi-be-alpro/pkg/ragclient"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}

	// Connect to database
	db := config.SetupDatabase()

	// Initialize Gin app
	r := gin.Default()

	// Setup CORS
	r.Use(middlewares.CORS())

	// Setup routes group
	api := r.Group("/api")

	// Setup Clients
	ragClient := ragclient.NewRAGClient()

	// Setup Repositories
	userRepo := userRepository.NewUserRepository(db)
	mpRepo := mpRepository.NewMealPlanRepository(db)
	cRepo := chatRepository.NewChatRepository(db)

	// Setup Services
	jwtSvc := authService.NewJWTService()
	uService := userService.NewUserService(userRepo)
	aService := authService.NewAuthService(userRepo, uService, jwtSvc)
	mService := mpService.NewMealPlanService(mpRepo, userRepo, ragClient)

	iService := chatService.NewIntentService()
	cService := chatService.NewChatService(cRepo, iService, mpRepo, userRepo, ragClient)

	// Setup Controllers
	userCtrl := userController.NewUserController(uService)
	authCtrl := authController.NewAuthController(aService)
	mpCtrl := mpController.NewMealPlanController(mService)
	cCtrl := chatController.NewChatController(cService)

	// Register Routes
	auth.RegisterAuthRoutes(api, authCtrl)
	user.RegisterUserRoutes(api, userCtrl, jwtSvc)
	meal_plan.RegisterMealPlanRoutes(api, mpCtrl, jwtSvc)
	chat.RegisterChatRoutes(api, cCtrl, jwtSvc)

	// Start App
	r.Run(":8080")
}
