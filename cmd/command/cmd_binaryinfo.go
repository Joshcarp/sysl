package command

import (
	"fmt"
	"github.com/anz-bank/sysl/pkg/parse"
	"github.com/anz-bank/sysl/pkg/sysl"
	"github.com/anz-bank/sysl/pkg/syslutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
	"runtime"
)

type infoCmd struct{}

func (p *infoCmd) Name() string            { return "info" }
func (p *infoCmd) RequireSyslModule() bool { return false }

func (p *infoCmd) Configure(app *kingpin.Application) *kingpin.CmdClause {
	cmd := app.Command(p.Name(), "Show binary information")
	return cmd
}

func findRootFromSyslModule(modulePath string, fs afero.Fs) (string, error) {
	currentPath, err := filepath.Abs(modulePath)
	if err != nil {
		return "", err
	}

	systemRoot, err := filepath.Abs(string(os.PathSeparator))
	if err != nil {
		return "", err
	}

	// Keep walking up the directories to find nearest root marker
	for {
		currentPath = filepath.Dir(currentPath)
		exists, err := afero.Exists(fs, filepath.Join(currentPath, SyslRootMarker))
		reachedRoot := currentPath == systemRoot || (err != nil && os.IsPermission(err))
		switch {
		case exists:
			return currentPath, nil
		case reachedRoot:
			return "", nil
		case err != nil:
			return "", err
		}
	}
}

func (p *infoCmd) Execute(args ExecuteArgs) error {
	fmt.Printf("Version    : %s\n", Version)
	fmt.Printf("Commit ID  : %s\n", GitCommit)
	fmt.Printf("Build date : %s\n", BuildDate)
	fmt.Printf("GOOS       : %s\n", runtime.GOOS)
	fmt.Printf("GOARCH     : %s\n", runtime.GOARCH)
	fmt.Printf("Go Version : %s\n", runtime.Version())
	fmt.Printf("Build OS   : %s\n", BuildOS)
	return nil
}

// Version   - Binary version
// GitCommit - Commit SHA of the source code
// BuildDate - Binary build date
// BuildOS   - Operating System used to build binary
//nolint:gochecknoglobals
var (
	Version   = "unspecified"
	GitCommit = "unspecified"
	BuildDate = "unspecified"
	BuildOS   = "unspecified"
)

const Debug string = "debug"

type projectConfiguration struct {
	Module, Root string
	RootIsFound  bool
	fs           afero.Fs
}

func LoadSyslModule(root, filename string, fs afero.Fs, logger *logrus.Logger) (*sysl.Module, string, error) {
	logger.Debugf("Attempting to load module:%s (root:%s)", filename, root)
	projectConfig := NewProjectConfiguration()
	if err := projectConfig.ConfigureProject(root, filename, fs, logger); err != nil {
		return nil, "", err
	}

	modelParser := parse.NewParser()
	if !projectConfig.RootIsFound {
		modelParser.RestrictToLocalImport()
	}
	return parse.LoadAndGetDefaultApp(projectConfig.Module, projectConfig.fs, modelParser)
}

func NewProjectConfiguration() *projectConfiguration {
	return &projectConfiguration{
		Root:        "",
		Module:      "",
		RootIsFound: false,
		fs:          nil,
	}
}

func (pc *projectConfiguration) ConfigureProject(root, module string, fs afero.Fs, logger *logrus.Logger) error {
	rootIsDefined := root != ""

	modulePath := module
	if rootIsDefined {
		modulePath = filepath.Join(root, module)
	}

	syslRootPath, err := findRootFromSyslModule(modulePath, fs)
	if err != nil {
		return err
	}

	rootMarkerExists := syslRootPath != ""

	if rootIsDefined {
		pc.RootIsFound = true
		pc.Root = root
		pc.Module = module
		if rootMarkerExists {
			logger.Warningf("%s found in %s but will use %s instead",
				SyslRootMarker, syslRootPath, pc.Root)
		} else {
			logger.Warningf("%s is not defined but root flag is defined in %s",
				SyslRootMarker, pc.Root)
		}
	} else {
		if rootMarkerExists {
			pc.Root = syslRootPath

			// module has to be relative to the root
			absModulePath, err := filepath.Abs(module)
			if err != nil {
				return err
			}
			pc.Module, err = filepath.Rel(pc.Root, absModulePath)
			if err != nil {
				return err
			}
			pc.RootIsFound = true
		} else {
			// uses the module directory as the root, changing the module to be relative to the root
			pc.Root = filepath.Dir(module)
			pc.Module = filepath.Base(module)
			pc.RootIsFound = false
			logger.Warningf("root and %s are undefined, %s will be used instead",
				SyslRootMarker, pc.Root)
		}
	}

	pc.fs = syslutil.NewChrootFs(fs, pc.Root)
	return nil
}
