package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type GeolocationPosition struct {
	Address   string    `json:"address"`
	Date      time.Time `json:"date"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp int64     `json:"timestamp"`
}

type TinyMDMAppVersion string

func (v *TinyMDMAppVersion) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*v = TinyMDMAppVersion(s)
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		*v = TinyMDMAppVersion(fmt.Sprintf("%g", f))
		return nil
	}
	return fmt.Errorf("TinyMDMAppVersion: cannot unmarshal %s", string(data))
}

// NullableInt64 handles int64 fields that may be empty strings or numbers in JSON.
type NullableInt64 int64

func (n *NullableInt64) UnmarshalJSON(data []byte) error {
	if string(data) == "\"\"" || string(data) == "null" {
		*n = 0
		return nil
	}
	var i int64
	if err := json.Unmarshal(data, &i); err == nil {
		*n = NullableInt64(i)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "" {
			*n = 0
			return nil
		}
		var parsed int64
		if _, err := fmt.Sscanf(s, "%d", &parsed); err == nil {
			*n = NullableInt64(parsed)
			return nil
		}
	}
	return fmt.Errorf("NullableInt64: cannot unmarshal %s", string(data))
}

// NullableTime handles time fields that may be null, empty, or in various string formats in JSON.
type NullableTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullableTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "\"\"" {
		nt.Valid = false
		nt.Time = time.Time{}
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "" {
			nt.Valid = false
			nt.Time = time.Time{}
			return nil
		}
		// Try "2006-01-02 15:04:05" format
		t, err := time.Parse("2006-01-02 15:04:05", s)
		if err == nil {
			nt.Time = t
			nt.Valid = true
			return nil
		}
		// Try RFC3339
		t, err = time.Parse(time.RFC3339, s)
		if err == nil {
			nt.Time = t
			nt.Valid = true
			return nil
		}
		return fmt.Errorf("NullableTime: cannot parse time string '%s'", s)
	}
	return fmt.Errorf("NullableTime: cannot unmarshal %s", string(data))
}

func (nt NullableTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time.Format("2006-01-02 15:04:05"))
}

type Device struct {
	ID                             string                `json:"id"`
	Name                           string                `json:"name"`
	Nickname                       string                `json:"nickname"`
	PhoneNumber                    string                `json:"phone_number"`
	EnrollmentType                 string                `json:"enrollment_type"`
	OSVersion                      string                `json:"os_version"`
	BatteryLevel                   int                   `json:"battery_level"`
	BatteryStatus                  string                `json:"battery_status"`
	BatteryHealth                  string                `json:"battery_health"`
	GeolocationActivated           bool                  `json:"geolocation_activated"`
	GPSActivated                   bool                  `json:"gps_activated"`
	GeolocationPositions           []GeolocationPosition `json:"geolocation_positions"`
	IMEI                           string                `json:"imei"`
	ICCID                          string                `json:"iccid"`
	SSAID                          string                `json:"ssaid"`
	SerialNumber                   string                `json:"serial_number"`
	BuildID                        string                `json:"build_id"`
	Manufacturer                   string                `json:"manufacturer"`
	TinyMDMAppVersion              TinyMDMAppVersion     `json:"tinymdm_app_version"`
	EnrollmentTimestamp            int64                 `json:"enrollment_timestamp"`
	LastLockRequestDate            NullableTime          `json:"last_lock_request_date"`
	LastRebootRequestDate          NullableTime          `json:"last_reboot_request_date"`
	LastChangePasswordRequestDate  NullableTime          `json:"last_change_password_request_date"`
	LastDeletePasswordRequestDate  NullableTime          `json:"last_delete_password_request_date"`
	LastMessageSentRequestDate     NullableTime          `json:"last_message_sent_request_date"`
	LastWipeRequestDate            NullableTime          `json:"last_wipe_request_date"`
	LastSyncTimestamp              int64                 `json:"last_sync_timestamp"`
	LockAcknowledgeTime            NullableInt64         `json:"lock_acknowledge_time"`
	RebootAcknowledgeTime          NullableInt64         `json:"reboot_acknowledge_time"`
	ChangePasswordAcknowledgeTime  NullableInt64         `json:"change_password_acknowledge_time"`
	DeletePasswordAcknowledgeTime  NullableInt64         `json:"delete_password_acknowledge_time"`
	MessageReceivedAcknowledgeTime NullableInt64         `json:"message_received_acknowledge_time"`
	LastChangeUserRequestTimestamp NullableInt64         `json:"last_change_user_request_timestamp"`
	PolicyID                       string                `json:"policy_id"`
	UserID                         string                `json:"user_id"`
	GroupID                        string                `json:"group_id"`
	TransferStatus                 string                `json:"transfer_status"`
	TransferStatusMessage          string                `json:"transfer_status_message"`
}
