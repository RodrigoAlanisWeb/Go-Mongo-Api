package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

// Product product
type Product struct {
	Id    int
	Name  string
	Count int
}

func connect() {
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

func createProduct(c *gin.Context) {
	name := c.PostForm("name")
	count, err := strconv.Atoi(c.PostForm("count"))

	if err != nil {
		c.JSON(500, gin.H{"msg": "Error By Creating The Product"})
		fmt.Println(err)
		return
	}

	pds := db.Database("go-mongo").Collection("products")

	pd := Product{Name: name, Count: count, Id: rand.Int()}
	res, err := pds.InsertOne(context.TODO(), pd)

	if err != nil {
		c.JSON(500, gin.H{"msg": "Error By Creating The Product"})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"msg": "Product Created", "id": res.InsertedID})
}

func getProducts(c *gin.Context) {
	pds := db.Database("go-mongo").Collection("products")

	var data []Product

	cur, err := pds.Find(context.TODO(), bson.D{{}}, options.Find())

	fmt.Println(cur)

	if err != nil {
		c.JSON(500, gin.H{"msg": "Error By Getting The Products"})
		fmt.Println(err)
		return
	}
	for cur.Next(context.TODO()) {
		fmt.Println(cur)
		var el Product
		err := cur.Decode(&el)

		if err != nil {
			c.JSON(500, gin.H{"msg": "Error By Getting The Products"})
			fmt.Println(err)
			return
		}

		data = append(data, el)
	}

	cur.Close(context.TODO())

	c.JSON(200, gin.H{"data": data})
}

func getProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"msg": "Id Invalid"})
		fmt.Println(err)
		return
	}
	var product Product
	pds := db.Database("go-mongo").Collection("products")
	err = pds.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&product)
	if err != nil {
		c.JSON(500, gin.H{"msg": "Error"})
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"data": product})
}

func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"msg": "Id Invalid"})
		fmt.Println(err)
		return
	}
	pds := db.Database("go-mongo").Collection("products")
	_, err = pds.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if err != nil {
		c.JSON(500, gin.H{"msg": "Error"})
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"deleted": "Delete Successfully"})
}

func updatedProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"msg": "Id Invalid"})
		fmt.Println(err)
		return
	}
	name := c.PostForm("name")
	count, err := strconv.Atoi(c.PostForm("count"))
	pds := db.Database("go-mongo").Collection("products")
	update := bson.M{
		"$set": bson.M{"name": name, "count": count},
	}
	_, err = pds.UpdateOne(context.TODO(), bson.D{{"id", id}}, update)
	if err != nil {
		c.JSON(500, gin.H{"msg": "Error"})
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"updated": "Updated Successfully"})
}

func main() {
	router := gin.Default()
	connect()
	router.POST("/create", createProduct)
	router.GET("/products", getProducts)
	router.GET("/get/:id", getProduct)
	router.DELETE("/delete/:id", deleteProduct)
	router.PUT("/update/:id", updatedProduct)

	router.Run()
}
