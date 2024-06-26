package main

import (
	"client/backend/config"
	"embed"
	"github.com/getlantern/elevate"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"os"
	exec2 "os/exec"
	"path/filepath"
	"strings"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if len(os.Args) == 1 && strings.Contains(os.Args[0], "gpp") {
		var command *exec2.Cmd
		if _, err := os.Stat(".dev"); err == nil {
			command = exec2.Command(os.Args[0], filepath.Dir(os.Args[0]))
		} else {
			command = elevate.Command(os.Args[0], filepath.Dir(os.Args[0]))
		}
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		command.Stdin = os.Stdin
		_ = command.Start()
		_ = command.Wait()
		os.Exit(0)
	} else if !strings.Contains(os.Args[0], "gpp") {
		return
	}
	_ = os.Chdir(os.Args[1])
	if strings.Contains(os.Args[0], "gpp") {
		config.InitConfig()
	}
	// Create an instance of the app structure
	app := NewApp()
	defer app.Stop()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "gpp",
		Width:         360,
		Height:        520,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
