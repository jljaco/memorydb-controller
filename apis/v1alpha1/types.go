// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Returns the updates being applied to the ACL.
type ACLPendingChanges struct {
	UserNamesToAdd    []*string `json:"userNamesToAdd,omitempty"`
	UserNamesToRemove []*string `json:"userNamesToRemove,omitempty"`
}

// An Access Control List. You can authenticate users with Access Contol Lists.
// ACLs enable you to control cluster access by grouping users. These Access
// control lists are designed as a way to organize access to clusters.
type ACL_SDK struct {
	ARN                  *string   `json:"arn,omitempty"`
	Clusters             []*string `json:"clusters,omitempty"`
	MinimumEngineVersion *string   `json:"minimumEngineVersion,omitempty"`
	Name                 *string   `json:"name,omitempty"`
	// Returns the updates being applied to the ACL.
	PendingChanges *ACLPendingChanges `json:"pendingChanges,omitempty"`
	Status         *string            `json:"status,omitempty"`
	UserNames      []*string          `json:"userNames,omitempty"`
}

// The status of the ACL update
type ACLsUpdateStatus struct {
	ACLToApply *string `json:"aclToApply,omitempty"`
}

// Denotes the user's authentication properties, such as whether it requires
// a password to authenticate. Used in output responses.
type Authentication struct {
	PasswordCount *int64  `json:"passwordCount,omitempty"`
	Type          *string `json:"type_,omitempty"`
}

// Denotes the user's authentication properties, such as whether it requires
// a password to authenticate. Used in output responses.
type AuthenticationMode struct {
	Passwords []*ackv1alpha1.SecretKeyReference `json:"passwords,omitempty"`
	Type      *string                           `json:"type_,omitempty"`
}

// Indicates if the cluster has a Multi-AZ configuration (multiaz) or not (singleaz).
type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

// A list of cluster configuration options.
type ClusterConfiguration struct {
	Description            *string        `json:"description,omitempty"`
	EngineVersion          *string        `json:"engineVersion,omitempty"`
	MaintenanceWindow      *string        `json:"maintenanceWindow,omitempty"`
	Name                   *string        `json:"name,omitempty"`
	NodeType               *string        `json:"nodeType,omitempty"`
	NumShards              *int64         `json:"numShards,omitempty"`
	ParameterGroupName     *string        `json:"parameterGroupName,omitempty"`
	Port                   *int64         `json:"port,omitempty"`
	Shards                 []*ShardDetail `json:"shards,omitempty"`
	SnapshotRetentionLimit *int64         `json:"snapshotRetentionLimit,omitempty"`
	SnapshotWindow         *string        `json:"snapshotWindow,omitempty"`
	SubnetGroupName        *string        `json:"subnetGroupName,omitempty"`
	TopicARN               *string        `json:"topicARN,omitempty"`
	VPCID                  *string        `json:"vpcID,omitempty"`
}

// A list of updates being applied to the cluster
type ClusterPendingUpdates struct {
	// The status of the ACL update
	ACLs *ACLsUpdateStatus `json:"acls,omitempty"`
	// The status of the online resharding
	Resharding     *ReshardingStatus               `json:"resharding,omitempty"`
	ServiceUpdates []*PendingModifiedServiceUpdate `json:"serviceUpdates,omitempty"`
}

// Contains all of the attributes of a specific cluster.
type Cluster_SDK struct {
	ACLName                 *string `json:"aclName,omitempty"`
	ARN                     *string `json:"arn,omitempty"`
	AutoMinorVersionUpgrade *bool   `json:"autoMinorVersionUpgrade,omitempty"`
	AvailabilityMode        *string `json:"availabilityMode,omitempty"`
	// Represents the information required for client programs to connect to the
	// cluster and its nodes.
	ClusterEndpoint      *Endpoint `json:"clusterEndpoint,omitempty"`
	Description          *string   `json:"description,omitempty"`
	EnginePatchVersion   *string   `json:"enginePatchVersion,omitempty"`
	EngineVersion        *string   `json:"engineVersion,omitempty"`
	KMSKeyID             *string   `json:"kmsKeyID,omitempty"`
	MaintenanceWindow    *string   `json:"maintenanceWindow,omitempty"`
	Name                 *string   `json:"name,omitempty"`
	NodeType             *string   `json:"nodeType,omitempty"`
	NumberOfShards       *int64    `json:"numberOfShards,omitempty"`
	ParameterGroupName   *string   `json:"parameterGroupName,omitempty"`
	ParameterGroupStatus *string   `json:"parameterGroupStatus,omitempty"`
	// A list of updates being applied to the cluster
	PendingUpdates         *ClusterPendingUpdates     `json:"pendingUpdates,omitempty"`
	SecurityGroups         []*SecurityGroupMembership `json:"securityGroups,omitempty"`
	Shards                 []*Shard                   `json:"shards,omitempty"`
	SnapshotRetentionLimit *int64                     `json:"snapshotRetentionLimit,omitempty"`
	SnapshotWindow         *string                    `json:"snapshotWindow,omitempty"`
	SNSTopicARN            *string                    `json:"snsTopicARN,omitempty"`
	SNSTopicStatus         *string                    `json:"snsTopicStatus,omitempty"`
	Status                 *string                    `json:"status,omitempty"`
	SubnetGroupName        *string                    `json:"subnetGroupName,omitempty"`
	TLSEnabled             *bool                      `json:"tlsEnabled,omitempty"`
}

