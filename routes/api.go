package routes

import (
	"fmt"
	"myadmin/app/controller"
	"myadmin/app/repository"
	"myadmin/app/service"
	"myadmin/config"
	"myadmin/security"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	docs "myadmin/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApiRouter(db config.Database) {
	// Swagger docs init
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "localhost:7000"

	// Repository Asset
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	privilegeRepo := repository.NewPrivilegeRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	roleMenuRepo := repository.NewRoleMenuRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	rolePermissionRepo := repository.NewRolePermissionRepository(db)

	// Service Asset
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	roleService := service.NewRoleService(roleRepo)
	menuService := service.NewMenuService(menuRepo)
	privilegeService := service.NewPrivilegeService(privilegeRepo)
	userRoleService := service.NewUserRoleService(userRoleRepo)
	roleMenuService := service.NewRoleMenuService(roleMenuRepo)
	permissionService := service.NewPermissionService(permissionRepo)
	rolePermissionService := service.NewRolePermissionService(rolePermissionRepo)

	//Controller Asset
	userController := controller.NewUserController(userService, userRoleService)
	authController := controller.NewAuthController(userService, authService)
	roleController := controller.NewRoleController(roleService, roleMenuService, rolePermissionService)
	menuController := controller.NewMenuController(menuService)
	privilegeController := controller.NewPrivilegeController(privilegeService)
	permController := controller.NewPermissionController(permissionService)

	// Route
	httpRouter := gin.Default()

	// Register routing
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	v1 := httpRouter.Group("/api/v1")

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Testing  connection
	v1.GET("/status-check", func(ctx *gin.Context) {
		now := time.Now()

		message := fmt.Sprintf("Service ✅ API Up and Running %s %d, %d", now.Month(), now.Day(), now.Year())

		ctx.JSON(http.StatusOK, gin.H{"data": message})
	})

	noAuth := v1 // Grouping routes

	// Auth routes
	noAuth.POST("/auth/register", authController.Register)
	noAuth.POST("/auth/login", authController.Login)

	auth := v1.Use(security.Middleware()) // Grouping routes

	// Users routes
	auth.GET("/users", userController.Index)
	auth.GET("/users/:id", userController.Show)
	auth.PUT("/users/:id", userController.Update)
	auth.DELETE("/users/:id", userController.Delete)

	// User role routes
	auth.POST("/users/roles/assign", userController.AttachRole)
	auth.DELETE("/users/roles/unassign", userController.DetachRole)

	// Roles routes
	auth.POST("/roles", roleController.Store)
	auth.GET("/roles", roleController.Index)
	auth.GET("/roles/:id", roleController.Show)
	auth.PUT("/roles/:id", roleController.Update)
	auth.DELETE("/roles/:id", roleController.Delete)

	// Role menu routes
	auth.POST("/roles/menus", roleController.AttachMenu)
	auth.DELETE("/roles/menus/:id", roleController.DetachMenu)

	// Roles routes
	auth.GET("/permissions", permController.Index)
	auth.POST("/permissions", permController.Store)
	auth.GET("/permissions/:id", permController.Show)
	auth.PUT("/permissions/:id", permController.Update)
	auth.DELETE("/permissions/:id", permController.Delete)

	// Role permission routes
	auth.POST("/roles/permissions", roleController.AttachPermission)
	auth.DELETE("/roles/permissions/:id", roleController.DetachPermission)

	// Menus routes
	auth.POST("/menus", menuController.Store)
	auth.GET("/menus", menuController.Index)
	auth.GET("/menus/:id", menuController.Show)
	auth.PUT("/menus/:id", menuController.Update)
	auth.DELETE("/menus/:id", menuController.Delete)

	// Privilege routes
	auth.GET("/privileges", privilegeController.Index)
	auth.GET("/privileges/:id", privilegeController.Show)

	httpRouter.Run(":" + os.Getenv("APP_PORT")) // Routes
}
