// This file is part of Vassago.
// See LICENSE-Apache-2.0 for license information.

// Package client provides a gRPC client for the Vassago memory daemon.
package client

import (
	"context"

	pb "github.com/aj-nt/vassago-sdk/proto"
)


// TaskEntry wraps a proto TaskEntry for the SDK API.
type TaskEntry struct {
	Id            string
	ParentIds     string
	Status        string
	AgentType     string
	AssignedAgent string
	Goal          string
	Context       string
	ResultKey     string
	Priority      int32
	CreatedAt     int64
	ClaimedAt     int64
	CompletedAt   int64
	TtlSeconds    int32
	RetryCount    int32
	MaxRetries    int32
}

// ResetTimedOutResp holds the response from ResetTimedOutTasks.
type ResetTimedOutResp struct {
	ResetIds []string
	Count    int32
}

// MnemoClient is an interface for interacting with the Vassago daemon.
type MnemoClient interface {
	// IsConnected returns true if connected to the daemon.
	IsConnected() bool

	// RegisterAgent registers the agent with the daemon.
	RegisterAgent(ctx context.Context, agentID, name, role string) (returnedAgentID string, isNew bool, err error)

	// Heartbeat sends a keep-alive to the daemon.
	Heartbeat(ctx context.Context, agentID string) error

	// Subscribe opens a server-streaming subscription for memory change events.
	Subscribe(ctx context.Context, req *pb.SubscribeRequest) (pb.Vassago_SubscribeClient, error)

	// AddMemory adds or updates a memory entry.
	AddMemory(ctx context.Context, target, category, key, content string, priority int32, sourceAgent string) (*pb.MemoryEntry, error)

	// GetMemory retrieves a memory by ID.
	GetMemory(ctx context.Context, id string) (*pb.MemoryEntry, error)

	// RemoveMemory deletes a memory by ID.
	RemoveMemory(ctx context.Context, id string) error

	// SearchMemories searches memories using FTS5.
	SearchMemories(ctx context.Context, query string, target string, limit int32) ([]*pb.MemoryEntry, error)

	// ListMemories lists memories with optional filtering.
	ListMemories(ctx context.Context, target string, category *string, minPriority *int32) ([]*pb.MemoryEntry, error)

	// GetHotBlock retrieves the computed hot block for an agent.
	GetHotBlock(ctx context.Context, target, agentID string, charLimit int32, boostCategories, suppressCategories []string, boostPriority int32) (string, error)

	// CreateSession creates a new session.
	CreateSession(ctx context.Context, agentID, source, title string) (string, error)

	// AddMessages adds messages to a session.
	AddMessages(ctx context.Context, sessionID string, messages []Message) error

	// EndSession ends a session.
	EndSession(ctx context.Context, sessionID string) error

	// Consolidate triggers memory consolidation with the given scope.
	Consolidate(ctx context.Context, scope pb.ConsolidateScope, batchSize int32) (*pb.ConsolidateResponse, error)

	// HealthCheck verifies the daemon is reachable and responding.
	HealthCheck(ctx context.Context) error

	// SearchSessions searches session titles matching the query.
	SearchSessions(ctx context.Context, query string, limit int32) ([]SessionInfo, error)

	// ListRecentSessions returns the most recently active sessions.
	ListRecentSessions(ctx context.Context, limit int32) ([]SessionInfo, error)

	// GetSession retrieves a session with its messages.
	GetSession(ctx context.Context, sessionID string) (*SessionDetail, error)

	// SendMessage sends a message via the messaging service.
	SendMessage(ctx context.Context, platform, target, message string) (string, error)

	// ListChannels lists available channels for a platform.
	ListChannels(ctx context.Context, platform string) ([]*pb.Channel, error)

	// AddTodo creates a new todo item.
	AddTodo(ctx context.Context, content string, priority int32, sourceAgent string) (*TodoItem, error)

	// ListTodos returns todos, optionally filtering by completion status and agent.
	ListTodos(ctx context.Context, includeCompleted bool, sourceAgent string) ([]TodoItem, error)

	// CompleteTodo marks a todo as completed.
	CompleteTodo(ctx context.Context, id string) (*TodoItem, error)

	// RemoveTodo removes a todo item.
	RemoveTodo(ctx context.Context, id string) error

	// Cron job management
	CreateCronJob(ctx context.Context, name, schedule, agentID, message string, enabled *bool) (*CronJobInfo, error)
	ListCronJobs(ctx context.Context, agentID string, includeDisabled bool) ([]CronJobInfo, error)
	UpdateCronJob(ctx context.Context, id string, name, schedule, message *string, enabled *bool) (*CronJobInfo, error)
	DeleteCronJob(ctx context.Context, id string) error

	// Skill management
	AddSkill(ctx context.Context, name, description string, triggers []string, category, content, sourceAgent string, priority int32) (*pb.SkillEntry, error)
	GetSkill(ctx context.Context, id string) (*pb.SkillEntry, error)
	ListSkills(ctx context.Context, category, sourceAgent string, limit int32) ([]*pb.SkillEntry, error)
	RemoveSkill(ctx context.Context, id string) (bool, error)
	SearchSkills(ctx context.Context, query, category string, limit int32) ([]*pb.SkillEntry, error)

	// --- Saved Tools ---

	// AddSavedTool adds a new saved tool.
	AddSavedTool(ctx context.Context, name, description, toolType, parameters, script, category string, triggers []string, author string, priority int32) (*pb.SavedToolEntry, error)

	// GetSavedTool retrieves a saved tool by ID.
	GetSavedTool(ctx context.Context, id string) (*pb.SavedToolEntry, error)

	// ListSavedTools lists saved tools with optional filters.
	ListSavedTools(ctx context.Context, category, author string, limit int32) ([]*pb.SavedToolEntry, error)

	// SearchSavedTools searches saved tools via FTS5.
	SearchSavedTools(ctx context.Context, query, category string, limit int32) ([]*pb.SavedToolEntry, error)

	// UpdateSavedTool updates fields on a saved tool.
	UpdateSavedTool(ctx context.Context, id string, updates map[string]interface{}) (*pb.SavedToolEntry, error)

	// RemoveSavedTool deletes a saved tool.
	RemoveSavedTool(ctx context.Context, id string) (bool, error)

	// --- Task Management (Multi-Agent Orchestration) ---

	// AddTask creates a new orchestration task.
	AddTask(ctx context.Context, id, agentType, goal, context string, priority, ttlSeconds, maxRetries int32) (*TaskEntry, error)

	// GetTask retrieves a task by ID.
	GetTask(ctx context.Context, id string) (*TaskEntry, error)

	// ClaimTask atomically claims a ready task.
	ClaimTask(ctx context.Context, taskID, agentID string) (*TaskEntry, error)

	// CompleteTask marks a task as done.
	CompleteTask(ctx context.Context, taskID, resultKey string) (*TaskEntry, error)

	// FailTask marks a claimed task as failed.
	FailTask(ctx context.Context, taskID string) (*TaskEntry, error)

	// FindReadyTasks returns ready tasks, optionally filtered by agent_type.
	FindReadyTasks(ctx context.Context, agentType string, limit int32) ([]*TaskEntry, error)

	// ListTasksByStatus returns tasks filtered by optional status.
	ListTasksByStatus(ctx context.Context, status string, limit int32) ([]*TaskEntry, error)

	// ResetTimedOutTasks resets claimed tasks past their TTL.
	ResetTimedOutTasks(ctx context.Context) (*ResetTimedOutResp, error)

	// DeleteTask removes a task.
	DeleteTask(ctx context.Context, id string) (bool, error)

	// Close closes the connection to the daemon.
	Close() error
}

// Ensure that Client implements the MnemoClient interface.
var _ MnemoClient = (*Client)(nil)
