package main

import (
	"log"
	"posts_api/config"
	"posts_api/controller"
	"posts_api/repository"
	"posts_api/router"
	"posts_api/service"
)

const port = ":8000"

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

	log.Println("Server listening on port", port)
	if err := rout.ListenAndServe(port); err != nil {
		log.Fatalf("err: %s\n", err)
	}

}
