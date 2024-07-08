package repository

import (
	"context"

	"personal_page/model"
	"personal_page/repository/elastic"

	"github.com/elastic/go-elasticsearch/v8"
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
	return nil
	panic("GetMovieRepository: no avaliable database")
}
