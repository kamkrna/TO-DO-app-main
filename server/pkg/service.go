package pkg

import (
	"context"
	"fmt"
	"log"
	"my-app-server/server/auth"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	db *mongo.Client
}

func NewService(client *mongo.Client) *Service {

	return &Service{
		db: client,
	}
}

func (s *Service) CreateUserdata(username, password string) (UserDB, error){
	coll := s.db.Database("Todo_app").Collection("Users")

	hashedpassword, err := auth.HashedPassword(password)
	if err != nil {
		return UserDB{},fmt.Errorf("wrong password")
	}
	user := UserDB{
		Username : username,
		Password : hashedpassword,
		// Email : email,
	}

	result,err := coll.InsertOne(context.TODO(),user)
	if err != nil {
		log.Println("Error creating user!", err)
		return UserDB{}, err
	}
	
	id, ok := result.InsertedID.(primitive.ObjectID)
    if !ok {
        return UserDB{}, fmt.Errorf("unexpected InsertedID type: %T", result.InsertedID)
    }
    user.ID = id
	return user, nil
}

func (s *Service) FindUserbyusername(username string) (UserDB, error) {
    coll := s.db.Database("Todo_app").Collection("Users")
    var user UserDB
    err := coll.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return UserDB{}, fmt.Errorf("user not found")
        }
        log.Printf("Can't find the user! Username: %s, Error: %v", username, err)
        return UserDB{}, err
    }
	log.Printf("found user: %+v",user)
    return user, nil
}

func (s *Service) CreateTaskdata(title, description string, due_date time.Time, userId primitive.ObjectID)(TaskDB,error){
	coll := s.db.Database("Todo_app").Collection("tasklist")
	log.Printf("taskcreate logic hit")
	task:= TaskDB{
		Title : title,
		Description : description,
		Due_date : due_date,
		UserId: userId,
	}
	result,err := coll.InsertOne(context.TODO(),task)
	if err!=nil{
		log.Println("Error Creating the task!", err)
		return TaskDB{}, err
	}
	log.Printf("Inserted Task ID: %v", result.InsertedID)
	log.Printf("task data created")
	return task, nil
}

func (s *Service) FindTaskbyUid(userID primitive.ObjectID) ([]TaskDB, error){
	coll := s.db.Database("Todo_app").Collection("tasklist")
	filter := bson.M{"user_id" : userID}
	cursor,err := coll.Find(context.TODO(),filter)
	if err != nil{
		log.Println("Can't finding the task!", err)
		 return nil, err 
	}
	var tasks []TaskDB
	if err = cursor.All(context.TODO(), &tasks); err != nil{
		log.Println("Error retrieving the task!")
		return nil, err
	}
	return tasks, nil
}

func (s *Service) Deletetaskdata(taskID primitive.ObjectID) error{
	coll := s.db.Database("Todo_app").Collection("tasklist")
	filter := bson.M{"_id" : taskID}
	result,err := coll.DeleteOne(context.TODO(),filter)
	if err != nil{
		log.Println("Can't delete the task!", err)
		return err
	}
	if result.DeletedCount == 0{
		log.Println("No task was deleted!")
		return mongo.ErrNoDocuments
	}
	return nil
}

func (s *Service) FindUser(username string) (UserDB, error){
	coll:= s.db.Database("Todo_app").Collection("Users")
	var user UserDB
	err := coll.FindOne(context.TODO(),bson.M{"username": username}).Decode(&user)
	if err != nil {
        log.Printf("Can't find the user: %v", err)
        return UserDB{}, err
    }
	return user, err
}