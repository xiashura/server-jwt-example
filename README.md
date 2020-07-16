# server jwt example


## request
```curl
curl --location --request GET 'http://0.0.0.0:8080/Registration' \
--header 'Content-Type: application/json' \
--data-raw '{
    "User":{
        "Email":"Email",
        "Name":"NAme",
        "Password":"Password"
    }
}'

output:{
    "Token": {
        "Time": "2020-07-18T00:21:23.371367012+05:00",
        "Key": "youtoken"
    }
}
```

```curl
curl --location --request GET 'http://0.0.0.0:8080/Authentication' \
--header 'Content-Type: application/json' \
--data-raw '{
    "User":{
        "Email":"Email",
        "Name":"NAme",
        "Password":"Password"
    },
    "Token": {
        "Time": "2020-07-17T06:02:23.585851813+05:00",
        "Key": "youtoken
    }
    
}'

output:{
    "client": {
        "User": {
            "Email": "Email",
            "Name": "NAme",
            "Password": "Password"
        },
        "Token": {
            "Time": "2020-07-17T06:02:23.585851813+05:00",
            "Authorized": true,
            "Key": "youtoken"
        }
    }
}
```

```curl
curl --location --request GET 'http://0.0.0.0:8080/Unauthenticated' \
--header 'Content-Type: application/json' \
--data-raw '{

        "User": {
            "Email": "Email",
            "Name": "NAme",
            "Password": "Password"
        },
        "Token": {
            "Time": "2020-07-17T06:02:23.585851813+05:00",
            "Authorized": true,
            "Key": "youtoken"
        }
    }'
output:{
    "client": {
        "User": {
            "Email": "Email",
            "Name": "NAme",
            "Password": "Password"
        },
        "Token": {
            "Time": "2020-07-17T06:02:23.585851813+05:00",
            "Key": "youtoken"
        }
    }
}
```
```curl
curl --location --request GET 'http://0.0.0.0:8080/Expired' \
--header 'Content-Type: application/json' \
--data-raw '{

    "User": {
        "Email": "Email",
        "Name": "NAme",
        "Password": "Password"
    },
        "Token": {
            "Time": "2020-07-17T06:02:23.585851813+05:00",
            "Authorized": true,
            "Key": "youtoken"
    }
}'
output:{
    "Token": {
        "Time": "2020-07-17T06:02:23.585851813+05:00",
        "Authorized": true,
        "Key": "youtoken"
    }
}
```
