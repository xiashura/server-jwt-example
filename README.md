# server jwt example heroku

## request
```curl
curl --location --request GET 'https://server-jwt-example.herokuapp.com/api/v1/Registration' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Email":"Test@aaaa.aaa",
    "Name":"11111",
    "Password":"asdsad",
    "Authorized":false
}    
'
```
## response
```
{
    "Access": "token",
    "Refresh": "token"
}
```

## request
```curl
curl --location --request GET 'https://server-jwt-example.herokuapp.com/api/v1/Authentication' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Access": "token",
    "Refresh": "token"
}
'
```
## response
```
<h1>Auth</h1>
```

## request
```
curl --location --request GET 'https://server-jwt-example.herokuapp.com/api/v2/Updatetoken' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Email":"Test@aaaa.aaa",
    "Name":"11111",
    "Password":"asdsad",
    "Authorized":false
}    
'
```
## response
```
{
    "Access": "token",
    "Refresh": "token"
}
```

## request
```
curl --location --request GET 'https://server-jwt-example.herokuapp.com/api/v3/Refresh/delete/one' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Access": "token",
    "Refresh": "token"
}''
```
## response
```
{
    "MatchedCount": 1,
    "ModifiedCount": 1,
    "UpsertedCount": 0,
    "UpsertedID": null
}
```
## request
```
curl --location --request GET 'https://server-jwt-example.herokuapp.com/api/v3/Refresh/delete/many' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Email":"Test@aaaa.aaa",
    "Name":"11111",
    "Password":"asdsad",
    "Authorized":false
}    
'
```
## response
```
{
    "MatchedCount": 1,
    "ModifiedCount": 1,
    "UpsertedCount": 0,
    "UpsertedID": null
}
```
