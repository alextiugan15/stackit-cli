package client

import (
	"github.com/spf13/viper"
	"github.com/stackitcloud/stackit-cli/internal/pkg/config"
	genericclient "github.com/stackitcloud/stackit-cli/internal/pkg/generic-client"
	"github.com/stackitcloud/stackit-cli/internal/pkg/print"
	ufw "github.com/stackitcloud/stackit-sdk-go/services/ufw/v1api"
)

func ConfigureClient(p *print.Printer, cliVersion string) (*ufw.APIClient, error) {
	return genericclient.ConfigureClientGeneric(p, cliVersion, viper.GetString(config.UfwCustomEndpointKey), false, ufw.NewAPIClient)
}
