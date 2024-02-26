package models

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent"
	"poll-app/ent/poll"
	"poll-app/ent/polloption"
)

// Option represents a poll option entity in the system
type PollOption struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Votes    int    `json:"votes"`
	IsActive bool   `json:"-"`
}

// Poll represents a poll entity in the system
type Poll struct {
	ID          int           `json:"id"`
	Question    string        `json:"question"`
	PollOptions []*PollOption `json:"options"`
	CreatedBy   string        `json:"created_by"`
	IsActive    bool          `json:"-"`
}

// CreatePoll creates a new poll in the database
func CreatePoll(ctx context.Context, question string, createdBy string, optionTexts []string) (*Poll, error) {
	client := ctx.Value("client").(*ent.Client)
	if client == nil {
		return nil, errors.New("ent client is nil in context")
	}
	newPoll, err := client.Poll.
		Create().
		SetQuestion(question).
		SetCreatedBy(createdBy).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create poll: %v", err)
	}
	var options []*PollOption
	for _, text := range optionTexts {
		opt, err := client.PollOption.
			Create().
			SetText(text).
			SetPoll(newPoll).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create option for poll: %v", err)
		}
		options = append(options, &PollOption{
			ID:    opt.ID,
			Text:  opt.Text,
			Votes: 0,
		})
	}
	return &Poll{
		ID:          newPoll.ID,
		Question:    newPoll.Question,
		PollOptions: options,
		CreatedBy:   createdBy,
	}, nil
}

// ListPolls fetches all the active polls from the database
func ListPolls(ctx context.Context) ([]*Poll, error) {
	client := ctx.Value("client").(*ent.Client)
	polls, err := client.Poll.Query().
		Where(poll.IsActive(true)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*Poll
	for _, p := range polls {
		options, err := getOptionsForPoll(ctx, p.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch options for poll: %v", err)
		}

		result = append(result, &Poll{
			ID:          p.ID,
			Question:    p.Question,
			PollOptions: options,
			CreatedBy:   p.CreatedBy,
		})
	}
	return result, nil
}

// GetPoll retrieves a poll from the database based on the provided poll ID
func GetPoll(ctx context.Context, pollID int) (*Poll, error) {
	client := ctx.Value("client").(*ent.Client)
	p, err := client.Poll.Query().Where(poll.ID(pollID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	options, err := getOptionsForPoll(ctx, p.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch options for poll: %v", err)
	}
	return &Poll{
		ID:          p.ID,
		Question:    p.Question,
		PollOptions: options,
		CreatedBy:   p.CreatedBy,
	}, nil
}

// DeletePoll soft deletes an existing poll from the database along with its associated options
func DeletePoll(ctx context.Context, pollID int) error {
	client := ctx.Value("client").(*ent.Client)
	poll, err := client.Poll.Query().
		Where(poll.ID(pollID)).
		WithOptions().
		Only(ctx)
	if err != nil {
		return err
	}
	for _, option := range poll.Edges.Options {
		if _, err := client.PollOption.UpdateOne(option).SetIsActive(false).Save(ctx); err != nil {
			return err
		}
	}
	if _, err := client.Poll.UpdateOne(poll).SetIsActive(false).Save(ctx); err != nil {
		return err
	}
	return nil
}

// getOptionsForPoll retrieves the active options associated with a poll from the database
func getOptionsForPoll(ctx context.Context, pollID int) ([]*PollOption, error) {
	client := ctx.Value("client").(*ent.Client)
	options, err := client.PollOption.Query().
		Where(
			polloption.HasPollWith(poll.ID(pollID)),
			polloption.IsActiveEQ(true),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*PollOption
	for _, o := range options {
		result = append(result, &PollOption{
			ID:    o.ID,
			Text:  o.Text,
			Votes: o.Votes,
		})
	}
	return result, nil
}

// UpdatePoll updates an existing poll in the database
func UpdatePoll(ctx context.Context, pollID int, updatedPoll Poll) error {
	client := ctx.Value("client").(*ent.Client)
	poll, err := client.Poll.Query().
		Where(poll.ID(pollID)).
		WithOptions().
		Only(ctx)
	if err != nil {
		return err
	}
	if updatedPoll.Question != "" {
		if _, err := poll.Update().SetQuestion(updatedPoll.Question).Save(ctx); err != nil {
			return err
		}
	}
	for _, option := range poll.Edges.Options {
		if _, err := client.PollOption.UpdateOne(option).SetIsActive(false).Save(ctx); err != nil {
			return err
		}
	}
	for _, option := range updatedPoll.PollOptions {
		existingOption, err := client.PollOption.Query().
			Where(polloption.Text(option.Text)).
			Only(ctx)
		if err == nil {
			if _, err := client.PollOption.UpdateOne(existingOption).SetIsActive(true).Save(ctx); err != nil {
				return err
			}
		} else if ent.IsNotFound(err) {
			if _, err := client.PollOption.Create().
				SetPoll(poll).
				SetText(option.Text).
				SetIsActive(true).
				Save(ctx); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
