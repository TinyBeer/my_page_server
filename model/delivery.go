package model

const (
	StatusOk    = "ok"
	StatusError = "error"
)

const (
	ErrorNoToken = "no token"
)

type Base struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username,omitempty" binding:"min=4,max=10"`
	Password string `json:"password,omitempty" binding:"min=3,max=18"`
}

type TokenResponse struct {
	Base
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type VideoListResponse struct {
	Base
	Videoes []VideoItem `json:"videoes,omitempty"`
}

type VideoItem struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Title string `json:"title"`
	Time  string `json:"time"`
	Name  string `json:"name"`
	Intro string `json:"intro"`
}

type MovieListResponse struct {
	Base
	Movies []MovieItem `json:"movies,omitempty"`
}

type MovieItem struct {
	ID     string   `json:"id,omitempty"`
	Post   string   `json:"post,omitempty"`
	Title  string   `json:"title,omitempty"`
	Tags   []string `json:"tags,omitempty"`
	Desc   string   `json:"desc,omitempty"`
	Source string   `json:"source,omitempty"`
}

type MovieCreateRequest struct {
	Post   string   `json:"post,omitempty"`
	Title  string   `json:"title,omitempty"`
	Tags   []string `json:"tags,omitempty"`
	Desc   string   `json:"desc,omitempty"`
	Source string   `json:"source,omitempty"`
}

type MovieDeleteRequest struct {
	ID string `json:"id,omitempty"`
}
