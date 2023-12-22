package interfaces

import (
	"gorm.io/gorm"

	"github.com/flyteorg/flyte/flyteadmin/pkg/repositories/errors"
	schedulerInterfaces "github.com/flyteorg/flyte/flyteadmin/scheduler/repositories/interfaces"
	"github.com/flyteorg/flyte/flytestdlib/promutils"
)

// The Repository indicates the methods that each Repository must support.
// A Repository indicates a Database which is collection of Tables/models.
// The goal is allow databases to be Plugged in easily.
type Repository interface {
	TaskRepo() TaskRepoInterface
	WorkflowRepo() WorkflowRepoInterface
	LaunchPlanRepo() LaunchPlanRepoInterface
	ExecutionRepo() ExecutionRepoInterface
	ExecutionEventRepo() ExecutionEventRepoInterface
	ProjectRepo() ProjectRepoInterface
	ResourceRepo() ResourceRepoInterface
	NodeExecutionRepo() NodeExecutionRepoInterface
	NodeExecutionEventRepo() NodeExecutionEventRepoInterface
	TaskExecutionRepo() TaskExecutionRepoInterface
	NamedEntityRepo() NamedEntityRepoInterface
	DescriptionEntityRepo() DescriptionEntityRepoInterface
	SchedulableEntityRepo() schedulerInterfaces.SchedulableEntityRepoInterface
	ScheduleEntitiesSnapshotRepo() schedulerInterfaces.ScheduleEntitiesSnapShotRepoInterface
	SignalRepo() SignalRepoInterface

	GetGormDB() *gorm.DB
}

type NewRepositoryFunc = func(db *gorm.DB, errorTransformer errors.ErrorTransformer, scope promutils.Scope) Repository
