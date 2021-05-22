package models

type Target interface {
	HasService(map[string]CommandService) map[string]string
	HasSubdomain() bool
	GetTarget() string
}
