package model

type LoginResult struct {
	AccessToken  string
	RefreshToken string
}

type User struct {
	ID       string
	Name     string
	Password string
}

type Memo struct {
	ID        string
	Content   string
	Completed bool
}

func (m *Memo) ToMemoItem() MemoItem {
	return MemoItem{
		ID:        m.ID,
		Content:   m.Content,
		Completed: m.Completed,
	}
}

type Movie struct {
	ID     string
	Post   string
	Title  string
	Tags   []string
	Desc   string
	Source string
}

func (m *Movie) ToMovieItem() MovieItem {
	return MovieItem{
		ID:     m.ID,
		Post:   m.Post,
		Title:  m.Title,
		Tags:   m.Tags,
		Desc:   m.Desc,
		Source: m.Source,
	}
}
