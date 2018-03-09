package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DBInstanceAttribute is a nested struct in rds response
type DBInstanceAttribute struct {
	InsId                             int                                                `json:"InsId" xml:"InsId"`
	ReplicateId                       string                                             `json:"ReplicateId" xml:"ReplicateId"`
	PayType                           string                                             `json:"PayType" xml:"PayType"`
	DBInstanceDiskUsed                string                                             `json:"DBInstanceDiskUsed" xml:"DBInstanceDiskUsed"`
	VpcCloudInstanceId                string                                             `json:"VpcCloudInstanceId" xml:"VpcCloudInstanceId"`
	InstanceNetworkType               string                                             `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	EngineVersion                     string                                             `json:"EngineVersion" xml:"EngineVersion"`
	ConnectionMode                    string                                             `json:"ConnectionMode" xml:"ConnectionMode"`
	ZoneId                            string                                             `json:"ZoneId" xml:"ZoneId"`
	TempUpgradeTimeStart              string                                             `json:"TempUpgradeTimeStart" xml:"TempUpgradeTimeStart"`
	MaxIOPS                           int                                                `json:"MaxIOPS" xml:"MaxIOPS"`
	DBInstanceStatus                  string                                             `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	TempUpgradeRecoveryMaxConnections string                                             `json:"TempUpgradeRecoveryMaxConnections" xml:"TempUpgradeRecoveryMaxConnections"`
	MaxConnections                    int                                                `json:"MaxConnections" xml:"MaxConnections"`
	AccountMaxQuantity                int                                                `json:"AccountMaxQuantity" xml:"AccountMaxQuantity"`
	DBInstanceStorage                 int                                                `json:"DBInstanceStorage" xml:"DBInstanceStorage"`
	TempUpgradeTimeEnd                string                                             `json:"TempUpgradeTimeEnd" xml:"TempUpgradeTimeEnd"`
	DBInstanceMemory                  int                                                `json:"DBInstanceMemory" xml:"DBInstanceMemory"`
	CreationTime                      string                                             `json:"CreationTime" xml:"CreationTime"`
	ResourceGroupId                   string                                             `json:"ResourceGroupId" xml:"ResourceGroupId"`
	TempUpgradeRecoveryMemory         int                                                `json:"TempUpgradeRecoveryMemory" xml:"TempUpgradeRecoveryMemory"`
	LockReason                        string                                             `json:"LockReason" xml:"LockReason"`
	AccountType                       string                                             `json:"AccountType" xml:"AccountType"`
	GuardDBInstanceName               string                                             `json:"GuardDBInstanceName" xml:"GuardDBInstanceName"`
	LockMode                          string                                             `json:"LockMode" xml:"LockMode"`
	ConnectionString                  string                                             `json:"ConnectionString" xml:"ConnectionString"`
	Port                              string                                             `json:"Port" xml:"Port"`
	TempDBInstanceId                  string                                             `json:"TempDBInstanceId" xml:"TempDBInstanceId"`
	DBInstanceType                    string                                             `json:"DBInstanceType" xml:"DBInstanceType"`
	VSwitchId                         string                                             `json:"VSwitchId" xml:"VSwitchId"`
	AvailabilityValue                 string                                             `json:"AvailabilityValue" xml:"AvailabilityValue"`
	MaintainTime                      string                                             `json:"MaintainTime" xml:"MaintainTime"`
	Category                          string                                             `json:"Category" xml:"Category"`
	TempUpgradeRecoveryTime           string                                             `json:"TempUpgradeRecoveryTime" xml:"TempUpgradeRecoveryTime"`
	IncrementSourceDBInstanceId       string                                             `json:"IncrementSourceDBInstanceId" xml:"IncrementSourceDBInstanceId"`
	TempUpgradeRecoveryCpu            int                                                `json:"TempUpgradeRecoveryCpu" xml:"TempUpgradeRecoveryCpu"`
	DBInstanceNetType                 string                                             `json:"DBInstanceNetType" xml:"DBInstanceNetType"`
	Tags                              string                                             `json:"Tags" xml:"Tags"`
	DBInstanceCPU                     string                                             `json:"DBInstanceCPU" xml:"DBInstanceCPU"`
	TempUpgradeRecoveryClass          string                                             `json:"TempUpgradeRecoveryClass" xml:"TempUpgradeRecoveryClass"`
	GuardDBInstanceId                 string                                             `json:"GuardDBInstanceId" xml:"GuardDBInstanceId"`
	DBInstanceDescription             string                                             `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	Engine                            string                                             `json:"Engine" xml:"Engine"`
	ReadDelayTime                     string                                             `json:"ReadDelayTime" xml:"ReadDelayTime"`
	TempUpgradeRecoveryMaxIOPS        string                                             `json:"TempUpgradeRecoveryMaxIOPS" xml:"TempUpgradeRecoveryMaxIOPS"`
	DBMaxQuantity                     int                                                `json:"DBMaxQuantity" xml:"DBMaxQuantity"`
	MasterInstanceId                  string                                             `json:"MasterInstanceId" xml:"MasterInstanceId"`
	VpcId                             string                                             `json:"VpcId" xml:"VpcId"`
	DBInstanceClass                   string                                             `json:"DBInstanceClass" xml:"DBInstanceClass"`
	SupportUpgradeAccountType         string                                             `json:"SupportUpgradeAccountType" xml:"SupportUpgradeAccountType"`
	DBInstanceId                      string                                             `json:"DBInstanceId" xml:"DBInstanceId"`
	CanTempUpgrade                    bool                                               `json:"CanTempUpgrade" xml:"CanTempUpgrade"`
	ExpireTime                        string                                             `json:"ExpireTime" xml:"ExpireTime"`
	RegionId                          string                                             `json:"RegionId" xml:"RegionId"`
	AdvancedFeatures                  string                                             `json:"AdvancedFeatures" xml:"AdvancedFeatures"`
	DBInstanceClassType               string                                             `json:"DBInstanceClassType" xml:"DBInstanceClassType"`
	SecurityIPList                    string                                             `json:"SecurityIPList" xml:"SecurityIPList"`
	ReadOnlyDBInstanceIds             ReadOnlyDBInstanceIdsInDescribeDBInstanceAttribute `json:"ReadOnlyDBInstanceIds" xml:"ReadOnlyDBInstanceIds"`
}
