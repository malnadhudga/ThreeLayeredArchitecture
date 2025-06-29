package main

import (
	"fmt"
	"log"
	"net/http"

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

func main() {
	// --- TASK DB ---
	taskDB, err := task.InitDB()
	if err != nil {
		log.Fatal("Task DB connection failed:", err)
	}
	defer func() {
		if err := taskDB.Close(); err != nil {
			log.Printf("Error closing Task DB: %v", err)
		}
	}()

	taskStore := taskstore.NewTaskStore(taskDB)
	taskService := taskservice.NewTaskService(taskStore)
	taskHandler := taskhandler.NewTaskHandler(taskService)

	// --- USER DB ---
	userDB, err := userds.InitDB()
	if err != nil {
		log.Fatal("User DB connection failed:", err)
	}
	defer func() {
		if err := userDB.Close(); err != nil {
			log.Printf("Error closing User DB: %v", err)
		}
	}()

	userStore := userstore.NewUserStore(userDB)
	userService := userservice.NewUserService(userStore)
	userHandler := userhandler.NewUserHandler(userService)

	//  TASK
	http.HandleFunc("/task", taskHandler.HandleTasks)
	http.HandleFunc("/task/{id}", taskHandler.HandleTaskByID) // GET by ID

	// USER
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			userHandler.GetAllUsers(w, r)
		case "POST":
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/user/{id}", userHandler.GetUserbyID)

	fmt.Println("Server on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
