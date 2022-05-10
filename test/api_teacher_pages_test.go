package test

import (
	myhwproj "github.com/KaperD/HSE-SD-MyHwProj/internal"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestGetHomeworkPageTeacherOk(t *testing.T) {
	timeExample, err := time.Parse(time.RFC3339, "2022-05-01T17:32:28+03:00")
	require.Nil(t, err)
	initialHomeworks := []myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: timeExample.Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Task",
			Deadline:            timeExample.Add(3 * time.Hour),
		},
	}
	initialSubmissions := []myhwproj.Submission{
		{
			Id:         0,
			HomeworkId: 0,
			Datetime:   timeExample,
			Solution:   "1",
			Mark:       0,
			Comment:    "",
		},
		{
			Id:         1,
			HomeworkId: 0,
			Datetime:   timeExample.Add(time.Hour),
			Solution:   "2",
			Mark:       3,
			Comment:    "Hello",
		},
	}
	server := CreateServer(initialHomeworks, initialSubmissions)
	defer server.Close()

	data, code, err := Get(server, "/teacher/homeworks/0")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	expectedHtmlData, err := ioutil.ReadFile("./test/resources/get_homework_teacher.html")
	require.Nil(t, err)
	expectedHtml := string(expectedHtmlData)

	require.Equal(t, expectedHtml, string(data))
}

func TestGetHomeworksPageTeacherOk(t *testing.T) {
	timeExample, err := time.Parse(time.RFC3339, "2022-05-01T17:32:28+03:00")
	require.Nil(t, err)
	initialHomeworks := []myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: timeExample.Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Task",
			Deadline:            timeExample.Add(3 * time.Hour),
		},
		{
			Id:                  1,
			Name:                "Test2",
			PublicationDatetime: timeExample.Add(-time.Hour),
			Check:               "./gradlew test2",
			Task:                "Task2",
			Deadline:            timeExample.Add(4 * time.Hour),
		},
	}
	var initialSubmissions []myhwproj.Submission
	server := CreateServer(initialHomeworks, initialSubmissions)
	defer server.Close()

	data, code, err := Get(server, "/teacher/homeworks")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	expectedHtmlData, err := ioutil.ReadFile("./test/resources/get_homeworks_teacher.html")
	require.Nil(t, err)
	expectedHtml := string(expectedHtmlData)

	require.Equal(t, expectedHtml, string(data))
}

func TestGetSubmissionPageTeacherOk(t *testing.T) {
	timeExample, err := time.Parse(time.RFC3339, "2022-05-01T17:32:28+03:00")
	require.Nil(t, err)
	initialHomeworks := []myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: timeExample.Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Task",
			Deadline:            timeExample.Add(3 * time.Hour),
		},
	}
	initialSubmissions := []myhwproj.Submission{
		{
			Id:         0,
			HomeworkId: 0,
			Datetime:   timeExample,
			Solution:   "1",
			Mark:       0,
			Comment:    "",
		},
		{
			Id:         1,
			HomeworkId: 0,
			Datetime:   timeExample.Add(time.Hour),
			Solution:   "2",
			Mark:       3,
			Comment:    "Hello",
		},
	}
	server := CreateServer(initialHomeworks, initialSubmissions)
	defer server.Close()

	data, code, err := Get(server, "/teacher/submissions/1")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	expectedHtmlData, err := ioutil.ReadFile("./test/resources/get_submission_teacher.html")
	require.Nil(t, err)
	expectedHtml := string(expectedHtmlData)

	require.Equal(t, expectedHtml, string(data))
}

func TestCreateHomeworkPageTeacherOk(t *testing.T) {
	var initialHomeworks []myhwproj.Homework
	var initialSubmissions []myhwproj.Submission
	server := CreateServer(initialHomeworks, initialSubmissions)
	defer server.Close()

	data, code, err := Get(server, "/teacher/homeworks/create")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	expectedHtmlData, err := ioutil.ReadFile("./test/resources/get_create_homework_teacher.html")
	require.Nil(t, err)
	expectedHtml := string(expectedHtmlData)

	require.Equal(t, expectedHtml, string(data))
}
