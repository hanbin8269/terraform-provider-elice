// Copyright (c) HashiCorp, Inc.

package api

type Instance struct {
	Id             string `json:"id,omitempty"`
	Title          string `json:"title,omitempty"`
	ImageId        string `json:"image_id,omitempty"`
	InstanceTypeId string `json:"instance_type_id,omitempty"`
	Disk           int    `json:"disk,omitempty"`
}
