package test

import (
	"encoding/json"
	myhwproj "github.com/KaperD/HSE-SD-MyHwProj/internal"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestAddHomeworkTeacherOk(t *testing.T) {
	server := CreateServer([]myhwproj.Homework{}, []myhwproj.Submission{})
	defer server.Close()

	newHomework := myhwproj.NewHomework{
		Name:                "Test",
		PublicationDatetime: time.Now().Add(-time.Hour),
		Check:               "./gradlew test",
		Task:                "Test",
		Deadline:            time.Now().Add(time.Hour),
	}

	data, code, err := Post(server, "/v1/teacher/homeworks", newHomework)
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	var homework myhwproj.Homework
	require.Nil(t, json.Unmarshal(data, &homework))

	expectedHomework := myhwproj.Homework{
		Id:                  0,
		Name:                newHomework.Name,
		PublicationDatetime: newHomework.PublicationDatetime,
		Check:               newHomework.Check,
		Task:                newHomework.Task,
		Deadline:            newHomework.Deadline,
	}

	HomeworksEquals(t, expectedHomework, homework)
}

func TestAddHomeworkTeacherWrongJson(t *testing.T) {
	server := CreateServer([]myhwproj.Homework{}, []myhwproj.Submission{})
	defer server.Close()

	_, code, err := Post(server, "/v1/teacher/homeworks", "wrong")
	require.Nil(t, err)

	require.Equal(t, http.StatusBadRequest, code)
}

func TestGetHomeworksTeacherOk(t *testing.T) {
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

	data, code, err := Get(server, "/v1/teacher/homeworks?offset=0&limit=10")
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

func TestGetHomeworkTeacherOk(t *testing.T) {
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

	data, code, err := Get(server, "/v1/teacher/homeworks/0")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code)

	var homework myhwproj.Homework
	err = json.Unmarshal(data, &homework)
	require.Nil(t, err)

	expectedHomework := initialHomeworks[0]
	HomeworksEquals(t, expectedHomework, homework)
}

func TestGetHomeworkTeacherNotFound(t *testing.T) {
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

	_, code, err := Get(server, "/v1/teacher/homeworks/1")
	require.Nil(t, err)

	require.Equal(t, http.StatusNotFound, code)
}

func TestGetHomeworkSubmissionsTeacherOk(t *testing.T) {
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

	data, code, err := Get(server, "/v1/teacher/homeworks/0/submissions?offset=0&limit=10")
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

func TestGetSubmissionTeacherOk(t *testing.T) {
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

	data, code, err := Get(server, "/v1/teacher/submissions/1")
	require.Nil(t, err)

	require.Equal(t, http.StatusOK, code, string(data))

	var submission myhwproj.Submission
	err = json.Unmarshal(data, &submission)
	require.Nil(t, err)

	expectedSubmission := initialSubmissions[1]
	SubmissionsEquals(t, expectedSubmission, submission)
}
