package main

import (
	"github.com/marcuschui2022/blog-aggregator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	//init cfg
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	appState := &state{
		cfg: &cfg,
	}

	//init cmds
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	//get cli Args
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	//init cmd
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	//execute cmd
	err = cmds.run(appState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
