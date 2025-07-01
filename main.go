package main

import (
	"fmt"
	"log"
	"net/http"

	_ "ThreeLayeredArchitecture/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	// Task
	"ThreeLayeredArchitecture/datasource/task"
	taskhandler "ThreeLayeredArchitecture/handler/task"
	taskservice "ThreeLayeredArchitecture/services/task"
	taskstore "ThreeLayeredArchitecture/store/task"

	// User
	userds "ThreeLayeredArchitecture/datasource/user"
	userhandler "ThreeLayeredArchitecture/handler/user"
	userservice "ThreeLayeredArchitecture/services/user"
	userstore "ThreeLayeredArchitecture/store/user"
)

// @title           Task/User API
// @version         1.0
// @description     This is a sample server for managing tasks and users.
// @host            localhost:8000
// @BasePath        /
func main() {
	// Task DB init...
	taskDB, err := task.InitDB()
	if err != nil {
		log.Fatal("Task DB connection failed:", err)
	}
	defer taskDB.Close()

	taskStore := taskstore.NewTaskStore(taskDB)
	taskService := taskservice.NewTaskService(taskStore)
	taskHandler := taskhandler.NewTaskHandler(taskService)

	// User DB init...
	userDB, err := userds.InitDB()
	if err != nil {
		log.Fatal("User DB connection failed:", err)
	}
	defer userDB.Close()

	userStore := userstore.NewUserStore(userDB)
	userService := userservice.NewUserService(userStore)
	userHandler := userhandler.NewUserHandler(userService)

	// Routes
	http.HandleFunc("/task", taskHandler.HandleTasks)
	http.HandleFunc("/task/{id}", taskHandler.HandleTaskByID)
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetAllUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/user/{id}", userHandler.GetUserbyID)

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server running at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
