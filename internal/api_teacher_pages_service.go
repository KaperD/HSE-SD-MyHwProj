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

// TeacherPagesApiService is a service that implements the logic for the TeacherPagesApiServicer
// This service should implement the business logic for every endpoint for the TeacherPagesApi API.
// Include any external packages or services that will be required by this service.
type TeacherPagesApiService struct {
}

// NewTeacherPagesApiService creates a default api service
func NewTeacherPagesApiService() TeacherPagesApiServicer {
	return &TeacherPagesApiService{}
}

// CreateHomeworkPageTeacher - Get creating homework page
func (s *TeacherPagesApiService) CreateHomeworkPageTeacher(ctx context.Context) (ImplResponse, error) {
	// TODO - update CreateHomeworkPageTeacher with the required logic for this service method.
	// Add api_teacher_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateHomeworkPageTeacher method not implemented")
}

// GetHomeworkPageTeacher - Get homework page
func (s *TeacherPagesApiService) GetHomeworkPageTeacher(ctx context.Context, homeworkId int64, page int32) (ImplResponse, error) {
	// TODO - update GetHomeworkPageTeacher with the required logic for this service method.
	// Add api_teacher_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHomeworkPageTeacher method not implemented")
}

// GetHomeworksPageTeacher - Get homeworks page
func (s *TeacherPagesApiService) GetHomeworksPageTeacher(ctx context.Context, page int32) (ImplResponse, error) {
	// TODO - update GetHomeworksPageTeacher with the required logic for this service method.
	// Add api_teacher_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHomeworksPageTeacher method not implemented")
}

// GetSubmissionPageTeacher - Get submission page
func (s *TeacherPagesApiService) GetSubmissionPageTeacher(ctx context.Context, submissionId int64) (ImplResponse, error) {
	// TODO - update GetSubmissionPageTeacher with the required logic for this service method.
	// Add api_teacher_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSubmissionPageTeacher method not implemented")
}
