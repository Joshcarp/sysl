package main

import (
        "os"

        "github.com/Joshcarp/sysl/pkg/command"
        "github.com/sirupsen/logrus"
        "github.com/spf13/afero"
)

// main is as small as possible to minimise its no-coverage footprint.
func main() {
        if rc := command.Main2(os.Args, afero.NewOsFs(), logrus.StandardLogger(), command.Main3); rc != 0 {
                os.Exit(rc)
        }
}