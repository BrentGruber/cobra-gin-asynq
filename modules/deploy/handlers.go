package deploy

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"releasr/internal/tasks"

	"github.com/hibiken/asynq"
)

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p tasks.EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%s, template_id=%s", p.UserID, p.TemplateID)
	// Email delivery code ...
	return nil
}
