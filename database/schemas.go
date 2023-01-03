package database

import (
	"time"

	"github.com/lib/pq"
)

type Quote struct {
	ID   string `gorm:"primaryKey"`
	Date *time.Time
	Tags pq.StringArray `gorm:"type:varchar(64)[]"`
}

type QuoteLine struct {
	Author  string
	Content string
}
type Image struct {
	Id          string
	Title       string
	Description string
	Date        *time.Time
	Place       string
	Link        string
}

type Admincard struct {
	Id           string
	Name         string
	Desc         string
	Img          string
	DevBadge     bool
	Admincard    bool
	AssignedUser string
}
type Splash struct {
	Id     string
	Splash string
}
