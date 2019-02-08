/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package telemetry

import "time"

type Telemetry interface {
	PollCpu() (cpuSnapshot CpuUsage)
	AvgCpuUsage(init, final CpuUsage) (avg float64)
}

type SystemUsage struct {
	Memory memoryUsage
	CpuBusyAvg float64
}

type memoryUsage struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects uint64
}

type CpuUsage struct {
	Busy, // time used by all processes. this ideally does not include system processes.
	Idle, // time used by the idle process
	Total uint64 // reported sum total of all usage
}

func GetCpuUsageAverage(usageAvg *float64, lastSample *CpuUsage) {
	for {
		nextUsage := PollCpu()
		*usageAvg = AvgCpuUsage(*lastSample, nextUsage)
		*lastSample = nextUsage

		time.Sleep(time.Second * 30)
	}
}
