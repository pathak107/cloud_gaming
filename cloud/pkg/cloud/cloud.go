package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type CloudType int

const (
	AWS CloudType = iota
	GCP
	AZURE
)

var (
	AppsConfigBasePath = "/home/pathak107/Documents/webDev/cloudGaming/pkg/cloud/apps_config/"
	AWSConfigPath      = "/home/pathak107/Documents/webDev/cloudGaming/pkg/cloud/config/aws_config.json"
)

func NewCloudProvider(cloudType CloudType) (CloudProvider, error) {
	if cloudType == AWS {
		var cfg awsConfig
		configJSON, err := ioutil.ReadFile(AWSConfigPath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(configJSON, &cfg)
		if err != nil {
			return nil, err
		}
		return initAWS(context.Background(), cfg)
	}
	return nil, errors.New("invalid cloud type specified")
}

func CreateInstanceParamsFromAppCfg(cloudType CloudType, appConfigName string) (CreateInstanceParams, error) {
	var cfg AppConfigJSON
	appCfgJson, err := ioutil.ReadFile(AppsConfigBasePath + appConfigName + ".json")
	if err != nil {
		return CreateInstanceParams{}, err
	}

	err = json.Unmarshal(appCfgJson, &cfg)
	if err != nil {
		return CreateInstanceParams{}, err
	}

	if cloudType == AWS {
		return CreateInstanceParams{
			Image:    cfg.Aws.AmiID,
			Hardware: cfg.Aws.InstanceType,
		}, nil
	}

	return CreateInstanceParams{}, errors.New("invalid cloud type specified")
}

func LaunchVM(ctx context.Context, c CloudProvider, cloudType CloudType, vmName string, appConfigName string) (Instance, error) {
	cfg, err := CreateInstanceParamsFromAppCfg(cloudType, appConfigName)
	if err != nil {
		return Instance{}, err
	}

	cfg.Name = vmName
	return c.CreateVM(ctx, cfg)
}

// func SetupRDP(ctx context.Context, vm Instance) error {}

// func StartRDP(ctx context.Context, vm Instance, rdpType RDPType) error {}

// func StopVM() {}

// func TerminateVM() {}
