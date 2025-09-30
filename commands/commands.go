package commands

import (
	"web/commands/generate"
	"web/commands/gorm"
	"web/commands/migrate"

	"github.com/urfave/cli/v2"
)

func All() []*cli.Command {
	commands := []*cli.Command{
		migrate.Command(),
		generate.Command(),
		gorm.Command(),
	}
	return commands
}
