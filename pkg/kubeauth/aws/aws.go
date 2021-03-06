package aws

import (
	"context"
	"fmt"

	api "github.com/portworx/px-backup-api/pkg/apis/v1"
	"github.com/portworx/px-backup-api/pkg/kubeauth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	pluginName = "aws"
)

type aws struct {
}

// Init initializes the gcp auth plugin
func (a *aws) Init() error {
	return nil
}

func (a *aws) UpdateClient(
	conn *grpc.ClientConn,
	ctx context.Context,
	cloudCredentialName string,
	orgID string,
	client *rest.Config,
) (bool, error) {
	if client.ExecProvider != nil {
		if client.ExecProvider.Command == "aws-iam-authenticator" || client.ExecProvider.Command == "aws" {
			if cloudCredentialName == "" {
				return false, fmt.Errorf("CloudCredential not provided for EKS cluster")
			}

			cloudCredentialClient := api.NewCloudCredentialClient(conn)
			resp, err := cloudCredentialClient.Inspect(
				ctx,
				&api.CloudCredentialInspectRequest{
					Name:           cloudCredentialName,
					OrgId:          orgID,
					IncludeSecrets: true,
				},
			)
			if err != nil {
				return false, err
			}
			cloudCredential := resp.GetCloudCredential()
			if cloudCredential.GetCloudCredentialInfo().GetType() != api.CloudCredentialInfo_AWS {
				return false, fmt.Errorf("need AWS CloudCredential for EKS cluster. Provided %v", cloudCredential.GetCloudCredentialInfo().GetType())
			}
			if client.ExecProvider.Env == nil {
				client.ExecProvider.Env = make([]clientcmdapi.ExecEnvVar, 0)
			}
			client.ExecProvider.Env = append(client.ExecProvider.Env, clientcmdapi.ExecEnvVar{
				Name:  "AWS_ACCESS_KEY",
				Value: cloudCredential.GetCloudCredentialInfo().GetAwsConfig().GetAccessKey(),
			})
			client.ExecProvider.Env = append(client.ExecProvider.Env, clientcmdapi.ExecEnvVar{
				Name:  "AWS_ACCESS_KEY_ID",
				Value: cloudCredential.GetCloudCredentialInfo().GetAwsConfig().GetAccessKey(),
			})
			client.ExecProvider.Env = append(client.ExecProvider.Env, clientcmdapi.ExecEnvVar{
				Name:  "AWS_SECRET_KEY",
				Value: cloudCredential.GetCloudCredentialInfo().GetAwsConfig().GetSecretKey(),
			})
			client.ExecProvider.Env = append(client.ExecProvider.Env, clientcmdapi.ExecEnvVar{
				Name:  "AWS_SECRET_ACCESS_KEY",
				Value: cloudCredential.GetCloudCredentialInfo().GetAwsConfig().GetSecretKey(),
			})

			// Remove the profile env if present since we are passing in the creds through env
			tempEnv := make([]clientcmdapi.ExecEnvVar, 0)
			for _, env := range client.ExecProvider.Env {
				if env.Name == "AWS_PROFILE" {
					continue
				}
				tempEnv = append(tempEnv, env)
			}
			client.ExecProvider.Env = tempEnv
		}
		return true, nil
	}
	return false, nil
}

func init() {
	if err := kubeauth.Register(pluginName, &aws{}); err != nil {
		logrus.Panicf("Error registering aws auth plugin: %v", err)
	}
}
