package extensions

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bishopfox/sliver/client/assets"
	"github.com/bishopfox/sliver/client/console"
	"github.com/desertbit/grumble"
)

// ExtensionsRemoveCmd - Remove an extension
func ExtensionsRemoveCmd(ctx *grumble.Context, con *console.SliverConsoleClient) {
	name := ctx.Args.String("name")
	if name == "" {
		con.Println("Extension name is required")
		return
	}
	confirm := false
	prompt := &survey.Confirm{Message: fmt.Sprintf("Remove '%s' extension?", name)}
	survey.AskOne(prompt, &confirm)
	if !confirm {
		return
	}
	err := RemoveExtensionByCommandName(name, con)
	if err != nil {
		con.PrintErrorf("Error removing extension: %s", err)
		return
	} else {
		con.PrintInfof("Extension '%s' removed\n", name)
	}
}

// RemoveExtensionByCommandName - Remove an extension by command name
func RemoveExtensionByCommandName(commandName string, con *console.SliverConsoleClient) error {
	if commandName == "" {
		return errors.New("command name is required")
	}
	if _, ok := loadedExtensions[commandName]; !ok {
		return errors.New("extension not loaded")
	}
	delete(loadedExtensions, commandName)

	allCommands := con.App.Commands().All()
	var index int
	var cmd *grumble.Command
	for index, cmd = range allCommands {
		if cmd.Name == commandName {
			break
		}
	}
	//lint:ignore SA4006 - false positive
	allCommands = append(allCommands[:index], allCommands[index+1:]...)

	extPath := filepath.Join(assets.GetExtensionsDir(), filepath.Base(commandName))
	if _, err := os.Stat(extPath); os.IsNotExist(err) {
		return nil
	}
	err := os.RemoveAll(extPath)
	if err != nil {
		return err
	}

	return nil
}
