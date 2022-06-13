package myhwproj

import (
	"gorm.io/gorm"
	"time"
)

// SubmissionDao - gives access to Submission objects
type SubmissionDao interface {
	AddSubmission(homeworkId int64, newSubmission NewSubmission) Submission
	GetHomeworkSubmissions(homeworkId int64, offset int32, limit int32) []Submission
	GetSubmissionById(submissionId int64) *Submission
	UpdateSubmission(submission Submission)
}

// PostgresSubmissionDao - gives access to Submission objects stored in Postgres DB
type PostgresSubmissionDao struct {
	db *gorm.DB
}

// NewPostgresSubmissionDao creates a default postgres submission dao
func NewPostgresSubmissionDao(db *gorm.DB) SubmissionDao {
	return &PostgresSubmissionDao{db: db}
}

// AddSubmission adds new Submission to the system
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

// GetHomeworkSubmissions returns Submission list of Homework with specified id
// Returns not more than limit elements skipping first offset elements
func (p *PostgresSubmissionDao) GetHomeworkSubmissions(homeworkId int64, offset int32, limit int32) []Submission {
	var submissions []Submission
	p.db.Limit(int(limit)).Offset(int(offset)).Where("homework_id = ?", homeworkId).Order("datetime desc, id").Find(&submissions)
	return submissions
}

// GetSubmissionById returns Submission by its id, or nil if there is no Submission with this id
func (p *PostgresSubmissionDao) GetSubmissionById(submissionId int64) *Submission {
	var submission Submission
	result := p.db.First(&submission, submissionId)
	if result.Error != nil {
		return nil
	}
	return &submission
}

// UpdateSubmission updates given Submission in the system (finds old version by id)
func (p *PostgresSubmissionDao) UpdateSubmission(submission Submission) {
	p.db.Save(&submission)
}
