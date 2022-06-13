package myhwproj

import (
	"encoding/json"
	"log"
	"math/rand"

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

type RabbitMQMessage struct {
	Submission Submission
	Homework   Homework
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

// CheckSubmission checks the submission async. After check calls handler on the result
func (r *RabbitMQWorkersService) CheckSubmission(submission Submission) {
	go func() {
		homework := r.HomeworkDao.GetHomeworkById(submission.HomeworkId)
		message := RabbitMQMessage{submission, *homework}
		body, err := json.Marshal(message)
		failOnError(err, "Failed to serialize submission to JSON")

		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // noWait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		corrId := randomString(32)

		err = ch.Publish(
			"",          // exchange
			"rpc_queue", // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: corrId,
				ReplyTo:       q.Name,
				Body:          []byte(body),
			},
		)

		failOnError(err, "Failed to publish a message")

		for d := range msgs {
			if corrId == d.CorrelationId {
				err = json.Unmarshal(d.Body, &submission)
				failOnError(err, "Failed to parse response submission json")
				r.ResultHandler(submission)
				break
			}
		}
	}()
}
