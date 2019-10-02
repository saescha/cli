// +build V7

package rpc

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"sync"

	"code.cloudfoundry.org/cli/command"
	v7command "code.cloudfoundry.org/cli/command/v7"
	plugin "code.cloudfoundry.org/cli/plugin/v7"
	plugin_models "code.cloudfoundry.org/cli/plugin/v7/models"
	"code.cloudfoundry.org/cli/version"
	"github.com/blang/semver"
)

type CliRpcCmd struct {
	PluginMetadata       *plugin.PluginMetadata
	MetadataMutex        *sync.RWMutex
	Config               command.Config
	AppActor             v7command.AppActor
	outputCapture        OutputCapture
	terminalOutputSwitch TerminalOutputSwitch
	outputBucket         *bytes.Buffer
	stdout               io.Writer
}

func (cmd *CliRpcCmd) IsMinCliVersion(passedVersion string, retVal *bool) error {
	if version.VersionString() == version.DefaultVersion {
		*retVal = true
		return nil
	}

	actualVersion, err := semver.Make(version.VersionString())
	if err != nil {
		return err
	}

	requiredVersion, err := semver.Make(passedVersion)
	if err != nil {
		return err
	}

	*retVal = actualVersion.GTE(requiredVersion)

	return nil
}

func (cmd *CliRpcCmd) SetPluginMetadata(pluginMetadata plugin.PluginMetadata, retVal *bool) error {
	cmd.MetadataMutex.Lock()
	defer cmd.MetadataMutex.Unlock()

	cmd.PluginMetadata = &pluginMetadata
	*retVal = true
	return nil
}

func (cmd *CliRpcCmd) DisableTerminalOutput(disable bool, retVal *bool) error {
	cmd.terminalOutputSwitch.DisableTerminalOutput(disable)
	*retVal = true
	return nil
}

func (cmd *CliRpcCmd) CallCoreCommand(args []string, retVal *bool) error {
	return errors.New("unimplemented")
}

func (cmd *CliRpcCmd) GetOutputAndReset(args bool, retVal *[]string) error {
	return errors.New("unimplemented")
}

func (cmd *CliRpcCmd) GetApp(appName string, retVal *plugin_models.Application) error {
	spaceGUID := cmd.Config.TargetedSpace().GUID
	app, _, err := cmd.AppActor.GetDetailedAppSummary(appName, spaceGUID, true)
	assignableValue := plugin_models.Application(app)

	to := reflect.ValueOf(retVal).Elem()
	to.Set(reflect.ValueOf(assignableValue))

	return err
}
