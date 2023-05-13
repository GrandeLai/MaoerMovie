// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	AccessExpire string `json:"access_txpire"`
	RefreshAfter string `json:"refresh_after"`
}

type RegisterRequest struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	EmailCode string `json:"emailCode,optional"`
	PhoneCode string `json:"phoneCode,optional"`
}

type RegisterResponse struct {
	Status
	Id string `json:"id"`
}

type SendEmailCodeRequest struct {
	Email string `json:"email"`
}

type SendEmailCodeResponse struct {
	Status
}

type RefreshAuthResponse struct {
	Status
	AccessToken  string `json:"access_token"`
	AccessExpire string `json:"access_txpire"`
	RefreshAfter string `json:"refresh_after"`
}

type UpdatePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type UpdatePasswordResponse struct {
	Status
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type PageLimit struct {
	Page string `json:"page,optional"` //查询的第几页
	Size string `json:"size,optional"` //每页页数
}

type UserPreview struct {
	Id        string `json:"id"`
	Name      string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
}

type Role struct {
	RoleType    string `json:"role_type"`
	CommunityId string `json:"community_id,omitempty"`
}

type GetUserInfoResponse struct {
	Status
	Id        string `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
}

type UpdateUserInfoRequest struct {
	Name   string `form:"name"`
	Gender string `form:"gender"`
	Phone  string `form:"phone"`
}

type UpdateUserInfoResponse struct {
	Status
}

type SearchUserRequest struct {
	Keyword string    `json:"keyword"`
	Limit   PageLimit `json:"limit"`
}

type SearchUserResponse struct {
	Status
	Users []UserPreview `json:"users"`
	Total string        `json:"total"`
}

type FilmRequest struct {
	Id string `json:"id"`
}

type GetFilmResponse struct {
	Status
	FilmName        string `json:"film_name"`
	FilmEnglishName string `json:"film_english_name"`
	FilmType        string `json:"film_type"`
	FilmCover       string `json:"film_cover"`
	FilmLength      string `json:"film_length"`
	FilmCategory    string `json:"film_category"`
	FilmArea        string `json:"film_area"`
	FilmTime        string `json:"film_time"`
	Director        string `json:"director"`
	Biography       string `json:"biography"`
}

type FilmDetailResponse struct {
	Status
	FilmScore      float64 `json:"film_score"`
	FilmScoreNum   string  `json:"film_score_num"`
	FilmPreSaleNum string  `json:"film_pre_sale_num"`
	FilmBoxOffice  string  `json:"film_box_office"`
	FilmImgs       string  `json:"film_imgs"`
	ActorList      string  `json:"actor_list"`
}

type FilmResponse struct {
	Status
}

type FilmSearchRequest struct {
	PageLimit
	Category   string `json:"category"`
	Area       string `json:"area"`
	TimeStart  string `json:"time_start"`
	TimeEnd    string `json:"time_end"`
	Keyword    string `json:"keyword"`
	SortedType string `json:"sorted_type"` //0是热门，1是时间，2是评价
}

type FilmSearchResponse struct {
	Status
	InfoList []*FilmInfo `json:"info_list"`
	Count    string      `json:"count"`
}

type Film struct {
	FilmName        string `json:"film_name"`
	FilmEnglishName string `json:"film_english_name"`
	FilmType        string `json:"film_type"`
	FilmCover       string `json:"film_cover"`
	FilmLength      string `json:"film_length"`
	FilmCategory    string `json:"film_category"`
	FilmArea        string `json:"film_area"`
	FilmTime        string `json:"film_time"`
	Director        string `json:"director"`
}

type FilmInfo struct {
	FilmName      string `json:"film_name"`
	FilmTime      string `json:"film_time"`
	FilmCategory  string `json:"film_category"`
	FilmScore     string `json:"film_score"`
	FilmCover     string `json:"film_cover"`
	ActorNameList string `json:"actor_name_list"`
}

type FilmDeleteRequest struct {
	Id string `json:"id"`
}

type FilmNewRequest struct {
	FilmId          string `form:"film_id,optional"`
	FilmName        string `form:"film_name"`
	FilmEnglishName string `form:"film_english_name"`
	FilmType        string `form:"film_type"`
	FilmLength      string `form:"film_length"`
	FilmCategory    string `form:"film_category"`
	FilmArea        string `form:"film_area"`
	FilmTime        string `form:"film_time"`
	Director        string `form:"director"`
	Biography       string `form:"biography"`
	ActorList       string `form:"actor_list"`
	RoleList        string `form:"role_list"`
}

type FilmDetailAddRequest struct {
	FilmId         string `form:"film_id""`
	FilmPreSaleNum string `form:"film_pre_sale_num"`
}

type GetActorListResponse struct {
	Status
	ActorList []Actor `json:"actor_preview_list"`
	Count     string  `json:"count"`
}

type FilmCreateResponse struct {
	Status
	Id string `json:"id"`
}

type CommonResponse struct {
	Status
}

type FilmListRequest struct {
	Page string `json:"page,optional"` //查询的第几页
	Size string `json:"size,optional"` //每页页数
}

type AdminFilmListResponse struct {
	List  []*AdminFilm `json:"list"`
	Count string       `json:"count"`
	Status
}

type AdminFilm struct {
	FilmName        string `json:"film_name"`
	FilmEnglishName string `json:"film_english_name"`
	FilmType        string `json:"film_type"`
	FilmCover       string `json:"film_cover"`
	FilmLength      string `json:"film_length"`
	FilmCategory    string `json:"film_category"`
	FilmArea        string `json:"film_area"`
	FilmTime        string `json:"film_time"`
	Director        string `json:"director"`
}

type Actor struct {
	ActorId     string `json:"actor_id"`
	ActorName   string `json:"actor_name""`
	ActorAvatar string `json:"actor_avatar"`
}