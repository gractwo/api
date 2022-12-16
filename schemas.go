package main

import (
	"time"

	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	date time.Time
	tags []string
}

type QuoteLine struct {
	gorm.Model
	author  string
	content string
}
type Image struct {
	gorm.Model
	id          string
	title       string
	description string
	date        time.Time
	place       string
	link        string
}
