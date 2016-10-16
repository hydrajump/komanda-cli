package command

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/logger"
)

type RawCmd struct {
	*MetadataTmpl
}

func (e *RawCmd) Metadata() CommandMetadata {
	return e
}

func (e *RawCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		cmd := strings.Join(args[1:], " ")
		logger.Logger.Printf("RAW COMMAND %s\n", cmd)
		s.Client.Raw(cmd)
		return nil
	})

	return nil
}

func rawCmd() Command {
	return &RawCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "raw",
			args:        "<command> [data]",
			description: "raw command",
		},
	}
}
