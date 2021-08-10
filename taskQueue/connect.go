package taskQueue

import (
	"log"
	"github.com/lne-io/feedback/config"
	"github.com/lne-io/feedback/tasks"

	"github.com/hibiken/asynq"
)


func ConnectTaskQueueClient() {
	redisAddr := config.GetEnv("REDIS_ADDR", "127.0.0.1:6379")

	QueueClient = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
}

func ConnectTaskQueueWorker() {
	redisAddr := config.GetEnv("REDIS_ADDR", "127.0.0.1:6379")

	srv := asynq.NewServer(
        asynq.RedisClientOpt{Addr: redisAddr},
        asynq.Config{Concurrency: 1},
    )

    mux := asynq.NewServeMux()
    mux.HandleFunc(tasks.TypeRegisterFeedback, tasks.HandleRegisterFeedbackTask)
   
    if err := srv.Run(mux); err != nil {
        log.Fatalf("could not run server: %v", err)
    }
}