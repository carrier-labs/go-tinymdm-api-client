package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/carrier-labs/go-tinymdm-api-client/client"
	"github.com/carrier-labs/go-tinymdm-api-client/models"
)

type DeviceService struct {
	Client *client.Client
}

// DeviceListParams allows paginating the devices endpoint.
type DeviceListParams struct {
	PerPage int
}

func NewDeviceService(c *client.Client) *DeviceService {
	return &DeviceService{Client: c}
}

func (s *DeviceService) GetDevices(ctx context.Context, params *DeviceListParams) ([]models.Device, int, *string, *string, error) {
	query := url.Values{}
	if params != nil {
		perPage := params.PerPage
		if perPage > 1000 {
			perPage = 1000
		}
		if perPage > 0 {
			query.Set("per_page", fmt.Sprintf("%d", perPage))
		}
	}
	endpoint := "devices"
	if len(query) > 0 {
		endpoint += "?" + query.Encode()
	}
	respBody, err := s.Client.DoRequest(ctx, "GET", endpoint, nil)
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

// GetDevice fetches a single device by ID.
func (s *DeviceService) GetDevice(ctx context.Context, deviceID string) (*models.Device, error) {
	respBody, err := s.Client.DoRequest(ctx, "GET", "devices/"+deviceID, nil)
	if err != nil {
		return nil, err
	}
	var device models.Device
	if err := json.Unmarshal(respBody, &device); err != nil {
		return nil, fmt.Errorf("decode device: %w", err)
	}
	return &device, nil
}

// RefreshLocation asks TinyMDM to have the device report its latest location
// points. The response has no content; the updated points appear on a
// subsequent GetDevice/GetDevices call. TinyMDM rate-limits this to 1 call per
// hour per device, and requires geolocation to be enabled on the device.
func (s *DeviceService) RefreshLocation(ctx context.Context, deviceID string) error {
	_, err := s.Client.DoRequest(ctx, "GET", "devices/"+deviceID+"/refresh_location", nil)
	return err
}
