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

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.LoadConfig(".profile")
	systConfig, err := config.GetSystemConfig("system_config.yaml")
	if err != nil {
		log.Fatalf("Failed to get system config: %s", err)
	}

	repo := getRepository(systConfig)

	postService := service.NewPostService(repo)
	postController := controller.NewPostController(postService)

	rout := getRouter(systConfig)

	rout.Get("/", simplePrt)
	rout.Get("/posts", postController.GetPosts)
	rout.Post("/posts", postController.AddPost)

	port := getPort()
	log.Printf("Server listening on port %s.\n", port)
	if err := rout.ListenAndServe(port); err != nil {
		log.Fatalf("err: %s\n", err)
	}

}

func getRouter(systConfig config.SystemConfig) router.Router {
	var rout router.Router
	switch systConfig.Router {
	case config.FIBER:
		rout = router.NewFiberRouter()
	case config.CHI:
		rout = router.NewChiRouter()
	case config.MUX:
		rout = router.NewMuxRouter()
	case config.GIN:
		rout = router.NewGinRouter()
	}
	return rout
}

func getRepository(systConfig config.SystemConfig) repository.PostRepo {
	var repo repository.PostRepo
	var err error

	switch systConfig.DB {
	case config.FIREBASE:
		repo, err = repository.NewFirebaseRepo()
		if err != nil {
			log.Fatalf("Failed to create Firebase repo: %s", err)
		}
	case config.MEM:
		repo, err = repository.NewMemRepo()
		if err != nil {
			log.Fatalf("Failed to create Mem repo: %s", err)
		}
	case config.SQLITE:
		repo, err = repository.NewSQLiteRepo(true)
		if err != nil {
			log.Fatalf("Failed to create SQLite repo: %s", err)
		}
	}
	return repo
}

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		return ":" + "8080"
	}
	return ":" + port
}
