// Code generated by ent, DO NOT EDIT.

package ent

import (
	"poll-app/ent/poll"
	"poll-app/ent/polloption"
	"poll-app/ent/schema"
	"poll-app/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	pollFields := schema.Poll{}.Fields()
	_ = pollFields
	// pollDescIsActive is the schema descriptor for is_active field.
	pollDescIsActive := pollFields[2].Descriptor()
	// poll.DefaultIsActive holds the default value on creation for the is_active field.
	poll.DefaultIsActive = pollDescIsActive.Default.(bool)
	polloptionFields := schema.PollOption{}.Fields()
	_ = polloptionFields
	// polloptionDescVotes is the schema descriptor for votes field.
	polloptionDescVotes := polloptionFields[1].Descriptor()
	// polloption.DefaultVotes holds the default value on creation for the votes field.
	polloption.DefaultVotes = polloptionDescVotes.Default.(int)
	// polloptionDescIsActive is the schema descriptor for is_active field.
	polloptionDescIsActive := polloptionFields[2].Descriptor()
	// polloption.DefaultIsActive holds the default value on creation for the is_active field.
	polloption.DefaultIsActive = polloptionDescIsActive.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}