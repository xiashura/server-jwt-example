# server jwt example

install 
```bash
 go mod download
```

run 
```bash 
go run cmd/service/authentication/main.go
```

Docker 
```bash
docker-compose up -d
```

## request
```curl
curl --location --request GET '127.0.0.1:8080/Registration' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Email":"asddasad",
    "Name":"asda",
    "Password":"asdsad",
    "Authorized":false
}    
'
```
## response
```
{
    "Key": "youtoken"
}
```

## request
```curl
curl --location --request GET '127.0.0.1:8080/Authentication' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Key": "youtoken"
}'

```
## response
```
<h1>Home</h1>
```

## request
```
curl --location --request GET '127.0.0.1:8080/Unauthenticated' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Key": "youtoken"
}'
```
## response
```
<h1>Home</h1>
```

## request
```
curl --location --request GET '127.0.0.1:8080/Expired' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Key": "youtoken"
}'
```
## response
```
<h1>Home</h1>
```