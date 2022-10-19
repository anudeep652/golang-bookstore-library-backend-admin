package bookController

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID primitive.ObjectID `json:"_id"`
	jwt.StandardClaims
}

func GenerateJwt(_id primitive.ObjectID) (string, error) {
	expiration := time.Second * 1296000 //15days
	claims := &Claims{ID: _id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil

}

type Review struct {
	Reviewer string    `json:"reviewer" bson:"reviewer"`
	Subject  string    `json:"subject" bson:"subject"`
	Message  string    `json:"message" bson:"message"`
	Date     time.Time `json:"date" bson:"date"`
	Stars    int       `json:"stars" bson:"stars" default:"0"`
}

type User struct {
	ID        *primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string              `json:"username" bson:"username"`
	Email     string              `json:"email" bson:"email"`
	Password  string              `json:"password" size:"100" bson:"password"`
	CreatedAt time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt" bson:"updatedAt"`
	Role      string              `json:"role" bson:"role"`
}

type Book struct {
	ID          *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string              `json:"name" bson:"name"`
	Author      string              `json:"author" bson:"author"`
	PayAmount   int                 `json:"payAmount" bson:"payAmount"`
	RentAmount  int                 `json:"rentAmount" bson:"rentAmount"`
	ImageUrl    string              `json:"imageUrl" bson:"imageUrl"`
	Reviews     []Review            `json:"reviews" default:"[]" bson:"reviews"`
	Description string              `json:"description" bson:"description"`
}

type admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type returnJson struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

const dbName = "bookstore"
const colName = "users"

// * MOST IMPORTANT
var userCollection *mongo.Collection
var bookCollection *mongo.Collection

// go file execution nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/MyProgram/main.go
// connect with mongodb
func init() {

	connectionString := os.Getenv("MONGO_URI")
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")
	userCollection = client.Database(dbName).Collection(colName)
	bookCollection = client.Database(dbName).Collection("books")
	fmt.Println("Collection instance is ready")
}

func AdminLogin(c *fiber.Ctx) error {

	var enteredCredentials admin
	c.BodyParser(&enteredCredentials)

	var user User

	err := userCollection.FindOne(context.Background(), bson.M{"email": enteredCredentials.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusUnauthorized)
			return c.SendString("Invalid user")
		}
		log.Fatal(err)
	}
	if user.Role == "user" {
		c.SendStatus(400)
		return c.SendString("You are not admin")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(enteredCredentials.Password))
	if err != nil {
		log.Fatal(err)
	}

	token, _ := GenerateJwt(*user.ID)

	return c.JSON(returnJson{user.Username, token})

}

func RegisterBook(c *fiber.Ctx) error {
	fmt.Println("Are u there?")
	var book Book
	c.BodyParser(&book)
	_, err := bookCollection.InsertOne(context.TODO(), book)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.SendString("Successfully added book")
}
