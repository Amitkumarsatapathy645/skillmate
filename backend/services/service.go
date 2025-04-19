package services

import (
	"context"
	"time"

	"github.com/amit645/skillmate-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceService struct {
	db *mongo.Client
}

func NewServiceService(db *mongo.Client) *ServiceService {
	return &ServiceService{db: db}
}

func (s *ServiceService) CreateService(title, description, city string, price float64, tags []string, freelancerID string) (*models.Service, error) {
	freelancerObjID, err := primitive.ObjectIDFromHex(freelancerID)
	if err != nil {
		return nil, err
	}

	service := models.Service{
		Title:       title,
		Description: description,
		Price:       price,
		Tags:        tags,
		Freelancer:  freelancerObjID,
		City:        city,
		CreatedAt:   time.Now().Unix(),
	}

	collection := s.db.Database("skillmate").Collection("services")
	result, err := collection.InsertOne(context.TODO(), service)
	if err != nil {
		return nil, err
	}

	service.ID = result.InsertedID.(primitive.ObjectID)
	return &service, nil
}

func (s *ServiceService) GetServicesByFreelancer(freelancerID string) ([]models.Service, error) {
	freelancerObjID, err := primitive.ObjectIDFromHex(freelancerID)
	if err != nil {
		return nil, err
	}

	collection := s.db.Database("skillmate").Collection("services")
	cursor, err := collection.Find(context.TODO(), bson.M{"freelancer": freelancerObjID})
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
