package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        string `bson:"_id,omitempty"`
	Name      string `bson:"name,omitempty"`
	Password  string `bson:"password,omitempty"`
	CreatedAt int64  `bson:"created_at,omitempt "`
	UpdatedAt int64  `bson:"updated_at,omitempt "`
	DeletedAt int64  `bson:"deleted_at,omitempt "`
}

type UserRepo struct {
	col *mongo.Collection
}

func (r *UserRepo) CreateUser(name, password string) error {
	now := time.Now()
	u := User{
		ID:        "",
		Name:      name,
		Password:  password,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
		DeletedAt: 0,
	}
	_, err := r.col.InsertOne(context.Background(), u)
	return err
}

func (r *UserRepo) GetUserByName(name string) (*User, error) {
	result := r.col.FindOne(context.Background(), bson.D{
		{Key: "name", Value: name},
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "deleted_at", Value: bson.D{{"$exists", false}}}},
			bson.D{{Key: "deleted_at", Value: 0}},
		}},
	})
	if result.Err() != nil {
		return nil, result.Err()
	}
	u := &User{}
	err := result.Decode(u)
	return u, err
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	now := time.Now()
	admin := &User{
		ID:        "",
		Name:      "admin",
		Password:  "$2a$10$1iXYspdrld4iQ.W41m6iaOvbOgoyOncycJQ8pWRCdzyOWg8bsEMnq",
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
		DeletedAt: 0,
	}
	col := db.Collection("user")

	_, err := col.UpdateOne(context.Background(), bson.D{
		{Key: "name", Value: admin.Name},
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "deleted_at", Value: bson.D{{Key: "$exists", Value: false}}}},
			bson.D{{Key: "deleted_at", Value: 0}},
		}},
	},
		bson.D{{Key: "$set", Value: admin}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		panic(err)
	}

	return &UserRepo{
		col: col,
	}
}
