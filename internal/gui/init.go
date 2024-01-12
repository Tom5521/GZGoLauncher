package gui

import (
	"log"
	"reflect"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GZGoLauncher/internal/config"
	"github.com/Tom5521/GZGoLauncher/locales"
	"github.com/Tom5521/GZGoLauncher/pkg/gzrun"
	"github.com/Tom5521/GZGoLauncher/pkg/gzsave"
)

type Gui struct {
	ui
}

type ui struct {
	App           fyne.App
	MainWindow    fyne.Window
	WadList       *widget.List
	ModsList      *widget.List
	ZRunnerSelect *widget.Select
	CustomArgs    *widget.Entry
}

var (
	Runner   gzrun.Pars = gzsave.Read()
	settings            = &config.Settings
	po                  = locales.GetPo(settings.Lang)
)

func AutoSaver() {
	for {
		oldRunner := Runner
		oldConfig := config.Settings
		time.Sleep(50 * time.Millisecond)
		if !reflect.DeepEqual(oldRunner, Runner) {
			log.Println("Runner es diferente")
			gzsave.Save(Runner)
		}
		if !reflect.DeepEqual(oldConfig, config.Settings) {
			log.Println("Config es diferente")
			err := settings.Write()
			if err != nil {
				ErrWin(err)
			}
		}
	}
}

func Init() *Gui {
	app := app.New()
	ui := &Gui{}
	ui.App = app
	// Initialize auto saver
	go AutoSaver()
	return ui
}
