frontend
create super key for jwt token
method to create and use database docker mysql postgrest mongodb 
public github
monotring service 



HTTP 

curl -X POST  http://0.0.0.0:8080/authentication  \
-H 'Content-Type: application/json' \
-H 'Authorization:token'

curl -X GET  http://0.0.0.0:8080/authentication  \
-H 'Content-Type: application/json' \
-H 'Name: test_user' \               
-H 'Email:test@mail.com'\    
-H 'Password:12123123ee'\
     
curl -X POST  http://0.0.0.0:8080/registration  \
-H 'Content-Type: application/json' \
-H 'Name: test_user'  \
-H 'Email:test@mail.com' \
-H 'Password:12123123ee'
