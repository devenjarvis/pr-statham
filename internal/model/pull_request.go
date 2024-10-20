package model

import (
	"time"
)

type MergeableState string

const (
	MergeableStateMergeable   MergeableState = "MERGEABLE"
	MergeableStateConflicting MergeableState = "CONFLICTING"
	MergeableStateUnknown     MergeableState = "UNKNOWN"
)

type PullRequest struct {
	Title     string
	Permalink string
	Additions int
	Deletions int
	CreatedAt time.Time
	IsDraft   bool
	Mergeable MergeableState
	Merged    bool
	Closed    bool
}
