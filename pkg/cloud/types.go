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

type jsonConfigForApp struct {
	aws struct {
		ami_id        string `json:"ami_id"`
		instance_type string `json:"instance_type"`
	} `json:"aws"`
}

type Instance struct {
	VmID      string
	PublicIP  string
	PrivateIP string
}

type InstanceConfig struct {
	Name           string
	Image          string
	Hardware       string
	Storage        string
	PrivateNetwork string
}

type ImageConfig struct {
	VmID        string
	Name        string
	Description string
}

type Cloud interface {
	LaunchVM(ctx context.Context, instanceCfg InstanceConfig, jsonConfigPath JsonCfgForApp) (*Instance, error)
}

type CloudProvider interface {
	DestroyVm(ctx context.Context, vmID string) error
	Status(ctx context.Context, vmID string) (string, error)
	GetVmsUsage(ctx context.Context, tenantId string) (float64, error)
	StartInstance(ctx context.Context, vmID string) error
	StopInstance(ctx context.Context, vmID string) error
	CreateImage(ctx context.Context, imgParam ImageConfig) (string, string, error)
	DeleteImage(ctx context.Context, imageID string) error
	CreateVM(ctx context.Context, instanceCfg InstanceConfig) (Instance, error)
	ImageId(ctx context.Context, imgName string) (string, error)
}
