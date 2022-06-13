package test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http/httptest"
	"sort"
	"testing"
	"time"

	myhwproj "github.com/KaperD/HSE-SD-MyHwProj/internal"
	"github.com/stretchr/testify/assert"
)

func HomeworksEquals(t *testing.T, expectedHomework myhwproj.Homework, homework myhwproj.Homework) {
	assert.Equal(t, expectedHomework.Id, homework.Id)
	assert.Equal(t, expectedHomework.Name, homework.Name)
	assert.Equal(t, expectedHomework.Task, homework.Task)
	assert.Equal(t, expectedHomework.Deadline.Unix(), homework.Deadline.Unix())
	assert.Equal(t, expectedHomework.PublicationDatetime.Unix(), homework.PublicationDatetime.Unix())
}

func SubmissionsEquals(t *testing.T, expectedSubmission myhwproj.Submission, submission myhwproj.Submission) {
	assert.Equal(t, expectedSubmission.Id, submission.Id)
	assert.Equal(t, expectedSubmission.HomeworkId, submission.HomeworkId)
	assert.Equal(t, expectedSubmission.Solution, submission.Solution)
	assert.Equal(t, expectedSubmission.Mark, submission.Mark)
	assert.Equal(t, expectedSubmission.Comment, submission.Comment)
	assert.Equal(t, expectedSubmission.Datetime.Unix(), submission.Datetime.Unix())
}

func Post(server *httptest.Server, url string, message any) ([]byte, int, error) {
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		return nil, 0, err
	}

	resp, err := server.Client().Post(server.URL+url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return nil, 0, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return data, resp.StatusCode, nil
}

func Get(server *httptest.Server, url string) ([]byte, int, error) {
	resp, err := server.Client().Get(server.URL + url)
	if err != nil {
		return nil, 0, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return data, resp.StatusCode, nil
}

func CreateServer(homeworks []myhwproj.Homework, submissions []myhwproj.Submission) *httptest.Server {
	timezone, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = timezone

	TemplateCache, err := myhwproj.NewTemplateCache("./ui/html")
	if err != nil {
		log.Fatal(err)
	}

	HomeworkDao := homeworkDaoImpl{
		homeworks: make(map[int64]myhwproj.Homework),
		lastId:    0,
	}
	for _, homework := range homeworks {
		HomeworkDao.homeworks[homework.Id] = homework
	}
	SubmissionDao := submissionDaoImpl{
		submissions: make(map[int64]myhwproj.Submission),
		lastId:      0,
	}
	for _, submission := range submissions {
		SubmissionDao.submissions[submission.Id] = submission
	}

	WorkersService := rabbitMQWorkersServiceImpl{
		func(submission myhwproj.Submission) {},
	}

	StudentApiService := myhwproj.NewStudentApiService(&SubmissionDao, &HomeworkDao, &WorkersService)
	StudentApiController := myhwproj.NewStudentApiController(StudentApiService)

	StudentPagesApiService := myhwproj.NewStudentPagesApiService(StudentApiService, TemplateCache)
	StudentPagesApiController := myhwproj.NewStudentPagesApiController(StudentPagesApiService)

	TeacherApiService := myhwproj.NewTeacherApiService(&SubmissionDao, &HomeworkDao)
	TeacherApiController := myhwproj.NewTeacherApiController(TeacherApiService)

	TeacherPagesApiService := myhwproj.NewTeacherPagesApiService(TeacherApiService, TemplateCache)
	TeacherPagesApiController := myhwproj.NewTeacherPagesApiController(TeacherPagesApiService)

	router := myhwproj.NewRouter(StudentApiController, StudentPagesApiController, TeacherApiController, TeacherPagesApiController)

	server := httptest.NewServer(router)
	return server
}

type submissionDaoImpl struct {
	submissions map[int64]myhwproj.Submission
	lastId      int64
}

func (s *submissionDaoImpl) AddSubmission(homeworkId int64, newSubmission myhwproj.NewSubmission) myhwproj.Submission {
	submission := myhwproj.Submission{
		Id:         s.lastId,
		HomeworkId: homeworkId,
		Datetime:   time.Now(),
		Solution:   newSubmission.Solution,
		Mark:       0,
		Comment:    "",
	}
	s.lastId++
	s.submissions[submission.Id] = submission
	return submission
}

func (s *submissionDaoImpl) GetHomeworkSubmissions(homeworkId int64, offset int32, limit int32) []myhwproj.Submission {
	result := make([]myhwproj.Submission, 0)
	for _, v := range s.submissions {
		if v.HomeworkId == homeworkId {
			result = append(result, v)
		}
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Datetime.After(result[j].Datetime)
	})
	end := int(math.Min(float64(offset+limit), float64(len(result))))
	return result[offset:end]
}

func (s *submissionDaoImpl) GetSubmissionById(submissionId int64) *myhwproj.Submission {
	submission, ok := s.submissions[submissionId]
	if !ok {
		return nil
	}
	return &submission
}

func (s *submissionDaoImpl) UpdateSubmission(submission myhwproj.Submission) {
	s.submissions[submission.Id] = submission
}

type homeworkDaoImpl struct {
	homeworks map[int64]myhwproj.Homework
	lastId    int64
}

func (h *homeworkDaoImpl) AddHomework(newHomework myhwproj.NewHomework) myhwproj.Homework {
	homework := myhwproj.Homework{
		Id:                  h.lastId,
		Name:                newHomework.Name,
		PublicationDatetime: newHomework.PublicationDatetime,
		Check:               newHomework.Check,
		Task:                newHomework.Task,
		Deadline:            newHomework.Deadline,
	}
	h.lastId++
	h.homeworks[homework.Id] = homework
	return homework
}

func (h *homeworkDaoImpl) GetHomeworkById(homeworkId int64) *myhwproj.Homework {
	homework, ok := h.homeworks[homeworkId]
	if !ok {
		return nil
	}
	return &homework
}

func (h *homeworkDaoImpl) GetHomeworks(offset int32, limit int32, onlyPublished bool) []myhwproj.Homework {
	result := make([]myhwproj.Homework, 0)
	for _, v := range h.homeworks {
		if onlyPublished && v.PublicationDatetime.After(time.Now()) {
			continue
		}
		result = append(result, v)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Deadline.After(result[j].Deadline)
	})
	end := int(math.Min(float64(offset+limit), float64(len(result))))
	return result[offset:end]
}

type rabbitMQWorkersServiceImpl struct {
	handler func(submission myhwproj.Submission)
}

func (r *rabbitMQWorkersServiceImpl) SetHandler(f func(submission myhwproj.Submission)) {
	r.handler = f
}

func (r *rabbitMQWorkersServiceImpl) CheckSubmission(submission myhwproj.Submission) {
	submission.Mark = 10
	r.handler(submission)
}
