package job

import (
	"context"
	"time"
)

// 任务设计参考 - https://www.oschina.net/news/406862

// 执行策略
type ExecutionPolicy string

// 阻塞策略
type BlockingPolicy string

type JobStatus string

// 任务状态
const (
	JobStatusEnabled  JobStatus = "enabled"  // 启用
	JobStatusDisabled JobStatus = "disabled" // 禁用
)

// 执行策略
const (
	ExecutionPolicyOnce   ExecutionPolicy = "once"   // 单次执行
	ExecutionPolicyRepeat ExecutionPolicy = "repeat" // 重复执行
)

// 阻塞策略
const (
	BlockingPolicyDiscard  BlockingPolicy = "discard"  // 丢弃
	BlockingPolicyParallel BlockingPolicy = "parallel" // 并行
)

type Job struct {
	ID              string                 // 任务唯一标识
	Group           string                 // 任务分组
	Name            string                 // 任务名称
	Description     string                 // 任务描述
	ExecutorName    string                 // 执行器名称
	ExecutionPolicy ExecutionPolicy        // 执行策略：单次/重复
	Status          JobStatus              // 状态：启用/禁用
	CronExpression  string                 // Cron表达式
	Parameters      map[string]interface{} // 任务参数
	BlockingPolicy  BlockingPolicy         // 阻塞策略：丢弃/并行
	Timeout         time.Duration          // 超时时间
	MaxRetry        int                    // 最大重试次数
	RetryInterval   time.Duration          // 重试间隔
	ParallelNum     int                    // 并行数
	RunningCount    int                    // 当前运行数
}

type Executor interface {
	Execute(ctx context.Context, job *Job) error
	Name() string
}

type JobResult interface {
}

type JobScheduler interface {
	// 生命周期管理
	Start()
	Stop()

	// 执行器管理
	RegisterExecutor(executor Executor)
	ListExecutors() []Executor

	// 任务管理
	AddOrUpdateJob(job *Job) (string, error)
	EnableJob(jobID string) error
	DisableJob(jobID string) error
	DeleteJob(jobID string) error
	ExecuteNow(jobID string) error
	ListJobs() []*Job
	JobExists(jobID string) bool

	// 结果获取
	GetResults() <-chan *JobResult
}

func Retry(ctx context.Context, executor Executor, job *Job) {
	for i := 0; i <= job.MaxRetry; i++ {
		err := executor.Execute(ctx, job)
		if err == nil {
			return
		}

		if i < job.MaxRetry {
			time.Sleep(job.RetryInterval)
		}
	}
}
