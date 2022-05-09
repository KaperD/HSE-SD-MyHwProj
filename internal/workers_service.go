package myhwproj

import (
	"time"
)

type WorkersService interface {
	SetHandler(func(submission Submission))
	CheckSubmission(submission Submission)
}

type RabbitMQWorkersService struct {
	ResultHandler func(submission Submission)
}

func NewRabbitMQWorkersService() WorkersService {
	return &RabbitMQWorkersService{ResultHandler: func(submission Submission) {}}
}

func (r *RabbitMQWorkersService) SetHandler(f func(submission Submission)) {
	r.ResultHandler = f
}

func (r *RabbitMQWorkersService) CheckSubmission(submission Submission) {
	go func() {
		time.Sleep(time.Second * 4)
		submission.Mark = 5
		submission.Comment = "Checked"
		r.ResultHandler(submission)
	}()
}
