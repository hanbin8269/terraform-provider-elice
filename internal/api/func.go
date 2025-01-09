// Copyright (c) HashiCorp, Inc.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetInstance(ID string) (*Instance, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/instance/%s", c.HostURL, ID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	instance := Instance{}
	err = json.Unmarshal(body, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}

func (c *Client) CreateInstance(
	Title string,
	ImageId string,
	InstanceTypeId string,
	Disk int,
) (*Instance, error) {
	rb, err := json.Marshal(map[string]interface{}{
		"title":            Title,
		"image_id":         ImageId,
		"instance_type_id": InstanceTypeId,
		"disk":             Disk,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/instance", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	instance := Instance{}
	err = json.Unmarshal(body, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}
