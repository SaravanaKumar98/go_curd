package service

import (
	"context"
	"curd/database"
	"curd/model"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(user *model.User) error {
	collection := database.CrudDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	filter := bson.M{"email": user.Email}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count == 0 {
		dob, err := time.Parse("2006-01-02", user.DOB)
		if err != nil {
			return err
		}
		user.Age = calculateAge(dob)
		user.Uuid, err = NewUuid()
		if err != nil {
			return err
		}
		_, err = collection.InsertOne(ctx, user)
		if err != nil {
			return err
		}
	} else {
		return errors.New("given user is already exist")
	}
	return nil
}

func UserUpdate(user *model.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	collection := database.CrudDB.Collection("users")

	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		return err
	}
	user.Age = calculateAge(dob)
	fmt.Println(user)
	filter := bson.M{"uuid": user.Uuid}
	value := bson.M{"$set": user}
	res, err := collection.UpdateOne(ctx, filter, value)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("given data matches not found")
	}

	return nil
}

func FindUser(req *model.Request) (*model.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	collection := database.CrudDB.Collection("users")
	filter := bson.M{"uuid": req.Uuid}
	var user model.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return &model.User{}, err
	}
	fmt.Println(user)
	return &user, nil
}

func FindAllUser() ([]model.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	collection := database.CrudDB.Collection("users")
	var user []model.User
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []model.User{}, err
	}
	err = result.All(ctx, &user)
	if err != nil {
		return []model.User{}, err
	}
	return user, nil
}

func DeleteUser(req *model.Request) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	collection := database.CrudDB.Collection("users")
	filter := bson.M{"uuid": req.Uuid}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("given data not found")
	}
	return nil
}
