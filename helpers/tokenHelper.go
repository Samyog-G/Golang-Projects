package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Samyog-G/jwt-authentication/database"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	UserType  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens generates an access and refresh token for the user.
func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	// Define access token claims
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	// Define refresh token claims
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	// Generate signed JWT tokens using HS256 method
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	return token, refreshToken, nil
}

// UpdateAllTokens updates or inserts the JWT tokens for a user in the database.
func UpdateAllTokens(userID string, signedToken string, signedRefreshToken string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Prepare the update data
	var updateOBJ primitive.D
	updateOBJ = append(updateOBJ, primitive.E{Key: "token", Value: signedToken})
	updateOBJ = append(updateOBJ, primitive.E{Key: "refreshtoken", Value: signedRefreshToken})

	// Add timestamp for when the document was updated
	UpdatedAt := time.Now()
	updateOBJ = append(updateOBJ, primitive.E{Key: "updatedAt", Value: UpdatedAt})

	// Set upsert to true to create the user if they don't exist
	upsert := true
	filter := bson.M{"UserID": userID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	// Perform the update in the database
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateOBJ},
		},
		&opt,
	)
	if err != nil {
		log.Panic(err)
	}
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken, &SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("The token is invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("Token is expired")
		msg = err.Error()
		return
	}
	return
}
