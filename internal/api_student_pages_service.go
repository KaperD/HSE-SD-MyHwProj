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
	"net/http"
)

// StudentPagesApiService is a service that implements the logic for the StudentPagesApiServicer
// This service should implement the business logic for every endpoint for the StudentPagesApi API.
// Include any external packages or services that will be required by this service.
type StudentPagesApiService struct {
}

// NewStudentPagesApiService creates a default api service
func NewStudentPagesApiService() StudentPagesApiServicer {
	return &StudentPagesApiService{}
}

// CreateSubmissionPageStudent - Get creating submission page
func (s *StudentPagesApiService) CreateSubmissionPageStudent(ctx context.Context, homeworkId int64) (ImplResponse, error) {
	// TODO - update CreateSubmissionPageStudent with the required logic for this service method.
	// Add api_student_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateSubmissionPageStudent method not implemented")
}

// GetHomeworkPageStudent - Get homework page
func (s *StudentPagesApiService) GetHomeworkPageStudent(ctx context.Context, homeworkId int64, page int32) (ImplResponse, error) {
	// TODO - update GetHomeworkPageStudent with the required logic for this service method.
	// Add api_student_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHomeworkPageStudent method not implemented")
}

// GetHomeworksPageStudent - Get homeworks page
func (s *StudentPagesApiService) GetHomeworksPageStudent(ctx context.Context, page int32) (ImplResponse, error) {
	// TODO - update GetHomeworksPageStudent with the required logic for this service method.
	// Add api_student_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHomeworksPageStudent method not implemented")
}

// GetSubmissionPageStudent - Get submission page
func (s *StudentPagesApiService) GetSubmissionPageStudent(ctx context.Context, submissionId int64) (ImplResponse, error) {
	// TODO - update GetSubmissionPageStudent with the required logic for this service method.
	// Add api_student_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSubmissionPageStudent method not implemented")
}
