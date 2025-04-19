package services

import (
	"context"
	"time"

	"github.com/amit645/skillmate-backend/models"
	"github.com/amit645/skillmate-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *mongo.Client
}

func NewAuthService(db *mongo.Client) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(name, email, password, role string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := models.User{
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		Role:      role,
		CreatedAt: time.Now().Unix(),
	}

	collection := s.db.Database("skillmate").Collection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()
	token, err := utils.GenerateJWT(userID, role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	collection := s.db.Database("skillmate").Collection("users")
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.ID.Hex(), user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
