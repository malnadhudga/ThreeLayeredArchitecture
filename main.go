package main

import (
	_ "ThreeLayeredArchitecture/docs"
	taskhandler "ThreeLayeredArchitecture/handler/task"
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

	taskStore := taskstore.NewTaskStore()
	taskService := taskservice.NewTaskService(taskStore)
	taskHandler := taskhandler.NewTaskHandler(taskService)

	//userDB, err := userds.InitDB()
	//if err != nil {
	//	log.Fatal("User DB connection failed:", err)
	//}
	//defer userDB.Close()

	userStore := userstore.NewUserStore()
	userService := userservice.NewUserService(userStore)
	userHandler := userhandler.NewUserHandler(userService)

	app.GET("/task", taskHandler.GetAllTasks)
	app.POST("/task", taskHandler.CreateTask)
	app.DELETE("/task", taskHandler.DeleteTask)
	app.PATCH("/task", taskHandler.MarkTaskComplete)
	app.GET("/task/{id}", taskHandler.HandleTaskByID)

	app.GET("/user", userHandler.GetAllUsers)
	app.POST("/user", userHandler.CreateUser)
	app.GET("/user/{id}", userHandler.GetUserbyID)

	app.Run()

	// Swagger
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
