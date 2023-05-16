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

type CinemaListRequest struct {
	BrandId    string `json:"brand_id"`
	HallType   string `json:"hall_type"`
	DistrictId string `json:"district_id"`
	PageLimit
}

type CinemaListResponse struct {
	Status
	List  []CinemaPreview `json:"list"`
	Total string          `json:"total"`
}

type CinemaPreview struct {
	CinemaId   string `json:"cinema_id"`
	CinemaName string `json:"cinema_name"`
	Address    string `json:"address"`
	MinPrice   string `json:"min_price"`
}

type ConditionListRequest struct {
	CityName string `json:"city_name"`
}

type ConditionListResponse struct {
	Status
	BrandList    []BrandCondition    `json:"brand_list"`
	DistrictList []DistrictCondition `json:"district_list"`
	HallTypeList []HallType          `json:"hall_type_list"`
}

type BrandCondition struct {
	Id        string `json:"id"`
	BrandName string `json:"brand_name"`
}

type DistrictCondition struct {
	Id           string `json:"id"`
	DistrictName string `json:"district_name"`
}

type HallType struct {
	Id           string `json:"id"`
	HallTypeName string `json:"hall_type_name"`
}

type CinemaFilmListRequest struct {
	CinemaId string `json:"cinema_id"`
}

type CinemaFilmListResponse struct {
	Status
	FilmList      []CinemaFilm `json:"film_list"`
	CinemaName    string       `json:"cinema_name"`
	CinemaAddress string       `json:"cinema_address"`
	CinemaPhone   string       `json:"cinema_phone"`
	CinemaImgs    string       `json:"cinema_imgs"`
}

type CinemaFilm struct {
	FilmId       string `json:"film_d"`
	FilmName     string `json:"film_name"`
	FilmLength   string `json:"film_length"`
	FilmCover    string `json:"film_cover"`
	FilmCategory string `json:"film_category"`
	ActorList    string `json:"actor_list"`
}

type CinemaShowListRequest struct {
	CinemaId string `json:"cinema_id"`
	FilmId   string `json:"film_id"`
	Date     string `json:"date"`
}

type CinemaShowListResponse struct {
	Status
	ShowList []CinemaShow `json:"show_list"`
}

type CinemaShow struct {
	ShowId    string `json:"show_id"`
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
	Language  string `json:"language"`
	Price     string `json:"price"`
	HallName  string `json:"hall_name"`
}

type HallSeatInfoRequest struct {
	HallId string `json:"hall_id"`
}

type HallSeatInfoResponse struct {
	Status
	SeatInfo
	SoldSeats string `json:"sold_seats"`
}

type SeatInfo struct {
	AllSeats string   `json:"all_seats"`
	Single   [][]Seat `json:"single"`
	Couple   [][]Seat `json:"couple"`
}

type Seat struct {
	SeatId int `json:"seat_id"`
	Row    int `json:"row"`
	Column int `json:"column"`
}
