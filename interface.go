package backendplugin

import (
	"context"
	"time"
)

type BackendPlugin interface {
	// Configuration.
	Configure(_ context.Context, config map[string]string) error

	// Workspace operations.
	ListWorkspaces(_ context.Context) ([]string, error)
	DeleteWorkspace(_ context.Context, workspace string, force bool) error

	// State operations.
	GetStatePayload(_ context.Context, workspace string) (*Payload, error)
	PutState(_ context.Context, workspace string, data []byte) error
	DeleteState(_ context.Context, workspace string) error
	LockState(_ context.Context, workspace string, info *LockInfo) (string, error)
	UnlockState(_ context.Context, workspace, id string) error
}

// LockInfo is the information that is stored with a lock.
type LockInfo struct {
	// Unique ID for the lock. NewLockInfo provides a random ID, but this may
	// be overridden by the lock implementation. The final value of ID will be
	// returned by the call to Lock.
	ID string

	// Terraform operation, provided by the caller.
	Operation string

	// Extra information to store with the lock, provided by the caller.
	Info string

	// user@hostname when available
	Who string

	// Terraform version
	Version string

	// Time that the lock was taken.
	Created time.Time

	// Path to the state file when applicable. Set by the Lock implementation.
	Path string
}

// Payload is the data that is stored in the state file.
type Payload struct {
	MD5  []byte
	Data []byte
}
