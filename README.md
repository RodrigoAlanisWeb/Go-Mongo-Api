# Golang Mongodb Api
> ## Api REST With **Golang** gin Package

<img src="./Go.png" width="250"/> <img src="./mongo.svg" width="250"/>

# Packages

<br>

|routing|database      |utils  |
|-------|--------------|-------|
|  gin  | mongo driver |strconv|
|  http | context      |rand   |
<br>

# How To Use Gin

## Install 

```
go get -u github.com/gin-gonic/gin
```
<br>

## Basic Router

```go
func main() {
    // Create A Router
    router := gin.Default()
    // Home Route
    router.GET("/",func (c *gin.Context) {
        // Send A JSON
        c.JSON(200, gin.H{
            "msg": "Hi! Welcome To My API"
        })
    })

    // Run The Server
    router.RUN()

    // http://localhost:8080
}
```
<br>

## More Info In Gin Docs

[gin docs](https://github.com/gin-gonic/gin)

<br>

# MongoDb With Golang

<br>

## Install

```
go get go.mongodb.org/mongo-driver/mongo
```

## Connect

```go
// Global Db Variable
var db *mongo.Client

func connect() {
    // Connect With The Mongo Client
	con, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println(err)
		return
	}

	// Check the connection
	err = con.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db = con
}
```

## More Info In Docs

[mongo docs](https://github.com/mongodb/mongo-go-driver)

<br>

## Thanks For Reading

[github profile](https://github.com/RodrigoAlanisWeb)