package types

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

type TokenRefreshReq struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginResp struct {
	AccessToken  string              `json:"access_token"`
	RefreshToken string              `json:"refresh_token"`
	User         *UserWithParamsResp `json:"user"`
}

type JwtToken struct {
	UserID        int    `json:"uid"`
	AccessToken   string `json:"act"`
	RefreshToken  string `json:"rft"`
	AccessUuid    string `json:"aid"`
	RefreshUuid   string `json:"rid"`
	AccessExpiry  int64  `json:"axp"`
	RefreshExpiry int64  `json:"rxp"`
}
