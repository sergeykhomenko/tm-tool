package project

import (
	"gorm.io/gorm"
)

type TicketEntity struct {
	gorm.Model
	ProjectId   uint
	Author      string
	Title       string `json:"title"`
	Description string `json:"description"`
	Assigned    string
}

func (t *Ticket) ToORM() TicketEntity {
	return TicketEntity{
		Title:       t.Title,
		Author:      t.Author,
		Description: t.Description,
		Assigned:    t.Assigned,
	}
}

func (t *TicketEntity) Proto() *Ticket {
	return &Ticket{
		Title:       t.Title,
		Author:      t.Author,
		Description: t.Description,
		Assigned:    t.Assigned,
	}
}
