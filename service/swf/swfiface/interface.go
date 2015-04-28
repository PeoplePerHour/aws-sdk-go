package swfiface

import (
	"github.com/awslabs/aws-sdk-go/service/swf"
)

type SWFAPI interface {
	CountClosedWorkflowExecutions(*swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error)

	CountOpenWorkflowExecutions(*swf.CountOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error)

	CountPendingActivityTasks(*swf.CountPendingActivityTasksInput) (*swf.PendingTaskCount, error)

	CountPendingDecisionTasks(*swf.CountPendingDecisionTasksInput) (*swf.PendingTaskCount, error)

	DeprecateActivityType(*swf.DeprecateActivityTypeInput) (*swf.DeprecateActivityTypeOutput, error)

	DeprecateDomain(*swf.DeprecateDomainInput) (*swf.DeprecateDomainOutput, error)

	DeprecateWorkflowType(*swf.DeprecateWorkflowTypeInput) (*swf.DeprecateWorkflowTypeOutput, error)

	DescribeActivityType(*swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error)

	DescribeDomain(*swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error)

	DescribeWorkflowExecution(*swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error)

	DescribeWorkflowType(*swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error)

	GetWorkflowExecutionHistory(*swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error)

	ListActivityTypes(*swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error)

	ListClosedWorkflowExecutions(*swf.ListClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error)

	ListDomains(*swf.ListDomainsInput) (*swf.ListDomainsOutput, error)

	ListOpenWorkflowExecutions(*swf.ListOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error)

	ListWorkflowTypes(*swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error)

	PollForActivityTask(*swf.PollForActivityTaskInput) (*swf.PollForActivityTaskOutput, error)

	PollForDecisionTask(*swf.PollForDecisionTaskInput) (*swf.PollForDecisionTaskOutput, error)

	RecordActivityTaskHeartbeat(*swf.RecordActivityTaskHeartbeatInput) (*swf.RecordActivityTaskHeartbeatOutput, error)

	RegisterActivityType(*swf.RegisterActivityTypeInput) (*swf.RegisterActivityTypeOutput, error)

	RegisterDomain(*swf.RegisterDomainInput) (*swf.RegisterDomainOutput, error)

	RegisterWorkflowType(*swf.RegisterWorkflowTypeInput) (*swf.RegisterWorkflowTypeOutput, error)

	RequestCancelWorkflowExecution(*swf.RequestCancelWorkflowExecutionInput) (*swf.RequestCancelWorkflowExecutionOutput, error)

	RespondActivityTaskCanceled(*swf.RespondActivityTaskCanceledInput) (*swf.RespondActivityTaskCanceledOutput, error)

	RespondActivityTaskCompleted(*swf.RespondActivityTaskCompletedInput) (*swf.RespondActivityTaskCompletedOutput, error)

	RespondActivityTaskFailed(*swf.RespondActivityTaskFailedInput) (*swf.RespondActivityTaskFailedOutput, error)

	RespondDecisionTaskCompleted(*swf.RespondDecisionTaskCompletedInput) (*swf.RespondDecisionTaskCompletedOutput, error)

	SignalWorkflowExecution(*swf.SignalWorkflowExecutionInput) (*swf.SignalWorkflowExecutionOutput, error)

	StartWorkflowExecution(*swf.StartWorkflowExecutionInput) (*swf.StartWorkflowExecutionOutput, error)

	TerminateWorkflowExecution(*swf.TerminateWorkflowExecutionInput) (*swf.TerminateWorkflowExecutionOutput, error)
}