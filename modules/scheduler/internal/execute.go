package internal

import "github.com/rs/zerolog"

func Execute(taskId string, logger zerolog.Logger) error {
	logger.Info().Str("task_id", taskId).Str("event", "scheduler:execute").Msg("Starting task")
	return nil
}
