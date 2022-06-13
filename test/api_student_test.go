package test

import (
	"encoding/json"
	myhwproj "github.com/KaperD/HSE-SD-MyHwProj/internal"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestAddSubmissionStudentOk(t *testing.T) {
	server := CreateServer([]myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(time.Hour),
		},
	}, []myhwproj.Submission{})
	defer server.Close()

	solution := "https://github.com/KaperD/HSE-SD-Course-Roguelike"
	data, code, err := Post(server, "/v1/student/homeworks/0/submissions", myhwproj.NewSubmission{Solution: solution})
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	var submission myhwproj.Submission
	require.Nil(t, json.Unmarshal(data, &submission))

	require.Equal(t, int64(0), submission.Id)
	require.Equal(t, solution, submission.Solution)
	require.Equal(t, int64(0), submission.HomeworkId)
}

func TestAddSubmissionStudentHomeworkNotFound(t *testing.T) {
	server := CreateServer([]myhwproj.Homework{}, []myhwproj.Submission{})
	defer server.Close()

	_, code, err := Post(server, "/v1/student/homeworks/0/submissions", myhwproj.NewSubmission{Solution: "https://github.com/KaperD/HSE-SD-Course-Roguelike"})
	require.Nil(t, err)

	require.Equal(t, http.StatusNotFound, code)
}

func TestAddSubmissionStudentWrongJson(t *testing.T) {
	server := CreateServer([]myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(time.Hour),
		},
	}, []myhwproj.Submission{})
	defer server.Close()

	_, code, err := Post(server, "/v1/student/homeworks/0/submissions", "wrong")
	require.Nil(t, err)

	require.Equal(t, http.StatusBadRequest, code)
}

func TestGetHomeworksStudentOk(t *testing.T) {
	initialHomeworks := []myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(time.Hour),
		},
		{
			Id:                  1,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(3 * time.Hour),
		},
		{
			Id:                  2,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(2 * time.Hour),
		},
	}
	server := CreateServer(initialHomeworks, []myhwproj.Submission{})
	defer server.Close()

	data, code, err := Get(server, "/v1/student/homeworks?offset=0&limit=10")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	homeworks := make([]myhwproj.Homework, 0, 10)
	err = json.Unmarshal(data, &homeworks)
	require.Nil(t, err)

	expectedHomeworks := []myhwproj.Homework{initialHomeworks[1], initialHomeworks[2], initialHomeworks[0]}
	require.Equal(t, len(expectedHomeworks), len(homeworks))
	for i := 0; i < len(expectedHomeworks); i++ {
		HomeworksEquals(t, expectedHomeworks[i], homeworks[i])
	}
}

func TestGetHomeworkStudentOk(t *testing.T) {
	initialHomeworks := []myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(time.Hour),
		},
	}
	server := CreateServer(initialHomeworks, []myhwproj.Submission{})
	defer server.Close()

	data, code, err := Get(server, "/v1/student/homeworks/0")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	var homework myhwproj.Homework
	err = json.Unmarshal(data, &homework)
	require.Nil(t, err)

	expectedHomework := initialHomeworks[0]
	HomeworksEquals(t, expectedHomework, homework)
}

func TestGetHomeworkStudentNotFound(t *testing.T) {
	initialHomeworks := []myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(time.Hour),
		},
	}
	server := CreateServer(initialHomeworks, []myhwproj.Submission{})
	defer server.Close()

	_, code, err := Get(server, "/v1/student/homeworks/1")
	require.Nil(t, err)

	require.Equal(t, http.StatusNotFound, code)
}

func TestGetHomeworkSubmissionsStudentOk(t *testing.T) {
	initialHomeworks := []myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(time.Hour),
		},
	}
	initialSubmissions := []myhwproj.Submission{
		{
			Id:         0,
			HomeworkId: 0,
			Datetime:   time.Now().Add(-time.Hour),
			Solution:   "https://yandex.ru",
			Mark:       0,
			Comment:    "",
		},
		{
			Id:         1,
			HomeworkId: 0,
			Datetime:   time.Now(),
			Solution:   "https://yandex.ru",
			Mark:       0,
			Comment:    "",
		},
	}
	server := CreateServer(initialHomeworks, initialSubmissions)
	defer server.Close()

	data, code, err := Get(server, "/v1/student/homeworks/0/submissions?offset=0&limit=10")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code, string(data))

	submissions := make([]myhwproj.Submission, 0, 10)
	err = json.Unmarshal(data, &submissions)
	require.Nil(t, err)

	expectedSubmissions := []myhwproj.Submission{initialSubmissions[1], initialSubmissions[0]}
	require.Equal(t, len(expectedSubmissions), len(submissions))
	for i := 0; i < len(expectedSubmissions); i++ {
		SubmissionsEquals(t, expectedSubmissions[i], submissions[i])
	}
}

func TestGetSubmissionStudentOk(t *testing.T) {
	initialHomeworks := []myhwproj.Homework{
		{
			Id:                  0,
			Name:                "Test",
			PublicationDatetime: time.Now().Add(-time.Hour),
			Check:               "./gradlew test",
			Task:                "Test",
			Deadline:            time.Now().Add(time.Hour),
		},
	}
	initialSubmissions := []myhwproj.Submission{
		{
			Id:         0,
			HomeworkId: 0,
			Datetime:   time.Now().Add(-time.Hour),
			Solution:   "https://yandex.ru",
			Mark:       0,
			Comment:    "",
		},
		{
			Id:         1,
			HomeworkId: 0,
			Datetime:   time.Now(),
			Solution:   "https://yandex.ru",
			Mark:       0,
			Comment:    "",
		},
	}
	server := CreateServer(initialHomeworks, initialSubmissions)
	defer server.Close()

	data, code, err := Get(server, "/v1/student/submissions/1")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code, string(data))

	var submission myhwproj.Submission
	err = json.Unmarshal(data, &submission)
	require.Nil(t, err)

	expectedSubmission := initialSubmissions[1]
	SubmissionsEquals(t, expectedSubmission, submission)
}
