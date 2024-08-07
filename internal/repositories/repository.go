package repositories

import (
	"context"
	"github.com/nelsonalves117/go-orders-api/internal/canonical"
	"github.com/nelsonalves117/go-orders-api/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetAllOrders() ([]canonical.Order, error)
	GetOrderById(id string) (canonical.Order, error)
	CreateOrder(order canonical.Order) (canonical.Order, error)
	UpdateOrder(id string, order canonical.Order) (canonical.Order, error)
	DeleteOrder(id string) error
}

type repository struct {
	collection *mongo.Collection
}

func New() Repository {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Get().ConnectionString))
	if err != nil {
		panic(err)
	}

	return &repository{
		collection: client.Database("order_db").Collection("orderSlice"),
	}
}

func (repo *repository) GetAllOrders() ([]canonical.Order, error) {
	var orderSlice []canonical.Order

	res, err := repo.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for res.Next(context.Background()) {
		var order canonical.Order

		err := res.Decode(&order)
		if err != nil {
			return nil, err
		}

		orderSlice = append(orderSlice, order)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return orderSlice, nil
}

func (repo *repository) GetOrderById(id string) (canonical.Order, error) {
	var order canonical.Order

	err := repo.collection.FindOne(context.Background(), bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}).Decode(&order)

	if err != nil {
		return canonical.Order{}, err
	}

	return order, nil
}

func (repo *repository) CreateOrder(order canonical.Order) (canonical.Order, error) {
	_, err := repo.collection.InsertOne(context.Background(), order)
	if err != nil {
		return canonical.Order{}, err
	}

	return order, nil
}

func (repo *repository) UpdateOrder(id string, order canonical.Order) (canonical.Order, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	fields := bson.M{
		"$set": bson.M{
			"products": order.Products,
			"total":    order.Total,
			"status":   order.Status,
		},
	}

	_, err := repo.collection.UpdateOne(context.Background(), filter, fields)

	if err != nil {
		return canonical.Order{}, err
	}

	return order, nil
}

func (repo *repository) DeleteOrder(id string) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
