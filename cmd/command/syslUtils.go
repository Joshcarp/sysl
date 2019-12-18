package command

import (
	"github.com/anz-bank/sysl/pkg/parse"
	sysl "github.com/anz-bank/sysl/pkg/sysl"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

func LoadSyslModule(root, filename string, fs afero.Fs, logger *logrus.Logger) (*sysl.Module, string, error) {
	logger.Debugf("Attempting to load module:%s (root:%s)", filename, root)
	projectConfig := newProjectConfiguration()
	if err := projectConfig.configureProject(root, filename, fs, logger); err != nil {
		return nil, "", err
	}

	modelParser := parse.NewParser()
	if !projectConfig.rootIsFound {
		modelParser.RestrictToLocalImport()
	}
	return parse.LoadAndGetDefaultApp(projectConfig.module, projectConfig.fs, modelParser)
}

func newProjectConfiguration() *projectConfiguration {
	return &projectConfiguration{
		root:        "",
		module:      "",
		rootIsFound: false,
		fs:          nil,
	}
}
