package repository

import (
	"fmt"
	"time"

	"github.com/sraynitjsr/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

var studentCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	studentCollection = client.Database("students").Collection("students")
}

func AddStudent(student model.Student) error {
	_, err := studentCollection.InsertOne(context.TODO(), student)
	return err
}

func DeleteStudent(id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := studentCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}

func FindStudent(id string) (model.Student, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	var student model.Student
	err := studentCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&student)
	return student, err
}

func GetAllStudents() []model.Student {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	cursor, err := studentCollection.Find(ctx, bson.D{})
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("Error: Operation timed out while fetching students.")
		} else {
			fmt.Println("Error: Unable to execute Find query. ", err)
		}
		return nil
	}

	var students []model.Student
	err = cursor.All(ctx, &students)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("Error: Operation timed out while reading students.")
		} else {
			fmt.Println("Error: Unable to read data from cursor. ", err)
		}
		return nil
	}

	return students
}

func FindStudentsByName(name string) []model.Student {
	cursor, _ := studentCollection.Find(context.TODO(), bson.M{"name": name})
	var students []model.Student
	_ = cursor.All(context.TODO(), &students)
	return students
}

func FindStudentByRoll(roll string) (model.Student, error) {
	var student model.Student
	err := studentCollection.FindOne(context.TODO(), bson.M{"roll": roll}).Decode(&student)
	return student, err
}

func RollExists(roll string) bool {
	count, _ := studentCollection.CountDocuments(context.TODO(), bson.M{"roll": roll})
	return count > 0
}
