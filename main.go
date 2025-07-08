package main

import (
	_ "ThreeLayeredArchitecture/docs"
	taskhandler "ThreeLayeredArchitecture/handler/task"
	"ThreeLayeredArchitecture/migrations"
	taskservice "ThreeLayeredArchitecture/services/task"
	taskstore "ThreeLayeredArchitecture/store/task"
	httpSwagger "github.com/swaggo/http-swagger"
	"gofr.dev/pkg/gofr"
	"log"
	"net/http"

	// User.
	userhandler "ThreeLayeredArchitecture/handler/user"
	userservice "ThreeLayeredArchitecture/services/user"
	userstore "ThreeLayeredArchitecture/store/user"
)

// @title           Task/User API
// @version         1.0
// @description     This is a sample server for managing tasks and users.
// @host            localhost:8000
// @BasePath        /.
func main() {
	app := gofr.New()

	// Run DB migrations
	app.Migrate(migrations.All())

	// Setup Task routes
	taskStore := taskstore.NewTaskStore()
	taskService := taskservice.NewTaskService(taskStore)
	taskHandler := taskhandler.NewTaskHandler(taskService)

	app.GET("/task", taskHandler.GetAllTasks)
	app.POST("/task", taskHandler.CreateTask)
	app.DELETE("/task", taskHandler.DeleteTask)
	app.PATCH("/task", taskHandler.MarkTaskComplete)
	app.GET("/task/{id}", taskHandler.HandleTaskByID)

	// Setup User routes
	userStore := userstore.NewUserStore()
	userService := userservice.NewUserService(userStore)
	userHandler := userhandler.NewUserHandler(userService)

	app.GET("/user", userHandler.GetAllUsers)
	app.POST("/user", userHandler.CreateUser)
	app.GET("/user/{id}", userHandler.GetUserbyID)

	// Start a separate HTTP server for Swagger docs
	go func() {
		http.Handle("/swagger/", httpSwagger.WrapHandler)
		log.Println("Swagger docs available at http://localhost:8001/swagger/index.html")
		log.Fatal(http.ListenAndServe(":8001", nil))
	}()

	app.Run()
}
