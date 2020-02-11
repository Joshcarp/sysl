package implementation

import (
	"github.service.anz/sysl/server-lib/config"
	"github.service.anz/sysl/server-lib/handlerinitialiser"
)

type Handler struct {
	Router             []handlerinitialiser.HandlerInitialiser
	ServerConfig       *config.CommonHTTPServerConfig
	MAdminServerConfig *config.CommonHTTPServerConfig
	MLibraryConfig     *config.LibraryConfig
}

// impl 'h Handler' github.service.anz/sysl/server-lib/core.Manager

func (r Handler) EnabledHandlers() []handlerinitialiser.HandlerInitialiser {
	return r.Router
}

func (r Handler) LibraryConfig() *config.LibraryConfig {
	return r.MLibraryConfig
}

func (r Handler) AdminServerConfig() *config.CommonHTTPServerConfig {
	return r.MAdminServerConfig
}

func (r Handler) PublicServerConfig() *config.CommonHTTPServerConfig {
	return r.ServerConfig
}
