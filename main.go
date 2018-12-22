// OnePort-go is a super port proxy
package main

import (
	"fmt"
	"os"

	"github.com/ourCloudSky/OnePort-go/constant"
	"github.com/ukautz/clif"
)

func main() {
	app := clif.New("OnePort", constant.Version, constant.Description)
	if len(os.Args) == 1 || os.Args[1] == "list" {
		fmt.Print(constant.CharacterDrawing)
	}
	app.Add(
		clif.NewCommand(
			"server",
			"Run the OnePort server process",
			func(out clif.Output, cmd *clif.Command) {

			},
		).AddOption(
			clif.NewOption(
				"configure",
				"c",
				"The configure file path",
				"default",
				true,
				false,
			),
		),
	)
	app.Run()
}
