package elastic

import (
	"context"
	"encoding/json"
	"errors"

	"personal_page/model"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	indexName = "movie"
)

type RepoMovie struct {
	ID     string   `json:"_id,omitempty"`
	Post   string   `json:"post,omitempty"`
	Title  string   `json:"title,omitempty"`
	Tags   []string `json:"tags,omitempty"`
	Desc   string   `json:"desc,omitempty"`
	Source string   `json:"source,omitempty"`
}

func (m *RepoMovie) toModelMovie() *model.Movie {
	return &model.Movie{
		ID:     m.ID,
		Post:   m.Post,
		Title:  m.Title,
		Tags:   m.Tags,
		Desc:   m.Desc,
		Source: m.Source,
	}
}

func fromModelMovietoRepoMovie(m *model.Movie) *RepoMovie {
	return &RepoMovie{
		ID:     m.ID,
		Post:   m.Post,
		Title:  m.Title,
		Tags:   m.Tags,
		Desc:   m.Desc,
		Source: m.Source,
	}
}

type MovieRepo struct {
	client *elasticsearch.TypedClient
}

func NewMovieRepo(client *elasticsearch.TypedClient) *MovieRepo {
	return &MovieRepo{
		client: client,
	}
}

// Create implements repository.MovieRepository.
func (m *MovieRepo) Create(ctx context.Context, mm *model.Movie) error {
	movie := fromModelMovietoRepoMovie(mm)
	countResp, err := m.client.Count().Index(indexName).
		Query(&types.Query{
			MatchPhrase: map[string]types.MatchPhraseQuery{
				"source": {Query: movie.Source},
			},
		}).Do(ctx)

	if err == nil {
		if countResp.Count != 0 {
			return errors.New("duplicate source")
		}
	}
	if err != nil {
		if e, ok := err.(*types.ElasticsearchError); ok {
			if e.Status != 404 {
				return err
			}
		} else {
			return err
		}
	}

	_, err = m.client.
		Index(indexName).
		Document(movie).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID implements repository.MovieRepository.
func (m *MovieRepo) DeleteByID(ctx context.Context, id string) error {
	_, err := m.client.Delete(indexName, id).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

// List implements repository.MovieRepository.
func (m *MovieRepo) List(ctx context.Context) ([]*model.Movie, error) {
	resp, err := m.client.
		Search().
		Index(indexName).
		Do(ctx)
	if err != nil {
		if e, ok := err.(*types.ElasticsearchError); ok {
			if e.Status == 404 {
				return nil, nil
			}
		}
		return nil, err
	}
	list := make([]*model.Movie, 0, len(resp.Hits.Hits))

	for _, hit := range resp.Hits.Hits {
		mv := new(RepoMovie)
		err = json.Unmarshal(hit.Source_, mv)
		if err != nil {
			return nil, err
		}
		mv.ID = *hit.Id_
		list = append(list, mv.toModelMovie())
	}
	return list, nil
}
