package models

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent"
	"poll-app/ent/poll"
	"poll-app/ent/polloption"
	"poll-app/ent/vote"
)

// Vote represents a vote entity in the system
type Vote struct {
	ID       int    `json:"id"`
	PollID   int    `json:"poll_id"`
	UserID   string `json:"user_id"`
	OptionID int    `json:"option_id"`
}

// VoteOnPoll allows a user to vote on a poll by providing the poll ID, user ID, and the ID of the option they are voting for.
func VoteOnPoll(ctx context.Context, pollID int, userID string, optionID int) error {
	client := ctx.Value("client").(*ent.Client)
	if client == nil {
		return errors.New("ent client is nil in context")
	}

	hasVoted, err := HasUserVoted(ctx, pollID, userID)
	if err != nil {
		return err
	}
	if hasVoted {
		return errors.New("user has already voted on this poll")
	}

	exists, err := client.Poll.Query().
		Where(poll.ID(pollID)).
		QueryOptions().
		Where(polloption.ID(optionID)).
		Exist(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("option with given ID not found for the poll")
	}

	_, err = client.Vote.Create().
		SetPollID(pollID).
		SetUserID(userID).
		SetOptionID(optionID).
		Save(ctx)
	if err != nil {
		return err
	}

	_, err = client.PollOption.UpdateOneID(optionID).
		AddVotes(1).
		Save(ctx)
	return err
}

// GetVotes retrieves votes for a specific poll option from the database
func GetVotes(ctx context.Context, pollID, optionID int) ([]*Vote, error) {
	client := ctx.Value("client").(*ent.Client)

	votes, err := client.Vote.Query().
		Where(vote.OptionID(optionID)).
		Where(vote.PollID(pollID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch votes for poll option: %v", err)
	}

	var voteModels []*Vote
	for _, v := range votes {
		voteModels = append(voteModels, &Vote{
			ID:       v.ID,
			PollID:   v.PollID,
			UserID:   v.UserID,
			OptionID: v.OptionID,
		})
	}

	return voteModels, nil
}

func HasUserVoted(ctx context.Context, pollID int, userID string) (bool, error) {
	client := ctx.Value("client").(*ent.Client)
	if client == nil {
		return false, errors.New("ent client is nil in context")
	}

	exists, err := client.Vote.Query().
		Where(vote.PollID(pollID)).
		Where(vote.UserID(userID)).
		Exist(ctx)
	if err != nil {
		return false, err
	}
	return exists, nil
}
