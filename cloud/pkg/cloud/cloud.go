package cloud

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pathak107/cloudesk/pkg/utils"
)

var (
	configPathAWS = "/home/pathak107/Documents/webDev/cloudGaming/pkg/cloud/config/aws_config.json"
)

type CloudService struct {
	cloud CloudProvider
}

func NewCloudService() (*CloudService, error) {
	var cfg awsConfig
	configJSON, err := ioutil.ReadFile(configPathAWS)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configJSON, &cfg)
	if err != nil {
		return nil, err
	}
	awsProvider, err := initAWS(context.Background(), cfg)
	if err != nil {
		return nil, err
	}
	return &CloudService{
		cloud: awsProvider,
	}, nil
}

func (c *CloudService) LaunchVM(ctx context.Context, cfg *CreateInstanceParams) (Instance, error) {
	vm, err := c.cloud.CreateVM(ctx, *cfg)
	if err != nil {
		log.Println(err)
		return Instance{}, utils.NewUnexpectedServerError()
	}
	return vm, nil
}

// func SetupRDP(ctx context.Context, vm Instance) error {}

// func StartRDP(ctx context.Context, vm Instance, rdpType RDPType) error {}

// func StopVM() {}

// func TerminateVM() {}
