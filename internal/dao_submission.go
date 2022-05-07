package myhwproj

import (
	"gorm.io/gorm"
	"time"
)

type SubmissionDao interface {
	AddSubmission(homeworkId int64, newSubmission NewSubmission) Submission
	GetHomeworkSubmissions(homeworkId int64, offset int32, limit int32) []Submission
	GetSubmissionById(submissionId int64) *Submission
	UpdateSubmission(submission Submission)
}

type PostgresSubmissionDao struct {
	db *gorm.DB
}

// NewPostgresSubmissionDao creates a default postgres submission dao
func NewPostgresSubmissionDao(db *gorm.DB) SubmissionDao {
	return &PostgresSubmissionDao{db: db}
}

func (p *PostgresSubmissionDao) AddSubmission(homeworkId int64, newSubmission NewSubmission) Submission {
	submission := Submission{
		Id:         0,
		HomeworkId: homeworkId,
		Datetime:   time.Now(),
		Solution:   newSubmission.Solution,
		Mark:       0,
		Comment:    "",
	}
	p.db.Create(&submission)
	return submission
}

func (p *PostgresSubmissionDao) GetHomeworkSubmissions(homeworkId int64, offset int32, limit int32) []Submission {
	var submissions []Submission
	p.db.Limit(int(limit)).Offset(int(offset)).Where("homework_id = ?", homeworkId).Order("datetime desc, id").Find(&submissions)
	return submissions
}

func (p *PostgresSubmissionDao) GetSubmissionById(submissionId int64) *Submission {
	var submission Submission
	result := p.db.First(&submission, submissionId)
	if result.Error != nil {
		return nil
	}
	return &submission
}

func (p *PostgresSubmissionDao) UpdateSubmission(submission Submission) {
	p.db.Save(&submission)
}
