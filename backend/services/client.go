package services

import (
	"context"
	"time"

	"github.com/amit645/skillmate-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientService struct {
	db *mongo.Client
}

func NewClientService(db *mongo.Client) *ClientService {
	return &ClientService{db: db}
}

func (s *ClientService) CreateRequest(title, description, city string, budget float64, skills []string, clientID string) (*models.Request, error) {
	clientObjID, err := primitive.ObjectIDFromHex(clientID)
	if err != nil {
		return nil, err
	}

	request := models.Request{
		Title:       title,
		Description: description,
		Skills:      skills,
		Budget:      budget,
		City:        city,
		Client:      clientObjID,
		CreatedAt:   time.Now().Unix(),
	}

	collection := s.db.Database("skillmate").Collection("requests")
	result, err := collection.InsertOne(context.TODO(), request)
	if err != nil {
		return nil, err
	}

	request.ID = result.InsertedID.(primitive.ObjectID)
	return &request, nil
}

func (s *ClientService) BrowseServices(skill, city string, minPrice, maxPrice float64) ([]models.Service, error) {
	collection := s.db.Database("skillmate").Collection("services")
	filter := bson.M{}

	if skill != "" {
		filter["tags"] = bson.M{"$in": []string{skill}}
	}
	if city != "" {
		filter["city"] = city
	}
	if minPrice > 0 || maxPrice > 0 {
		filter["price"] = bson.M{"$gte": minPrice, "$lte": maxPrice}
	}

	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetSort(bson.M{"created_at": -1}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var services []models.Service
	if err = cursor.All(context.TODO(), &services); err != nil {
		return nil, err
	}

	return services, nil
}
