package configserviceiface

import (
	"github.com/awslabs/aws-sdk-go/service/configservice"
)

type ConfigServiceAPI interface {
	DeleteDeliveryChannel(*configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error)

	DeliverConfigSnapshot(*configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error)

	DescribeConfigurationRecorderStatus(*configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error)

	DescribeConfigurationRecorders(*configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error)

	DescribeDeliveryChannelStatus(*configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error)

	DescribeDeliveryChannels(*configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error)

	GetResourceConfigHistory(*configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error)

	PutConfigurationRecorder(*configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error)

	PutDeliveryChannel(*configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error)

	StartConfigurationRecorder(*configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error)

	StopConfigurationRecorder(*configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error)
}