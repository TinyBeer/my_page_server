package mongo

import (
	"context"
	"time"

	"personal_page/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Memo struct {
	ID        string `bson:"_id,omitempty"`
	Content   string `bson:"content,omitempty"`
	Completed bool   `bson:"completed,omitempty"`
	CreatedAt int64  `bson:"created_at,omitempty"`
	UpdatedAt int64  `bson:"updated_at,omitempty"`
	DeletedAt int64  `bson:"deleted_at,omitempty"`
}

func (m Memo) toModelMemo() *model.Memo {
	return &model.Memo{
		ID:        m.ID,
		Content:   m.Content,
		Completed: m.Completed,
	}
}

type MemoRepo struct {
	col *mongo.Collection
}

// CompleteWithId implements MemoRepository.
func (r *MemoRepo) CompleteWithId(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: oid},
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "deleted_at", Value: bson.D{{Key: "$exists", Value: false}}}},
				bson.D{{Key: "deleted_at", Value: 0}},
			}},
		},

		bson.M{"$set": bson.M{"completed": true, "updated_at": time.Now().Unix()}})
	return err
}

// ListMemo implements MemoRepository.
func (r *MemoRepo) ListMemo() ([]*model.Memo, error) {
	cursor, err := r.col.Find(context.Background(),
		bson.D{{"$or", bson.A{
			bson.D{{"deleted_at", bson.D{{"$exists", false}}}},
			bson.D{{"deleted_at", 0}},
		}}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	var list []*model.Memo
	memo := &Memo{}
	for cursor.Next(context.Background()) {
		err = cursor.Decode(memo)
		if err != nil {
			return nil, err
		}
		list = append(list, memo.toModelMemo())
	}
	return list, nil
}

// CreateMemo implements MemoRepository.
func (r *MemoRepo) CreateMemo(content string) error {
	now := time.Now()
	m := Memo{
		Content:   content,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}
	_, err := r.col.InsertOne(context.Background(), m)
	return err
}

// DeleteMemoById implements MemoRepository.
func (r *MemoRepo) DeleteMemoById(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.col.UpdateByID(context.Background(),
		oid,
		bson.M{"$set": bson.M{"deleted_at": time.Now().Unix()}},
	)
	return err
}

func NewMemoRepo(db *mongo.Database) *MemoRepo {
	return &MemoRepo{
		col: db.Collection("memo"),
	}
}
