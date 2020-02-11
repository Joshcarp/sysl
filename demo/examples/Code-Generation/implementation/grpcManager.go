package implementation

import (
	"github.service.anz/sysl/server-lib/config"
	"github.service.anz/sysl/server-lib/handlerinitialiser"
)

type GRPCHandler struct{}

func (h GRPCHandler) EnabledGrpcHandlers() []handlerinitialiser.GrpcHandlerInitialiser {
	return nil
}

func (h GRPCHandler) GrpcAdminServerConfig() *config.CommonServerConfig {
	return nil
}

func (h GRPCHandler) GrpcPublicServerConfig() *config.CommonServerConfig {
	return nil
}