// Represents the information required for client programs to connect to the
// cluster and its nodes.
type Endpoint struct {
	Address *string `json:"address,omitempty"`
	Port    *int64  `json:"port,omitempty"`
}

// Provides details of the Redis engine version
type EngineVersionInfo struct {
	EnginePatchVersion   *string `json:"enginePatchVersion,omitempty"`
	EngineVersion        *string `json:"engineVersion,omitempty"`
	ParameterGroupFamily *string `json:"parameterGroupFamily,omitempty"`
}

// Represents a single occurrence of something interesting within the system.
// Some examples of events are creating a cluster or adding or removing a node.
type Event struct {
	Date       *metav1.Time `json:"date,omitempty"`
	Message    *string      `json:"message,omitempty"`
	SourceName *string      `json:"sourceName,omitempty"`
}

// Used to streamline results of a search based on the property being filtered.
type Filter struct {
	Name   *string   `json:"name,omitempty"`
	Values []*string `json:"values,omitempty"`
}

// Represents an individual node within a cluster. Each node runs its own instance
// of the cluster's protocol-compliant caching software.
type Node struct {
	AvailabilityZone *string      `json:"availabilityZone,omitempty"`
	CreateTime       *metav1.Time `json:"createTime,omitempty"`
	// Represents the information required for client programs to connect to the
	// cluster and its nodes.
	Endpoint *Endpoint `json:"endpoint,omitempty"`
	Name     *string   `json:"name,omitempty"`
	Status   *string   `json:"status,omitempty"`
}

// Describes an individual setting that controls some aspect of MemoryDB behavior.
type Parameter struct {
	AllowedValues        *string `json:"allowedValues,omitempty"`
	DataType             *string `json:"dataType,omitempty"`
	Description          *string `json:"description,omitempty"`
	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Value                *string `json:"value,omitempty"`
}

// Represents the output of a CreateParameterGroup operation. A parameter group
// represents a combination of specific values for the parameters that are passed
// to the engine software during startup.
type ParameterGroup_SDK struct {
	ARN         *string `json:"arn,omitempty"`
	Description *string `json:"description,omitempty"`
	Family      *string `json:"family,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// Describes a name-value pair that is used to update the value of a parameter.
type ParameterNameValue struct {
	ParameterName  *string `json:"parameterName,omitempty"`
	ParameterValue *string `json:"parameterValue,omitempty"`
}

// Update action that has yet to be processed for the corresponding apply/stop
// request
type PendingModifiedServiceUpdate struct {
	ServiceUpdateName *string `json:"serviceUpdateName,omitempty"`
	Status            *string `json:"status,omitempty"`
}

// A request to configure the number of replicas in a shard
type ReplicaConfigurationRequest struct {
	ReplicaCount *int64 `json:"replicaCount,omitempty"`
}

// The status of the online resharding
type ReshardingStatus struct {
	// Represents the progress of an online resharding operation.
	SlotMigration *SlotMigration `json:"slotMigration,omitempty"`
}

// Represents a single security group and its status.
type SecurityGroupMembership struct {
	SecurityGroupID *string `json:"securityGroupID,omitempty"`
	Status          *string `json:"status,omitempty"`
}

// An update that you can apply to your MemoryDB clusters.
type ServiceUpdate struct {
	AutoUpdateStartDate *metav1.Time `json:"autoUpdateStartDate,omitempty"`
	ClusterName         *string      `json:"clusterName,omitempty"`
	Description         *string      `json:"description,omitempty"`
	NodesUpdated        *string      `json:"nodesUpdated,omitempty"`
	ReleaseDate         *metav1.Time `json:"releaseDate,omitempty"`
	ServiceUpdateName   *string      `json:"serviceUpdateName,omitempty"`
	Status              *string      `json:"status,omitempty"`
}

// A request to apply a service update
type ServiceUpdateRequest struct {
	ServiceUpdateNameToApply *string `json:"serviceUpdateNameToApply,omitempty"`
}

// Represents a collection of nodes in a cluster. One node in the node group
// is the read/write primary node. All the other nodes are read-only Replica
// nodes.
type Shard struct {
	Name          *string `json:"name,omitempty"`
	Nodes         []*Node `json:"nodes,omitempty"`
	NumberOfNodes *int64  `json:"numberOfNodes,omitempty"`
	Slots         *string `json:"slots,omitempty"`
	Status        *string `json:"status,omitempty"`
}

// Shard configuration options. Each shard configuration has the following:
// Slots and ReplicaCount.
type ShardConfiguration struct {
	ReplicaCount *int64  `json:"replicaCount,omitempty"`
	Slots        *string `json:"slots,omitempty"`
}

// A request to configure the sharding properties of a cluster
type ShardConfigurationRequest struct {
	ShardCount *int64 `json:"shardCount,omitempty"`
}

// Provides details of a shard in a snapshot
type ShardDetail struct {
	// Shard configuration options. Each shard configuration has the following:
	// Slots and ReplicaCount.
	Configuration        *ShardConfiguration `json:"configuration,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	Size                 *string             `json:"size,omitempty"`
	SnapshotCreationTime *metav1.Time        `json:"snapshotCreationTime,omitempty"`
}

