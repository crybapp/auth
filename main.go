package main

/**
 * Go packages
 */
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

/**
 * Vendor packages
 */
import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Resource : http response for authenticated property
type Resource struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Resource bson.M `json:"resource"`
}

// AuthenticationRequest : http request for authentication
type AuthenticationRequest struct {
	Token string `json:"token"`
}

func createDatabaseClient() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGODB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	cancel()

	if err != nil {
		return nil, err
	}

	return client, nil
}

func extractJwtToken(req *http.Request) (string, error) {
	var token string
	var request AuthenticationRequest

	if req.Body == nil {
		return token, errors.New("no request body found")
	}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return token, err
	}

	token = request.Token

	return token, nil
}

func getCollectionName(_type string) string {
	switch _type {
	case "user":
		return "users"
	case "vm":
		return "vms"
	default:
		return _type + "s"
	}
}

func handleError(w http.ResponseWriter, e error) {
	http.Error(w, e.Error(), 500)
}

func fetchResourceFromToken(tokenString string) (Resource, error) {
	var resource Resource
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return resource, err
	}

	resourceID := fmt.Sprintf("%v", claims["id"])
	resourceType := fmt.Sprintf("%v", claims["type"])

	fmt.Println("authenticated", resourceType, resourceID)

	document, err := fetchDocument(getCollectionName(resourceType), resourceID)
	if err != nil {
		return resource, err
	}

	resource = Resource{ID: resourceID, Type: resourceType, Resource: document}

	return resource, nil
}

func fetchDocument(_collection string, id string) (bson.M, error) {
	client, err := createDatabaseClient()

	result := bson.M{}
	filter := bson.M{"info.id": &id}

	collection := client.Database("cryb").Collection(_collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = collection.FindOne(ctx, filter).Decode(&result)

	cancel()

	if err != nil {
		return result, err
	}

	return result, nil
}

func authenticate(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprintf(w, "method not acceptable\n")
		return
	}

	tokenString, err := extractJwtToken(req)
	if err != nil {
		handleError(w, err)
		return
	}

	resource, err := fetchResourceFromToken(tokenString)
	if err != nil {
		handleError(w, err)
		return
	}

	out, err := json.Marshal(resource)
	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Fprintf(w, string(out))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	http.HandleFunc("/", authenticate)

	fmt.Println("listening on :4500")
	http.ListenAndServe(":4500", nil)
}
