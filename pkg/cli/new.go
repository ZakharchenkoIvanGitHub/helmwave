package cli

import (
	"github.com/zhilyaev/helmwave/pkg/helmwave"
	"helm.sh/helm/v3/pkg/cli"
)

func New() *helmwave.Config {
	return &helmwave.Config{
		Version: "0.3.3",
		Helm:    cli.New(),
	}
}
