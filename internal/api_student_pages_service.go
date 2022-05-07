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
	"bytes"
	"context"
	"errors"
	"html/template"
	"log"
	"net/http"
)

// StudentPagesApiService is a service that implements the logic for the StudentPagesApiServicer
// This service should implement the business logic for every endpoint for the StudentPagesApi API.
// Include any external packages or services that will be required by this service.
type StudentPagesApiService struct {
	StudentApiService StudentApiServicer
}

// NewStudentPagesApiService creates a default api service
func NewStudentPagesApiService(studentApiService StudentApiServicer) StudentPagesApiServicer {
	return &StudentPagesApiService{StudentApiService: studentApiService}
}

// CreateSubmissionPageStudent - Get creating submission page
func (s *StudentPagesApiService) CreateSubmissionPageStudent(ctx context.Context, homeworkId int64) (ImplResponse[string], error) {
	result, err := s.StudentApiService.GetHomeworkByIdStudent(ctx, homeworkId)
	var homework *Homework
	if err == nil {
		homework = &result.Body
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = ts.Execute(buf, homework)
	if err != nil {
		log.Fatal(err)
	}
	return Response(http.StatusOK, buf.String()), nil
}

// GetHomeworkPageStudent - Get homework page
func (s *StudentPagesApiService) GetHomeworkPageStudent(ctx context.Context, homeworkId int64, page int32) (ImplResponse[string], error) {
	// TODO - update GetHomeworkPageStudent with the required logic for this service method.
	// Add api_student_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(http.StatusOK, {}) or use other options such as http.Ok ...
	//return Response(http.StatusOK, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, ""), errors.New("GetHomeworkPageStudent method not implemented")
}

// GetHomeworksPageStudent - Get homeworks page
func (s *StudentPagesApiService) GetHomeworksPageStudent(ctx context.Context, page int32) (ImplResponse[string], error) {
	// TODO - update GetHomeworksPageStudent with the required logic for this service method.
	// Add api_student_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(http.StatusOK, {}) or use other options such as http.Ok ...
	//return Response(http.StatusOK, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, ""), errors.New("GetHomeworksPageStudent method not implemented")
}

// GetSubmissionPageStudent - Get submission page
func (s *StudentPagesApiService) GetSubmissionPageStudent(ctx context.Context, submissionId int64) (ImplResponse[string], error) {
	// TODO - update GetSubmissionPageStudent with the required logic for this service method.
	// Add api_student_pages_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(http.StatusOK, {}) or use other options such as http.Ok ...
	//return Response(http.StatusOK, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, ""), errors.New("GetSubmissionPageStudent method not implemented")
}
