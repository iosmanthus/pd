// Copyright 2017 TiKV Project Authors.
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

package constant

// PriorityLevel lower level means higher priority
type PriorityLevel int

// Built-in priority level
const (
	Low PriorityLevel = iota
	Medium
	High
	Urgent

	PriorityLevelLen
)

func (p PriorityLevel) String() string {
	switch p {
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	case Urgent:
		return "urgent"
	default:
		return "unknown"
	}
}

// ScheduleKind distinguishes resources and schedule policy.
type ScheduleKind struct {
	Resource ResourceKind
	Policy   SchedulePolicy
}

// NewScheduleKind creates a schedule kind with resource kind and schedule policy.
func NewScheduleKind(resource ResourceKind, policy SchedulePolicy) ScheduleKind {
	return ScheduleKind{
		Resource: resource,
		Policy:   policy,
	}
}

// ResourceKind distinguishes different kinds of resources.
type ResourceKind int

const (
	// LeaderKind indicates the leader kind resource
	LeaderKind ResourceKind = iota
	// RegionKind indicates the region kind resource
	RegionKind
	// WitnessKind indicates the witness kind resource
	WitnessKind
)

func (k ResourceKind) String() string {
	switch k {
	case LeaderKind:
		return "leader"
	case RegionKind:
		return "region"
	case WitnessKind:
		return "witness"
	default:
		return "unknown"
	}
}

// SchedulePolicy distinguishes different kinds of schedule policies.
type SchedulePolicy int

const (
	// ByCount indicates that balance by count
	ByCount SchedulePolicy = iota
	// BySize indicates that balance by size
	BySize
)

func (k SchedulePolicy) String() string {
	switch k {
	case ByCount:
		return "count"
	case BySize:
		return "size"
	default:
		return "unknown"
	}
}

// StringToSchedulePolicy creates a schedule policy with string.
func StringToSchedulePolicy(input string) SchedulePolicy {
	switch input {
	case BySize.String():
		return BySize
	case ByCount.String():
		return ByCount
	default:
		panic("invalid schedule policy: " + input)
	}
}

// KeyType distinguishes different kinds of key types
type KeyType int

const (
	// Table indicates that the key is table key
	Table KeyType = iota
	// Raw indicates that the key is raw key.
	Raw
	// Txn indicates that the key is txn key.
	Txn
)

func (k KeyType) String() string {
	switch k {
	case Table:
		return "table"
	case Raw:
		return "raw"
	case Txn:
		return "txn"
	default:
		return "unknown"
	}
}

// StringToKeyType creates a key type with string.
func StringToKeyType(input string) KeyType {
	switch input {
	case Table.String():
		return Table
	case Raw.String():
		return Raw
	case Txn.String():
		return Txn
	default:
		panic("invalid key type: " + input)
	}
}
