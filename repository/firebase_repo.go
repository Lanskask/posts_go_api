package repository

import (
	"config"
	"context"
	"entity"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/hashicorp/go-multierror"
	"google.golang.org/api/option"
)

type FirebaseRepo struct {
	client *firestore.Client
}

func NewFirebaseRepo() (PostRepo, error) {
	config.LoadConfig()

	credFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	opt := option.WithCredentialsFile(credFile)

	fApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	ctx := context.Background()
	client, err := fApp.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing a firebase client: %v", err)
	}

	return &FirebaseRepo{client: client}, nil
}

func (r *FirebaseRepo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	defer r.client.Close()

	_, _, err := r.client.Collection(tableName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		return nil, fmt.Errorf("error add data to collection: %v", err)
	}
	return post, nil
}

func (r *FirebaseRepo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	defer r.client.Close()

	var posts []entity.Post

	iter := r.client.Collection(tableName).Documents(ctx)
	var combErr error

	allDocs, err := iter.GetAll()
	if err != nil {
		combErr = multierror.Append(combErr, fmt.Errorf("err get doc from iterator: %s", err))
	}
	for _, doc := range allDocs {
		post := entity.Post{
			ID:    int(doc.Data()["ID"].(int64)),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	iter.Stop()

	//combErr, posts = IterOverDocs(iter, combErr, posts)

	if combErr != nil {
		return nil, combErr
	}
	return posts, nil
}

func (r *FirebaseRepo) Delete(post *entity.Post) (int64, error) {
	return 0, nil
}
