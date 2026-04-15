package models

import (
	"time"

	"gorm.io/gorm"
)

// ModelRegistry holds all registered models
var ModelRegistry = []interface{}{}

// RegisterModel adds a model to the registry
func RegisterModel(model interface{}) {
	ModelRegistry = append(ModelRegistry, model)
}

func init() {
	// Register all models
	RegisterModel(&News{})
	RegisterModel(&NewsLike{})
	RegisterModel(&NewsViewing{})
	RegisterModel(&NewsAnalytics{})
	RegisterModel(&User{})
	RegisterModel(&NewsGroup{})
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type NewsGroup struct {
	gorm.Model
	Title string `gorm:"type:text" json:"title"`
}
type News struct {
	gorm.Model
	Title            string `gorm:"type:text" json:"title"`
	MessageDate      string `gorm:"size:20" json:"messageDate"`
	Message          string `gorm:"type:text" json:"message"`
	Type             int    `gorm:"default:1" json:"type"`
	MediaLink        string `gorm:"size:255" json:"mediaLink"`
	DownloadLink     string `gorm:"size:255" json:"downloadLink"`
	StatusId         int    `gorm:"default:0" json:"statusId"`
	AuthorId         int    `gorm:"default:0" json:"authorId"`
	Liked            int    `gorm:"default:0" json:"liked"`
	Disliked         int    `gorm:"default:0" json:"disliked"`
	ViewCount        int    `gorm:"default:0" json:"viewCount"`
	PushNotification int    `gorm:"default:0" json:"pushNotification"`
	GroupId          int    `gorm:"default:0" json:"groupId"`
}

type NewsLike struct {
	gorm.Model
	NewsId int `gorm:"default:0" json:"newsId"`
	UserId int `gorm:"default:0" json:"userId"`
	Type   int `gorm:"default:1" json:"type"`
}

type NewsViewing struct {
	gorm.Model
	NewsId int `gorm:"default:0" json:"newsId"`
	UserId int `gorm:"default:0" json:"userId"`
}

type NewsAnalytics struct {
	NewsId   int `gorm:"default:0" json:"newsid"`
	Liked    int `gorm:"default:0" json:"liked"`
	Disliked int `gorm:"default:0" json:"disliked"`
	Viewed   int `gorm:"default:0" json:"viewed"`
}

type User struct {
	gorm.Model
	Name      string `gorm:"size:255" json:"name"`
	Login     string `gorm:"size:255" json:"login"`
	Password  string `gorm:"size:255" json:"password"`
	Role      string `gorm:"size:255" json:"role"`
	Email     string `gorm:"size:255" json:"email"`
	Status    string `gorm:"size:255" json:"status"`
	LastLogin time.Time
	ImageUrl  string `gorm:"size:255" json:"imageUrl"`
}

type CallProcedureResult struct {
	value int `gorm:"column:value" json:"value"`
}
