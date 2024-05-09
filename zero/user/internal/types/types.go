// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username" require:"true"`
	Password string `json:"password" require:"true"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token" require:"true"`
	RefreshToken string `json:"refresh_token" require:"true"`
}

type RefreshReq struct {
	RefreshToken string `json:"refresh_token" require:"true"`
}

type RefreshResp struct {
	AccessToken string `json:"access_token;omitempty" require:"true"`
}
