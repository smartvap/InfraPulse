package server

import (
	"context"
	"github.com/clarechu/infrapulse/pkg/server/router"
)

type CmdbConfig struct {
	DataRoot  string `yaml:"data_root"`
	Port      int32  `yaml:"port"`
	ProxyPort int32  `yaml:"proxy_port"`
}

func NewCmdb(config *CmdbConfig) (Bootstrap, error) {

	ctx, cancel := context.WithCancel(context.Background())
	return &CMDB{
		port:      config.Port,
		proxyPort: config.ProxyPort,
		server:    router.NewServer(),
		ctx:       ctx,
		cancel:    cancel,
	}, nil
}
