package models

type Target interface {
	HasService(map[string]CommandService) map[string]string
	GetSubdomain() []string
	GetTarget() string
}
