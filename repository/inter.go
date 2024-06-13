package repository

import (
	"context"

	"personal_page/model"
	"personal_page/repository/elastic"
	mongor "personal_page/repository/mongo"
	"personal_page/repository/mysql"

	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(name, password string) error
	GetUserByName(name string) (*model.User, error)
}

var (
	_ UserRepository = &mysql.UserRepo{}
	_ MemoRepository = &mysql.MemoRepo{}
)

type MemoRepository interface {
	ListMemo() ([]*model.Memo, error)
	CreateMemo(content string) error
	DeleteMemoById(id string) error
	CompleteWithId(id string) error
}

var (
	_ UserRepository = &mongor.UserRepo{}
	_ MemoRepository = &mongor.MemoRepo{}
)

type MovieRepository interface {
	List(ctx context.Context) ([]*model.Movie, error)
	Create(ctx context.Context, movie *model.Movie) error
	DeleteByID(ctx context.Context, id string) error
}

var _ MovieRepository = new(elastic.MovieRepo)

func GetMovieRepository(es *elasticsearch.TypedClient) MovieRepository {
	if es != nil {
		return elastic.NewMovieRepo(es)
	}
	panic("GetMovieRepository: no avaliable database")
}

func GetUserRepository(sql *gorm.DB, mongo *mongo.Database) UserRepository {
	if sql != nil {
		return mysql.NewUserRepo(sql)
	}

	if mongo != nil {
		return mongor.NewUserRepo(mongo)
	}

	panic("GetUserRepository: no avaliable database")
}

func GetMemoRepository(sql *gorm.DB, mongo *mongo.Database) MemoRepository {
	if sql != nil {
		return mysql.NewMemoRepo(sql)
	}

	if mongo != nil {
		return mongor.NewMemoRepo(mongo)
	}
	panic("GetMemoRepository: no avaliable database")
}
