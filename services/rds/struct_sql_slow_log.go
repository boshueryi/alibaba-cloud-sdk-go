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

// SQLSlowLog is a nested struct in rds response
type SQLSlowLog struct {
	SlowLogId                     int    `json:"SlowLogId" xml:"SlowLogId"`
	SQLId                         int    `json:"SQLId" xml:"SQLId"`
	DBName                        string `json:"DBName" xml:"DBName"`
	SQLText                       string `json:"SQLText" xml:"SQLText"`
	MySQLTotalExecutionCounts     int    `json:"MySQLTotalExecutionCounts" xml:"MySQLTotalExecutionCounts"`
	MySQLTotalExecutionTimes      int    `json:"MySQLTotalExecutionTimes" xml:"MySQLTotalExecutionTimes"`
	TotalLockTimes                int    `json:"TotalLockTimes" xml:"TotalLockTimes"`
	MaxLockTime                   int    `json:"MaxLockTime" xml:"MaxLockTime"`
	ParseTotalRowCounts           int    `json:"ParseTotalRowCounts" xml:"ParseTotalRowCounts"`
	ParseMaxRowCount              int    `json:"ParseMaxRowCount" xml:"ParseMaxRowCount"`
	ReturnTotalRowCounts          int    `json:"ReturnTotalRowCounts" xml:"ReturnTotalRowCounts"`
	ReturnMaxRowCount             int    `json:"ReturnMaxRowCount" xml:"ReturnMaxRowCount"`
	CreateTime                    string `json:"CreateTime" xml:"CreateTime"`
	SQLServerTotalExecutionCounts int    `json:"SQLServerTotalExecutionCounts" xml:"SQLServerTotalExecutionCounts"`
	SQLServerTotalExecutionTimes  int    `json:"SQLServerTotalExecutionTimes" xml:"SQLServerTotalExecutionTimes"`
	TotalLogicalReadCounts        int    `json:"TotalLogicalReadCounts" xml:"TotalLogicalReadCounts"`
	TotalPhysicalReadCounts       int    `json:"TotalPhysicalReadCounts" xml:"TotalPhysicalReadCounts"`
	ReportTime                    string `json:"ReportTime" xml:"ReportTime"`
	MaxExecutionTime              int    `json:"MaxExecutionTime" xml:"MaxExecutionTime"`
	AvgExecutionTime              int    `json:"AvgExecutionTime" xml:"AvgExecutionTime"`
}
