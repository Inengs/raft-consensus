package raft

import (
	"sync"
	"time"
)

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

	// --- Persistent state (must survive crash) ---
	currentTerm int
	votedFor int // -1 means none
	log []LogEntry

	// Volatile state (all nodes)
	commitIndex int
	lastApplied int
	state NodeState

	// Volatile state (leader only, reset after election)
	nextIndex []int // indexed by peer
	matchIndex []int

	// Timers
	electionTimer *time.Timer
	heartbeatTimer *time.Timer

	// Communication
	applyCh chan ApplyMsg // sends committed entries to the KV store

	// Persistence 
	persister Persister
}

// Apply message
type ApplyMsg struct {
	Command Command
	Index int
}

type Persister struct {
	currentTerm int
	log []LogEntry
}