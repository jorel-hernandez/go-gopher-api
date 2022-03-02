# Go mod init
```go
go mod init go-gopher-api
```

# Download dependencie
```go
go get -u github.com/gorilla/mux
```

# Run go api
```go
go run gopherapi/cmd/gopherapi/main.go -withData=true //with in memory data
go run gopherapi/cmd/gopherapi/main.go -withData=false //without in memory data
```

# GET all gophers
```curl
curl --location --request GET http://localhost:8080/gophers

# Respose:
[
    {"ID":"01D3XZ3ZHCP3KG9VT4FGAD8KDR","name":"Jenny","image":"https://storage.googleapis.com/gopherizeme.appspot.com/gophers/0ceb2c10fc0c30575c18ff1defa1ffd41501bc62.png","age":18},
    {"ID":"01D3XZ7CN92AKS9HAPSZ4D5DP9","name":"Billy","image":"https://storage.googleapis.com/gopherizeme.appspot.com/gophers/13c7d425111a501600db8587b52bb292836c5bee.png","age":24},
    {"ID":"01D3XZ89NFJZ9QT2DHVD462AC2","name":"Rainbow","image":"https://storage.googleapis.com/gopherizeme.appspot.com/gophers/b9e8d637c91c089fd56d7b159825fc9089377118.png","age":48},
    {"ID":"01D3XZ8JXHTDA6XY05EVJVE9Z2","name":"Bjorn","image":"https://storage.googleapis.com/gopherizeme.appspot.com/gophers/fd01b36091560c2a128b8fddfb2c627d8bb7417c.png","age":32}
]
```

# GET a gopher
```curl
curl --location --request GET http://localhost:8080/gophers/01D3XZ8JXHTDA6XY05EVJVE9Z2

# Respose:
{"ID":"01D3XZ8JXHTDA6XY05EVJVE9Z2","name":"Bjorn","image":"https://storage.googleapis.com/gopherizeme.appspot.com/gophers/fd01b36091560c2a128b8fddfb2c627d8bb7417c.png","age":32}
```

# Source
- https://blog.friendsofgo.tech/posts/como_crear_una_api_rest_en_golang/
- https://github.com/friendsofgo/gopherapi/blob/v0.1/pkg/server/api.go