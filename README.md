# Godemo

## Run project

To run the project, execute:

```
go run cmd/godemo-server/main.go --host 0.0.0.0 --port 8080
```

## Generate API

### Update and regenerate API

Routes can be changed in `docs/swagger.yml`. To regenerate the project with updated routes, execute:

```
swagger generate server -f docs/swagger.yml
```

## FAQ

### Swagger command error

When generating swagger command, it the error `target must reside inside a location in the $GOPATH/src or be a module` appaers, execute:

```
go mod init github.com/phoax/godemo
```
