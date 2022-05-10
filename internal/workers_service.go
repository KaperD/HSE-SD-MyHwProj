package myhwproj

import (
	"time"
)

// WorkersService defines the api for service which works with Workers who check Submissions
type WorkersService interface {
	SetHandler(func(submission Submission))
	CheckSubmission(submission Submission)
}

// RabbitMQWorkersService works with Workers through RabbitMQ
type RabbitMQWorkersService struct {
	ResultHandler func(submission Submission)
}

// NewRabbitMQWorkersService creates default RabbitMQWorkersService
func NewRabbitMQWorkersService() WorkersService {
	return &RabbitMQWorkersService{ResultHandler: func(submission Submission) {}}
}

// SetHandler changes the handler of check result
func (r *RabbitMQWorkersService) SetHandler(f func(submission Submission)) {
	r.ResultHandler = f
}

// CheckSubmission checks the submission async. After check calls handler on the result
func (r *RabbitMQWorkersService) CheckSubmission(submission Submission) {
	go func() {
		time.Sleep(time.Second * 4)
		submission.Mark = 5
		submission.Comment = "Checked"
		r.ResultHandler(submission)
	}()
}
