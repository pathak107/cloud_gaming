package cloud

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type awsProvider struct {
	vm     *ec2.Client
	config awsConfig
}

// DestroyVm completely deletes ec2 instance and its associated EBS volumes
func (a *awsProvider) DestroyVm(ctx context.Context, vmID string) error {
	if _, err := a.vm.TerminateInstances(ctx, &ec2.TerminateInstancesInput{InstanceIds: []string{vmID}}); err != nil {
		return err
	}
	return nil
}

func (a *awsProvider) Status(ctx context.Context, vmID string) (string, error) {
	status, err := a.vm.DescribeInstanceStatus(ctx, &ec2.DescribeInstanceStatusInput{InstanceIds: []string{vmID}})
	if err != nil {
		return "", err
	}
	return string(status.InstanceStatuses[0].InstanceState.Name.Values()[0]), nil
}

func (a *awsProvider) GetVmsUsage(ctx context.Context, tenantId string) (float64, error) {
	return 0, errors.New("not implemented")
}

func (a *awsProvider) StartInstance(ctx context.Context, vmID string) error {
	if _, err := a.vm.StartInstances(ctx, &ec2.StartInstancesInput{InstanceIds: []string{vmID}}); err != nil {
		return err
	}
	return nil
}

func (a *awsProvider) StopInstance(ctx context.Context, vmID string) error {
	if _, err := a.vm.StopInstances(ctx, &ec2.StopInstancesInput{InstanceIds: []string{vmID}}); err != nil {
		return err
	}
	return nil
}

func (a *awsProvider) CreateImage(ctx context.Context, imgParam ImageConfig) (string, string, error) {
	res, err := a.vm.CreateImage(ctx, &ec2.CreateImageInput{
		Description: aws.String(imgParam.Description),
		InstanceId:  aws.String(imgParam.VmID),
		Name:        aws.String(imgParam.Name),
	})
	if err != nil {
		return "", "", err
	}
	return aws.ToString(res.ImageId), imgParam.Name, nil
}

func (a *awsProvider) DeleteImage(ctx context.Context, imageID string) error {
	if _, err := a.vm.DeregisterImage(ctx, &ec2.DeregisterImageInput{ImageId: aws.String(imageID)}); err != nil {
		return err
	}
	return nil
}

//
func (a *awsProvider) CreateVM(ctx context.Context, instanceCfg InstanceConfig) (Instance, error) {
	log.Println("Creating ec2")
	initScript := fmt.Sprintf("#!/bin/bash\nhostnamectl set-hostname %s", instanceCfg.Name)
	userData := base64.StdEncoding.EncodeToString([]byte(initScript))
	log.Println("Check if the security group exists")
	securityGrpID := a.config.SecurityGrp //Ex- "sg-0ef576fe1f10689dd"
	_, err := a.vm.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
		GroupIds: []string{securityGrpID},
	})
	if err != nil {
		return Instance{}, err
	}

	// Flavor ID is instance type in AWS. Ex- t2-micro etc.
	instanceType := types.InstanceType(instanceCfg.Hardware)
	input := &ec2.RunInstancesInput{
		ImageId:          aws.String(instanceCfg.Image),
		InstanceType:     instanceType,
		MinCount:         aws.Int32(1),
		MaxCount:         aws.Int32(1),
		KeyName:          aws.String(a.config.KeyName), // Key pair name to ssh into ec2
		SecurityGroupIds: []string{securityGrpID},
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String(instanceCfg.Name),
					},
				},
			},
		},
		UserData: aws.String(userData),
	}

	ec2Instance, err := a.vm.RunInstances(ctx, input)
	if err != nil {
		return Instance{}, err
	}
	ec2InstanceID := aws.ToString(ec2Instance.Instances[0].InstanceId)

	log.Println("ec2 created, waiting for status")
	waiter := ec2.NewInstanceRunningWaiter(a.vm)

	// maxWaitTime is the maximum wait time, the waiter will wait for
	// the resource status.
	maxWaitTime := 5 * time.Minute
	// Wait will poll until it gets the resource status, or max wait time expires
	err = waiter.Wait(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{ec2InstanceID},
	}, maxWaitTime)

	if err != nil {
		return Instance{}, err
	}

	// ec2 in running state now, retrieve information
	res, err := a.vm.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{ec2InstanceID},
	})
	if err != nil {
		return Instance{}, err
	}
	ec2InstanceInfo := res.Reservations[0].Instances[0]
	log.Println("ec2 creation complete")
	return Instance{
		VmID:      aws.ToString(ec2InstanceInfo.InstanceId),
		PublicIP:  aws.ToString(ec2InstanceInfo.PublicIpAddress),
		PrivateIP: aws.ToString(ec2InstanceInfo.PrivateIpAddress),
	}, nil

}

func (a *awsProvider) ImageId(ctx context.Context, imgName string) (string, error) {
	res, err := a.vm.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{imgName},
			},
		},
	})
	if err != nil {
		return "", err
	}
	return aws.ToString(res.Images[0].ImageId), nil
}

// Aws credentails provider
type awsCredentialsProvider struct {
	Value aws.Credentials
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (s awsCredentialsProvider) Retrieve(_ context.Context) (aws.Credentials, error) {
	v := s.Value
	if v.AccessKeyID == "" || v.SecretAccessKey == "" {
		return aws.Credentials{
			Source: "StaticCredentials",
		}, errors.New("credentials are empty")
	}

	if len(v.Source) == 0 {
		v.Source = "StaticCredentials"
	}

	return v, nil
}

// InitAws initializes aws provider with client to interact with aws API
func initAWS(ctx context.Context, cfg awsConfig) (*awsProvider, error) {
	config, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(awsCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     cfg.AccessKey,
				SecretAccessKey: cfg.SecretKey,
			},
		}),
	)
	if err != nil {
		return &awsProvider{}, err
	}
	return &awsProvider{
		vm:     ec2.NewFromConfig(config),
		config: cfg,
	}, nil
}
