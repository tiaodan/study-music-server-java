package models

import "time"

type ConsumerRequest struct {
	ID          uint       `json:"id"`
	Username    string     `json:"username"`
	OldPassword string     `json:"old_password"`
	Password    string     `json:"password"`
	Sex         *uint8    `json:"sex"`
	PhoneNum    string     `json:"phone_num"`
	Email       string     `json:"email"`
	Birth       *time.Time `json:"birth"`
	Introduction string   `json:"introduction"`
	Location    string     `json:"location"`
	Avator      string     `json:"avator"`
	CreateTime  time.Time  `json:"create_time"`
}

type SingerRequest struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Sex         *uint8    `json:"sex"`
	Pic         string     `json:"pic"`
	Birth       *time.Time `json:"birth"`
	Location    string     `json:"location"`
	Introduction string   `json:"introduction"`
}

type SongRequest struct {
	ID          uint   `json:"id"`
	SingerId    uint   `json:"singer_id"`
	Name        string `json:"name"`
	Introduction string `json:"introduction"`
	Pic         string `json:"pic"`
	Lyric       string `json:"lyric"`
	Url         string `json:"url"`
}

type SongListRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Pic         string `json:"pic"`
	Introduction string `json:"introduction"`
	Style       string `json:"style"`
}

type CollectRequest struct {
	ID     uint `json:"id"`
	UserId uint `json:"user_id"`
	SongId uint `json:"song_id"`
	Type   *uint8 `json:"type"`
}

type CommentRequest struct {
	ID         uint   `json:"id"`
	UserId     uint   `json:"user_id"`
	SongId     uint   `json:"song_id"`
	SongListId *uint  `json:"song_list_id"`
	Content    string `json:"content"`
	Type       *uint8 `json:"type"`
}

type RankListRequest struct {
	ID         uint    `json:"id"`
	SongListId uint    `json:"song_list_id"`
	ConsumerId uint    `json:"consumer_id"`
	Score      float64 `json:"score"`
}

type AdminRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ListSongRequest struct {
	ID         uint `json:"id"`
	SongId     uint `json:"song_id"`
	SongListId uint `json:"song_list_id"`
}

type UserSupportRequest struct {
	ID        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	CommentId uint   `json:"comment_id"`
	Type      *uint8 `json:"type"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}
