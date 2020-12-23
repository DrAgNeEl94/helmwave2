package main

import (
	"github.com/urfave/cli/v2"
	"github.com/zhilyaev/helmwave/pkg/helmwave"
)

func flags(app *helmwave.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "tpl",
			Value:       "helmwave.yml.tpl",
			Usage:       "Main tpl file",
			EnvVars:     []string{"HELMWAVE_TPL_FILE"},
			Destination: &app.Tpl.From,
		},
		&cli.StringFlag{
			Name:        "file",
			Aliases:     []string{"f"},
			Value:       "helmwave.yml",
			Usage:       "Main yml file",
			EnvVars:     []string{"HELMWAVE_FILE", "HELMWAVE_YAML_FILE", "HELMWAVE_YML_FILE"},
			Destination: &app.Tpl.To,
		},
		&cli.StringFlag{
			Name:        "plan-dir",
			Value:       ".helmwave/",
			Usage:       "It keeps your state via planfile",
			EnvVars:     []string{"HELMWAVE_PLAN_DIR"},
			Destination: &app.PlanPath,
		},
		&cli.StringSliceFlag{
			Name:        "tags",
			Aliases:     []string{"t"},
			Usage:       "It allows you choose releases for sync. Example: -t tag1 -t tag3,tag4",
			EnvVars:     []string{"HELMWAVE_TAGS"},
			Destination: &app.Tags,
		},
		&cli.BoolFlag{
			Name:        "parallel",
			Usage:       "It allows you call `helm install` in parallel mode ",
			Value:       true,
			EnvVars:     []string{"HELMWAVE_PARALLEL"},
			Destination: &app.Parallel,
		},
		//&cli.BoolFlag{
		//	Name:        "force",
		//	Usage:       "It allows you call `helm install` in parallel mode ",
		//	Value:       true,
		//	EnvVars:     []string{"HELMWAVE_FORCE"},
		//	Destination: &app.Force,
		//},
		//
		//		LOGGER
		//
		&cli.StringFlag{
			Name:        "log-format",
			Usage:       "You can set: [ text | json | pad | emoji ]",
			Value:       "emoji",
			EnvVars:     []string{"HELMWAVE_LOG_FORMAT"},
			Destination: &app.Logger.Format,
		},
		&cli.StringFlag{
			Name:        "log-level",
			Usage:       "You can set: [ debug | info | warn | panic | fatal | trace ]",
			Value:       "info",
			EnvVars:     []string{"HELMWAVE_LOG_LEVEL", "HELMWAVE_LOG_LVL"},
			Destination: &app.Logger.Level,
		},
		&cli.BoolFlag{
			Name:        "log-color",
			Usage:       "Force color",
			Value:       true,
			EnvVars:     []string{"HELMWAVE_LOG_COLOR"},
			Destination: &app.Logger.Color,
		},
	}
}
