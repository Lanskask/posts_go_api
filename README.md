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


For multistage
```shell
docker build -t post_api -f multistage.Dockerfile .

docker run -p 8081:8081 --name=post_api_1 --env-file .profile post_api
docker run -p 8081:8081 --name sh_in_demo -it post_api /bin/sh  
```

For build image
```shell
docker build -t demo_proj_build -f build.Dockerfile .

docker run -p 8081:8081 demo_proj_build
docker run -it demo_proj_build /bin/sh
```

For build on golang docker image
```shell
docker build -t demo_proj_build_on_go -f build_on_go_image.Dockerfile .

docker run -p 8081:8081 demo_proj_build_on_go

docker run -p 8081:8081 --env-file .profile demo_proj_build_on_go 
docker run -p 8081:8081 --env-file .profile demo_proj_build_on_go env | grep -E 'GOOGLE_APPLICATION_CREDENTIALS|PORT' 

docker run -it demo_proj_build_on_go /bin/sh
```