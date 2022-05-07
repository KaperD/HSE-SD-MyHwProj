package myhwproj

import (
	"gorm.io/gorm"
)

type HomeworkDao interface {
	AddHomework(newHomework NewHomework) Homework
	GetHomeworkById(homeworkId int64) *Homework
	GetHomeworks(offset int32, limit int32, isPublished bool) []Homework
}

type PostgresHomeworkDao struct {
	db *gorm.DB
}

// NewPostgresHomeworkDao creates a default postgres homework dao
func NewPostgresHomeworkDao(db *gorm.DB) HomeworkDao {
	return &PostgresHomeworkDao{db: db}
}

func (p *PostgresHomeworkDao) AddHomework(newHomework NewHomework) Homework {
	homework := Homework{
		Id:                  0,
		Name:                newHomework.Name,
		PublicationDatetime: newHomework.PublicationDatetime,
		Check:               newHomework.Check,
		Task:                newHomework.Task,
		Deadline:            newHomework.Deadline,
	}
	p.db.Create(&homework)
	return homework
}

func (p *PostgresHomeworkDao) GetHomeworkById(homeworkId int64) *Homework {
	var homework Homework
	result := p.db.First(&homework, homeworkId)
	if result.Error != nil {
		return nil
	}
	return &homework
}

func (p *PostgresHomeworkDao) GetHomeworks(offset int32, limit int32, isPublished bool) []Homework {
	var homeworks []Homework
	if isPublished {
		p.db.Limit(int(limit)).Offset(int(offset)).Where("publication_datetime <= now()").Order("deadline desc, id").Find(&homeworks)
	} else {
		p.db.Limit(int(limit)).Offset(int(offset)).Order("deadline desc, id").Find(&homeworks)
	}
	return homeworks
}
