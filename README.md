To run the application, you have to run
```shell
docker compose up
```
- The application was designed to server gRPC and Rest API simultaneously, and we can configure through configuration which specifically to serve.
- Port 8001 works for Rest API, Port 8002 works for gRPC
- To test gRPC, you can  install grpcurl through ```brew install grpcurl ```

>Register Request
```sh
curl --location --request POST 'localhost:8001/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "mohammed",
    "email": "mohammedosama@ieee.org",
    "password": "hello",
    "password_confirmation": "hello"
}'
grpcurl -d '{                                       
    "name": "mohammed",
    "email": "mohammedosama@ieee.org",
    "password": "hello",
    "password_confirmation": "hello"
}' -plaintext localhost:8002 "Users.Register"


```
>Login Request

```shell
curl --location --request POST 'localhost:8001/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "mohammedosama@ieee.org",
    "password": "hello"
}'

grpcurl -d '{ 
        "email": "mohammedosama@ieee.org",                                                                                                                                                          
    "password": "hello"
}' -plaintext localhost:8002 "Users.Login"
```
>Get Companies
```shell
curl --location --request GET 'localhost:8001/api/companies?page=1'


grpcurl -d '{"page": 1}' -plaintext localhost:8002 "Companies.GetCompanyList"  
```

>Get Company By ID
```shell
curl --location --request GET 'localhost:8001/api/companies/823eac31-23c1-40d2-b7b2-d46da5e7ab8f'


grpcurl -d '{"id": "4a394913-2451-4104-a506-b4cbcb9391b6"}' -plaintext localhost:8002 "Companies.GetCompany" 
```
>Create Company
```shell
curl --location --request POST 'localhost:8001/api/companies' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MDE4OTgsImlhdCI6MTY3NjQ5ODI5OCwibmJmIjoxNjc2NDk4Mjk4LCJzdWIiOiJmNTE3YzAxZC05NGUyLTQxYzgtOTU4ZS0zYTZjMmFjMjA2ZDgifQ.ozaXoyuT2jrlU6Ig69XAhnsrMYBaspq7GClBJs1kAgk' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Company Name",
    "description": "Something really long goes here",
    "registered": true,
    "headcount": 123,
    "company_type": "NonProfit"
}'



grpcurl -rpc-header "authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MDE4OTgsImlhdCI6MTY3NjQ5ODI5OCwibmJmIjoxNjc2NDk4Mjk4LCJzdWIiOiJmNTE3YzAxZC05NGUyLTQxYzgtOTU4ZS0zYTZjMmFjMjA2ZDgifQ.ozaXoyuT2jrlU6Ig69XAhnsrMYBaspq7GClBJs1kAgk" -d '{                                                                                                                                                                                              
    "name": "Hello There",
    "description": "Something really long goes  here",
    "registered": true,
    "headcount": 123,
    "type": "Corporations"
}' -plaintext localhost:8002 "Companies.CreateCompany"
```
>Update Company
```shell
curl --location --request PATCH 'localhost:8001/api/companies/819c1ccc-08ba-4e15-a8c1-7ed31e490c20' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MDE4OTgsImlhdCI6MTY3NjQ5ODI5OCwibmJmIjoxNjc2NDk4Mjk4LCJzdWIiOiJmNTE3YzAxZC05NGUyLTQxYzgtOTU4ZS0zYTZjMmFjMjA2ZDgifQ.ozaXoyuT2jrlU6Ig69XAhnsrMYBaspq7GClBJs1kAgk' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Hello There",
    "description": "Something really long goes here",
    "registered": true,
    "headcount": 123,
    "company_type": "NonProfit"
}'
grpcurl -rpc-header "authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MDE4OTgsImlhdCI6MTY3NjQ5ODI5OCwibmJmIjoxNjc2NDk4Mjk4LCJzdWIiOiJmNTE3YzAxZC05NGUyLTQxYzgtOTU4ZS0zYTZjMmFjMjA2ZDgifQ.ozaXoyuT2jrlU6Ig69XAhnsrMYBaspq7GClBJs1kAgk" -d '{
    "name": "Hello There",
    "description": "Something really long goes  here",
    "registered": true,
    "headcount": 12, 
    "type": "Corporations",                                                                                                                                                                                                                              "id": "1f1529d6-0f21-4c4e-8a6b-2aa1fcde6fe2"
}' -plaintext localhost:8002 "Companies.UpdateCompany"
```
>Delete Company
```shell
curl --location --request DELETE 'localhost:8001/api/companies/819c1ccc-08ba-4e15-a8c1-7ed31e490c20' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MDE4OTgsImlhdCI6MTY3NjQ5ODI5OCwibmJmIjoxNjc2NDk4Mjk4LCJzdWIiOiJmNTE3YzAxZC05NGUyLTQxYzgtOTU4ZS0zYTZjMmFjMjA2ZDgifQ.ozaXoyuT2jrlU6Ig69XAhnsrMYBaspq7GClBJs1kAgk'


grpcurl -rpc-header "authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MDE4OTgsImlhdCI6MTY3NjQ5ODI5OCwibmJmIjoxNjc2NDk4Mjk4LCJzdWIiOiJmNTE3YzAxZC05NGUyLTQxYzgtOTU4ZS0zYTZjMmFjMjA2ZDgifQ.ozaXoyuT2jrlU6Ig69XAhnsrMYBaspq7GClBJs1kAgk" -d '{"id": "4a394913-2451-4104-a506-b4cbcb9391b6"}' -plaintext localhost:8002 "Companies.DeleteCompany"
```
```shell
After running docker compose up
You can create a test database
> createdb -h localhost -p 5432 -U postgres company_test
password: postgres
 
to start testing users suite
> ENV=test go test -v ./tests/integration/user/...
and for testing company suite
> ENV=test go test -v ./tests/integration/company/...
```
### For Kafka

- for each event that's dispatched we are going to invoke the corresponding handler (for company CRUD and login/register) 
