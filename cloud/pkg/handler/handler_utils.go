package handler

import "github.com/pathak107/cloudesk/pkg/cloud"

type Handler struct {
	cloudSvc *cloud.CloudService
}

func NewCloudHandler() (*Handler, error) {
	cSvc, err := cloud.NewCloudService()
	if err != nil {
		return nil, err
	}
	return &Handler{
		cloudSvc: cSvc,
	}, nil
}
