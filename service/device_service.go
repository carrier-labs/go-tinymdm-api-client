package service

import (
	"context"
	"encoding/json"

	"github.com/carrier-labs/go-tinymdm-api-client/client"
	"github.com/carrier-labs/go-tinymdm-api-client/models"
)

type DeviceService struct {
	Client *client.Client
}

func NewDeviceService(c *client.Client) *DeviceService {
	return &DeviceService{Client: c}
}

func (s *DeviceService) GetDevices(ctx context.Context) ([]models.Device, int, *string, *string, error) {
	respBody, err := s.Client.DoRequest(ctx, "GET", "devices", nil)
	if err != nil {
		return nil, 0, nil, nil, err
	}
	type deviceListResponse struct {
		Results  []models.Device `json:"results"`
		Count    int             `json:"count"`
		Previous *string         `json:"previous"`
		Next     *string         `json:"next"`
	}
	var resp deviceListResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, 0, nil, nil, err
	}
	return resp.Results, resp.Count, resp.Previous, resp.Next, nil
}
