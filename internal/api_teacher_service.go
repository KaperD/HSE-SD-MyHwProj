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
	"net/http"
	"errors"
)

// TeacherApiService is a service that implements the logic for the TeacherApiServicer
// This service should implement the business logic for every endpoint for the TeacherApi API.
// Include any external packages or services that will be required by this service.
type TeacherApiService struct {
}

// NewTeacherApiService creates a default api service
func NewTeacherApiService() TeacherApiServicer {
	return &TeacherApiService{}
}

// AddHomeworkTeacher - Add new homework
func (s *TeacherApiService) AddHomeworkTeacher(ctx context.Context, newHomework NewHomework) (ImplResponse, error) {
	// TODO - update AddHomeworkTeacher with the required logic for this service method.
	// Add api_teacher_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Homework{}) or use other options such as http.Ok ...
	//return Response(200, Homework{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddHomeworkTeacher method not implemented")
}

// GetHomeworkByIdTeacher - Get homework
func (s *TeacherApiService) GetHomeworkByIdTeacher(ctx context.Context, homeworkId int64) (ImplResponse, error) {
	// TODO - update GetHomeworkByIdTeacher with the required logic for this service method.
	// Add api_teacher_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Homework{}) or use other options such as http.Ok ...
	//return Response(200, Homework{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHomeworkByIdTeacher method not implemented")
}

// GetHomeworkSubmissionsTeacher - Get homework submissions
func (s *TeacherApiService) GetHomeworkSubmissionsTeacher(ctx context.Context, homeworkId int64, offset int32, limit int32) (ImplResponse, error) {
	// TODO - update GetHomeworkSubmissionsTeacher with the required logic for this service method.
	// Add api_teacher_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Submission{}) or use other options such as http.Ok ...
	//return Response(200, []Submission{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHomeworkSubmissionsTeacher method not implemented")
}

// GetHomeworksTeacher - Get homeworks
func (s *TeacherApiService) GetHomeworksTeacher(ctx context.Context, offset int32, limit int32) (ImplResponse, error) {
	// TODO - update GetHomeworksTeacher with the required logic for this service method.
	// Add api_teacher_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Homework{}) or use other options such as http.Ok ...
	//return Response(200, []Homework{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHomeworksTeacher method not implemented")
}

// GetSubmissionTeacher - Get submission
func (s *TeacherApiService) GetSubmissionTeacher(ctx context.Context, submissionId int64) (ImplResponse, error) {
	// TODO - update GetSubmissionTeacher with the required logic for this service method.
	// Add api_teacher_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Submission{}) or use other options such as http.Ok ...
	//return Response(200, Submission{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSubmissionTeacher method not implemented")
}
