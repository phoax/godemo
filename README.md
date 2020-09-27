# Godemo

## Run project

To set environment config, execute:

```
go run cmd/env/main.go
```

It will create a .env file with config example values.

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

## Demo

Make a transfer

```
curl --location --request POST 'localhost:8080/account/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "address": "0x940a7acD2A11846ba2F5Fc52f68EC69daFe8C550"
}'
```

Get balance

```
curl --location --request GET 'localhost:8080/account/balance/0x940a7acD2A11846ba2F5Fc52f68EC69daFe8C550'
```
