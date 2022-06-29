package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	"test.com/req/token"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var ctxG context.Context;
var clientG *mongo.Client;

var host = "192.168.43.142"

type userLog struct {
	Email string `json:"email"`
	Passsword string `json:"password"`
	Username string `json:"username"`
}
var bindUser userLog

func MongodbConnection(c *gin.Context){ 
    clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Minute)

	col := client.Database("Turism").Collection("users")
	 fmt.Println(ctx,col)
	 ctxG = ctx
	 clientG = client
}

func Login(c *gin.Context) {
	// """""""""""""""""get the form velue"""""""""""""""""
	c.ShouldBindJSON(&bindUser)
	// """""""""""""""""""find the user from db"""""""""""""""""""
	MongodbConnection(c)
	colection := clientG.Database("Turism").Collection("users")

	cur := colection.FindOne(ctxG, bson.M{ "email": bindUser.Email})
	var userOne userLog
	cur.Decode(&userOne)
	if userOne.Email != "" && bindUser.Passsword == userOne.Passsword && userOne.Email == userOne.Email{
		c.SetCookie("cookie",token.Genertetoken(),0,"/","",time.Now().Add(time.Minute*30),false,false)
		c.JSON(200, gin.H{
			"status":"valid",
			"username":userOne.Username,
		})
		// """""""""""""""""""""""send cookies"""""""""""""""""""""""
	}else{
		c.JSON(404, gin.H{
			"status":"invalid",
		})
	}
	fmt.Println(userOne)
	fmt.Println(bindUser.Passsword)
}
func Signin(c *gin.Context){
	// """""""""""""""""get the form velue"""""""""""""""""
	c.ShouldBindJSON(&bindUser)
	fmt.Println(bindUser.Email)
	// """""""""""""""""""find the user from db"""""""""""""""""""
	MongodbConnection(c)
	colection := clientG.Database("Turism").Collection("users")
	var userSignin userLog
	cur := colection.FindOne(ctxG, bson.M{"email":bindUser.Email})
	cur.Decode(&userSignin)
	fmt.Println(userSignin)
	if userSignin.Email != ""{
		c.JSON(404, gin.H{
			"status":"exist",
		})
	}else{
		userSignin.Email = bindUser.Email
		userSignin.Passsword = bindUser.Passsword
		userSignin.Username = bindUser.Username
		
		colection.InsertOne(ctxG, userSignin)
		c.JSON(200 , gin.H{
			"url":"http://192.168.43.142:5500/index.html?status=200",
		})
	}
}
func Logout(c *gin.Context){
	c.SetCookie("cookie",token.Genertetoken(),0,"/","",time.Now().Add(time.Minute*-30),false,false)
}
func Cors(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://" + host + ":5500")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	c.Next()
}