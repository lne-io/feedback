package tasks

import (
    "fmt"
    "context"
	"encoding/json"
	
	"github.com/lne-io/feedback/models"
	"github.com/lne-io/feedback/database"

    "github.com/hibiken/asynq"
)

// Feedback task types.
const (
    TypeRegisterFeedback   = "feedback:register"
)

// Register Feedback task
func NewRegisterFeedbackTask(feedback *models.Feedback) (*asynq.Task, error) {
    payload, err := json.Marshal(*feedback)
    if err != nil {
        return nil, err
    }
    return asynq.NewTask(TypeRegisterFeedback, payload), nil
}


// Register Feedback handler
func HandleRegisterFeedbackTask(ctx context.Context, t *asynq.Task) error {
    db := database.DB

	var feedback models.Feedback
    if err := json.Unmarshal(t.Payload(), &feedback); err != nil {
        return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
    }
    // save here
	if result := db.Create(&feedback); result.Error != nil {
		return result.Error
	}
    return nil
}