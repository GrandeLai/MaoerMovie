package kqueue

// 用户信息修改消息
type UserUpdateMessage struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type UserInsertMessage struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type FilmInsertMessage struct {
	FilmId        int64
	FilmName      string
	FilmTime      string
	FilmCategory  string
	FilmArea      string
	FilmCoverUrl  string
	ActorNameList string
}

type FilmUpdateMessage struct {
	FilmId        int64
	FilmName      string
	FilmTime      string
	FilmCategory  string
	FilmArea      string
	FilmCoverUrl  string
	ActorNameList string
	FilmScore     string
}

type ActorInsert struct {
	ActorList string
	FilmId    int64
	RoleList  string
}
