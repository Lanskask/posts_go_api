# Posts with Clean architecture

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

```shell
docker build -t demonstration_project -f multistage.Dockerfile .

docker run -p 8081:8081 demonstration_project
docker run  -p 8081:8081 demonstration_project:multistage
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

docker run -p 8081:8081 demo_proj_build_on_go --env-file

docker run -it demo_proj_build_on_go /bin/sh
```