// Represents the progress of an online resharding operation.
type SlotMigration struct {
	ProgressPercentage *float64 `json:"progressPercentage,omitempty"`
}

// Represents a copy of an entire cluster as of the time when the snapshot was
// taken.
type Snapshot_SDK struct {
	ARN *string `json:"arn,omitempty"`
	// A list of cluster configuration options.
	ClusterConfiguration *ClusterConfiguration `json:"clusterConfiguration,omitempty"`
	KMSKeyID             *string               `json:"kmsKeyID,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	Source               *string               `json:"source,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

// Represents the subnet associated with a cluster. This parameter refers to
// subnets defined in Amazon Virtual Private Cloud (Amazon VPC) and used with
// MemoryDB.
type Subnet struct {
	// Indicates if the cluster has a Multi-AZ configuration (multiaz) or not (singleaz).
	AvailabilityZone *AvailabilityZone `json:"availabilityZone,omitempty"`
	Identifier       *string           `json:"identifier,omitempty"`
}

// Represents the output of one of the following operations:
//
//   - CreateSubnetGroup
//
//   - UpdateSubnetGroup
//
// A subnet group is a collection of subnets (typically private) that you can
// designate for your clusters running in an Amazon Virtual Private Cloud (VPC)
// environment.
type SubnetGroup_SDK struct {
	ARN         *string   `json:"arn,omitempty"`
	Description *string   `json:"description,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Subnets     []*Subnet `json:"subnets,omitempty"`
	VPCID       *string   `json:"vpcID,omitempty"`
}

// A tag that can be added to an MemoryDB resource. Tags are composed of a Key/Value
// pair. You can use tags to categorize and track all your MemoryDB resources.
// When you add or remove tags on clusters, those actions will be replicated
// to all nodes in the cluster. A tag with a null Value is permitted. For more
// information, see Tagging your MemoryDB resources (https://docs.aws.amazon.com/MemoryDB/latest/devguide/tagging-resources.html)
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// A cluster whose updates have failed
type UnprocessedCluster struct {
	ClusterName  *string `json:"clusterName,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorType    *string `json:"errorType,omitempty"`
}

// You create users and assign them specific permissions by using an access
// string. You assign the users to Access Control Lists aligned with a specific
// role (administrators, human resources) that are then deployed to one or more
// MemoryDB clusters.
type User_SDK struct {
	ACLNames     []*string `json:"aclNames,omitempty"`
	ARN          *string   `json:"arn,omitempty"`
	AccessString *string   `json:"accessString,omitempty"`
	// Denotes the user's authentication properties, such as whether it requires
	// a password to authenticate. Used in output responses.
	Authentication       *Authentication `json:"authentication,omitempty"`
	MinimumEngineVersion *string         `json:"minimumEngineVersion,omitempty"`
	Name                 *string         `json:"name,omitempty"`
	Status               *string         `json:"status,omitempty"`
}
