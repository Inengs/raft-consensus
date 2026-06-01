package raft

import "sync"

type NodeState int

const (
	follower NodeState = iota
	Candidate
	Leader
)

type LogEntry struct {
	Term    int     // term when entry was received by leader
	Index   int     // position in the log (l-indexed)
	Command Command // the operation to apply to the KV store
}

type Command struct {
	Op    string // "put" "delete"
	Key   string
	Value string // empty for delete
}

type Node struct {
	mu sync.Mutex

	// Identity
	id int
	peers []string // peer HTTP addresses e.g. ["localhost:8001", "localhost:8002"]
}