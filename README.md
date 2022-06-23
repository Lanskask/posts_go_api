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