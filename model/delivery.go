package model

type SimpleResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username,omitempty" binding:"min=4,max=10"`
	Password string `json:"password,omitempty" binding:"min=3,max=18"`
}

type LoginResponse struct {
	Status       string `json:"status,omitempty"`
	Message      string `json:"message,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type VideoListResponse struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Videoes []VideoItem `json:"videoes,omitempty"`
}

type VideoItem struct {
	Image string `json:"image"`
	Title string `json:"title"`
	Time  string `json:"time"`
	Name  string `json:"name"`
	Intro string `json:"intro"`
}

type MemoListResponse struct {
	Status  string     `json:"status,omitempty"`
	Message string     `json:"message,omitempty"`
	Memoes  []MemoItem `json:"memoes,omitempty"`
}

type MemoItem struct {
	Id      uint   `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type MemoCreateRequest struct {
	Content string `json:"content,omitempty" binding:"min=2,max=255"`
}

type MemoUpdateRequest struct {
	Id      uint   `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type MemoDeleteRequest struct {
	Id uint `json:"id,omitempty"`
}
