package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/cmd/flower/rpc"
	"github.com/monzo/slog"
	"github.com/peterh/liner"
	"github.com/spf13/cobra"
)

type helpEntry struct {
	cmd, desc string
}

var help = []helpEntry{
	{
		cmd:  "rpc ",
		desc: "run an rpc in this env",
	},
	{
		cmd:  "env ",
		desc: fmt.Sprintf("switch environment (allowed: %v, %v)", aurora.Green("prod"), aurora.Green("local")),
	},
	{
		cmd:  "exit",
		desc: "exit this shell",
	},
	{
		cmd:  "help",
		desc: "show this help text",
	},
}

var shellCmd = &cobra.Command{
	Use:    "shell",
	Short:  "Launch an interactive RPC shell",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		state := liner.NewLiner()
		defer state.Close()

		lastLineErrored := false

		for {
			prompt := "🌸 > "
			if lastLineErrored {
				prompt = "😰 > "
			}

			prompt = fmt.Sprintf("[%v] %v", env, prompt)

			input, err := state.Prompt(prompt)
			if err == io.EOF {
				cmd.Println("Exiting 👋")
				return
			}
			if err != nil {
				slog.Error(nil, "Error getting prompt input: %v", err)
				return
			}

			// if there's no input, continue to the next line
			if input == "" {
				continue
			}

			// "tokenize" input and check the first word
			fields := strings.Fields(input)
			switch fields[0] {
			case "exit", "\\q":
				cmd.Println("Exiting 👋")
				return
			case "help", "h", "?":
				cmd.Printf("\n")

				for _, entry := range help {
					cmd.Printf("    %v  - %v\n", aurora.Blue(entry.cmd), entry.desc)
				}

				cmd.Printf("\n\n")
			case "env":
				if len(fields) < 2 {
					switch env {
					case "prod":
						cmd.Printf("current env: %v\n", aurora.Blue("prod"))
						cmd.Printf("  internal api: %v\n", aurora.Green(EnvProd))
						cmd.Printf("  external api: %v\n", aurora.Green(EnvProdAPI))
					case "local":
						cmd.Printf("current env: %v\n", aurora.Blue("local"))
						cmd.Printf("  internal api: %v\n", aurora.Green(EnvLocal))
						cmd.Printf("  external api: %v\n", aurora.Green(EnvLocalAPI))
					}

					lastLineErrored = false
					continue
				}

				switch fields[1] {
				case "prod", "production", "local":
					if !setEnvironment(fields[1]) {
						lastLineErrored = true
						continue
					}

					cmd.Printf("Now using %v\n", aurora.Blue(env))
				default:
					cmd.Println("Unknown environment.")
					lastLineErrored = true
					continue
				}

				lastLineErrored = false
			case "rpc":
				// trigger an RPC command with the remaining fields
				if len(fields) < 4 {
					cmd.Printf("Usage: rpc [%v] %v BODY\n", aurora.Blue("method"), aurora.Red("/service.foo"))
					lastLineErrored = true
					continue
				}

				result, err := rpc.InternalRequest(nil, fields[1], fields[2], strings.Join(fields[3:], " "))
				if err != nil {
					cmd.Println(err)
					lastLineErrored = true
				} else {
					lastLineErrored = false
				}

				cmd.Println(result)
			default:
				cmd.Printf("Command '%v' not found.", fields[0])
			}

			// add to history
			state.AppendHistory(input)
		}
	},
}
