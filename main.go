package main

import (
	"log"
	"os"
	"posts_api/config"
	"posts_api/controller"
	"posts_api/repository"
	"posts_api/router"
	"posts_api/service"
)

//const port = ":8000"

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
	log.Printf("Server listening on port %s.\n", port)
	if err := rout.ListenAndServe(port); err != nil {
		log.Fatalf("err: %s\n", err)
	}

}

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		return ":" + "8080"
	}
	return ":" + port
}
