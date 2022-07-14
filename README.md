# Posts with Clean architecture

Default port: 8080  
You can change env variables in `.profile` file

Connect your firebase project and   
put `firebase-adminsdk.json` key file to the root of this project.

Also add a `.profile` file with the following content  
```shell
export GOOGLE_APPLICATION_CREDENTIALS=./your-firebase-admin-sdk-key.json
```



## How to run test coverage

```shell
go test -coverprofile=coverage.out                                                                          21:38:53
go tool cover -html=coverage.out
```

## How to run Docker

Build and run 
```shell
docker build -t post_api -f multistage.Dockerfile .

docker run -p 8081:8081 --name=post_api_1 --env-file .profile post_api
```

To check env variables
```shell
docker run --env-file .profile post_api env | grep -E 'GOOGLE_APPLICATION_CREDENTIALS|PORT'
```

To `sh` in image
```shell
docker run --name sh_to_posts -it post_api /bin/sh
```