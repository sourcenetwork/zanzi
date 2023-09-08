package domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewPolicyRecord(policy *Policy, data []byte) *PolicyRecord {
	return &PolicyRecord{
		Policy:    policy,
		AppData:   data,
		CreatedAt: timestamppb.Now(),
	}
}
