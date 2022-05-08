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

type JsonCfgForApp string

const (
	HighEndGames JsonCfgForApp = "/home/pathak107/Documents/webDev/cloudGaming/pkg/cloud/apps_config/gaming/high-end-games.json"
	MidEndGames  JsonCfgForApp = "/home/pathak107/Documents/webDev/cloudGaming/pkg/cloud/apps_config/gaming/mid-end-games.json"
	Base         JsonCfgForApp = "/home/pathak107/Documents/webDev/cloudGaming/pkg/cloud/apps_config/base/base-image.json"
)

func GetProvider(cloudType CloudType) (CloudProvider, error) {
	if cloudType == AWS {
		var cfg awsConfig
		configJSON, err := ioutil.ReadFile("/home/pathak107/Documents/webDev/cloudGaming/pkg/cloud/config/aws-config.json")
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

func InstanceCfgFromJson(cloudType CloudType, jsonConfigPath JsonCfgForApp) (InstanceConfig, error) {
	var cfg jsonConfigForApp
	jsonConfig, err := ioutil.ReadFile(string(jsonConfigPath))
	if err != nil {
		return InstanceConfig{}, err
	}

	err = json.Unmarshal(jsonConfig, &cfg)
	if err != nil {
		return InstanceConfig{}, err
	}

	if cloudType == AWS {
		return InstanceConfig{
			Image:    cfg.aws.ami_id,
			Hardware: cfg.aws.instance_type,
		}, nil
	}

	return InstanceConfig{}, errors.New("invalid cloud type specified")
}

func LaunchVM(ctx context.Context, c CloudProvider, cloudType CloudType, instanceCfg InstanceConfig, jsonConfigPath JsonCfgForApp) (Instance, error) {
	cfg, err := InstanceCfgFromJson(cloudType, jsonConfigPath)
	if err != nil {
		return Instance{}, err
	}
	if jsonConfigPath == Base { //For cases when user wants to run their own setup with custom config
		instanceCfg.Image = cfg.Image
		return c.CreateVM(ctx, instanceCfg)
	}

	cfg.Name = instanceCfg.Name
	return c.CreateVM(ctx, cfg)
}

// func SetupRDP(ctx context.Context, vm Instance) error {}

// func StartRDP(ctx context.Context, vm Instance, rdpType RDPType) error {}

// func StopVM() {}

// func TerminateVM() {}
