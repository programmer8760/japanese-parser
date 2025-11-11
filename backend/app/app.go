package app

import (
	"fmt"
	"context"

	"github.com/programmer8760/japanese-parser/backend/parser"
	"github.com/programmer8760/japanese-parser/backend/types"
	"github.com/programmer8760/japanese-parser/backend/dictionary"
)

// App struct
type App struct {
	ctx context.Context
	parser *parser.Parser
	dictionary *dictionary.Dictionary
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	var err error
	a.parser, err = parser.NewParser()
	if err != nil {
		fmt.Println("Error initializing parser: ", err.Error())
	}

	a.dictionary, err = dictionary.NewDictionary()
	if err != nil {
		fmt.Println("Error initializing dictionary: ", err.Error())
	}
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

func (a *App) Tokenize(text string) ([]types.Token, error) {
	if a.parser == nil {
		return nil, fmt.Errorf("parser not initialized")
	}
	
	return a.parser.Tokenize(text)
}

func (a *App) Lookup(kanji string, reading string) ([]types.DictionaryEntry, error) {
	if a.dictionary == nil {
		return nil, fmt.Errorf("dictionary not initialized")
	}

	return a.dictionary.Lookup(kanji, reading)
}
