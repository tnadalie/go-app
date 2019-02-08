package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Client struct {
	gorm.Model
	Title    string `gorm:"unique" json:"title"`
	Archived bool   `json:"archived"`
}

func (p *Client) Archive() {
	p.Archived = true
}

func (p *Client) Restore() {
	p.Archived = false
}

type Task struct {
	gorm.Model
	Title    string     `json:"title"`
	Priority string     `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline *time.Time `gorm:"default:null" json:"deadline"`
	Done     bool       `json:"done"`
	ClientID uint       `json:"project_id"`
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Undo() {
	t.Done = false
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Client{}, &Task{})
	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
