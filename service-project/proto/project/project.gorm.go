package project

import (
	"gorm.io/gorm"
)

type ProjectEntity struct {
	gorm.Model
	ID          uint64         `gorm:"primarykey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Tickets     []TicketEntity `gorm:"foreignKey:ProjectId"`
}

func (p *Project) ToORM() (ProjectEntity, error) {
	model := ProjectEntity{
		Title:       p.Title,
		Description: p.Description,
		Tickets:     nil,
	}

	for _, ticket := range p.Tickets {
		model.Tickets = append(model.Tickets, ticket.ToORM())
	}

	return model, nil
}

func (p *ProjectEntity) Proto() *Project {
	proto := Project{
		Id:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Tickets:     nil,
	}

	for _, ticket := range p.Tickets {
		proto.Tickets = append(proto.Tickets, ticket.Proto())
	}

	return &proto
}
