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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// StudentPagesApiController binds http requests to an api service and writes the service results to the http response
type StudentPagesApiController struct {
	service      StudentPagesApiServicer
	errorHandler ErrorHandler
}

// StudentPagesApiOption for how the controller is set up.
type StudentPagesApiOption func(*StudentPagesApiController)

// WithStudentPagesApiErrorHandler inject ErrorHandler into controller
func WithStudentPagesApiErrorHandler(h ErrorHandler) StudentPagesApiOption {
	return func(c *StudentPagesApiController) {
		c.errorHandler = h
	}
}

// NewStudentPagesApiController creates a default api controller
func NewStudentPagesApiController(s StudentPagesApiServicer, opts ...StudentPagesApiOption) Router {
	controller := &StudentPagesApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the StudentPagesApiController
func (c *StudentPagesApiController) Routes() Routes {
	return Routes{
		{
			"CreateSubmissionPageStudent",
			strings.ToUpper("Get"),
			"/student/homeworks/{homeworkId}/submissions/create",
			c.CreateSubmissionPageStudent,
		},
		{
			"GetHomeworkPageStudent",
			strings.ToUpper("Get"),
			"/student/homeworks/{homeworkId}",
			c.GetHomeworkPageStudent,
		},
		{
			"GetHomeworksPageStudent",
			strings.ToUpper("Get"),
			"/student/homeworks",
			c.GetHomeworksPageStudent,
		},
		{
			"GetSubmissionPageStudent",
			strings.ToUpper("Get"),
			"/student/submissions/{submissionId}",
			c.GetSubmissionPageStudent,
		},
	}
}

// CreateSubmissionPageStudent - Get creating submission page
func (c *StudentPagesApiController) CreateSubmissionPageStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	homeworkIdParam, err := parseInt64Parameter(params["homeworkId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.CreateSubmissionPageStudent(r.Context(), homeworkIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeHTMLResponse(result.Body, &result.Code, w)

}

// GetHomeworkPageStudent - Get homework page
func (c *StudentPagesApiController) GetHomeworkPageStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	homeworkIdParam, err := parseInt64Parameter(params["homeworkId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	pageParam, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetHomeworkPageStudent(r.Context(), homeworkIdParam, pageParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeHTMLResponse(result.Body, &result.Code, w)

}

// GetHomeworksPageStudent - Get homeworks page
func (c *StudentPagesApiController) GetHomeworksPageStudent(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pageParam, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetHomeworksPageStudent(r.Context(), pageParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeHTMLResponse(result.Body, &result.Code, w)

}

// GetSubmissionPageStudent - Get submission page
func (c *StudentPagesApiController) GetSubmissionPageStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	submissionIdParam, err := parseInt64Parameter(params["submissionId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetSubmissionPageStudent(r.Context(), submissionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeHTMLResponse(result.Body, &result.Code, w)

}
