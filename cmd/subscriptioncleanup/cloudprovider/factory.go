package cloudprovider

import (
	"fmt"

	"github.com/kyma-project/kyma-environment-broker/cmd/subscriptioncleanup/model"
)

type ResourceCleaner interface {
	Do() error
}

//go:generate mockery --name=ProviderFactory
type ProviderFactory interface {
	New(hyperscalerType model.HyperscalerType, secretData map[string][]byte) (ResourceCleaner, error)
}

type providerFactory struct{}

func NewProviderFactory() ProviderFactory {
	return &providerFactory{}
}

func (pf *providerFactory) New(hyperscalerType model.HyperscalerType, secretData map[string][]byte) (ResourceCleaner, error) {
	switch hyperscalerType {
	case model.GCP:
		{
			return NewGCPeResourcesCleaner(secretData), nil
		}
	case model.Azure:
		{
			return NewAzureResourcesCleaner(secretData)
		}
	case model.AWS:
		{
			return NewAwsResourcesCleaner(secretData)
		}
	default:
		return nil, fmt.Errorf("unknown hyperscaler type")
	}
}
