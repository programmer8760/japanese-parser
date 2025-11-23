package app

import (
	"fmt"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/programmer8760/japanese-parser/backend/parser"
	"github.com/programmer8760/japanese-parser/backend/types"
	"github.com/programmer8760/japanese-parser/backend/utils"
	"github.com/programmer8760/japanese-parser/backend/exporter"
)

// App struct
type App struct {
	ctx context.Context
	parser *parser.Parser
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	go func() {
		var err error
		a.parser, err = parser.NewParser()
		if err != nil {
			fmt.Println("Error initializing parser: ", err.Error())
			return
		}

		runtime.EventsEmit(ctx, "parserReady", true)
	}()
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) Parse(text string) (types.ParserResult, error) {
	if a.parser == nil {
		return types.ParserResult{}, fmt.Errorf("parser not initialized")
	}

	return a.parser.Parse(text), nil
}

func (a *App) GetUniqueTokens(stats types.POSStats) map[string][]types.Token {
	return utils.GetUniqueTokens(stats)
}

func (a *App) SaveFile(extension string) (string, error) {
    path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
        Title:           "Сохранить как",
        DefaultFilename: "output" + extension,
        Filters: []runtime.FileFilter{
					{DisplayName: "*" + extension, Pattern: "*" + extension},
        },
    })
    if err != nil {
        return "", err
    }
    return path, nil
}

func (a *App) ExportTxt(parserResult types.ParserResult) error {
	path, err := a.SaveFile(".txt")
	if err != nil {
		return err
	}

	if path == "" {
		return nil
	}

	return exporter.ExportTxt(parserResult, path)
}

func (a *App) ExportCsv(parserResult types.ParserResult) error {
	path, err := a.SaveFile(".csv")
	if err != nil {
		return err
	}

	if path == "" {
		return nil
	}

	return exporter.ExportCsv(parserResult, path)
}

