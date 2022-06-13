/*
 * MyHwProj
 *
 * REST api for MyHwProj
 *
 * API version: 1.0.0
 * Contact: myhwproj@yandex.ru
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package myhwproj

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// StudentApiService is a service that implements the logic for the StudentApiServicer
// This service should implement the business logic for every endpoint for the StudentApi API.
// Include any external packages or services that will be required by this service.
type StudentApiService struct {
	SubmissionDao  SubmissionDao
	HomeworkDao    HomeworkDao
	WorkersService WorkersService
}

// NewStudentApiService creates a default api service
func NewStudentApiService(
	submissionDao SubmissionDao,
	homeworkDao HomeworkDao,
	workersService WorkersService,
) StudentApiServicer {
	result := StudentApiService{SubmissionDao: submissionDao, HomeworkDao: homeworkDao, WorkersService: workersService}
	result.WorkersService.SetHandler(func(submission Submission) {
		result.SubmissionDao.UpdateSubmission(submission)
	})
	return &result
}

// AddSubmissionStudent - Add new submission
func (s *StudentApiService) AddSubmissionStudent(_ context.Context, homeworkId int64, newSubmission NewSubmission) (ImplResponse[Submission], error) {
	if len(strings.TrimSpace(newSubmission.Solution)) == 0 {
		return Response(http.StatusBadRequest, Submission{}), errors.New("solution must not be blank")
	}
	if homework := s.HomeworkDao.GetHomeworkById(homeworkId); homework == nil || homework.PublicationDatetime.After(time.Now()) {
		return Response(http.StatusNotFound, Submission{}), errors.New(fmt.Sprintf("homework with id %d not found", homeworkId))
	}
	submission := s.SubmissionDao.AddSubmission(homeworkId, newSubmission)
	s.WorkersService.CheckSubmission(submission)
	return Response(http.StatusOK, submission), nil
}

// GetHomeworkByIdStudent - Get homework
func (s *StudentApiService) GetHomeworkByIdStudent(_ context.Context, homeworkId int64) (ImplResponse[Homework], error) {
	homework := s.HomeworkDao.GetHomeworkById(homeworkId)
	if homework == nil || homework.PublicationDatetime.After(time.Now()) {
		return Response(http.StatusNotFound, Homework{}), errors.New(fmt.Sprintf("homework with id %d not found", homeworkId))
	}
	return Response(http.StatusOK, *homework), nil
}

// GetHomeworkSubmissionsStudent - Get homework submissions
func (s *StudentApiService) GetHomeworkSubmissionsStudent(_ context.Context, homeworkId int64, offset int32, limit int32) (ImplResponse[[]Submission], error) {
	if offset < 0 || limit < 0 {
		return Response[[]Submission](http.StatusBadRequest, nil), errors.New("offset and limit must be non negative")
	}
	if homework := s.HomeworkDao.GetHomeworkById(homeworkId); homework == nil || homework.PublicationDatetime.After(time.Now()) {
		return Response[[]Submission](http.StatusNotFound, nil), errors.New(fmt.Sprintf("homework with id %d not found", homeworkId))
	}
	submissions := s.SubmissionDao.GetHomeworkSubmissions(homeworkId, offset, limit)
	return Response(http.StatusOK, submissions), nil
}

// GetHomeworksStudent - Get homeworks
func (s *StudentApiService) GetHomeworksStudent(_ context.Context, offset int32, limit int32) (ImplResponse[[]Homework], error) {
	if offset < 0 || limit < 0 {
		return Response[[]Homework](http.StatusBadRequest, nil), errors.New("offset and limit must be non negative")
	}
	homeworks := s.HomeworkDao.GetHomeworks(offset, limit, true)
	return Response(http.StatusOK, homeworks), nil
}

// GetSubmissionStudent - Get submission
func (s *StudentApiService) GetSubmissionStudent(_ context.Context, submissionId int64) (ImplResponse[Submission], error) {
	submission := s.SubmissionDao.GetSubmissionById(submissionId)
	if submission == nil {
		return Response(http.StatusNotFound, Submission{}), errors.New(fmt.Sprintf("submission with id %d not found", submissionId))
	}
	return Response(http.StatusOK, *submission), nil
}
