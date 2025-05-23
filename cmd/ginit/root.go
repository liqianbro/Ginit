package ginit

import (
	"fmt"

	"github.com/liqianbro/Ginit/config"
	"github.com/liqianbro/Ginit/internal/command/new"

	"github.com/spf13/cobra"
)

var CmdRoot = &cobra.Command{
	Use:     "ginit",
	Example: "ginit new demo-api",
	Short:   "\n  ________.__       .__  __   \n /  _____/|__| ____ |__|/  |_ \n/   \\  ___|  |/    \\|  \\   __\\\n\\    \\_\\  \\  |   |  \\  ||  |  \n \\______  /__|___|  /__||__|  \n        \\/        \\/          \n\n",
	Version: fmt.Sprintf("\n _   _                   \n| \\ | |_   _ _ __  _   _ \n|  \\| | | | | '_ \\| | | |\n| |\\  | |_| | | | | |_| |\n|_| \\_|\\__,_|_| |_|\\__,_| \n \nGinit %s - Copyright (c) 2023-2025 Ginit\nReleased under the MIT License.\n\n", config.Version),
}

func init() {
	CmdRoot.AddCommand(new.CmdNew)
}

// Execute executes the root command.
func Execute() error {
	return CmdRoot.Execute()
}
