// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package clouddms aliases all exported identifiers in package
// "cloud.google.com/go/clouddms/apiv1/clouddmspb".
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package clouddms

import (
	src "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/clouddms/apiv1/clouddmspb
const (
	CloudSqlSettings_ALWAYS                                                = src.CloudSqlSettings_ALWAYS
	CloudSqlSettings_MYSQL_5_6                                             = src.CloudSqlSettings_MYSQL_5_6
	CloudSqlSettings_MYSQL_5_7                                             = src.CloudSqlSettings_MYSQL_5_7
	CloudSqlSettings_MYSQL_8_0                                             = src.CloudSqlSettings_MYSQL_8_0
	CloudSqlSettings_NEVER                                                 = src.CloudSqlSettings_NEVER
	CloudSqlSettings_PD_HDD                                                = src.CloudSqlSettings_PD_HDD
	CloudSqlSettings_PD_SSD                                                = src.CloudSqlSettings_PD_SSD
	CloudSqlSettings_POSTGRES_10                                           = src.CloudSqlSettings_POSTGRES_10
	CloudSqlSettings_POSTGRES_11                                           = src.CloudSqlSettings_POSTGRES_11
	CloudSqlSettings_POSTGRES_12                                           = src.CloudSqlSettings_POSTGRES_12
	CloudSqlSettings_POSTGRES_13                                           = src.CloudSqlSettings_POSTGRES_13
	CloudSqlSettings_POSTGRES_9_6                                          = src.CloudSqlSettings_POSTGRES_9_6
	CloudSqlSettings_SQL_ACTIVATION_POLICY_UNSPECIFIED                     = src.CloudSqlSettings_SQL_ACTIVATION_POLICY_UNSPECIFIED
	CloudSqlSettings_SQL_DATABASE_VERSION_UNSPECIFIED                      = src.CloudSqlSettings_SQL_DATABASE_VERSION_UNSPECIFIED
	CloudSqlSettings_SQL_DATA_DISK_TYPE_UNSPECIFIED                        = src.CloudSqlSettings_SQL_DATA_DISK_TYPE_UNSPECIFIED
	ConnectionProfile_CREATING                                             = src.ConnectionProfile_CREATING
	ConnectionProfile_DELETED                                              = src.ConnectionProfile_DELETED
	ConnectionProfile_DELETING                                             = src.ConnectionProfile_DELETING
	ConnectionProfile_DRAFT                                                = src.ConnectionProfile_DRAFT
	ConnectionProfile_FAILED                                               = src.ConnectionProfile_FAILED
	ConnectionProfile_READY                                                = src.ConnectionProfile_READY
	ConnectionProfile_STATE_UNSPECIFIED                                    = src.ConnectionProfile_STATE_UNSPECIFIED
	ConnectionProfile_UPDATING                                             = src.ConnectionProfile_UPDATING
	DatabaseEngine_DATABASE_ENGINE_UNSPECIFIED                             = src.DatabaseEngine_DATABASE_ENGINE_UNSPECIFIED
	DatabaseEngine_MYSQL                                                   = src.DatabaseEngine_MYSQL
	DatabaseEngine_POSTGRESQL                                              = src.DatabaseEngine_POSTGRESQL
	DatabaseProvider_CLOUDSQL                                              = src.DatabaseProvider_CLOUDSQL
	DatabaseProvider_DATABASE_PROVIDER_UNSPECIFIED                         = src.DatabaseProvider_DATABASE_PROVIDER_UNSPECIFIED
	DatabaseProvider_RDS                                                   = src.DatabaseProvider_RDS
	MigrationJobVerificationError_AUTHENTICATION_FAILURE                   = src.MigrationJobVerificationError_AUTHENTICATION_FAILURE
	MigrationJobVerificationError_CANT_RESTART_RUNNING_MIGRATION           = src.MigrationJobVerificationError_CANT_RESTART_RUNNING_MIGRATION
	MigrationJobVerificationError_CONNECTION_FAILURE                       = src.MigrationJobVerificationError_CONNECTION_FAILURE
	MigrationJobVerificationError_CONNECTION_PROFILE_TYPES_INCOMPATIBILITY = src.MigrationJobVerificationError_CONNECTION_PROFILE_TYPES_INCOMPATIBILITY
	MigrationJobVerificationError_ERROR_CODE_UNSPECIFIED                   = src.MigrationJobVerificationError_ERROR_CODE_UNSPECIFIED
	MigrationJobVerificationError_INSUFFICIENT_MAX_REPLICATION_SLOTS       = src.MigrationJobVerificationError_INSUFFICIENT_MAX_REPLICATION_SLOTS
	MigrationJobVerificationError_INSUFFICIENT_MAX_WAL_SENDERS             = src.MigrationJobVerificationError_INSUFFICIENT_MAX_WAL_SENDERS
	MigrationJobVerificationError_INSUFFICIENT_MAX_WORKER_PROCESSES        = src.MigrationJobVerificationError_INSUFFICIENT_MAX_WORKER_PROCESSES
	MigrationJobVerificationError_INVALID_CONNECTION_PROFILE_CONFIG        = src.MigrationJobVerificationError_INVALID_CONNECTION_PROFILE_CONFIG
	MigrationJobVerificationError_INVALID_RDS_LOGICAL_REPLICATION          = src.MigrationJobVerificationError_INVALID_RDS_LOGICAL_REPLICATION
	MigrationJobVerificationError_INVALID_SHARED_PRELOAD_LIBRARY           = src.MigrationJobVerificationError_INVALID_SHARED_PRELOAD_LIBRARY
	MigrationJobVerificationError_INVALID_WAL_LEVEL                        = src.MigrationJobVerificationError_INVALID_WAL_LEVEL
	MigrationJobVerificationError_NO_PGLOGICAL_INSTALLED                   = src.MigrationJobVerificationError_NO_PGLOGICAL_INSTALLED
	MigrationJobVerificationError_PGLOGICAL_NODE_ALREADY_EXISTS            = src.MigrationJobVerificationError_PGLOGICAL_NODE_ALREADY_EXISTS
	MigrationJobVerificationError_UNSUPPORTED_DEFINER                      = src.MigrationJobVerificationError_UNSUPPORTED_DEFINER
	MigrationJobVerificationError_UNSUPPORTED_EXTENSIONS                   = src.MigrationJobVerificationError_UNSUPPORTED_EXTENSIONS
	MigrationJobVerificationError_UNSUPPORTED_GTID_MODE                    = src.MigrationJobVerificationError_UNSUPPORTED_GTID_MODE
	MigrationJobVerificationError_UNSUPPORTED_MIGRATION_TYPE               = src.MigrationJobVerificationError_UNSUPPORTED_MIGRATION_TYPE
	MigrationJobVerificationError_UNSUPPORTED_TABLE_DEFINITION             = src.MigrationJobVerificationError_UNSUPPORTED_TABLE_DEFINITION
	MigrationJobVerificationError_VERSION_INCOMPATIBILITY                  = src.MigrationJobVerificationError_VERSION_INCOMPATIBILITY
	MigrationJob_CDC                                                       = src.MigrationJob_CDC
	MigrationJob_COMPLETED                                                 = src.MigrationJob_COMPLETED
	MigrationJob_CONTINUOUS                                                = src.MigrationJob_CONTINUOUS
	MigrationJob_CREATING                                                  = src.MigrationJob_CREATING
	MigrationJob_DELETED                                                   = src.MigrationJob_DELETED
	MigrationJob_DELETING                                                  = src.MigrationJob_DELETING
	MigrationJob_DRAFT                                                     = src.MigrationJob_DRAFT
	MigrationJob_FAILED                                                    = src.MigrationJob_FAILED
	MigrationJob_FULL_DUMP                                                 = src.MigrationJob_FULL_DUMP
	MigrationJob_MAINTENANCE                                               = src.MigrationJob_MAINTENANCE
	MigrationJob_NOT_STARTED                                               = src.MigrationJob_NOT_STARTED
	MigrationJob_ONE_TIME                                                  = src.MigrationJob_ONE_TIME
	MigrationJob_PHASE_UNSPECIFIED                                         = src.MigrationJob_PHASE_UNSPECIFIED
	MigrationJob_PREPARING_THE_DUMP                                        = src.MigrationJob_PREPARING_THE_DUMP
	MigrationJob_PROMOTE_IN_PROGRESS                                       = src.MigrationJob_PROMOTE_IN_PROGRESS
	MigrationJob_RESTARTING                                                = src.MigrationJob_RESTARTING
	MigrationJob_RESUMING                                                  = src.MigrationJob_RESUMING
	MigrationJob_RUNNING                                                   = src.MigrationJob_RUNNING
	MigrationJob_STARTING                                                  = src.MigrationJob_STARTING
	MigrationJob_STATE_UNSPECIFIED                                         = src.MigrationJob_STATE_UNSPECIFIED
	MigrationJob_STOPPED                                                   = src.MigrationJob_STOPPED
	MigrationJob_STOPPING                                                  = src.MigrationJob_STOPPING
	MigrationJob_TYPE_UNSPECIFIED                                          = src.MigrationJob_TYPE_UNSPECIFIED
	MigrationJob_UPDATING                                                  = src.MigrationJob_UPDATING
	MigrationJob_WAITING_FOR_SOURCE_WRITES_TO_STOP                         = src.MigrationJob_WAITING_FOR_SOURCE_WRITES_TO_STOP
	SslConfig_SERVER_CLIENT                                                = src.SslConfig_SERVER_CLIENT
	SslConfig_SERVER_ONLY                                                  = src.SslConfig_SERVER_ONLY
	SslConfig_SSL_TYPE_UNSPECIFIED                                         = src.SslConfig_SSL_TYPE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/clouddms/apiv1/clouddmspb
var (
	CloudSqlSettings_SqlActivationPolicy_name              = src.CloudSqlSettings_SqlActivationPolicy_name
	CloudSqlSettings_SqlActivationPolicy_value             = src.CloudSqlSettings_SqlActivationPolicy_value
	CloudSqlSettings_SqlDataDiskType_name                  = src.CloudSqlSettings_SqlDataDiskType_name
	CloudSqlSettings_SqlDataDiskType_value                 = src.CloudSqlSettings_SqlDataDiskType_value
	CloudSqlSettings_SqlDatabaseVersion_name               = src.CloudSqlSettings_SqlDatabaseVersion_name
	CloudSqlSettings_SqlDatabaseVersion_value              = src.CloudSqlSettings_SqlDatabaseVersion_value
	ConnectionProfile_State_name                           = src.ConnectionProfile_State_name
	ConnectionProfile_State_value                          = src.ConnectionProfile_State_value
	DatabaseEngine_name                                    = src.DatabaseEngine_name
	DatabaseEngine_value                                   = src.DatabaseEngine_value
	DatabaseProvider_name                                  = src.DatabaseProvider_name
	DatabaseProvider_value                                 = src.DatabaseProvider_value
	File_google_cloud_clouddms_v1_clouddms_proto           = src.File_google_cloud_clouddms_v1_clouddms_proto
	File_google_cloud_clouddms_v1_clouddms_resources_proto = src.File_google_cloud_clouddms_v1_clouddms_resources_proto
	MigrationJobVerificationError_ErrorCode_name           = src.MigrationJobVerificationError_ErrorCode_name
	MigrationJobVerificationError_ErrorCode_value          = src.MigrationJobVerificationError_ErrorCode_value
	MigrationJob_Phase_name                                = src.MigrationJob_Phase_name
	MigrationJob_Phase_value                               = src.MigrationJob_Phase_value
	MigrationJob_State_name                                = src.MigrationJob_State_name
	MigrationJob_State_value                               = src.MigrationJob_State_value
	MigrationJob_Type_name                                 = src.MigrationJob_Type_name
	MigrationJob_Type_value                                = src.MigrationJob_Type_value
	SslConfig_SslType_name                                 = src.SslConfig_SslType_name
	SslConfig_SslType_value                                = src.SslConfig_SslType_value
)

// Specifies required connection parameters, and, optionally, the parameters
// required to create a Cloud SQL destination database instance.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type CloudSqlConnectionProfile = src.CloudSqlConnectionProfile

// Settings for creating a Cloud SQL database instance.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type CloudSqlSettings = src.CloudSqlSettings

// Specifies when the instance should be activated.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type CloudSqlSettings_SqlActivationPolicy = src.CloudSqlSettings_SqlActivationPolicy

// The storage options for Cloud SQL databases.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type CloudSqlSettings_SqlDataDiskType = src.CloudSqlSettings_SqlDataDiskType

// The database engine type and version.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type CloudSqlSettings_SqlDatabaseVersion = src.CloudSqlSettings_SqlDatabaseVersion

// A connection profile definition.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type ConnectionProfile = src.ConnectionProfile
type ConnectionProfile_Cloudsql = src.ConnectionProfile_Cloudsql
type ConnectionProfile_Mysql = src.ConnectionProfile_Mysql
type ConnectionProfile_Postgresql = src.ConnectionProfile_Postgresql

// The current connection profile state (e.g. DRAFT, READY, or FAILED).
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type ConnectionProfile_State = src.ConnectionProfile_State

// Request message for 'CreateConnectionProfile' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type CreateConnectionProfileRequest = src.CreateConnectionProfileRequest

// Request message to create a new Database Migration Service migration job in
// the specified project and region.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type CreateMigrationJobRequest = src.CreateMigrationJobRequest

// DataMigrationServiceClient is the client API for DataMigrationService
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type DataMigrationServiceClient = src.DataMigrationServiceClient

// DataMigrationServiceServer is the server API for DataMigrationService
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type DataMigrationServiceServer = src.DataMigrationServiceServer

// The database engine types.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type DatabaseEngine = src.DatabaseEngine

// The database providers.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type DatabaseProvider = src.DatabaseProvider

// A message defining the database engine and provider.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type DatabaseType = src.DatabaseType

// Request message for 'DeleteConnectionProfile' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type DeleteConnectionProfileRequest = src.DeleteConnectionProfileRequest

// Request message for 'DeleteMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type DeleteMigrationJobRequest = src.DeleteMigrationJobRequest

// Request message for 'GenerateSshScript' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type GenerateSshScriptRequest = src.GenerateSshScriptRequest
type GenerateSshScriptRequest_VmCreationConfig = src.GenerateSshScriptRequest_VmCreationConfig
type GenerateSshScriptRequest_VmSelectionConfig = src.GenerateSshScriptRequest_VmSelectionConfig

// Request message for 'GetConnectionProfile' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type GetConnectionProfileRequest = src.GetConnectionProfileRequest

// Request message for 'GetMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type GetMigrationJobRequest = src.GetMigrationJobRequest

// Request message for 'ListConnectionProfiles' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type ListConnectionProfilesRequest = src.ListConnectionProfilesRequest

// Response message for 'ListConnectionProfiles' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type ListConnectionProfilesResponse = src.ListConnectionProfilesResponse

// Retrieve a list of all migration jobs in a given project and location.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type ListMigrationJobsRequest = src.ListMigrationJobsRequest

// Response message for 'ListMigrationJobs' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type ListMigrationJobsResponse = src.ListMigrationJobsResponse

// Represents a Database Migration Service migration job object.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type MigrationJob = src.MigrationJob

// Error message of a verification Migration job.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type MigrationJobVerificationError = src.MigrationJobVerificationError

// A general error code describing the type of error that occurred.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type MigrationJobVerificationError_ErrorCode = src.MigrationJobVerificationError_ErrorCode

// The current migration job phase.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type MigrationJob_Phase = src.MigrationJob_Phase
type MigrationJob_ReverseSshConnectivity = src.MigrationJob_ReverseSshConnectivity

// The current migration job states.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type MigrationJob_State = src.MigrationJob_State
type MigrationJob_StaticIpConnectivity = src.MigrationJob_StaticIpConnectivity

// The type of migration job (one-time or continuous).
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type MigrationJob_Type = src.MigrationJob_Type
type MigrationJob_VpcPeeringConnectivity = src.MigrationJob_VpcPeeringConnectivity

// Specifies connection parameters required specifically for MySQL databases.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type MySqlConnectionProfile = src.MySqlConnectionProfile

// Represents the metadata of the long-running operation.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type OperationMetadata = src.OperationMetadata

// Specifies connection parameters required specifically for PostgreSQL
// databases.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type PostgreSqlConnectionProfile = src.PostgreSqlConnectionProfile

// Request message for 'PromoteMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type PromoteMigrationJobRequest = src.PromoteMigrationJobRequest

// Request message for 'RestartMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type RestartMigrationJobRequest = src.RestartMigrationJobRequest

// Request message for 'ResumeMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type ResumeMigrationJobRequest = src.ResumeMigrationJobRequest

// The details needed to configure a reverse SSH tunnel between the source and
// destination databases. These details will be used when calling the
// generateSshScript method (see
// https://cloud.google.com/database-migration/docs/reference/rest/v1/projects.locations.migrationJobs/generateSshScript)
// to produce the script that will help set up the reverse SSH tunnel, and to
// set up the VPC peering between the Cloud SQL private network and the VPC.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type ReverseSshConnectivity = src.ReverseSshConnectivity

// An entry for an Access Control list.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type SqlAclEntry = src.SqlAclEntry
type SqlAclEntry_ExpireTime = src.SqlAclEntry_ExpireTime
type SqlAclEntry_Ttl = src.SqlAclEntry_Ttl

// IP Management configuration.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type SqlIpConfig = src.SqlIpConfig

// Response message for 'GenerateSshScript' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type SshScript = src.SshScript

// SSL configuration information.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type SslConfig = src.SslConfig

// Specifies The kind of ssl configuration used.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type SslConfig_SslType = src.SslConfig_SslType

// Request message for 'StartMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type StartMigrationJobRequest = src.StartMigrationJobRequest

// The source database will allow incoming connections from the destination
// database's public IP. You can retrieve the Cloud SQL instance's public IP
// from the Cloud SQL console or using Cloud SQL APIs. No additional
// configuration is required.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type StaticIpConnectivity = src.StaticIpConnectivity

// Request message for 'StopMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type StopMigrationJobRequest = src.StopMigrationJobRequest

// UnimplementedDataMigrationServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type UnimplementedDataMigrationServiceServer = src.UnimplementedDataMigrationServiceServer

// Request message for 'UpdateConnectionProfile' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type UpdateConnectionProfileRequest = src.UpdateConnectionProfileRequest

// Request message for 'UpdateMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type UpdateMigrationJobRequest = src.UpdateMigrationJobRequest

// Request message for 'VerifyMigrationJob' request.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type VerifyMigrationJobRequest = src.VerifyMigrationJobRequest

// VM creation configuration message
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type VmCreationConfig = src.VmCreationConfig

// VM selection configuration message
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type VmSelectionConfig = src.VmSelectionConfig

// The details of the VPC where the source database is located in Google
// Cloud. We will use this information to set up the VPC peering connection
// between Cloud SQL and this VPC.
//
// Deprecated: Please use types in: cloud.google.com/go/clouddms/apiv1/clouddmspb
type VpcPeeringConnectivity = src.VpcPeeringConnectivity

// Deprecated: Please use funcs in: cloud.google.com/go/clouddms/apiv1/clouddmspb
func NewDataMigrationServiceClient(cc grpc.ClientConnInterface) DataMigrationServiceClient {
	return src.NewDataMigrationServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/clouddms/apiv1/clouddmspb
func RegisterDataMigrationServiceServer(s *grpc.Server, srv DataMigrationServiceServer) {
	src.RegisterDataMigrationServiceServer(s, srv)
}
