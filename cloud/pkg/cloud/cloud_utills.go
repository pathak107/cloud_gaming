package cloud

import (
	"context"
)

type awsConfig struct {
	AccessKey   string `json:"access_key"`
	SecretKey   string `json:"secret_key"`
	SecurityGrp string `json:"security_grp"`
	KeyName     string `json:"key_name"`
	Region      string `json:"region"`
}

type Instance struct {
	VmID      string
	PublicIP  string
	PrivateIP string
	Name      string
}

type CreateInstanceParams struct {
	Name           string
	Image          string
	Hardware       string
	Storage        string
	PrivateNetwork string
}

type CreateImageParams struct {
	VmID        string
	Name        string
	Description string
}

type CloudProvider interface {
	DestroyVm(ctx context.Context, vmID string) error
	Status(ctx context.Context, vmID string) (string, error)
	// GetVmsUsage(ctx context.Context, tenantId string) (float64, error)
	StartInstance(ctx context.Context, vmID string) error
	StopInstance(ctx context.Context, vmID string) error
	CreateImage(ctx context.Context, imgParam CreateImageParams) (string, string, error)
	DeleteImage(ctx context.Context, imageID string) error
	CreateVM(ctx context.Context, instanceParams CreateInstanceParams) (Instance, error)
	ImageId(ctx context.Context, imgName string) (string, error)
}
