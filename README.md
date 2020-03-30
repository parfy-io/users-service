# users-service
CRUD service for clients and its users

## API

### clients

#### create
Request:
```http request
POST /v1/clients

{"id": "test"}
```
Response created:
```http response
201 CREATED
Location: /v1/clients/test
```
Response client already exists:
```http response
409 CONFLICT

{"message":"client already exists"}
```

### internal

#### alive
Request:
```http request
GET /v1/internal/alive
```
Response:
```http response
200 OK 

{"alive":true}
```


