package myhwproj

import (
	"gorm.io/gorm"
)

// HomeworkDao - gives access to Homework objects
type HomeworkDao interface {
	AddHomework(newHomework NewHomework) Homework
	GetHomeworkById(homeworkId int64) *Homework
	GetHomeworks(offset int32, limit int32, onlyPublished bool) []Homework
}

// PostgresHomeworkDao - gives access to Homework objects stored in Postgres DB
type PostgresHomeworkDao struct {
	db *gorm.DB
}

// NewPostgresHomeworkDao creates a default postgres homework dao
func NewPostgresHomeworkDao(db *gorm.DB) HomeworkDao {
	return &PostgresHomeworkDao{db: db}
}

// AddHomework adds new Homework to the system
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

// GetHomeworkById returns Homework by its id, or nil if there is no Homework with this id
func (p *PostgresHomeworkDao) GetHomeworkById(homeworkId int64) *Homework {
	var homework Homework
	result := p.db.First(&homework, homeworkId)
	if result.Error != nil {
		return nil
	}
	return &homework
}

// GetHomeworks returns Homework list of not more than limit elements skipping first offset elements.
// If onlyPublished is true, works only with already published Homeworks
func (p *PostgresHomeworkDao) GetHomeworks(offset int32, limit int32, onlyPublished bool) []Homework {
	var homeworks []Homework
	if onlyPublished {
		p.db.Limit(int(limit)).Offset(int(offset)).Where("publication_datetime <= now()").Order("deadline desc, id").Find(&homeworks)
	} else {
		p.db.Limit(int(limit)).Offset(int(offset)).Order("deadline desc, id").Find(&homeworks)
	}
	return homeworks
}
