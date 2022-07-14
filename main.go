package main

import (
	"config"
	"controller"
	"log"
	"os"
	"repository"
	"router"
	"service"
)

const defaultPort = ":8000"

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = defaultPort
	}
	return ":" + port
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.LoadConfig()

	//repo, _ := repository.NewFirebaseRepo()
	repo, _ := repository.NewSQLiteRepo(false)

	postService := service.NewPostService(repo)
	postController := controller.NewPostController(postService)

	//rout := router.NewMuxRouter()
	//rout := router.NewChiRouter()
	//rout := router.NewGinRouter()
	rout := router.NewFiberRouter()

	rout.Get("/", simplePrt)
	rout.Get("/posts", postController.GetPosts)
	rout.Post("/posts", postController.AddPost)

	port := getPort()
	log.Println("Server listening on port", port)
	if err := rout.ListenAndServe(port); err != nil {
		log.Fatalf("err: %s\n", err)
	}

}
