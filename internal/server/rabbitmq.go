package server

import (
	"github.com/go-parrot/parrot-starter/internal/jobs"
	"github.com/go-parrot/parrot-starter/internal/tasks"
	rabbitmqConf "github.com/go-parrot/parrot/pkg/queue/rabbitmq"
	"github.com/go-parrot/parrot/pkg/transport/consumer/rabbitmq"
)

// NewRabbitmqConsumerServer create a redis server
func NewRabbitmqConsumerServer() *rabbitmq.Server {
	rabbitmqConf.Load()

	srv := rabbitmq.NewServer()

	// register handler
	srv.RegisterHandler(tasks.TypeEmailWelcome, jobs.SendWelcomeEmailHandler)
	// here register other handlers...

	return srv
}
