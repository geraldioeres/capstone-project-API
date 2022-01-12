package user

import (
	"context"
	"disspace/business/user"
	"disspace/helpers/messages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBUserRepository struct {
	Conn *mongo.Database
}

func NewMongoDBUserRepository(conn *mongo.Database) user.Repository {
	return &MongoDBUserRepository{
		Conn: conn,
	}
}

func (repository *MongoDBUserRepository) Register(ctx context.Context, data *user.UserDomain) (user.UserDomain, error) {
	newUser := UserFromDomain(*data)
	checkVar := []User{}
	filter := bson.D{{Key: "username", Value: newUser.Username}}
	check, errCheck := repository.Conn.Collection("users").Find(ctx, filter)
	if errCheck != nil {
		panic(errCheck)
	}
	if errCheck = check.All(ctx, &checkVar); errCheck != nil {
		panic(errCheck)
	}
	if len(checkVar) != 0 {
		return user.UserDomain{}, messages.ErrUsernameAlreadyExist
	}

	cursor, err := repository.Conn.Collection("users").InsertOne(ctx, newUser)
	userID := cursor.InsertedID.(primitive.ObjectID).Hex()
	newProfileUser := UserProfile{
		UserId:      userID,
		ProfilePict: "gs://disspace-250a1.appspot.com/profile_pict/profile_default.jpg",
		Bio:         " ",
		Following:   []string{},
		Followers:   []string{},
		Threads:     []string{},
		Reputation:  0,
	}
	_, errProfile := repository.Conn.Collection("user_profile").InsertOne(ctx, newProfileUser)
	if err != nil {
		return user.UserDomain{}, err
	}
	if errProfile != nil {
		return user.UserDomain{}, err
	}
	return user.UserDomain{}, nil

}

func (repository *MongoDBUserRepository) UserProfileGetByUserID(ctx context.Context, id string) (user.UserProfileDomain, error) {
	result := UserProfile{}
	filter := bson.D{{Key: "user_id", Value: id}}
	err := repository.Conn.Collection("user_profile").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.UserProfileToDomain(), nil
}

func (repository *MongoDBUserRepository) GetUserByID(ctx context.Context, id string) (user.UserDomain, error) {
	result := User{}
	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return user.UserDomain{}, errorConvert
	}
	filter := bson.D{{Key: "_id", Value: convert}}
	err := repository.Conn.Collection("users").FindOne(ctx, filter).Decode(result)
	if err != nil {
		panic(err)
	}
	return result.UserToDomain(), nil
}

func (repository *MongoDBUserRepository) Login(ctx context.Context, username string, password string) (user.UserDomain, error) {
	result := User{}
	filter := bson.D{{Key: "username", Value: username}}
	err := repository.Conn.Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return user.UserDomain{}, err
	}
	if password != result.Password {
		return user.UserDomain{}, messages.ErrInvalidCredentials
	}
	return result.UserToDomain(), nil
}

func (repository *MongoDBUserRepository) InsertSession(ctx context.Context, dataSession user.UserSessionDomain) error {
	newSession := SessionFromDomain(dataSession)
	_, err := repository.Conn.Collection("session").InsertOne(ctx, newSession)
	if err != nil {
		return err
	}

	return nil
}
