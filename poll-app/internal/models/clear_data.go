package models

import (
	"context"
	"poll-app/ent"
)

// ClearAllData clears all data from the database
func ClearAllData(ctx context.Context) error {
	client := ctx.Value("client").(*ent.Client)
	if _, err := client.Vote.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := client.PollOption.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := client.Poll.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := client.User.Delete().Exec(ctx); err != nil {
		return err
	}

	return nil
}
