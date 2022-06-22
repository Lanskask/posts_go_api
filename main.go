package main

import (
	"config"
	"controller"
	"log"
	"repository"
	"router"
	"service"
)

const port = ":8000"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.LoadConfig()

	repo := repository.NewFirebaseRepo()
	postService := service.NewPostService(repo)
	postController := controller.NewPostController(postService)

	//rout := router.NewMuxRouter()
	rout := router.NewChiRouter()

	rout.Get("/", simplePrt)
	rout.Get("/posts", postController.GetPosts)
	rout.Post("/posts", postController.AddPost)
	log.Println("Server listening on port", port)
	if err := rout.ListenAndServe(port); err != nil {
		log.Fatalf("err: %s\n", err)
	}

}
