package types

import "time"

// SLog -> Abbrevation for Structured Log
type SLog struct {
	Timestamp   time.Duration `json:"timestamp"`
	Level       string        `json:"level"`
	Message     string        `json:"message"`
	Code        string        `json:"code"`
	Application string        `json:"application"`
	Service     string        `json:"component"`
	IPAddress   string        `json:"ip"`
	Env         string        `json:"env"`
}
