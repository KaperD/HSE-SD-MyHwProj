package myhwproj

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// WorkersService defines the api for service which works with Workers who check Submissions
type WorkersService interface {
	SetHandler(func(submission Submission))
	CheckSubmission(submission Submission)
}

// Helper function to check rabbitmq errors
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// RabbitMQWorkersService works with Workers through RabbitMQ
type RabbitMQWorkersService struct {
	ResultHandler func(submission Submission)
	HomeworkDao   HomeworkDao
}

// NewRabbitMQWorkersService creates default RabbitMQWorkersService
func NewRabbitMQWorkersService(homeworkDao HomeworkDao) WorkersService {
	return &RabbitMQWorkersService{
		ResultHandler: func(submission Submission) {},
		HomeworkDao:   homeworkDao,
	}
}

// SetHandler changes the handler of check result
func (r *RabbitMQWorkersService) SetHandler(f func(submission Submission)) {
	r.ResultHandler = f
}

func (submission *Submission) SerializeToBase64() string {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(submission)
	failOnError(err, "Failed to encode Submission to Base64")
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

func DeserializeSubmissionFromBase64(str string) Submission {
	submission := Submission{}
	strBytes, err := base64.StdEncoding.DecodeString(str)
	failOnError(err, "Failed to decode Submission from Base64")
	buffer := bytes.Buffer{}
	buffer.Write(strBytes)
	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&submission)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
	return submission
}

type RabbitMQMessage struct {
	Submission Submission
	Howework   Homework
}

// CheckSubmission checks the submission async. After check calls handler on the result
func (r *RabbitMQWorkersService) CheckSubmission(submission Submission) {
	go func() {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		homework := r.HomeworkDao.GetHomeworkById(submission.HomeworkId)

		q, err := ch.QueueDeclare(
			"hello", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		failOnError(err, "Failed to declare a queue")
		message := RabbitMQMessage{submission, *homework}

		body, err := json.Marshal(message)
		failOnError(err, "Failed to serialize submission to JSON")

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", body)
		submission.Mark = 5
		submission.Comment = "Checked"
		r.ResultHandler(submission)
	}()
}
