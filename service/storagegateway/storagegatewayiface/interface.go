package storagegatewayiface

import (
	"github.com/awslabs/aws-sdk-go/service/storagegateway"
)

type StorageGatewayAPI interface {
	ActivateGateway(*storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error)

	AddCache(*storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error)

	AddUploadBuffer(*storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error)

	AddWorkingStorage(*storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error)

	CancelArchival(*storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error)

	CancelRetrieval(*storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error)

	CreateCachediSCSIVolume(*storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error)

	CreateSnapshot(*storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error)

	CreateSnapshotFromVolumeRecoveryPoint(*storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error)

	CreateStorediSCSIVolume(*storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error)

	CreateTapes(*storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error)

	DeleteBandwidthRateLimit(*storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error)

	DeleteChapCredentials(*storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error)

	DeleteGateway(*storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error)

	DeleteSnapshotSchedule(*storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error)

	DeleteTape(*storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error)

	DeleteTapeArchive(*storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error)

	DeleteVolume(*storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error)

	DescribeBandwidthRateLimit(*storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error)

	DescribeCache(*storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error)

	DescribeCachediSCSIVolumes(*storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error)

	DescribeChapCredentials(*storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error)

	DescribeGatewayInformation(*storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error)

	DescribeMaintenanceStartTime(*storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error)

	DescribeSnapshotSchedule(*storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error)

	DescribeStorediSCSIVolumes(*storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error)

	DescribeTapeArchives(*storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error)

	DescribeTapeRecoveryPoints(*storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error)

	DescribeTapes(*storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error)

	DescribeUploadBuffer(*storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error)

	DescribeVTLDevices(*storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error)

	DescribeWorkingStorage(*storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error)

	DisableGateway(*storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error)

	ListGateways(*storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error)

	ListLocalDisks(*storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error)

	ListVolumeRecoveryPoints(*storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error)

	ListVolumes(*storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error)

	ResetCache(*storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error)

	RetrieveTapeArchive(*storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error)

	RetrieveTapeRecoveryPoint(*storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error)

	ShutdownGateway(*storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error)

	StartGateway(*storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error)

	UpdateBandwidthRateLimit(*storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error)

	UpdateChapCredentials(*storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error)

	UpdateGatewayInformation(*storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error)

	UpdateGatewaySoftwareNow(*storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error)

	UpdateMaintenanceStartTime(*storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error)

	UpdateSnapshotSchedule(*storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error)

	UpdateVTLDeviceType(*storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error)
}