package gui

import (
	"os"
	"reflect"
	"time"

	"github.com/Tom5521/GZLauncher-gtk/internal/config"
	"github.com/Tom5521/GZLauncher-gtk/internal/tools"
	"github.com/Tom5521/GZLauncher-gtk/locales"
	"github.com/Tom5521/GZLauncher-gtk/pkg/zdoom/save"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

var (
	po       = locales.Current
	settings = &config.Settings
	runner   = save.Read()
)

type ui struct {
	Application *gtk.Application
	MainWindow  *gtk.ApplicationWindow

	WadsListView   *gtk.ListView
	WadsListModel  *gtk.StringList
	WadsListSModel *gtk.SingleSelection

	ModsListView   *gtk.ListView
	ModsListModel  *gtk.StringList
	ModsListSModel *gtk.NoSelection
}

func autoSaver() {
	for {
		oldSettings := config.Settings
		oldRunner := runner
		time.Sleep(50 * time.Millisecond)
		if !reflect.DeepEqual(config.Settings, oldSettings) {
			err := settings.Write()
			if err != nil {
				tools.ErrWin(err)
			}
		}
		if !reflect.DeepEqual(runner, oldRunner) {
			save.Save(runner)
		}
	}
}

const AppID string = "com.github.Tom5521.GZLauncher"

func InitGtk() {
	app := gtk.NewApplication(AppID, gio.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		activate(app)
	})
	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	go autoSaver()
	ui := &ui{Application: app}

	ui.MainWindow = gtk.NewApplicationWindow(app)
	ui.MainWindow.SetDefaultSize(1100, 500)
	ui.MainWindow.SetResizable(true)

	tools.ParentWin = &ui.MainWindow.Window

	tabs := gtk.NewNotebook()

	runBox := ui.StartTab()
	tabs.AppendPage(runBox, gtk.NewLabel(po.Get("Run")))
	runTab := tabs.Page(runBox)
	runTab.SetObjectProperty("tab-expand", true)

	settingsBox := ui.SettingsTab()
	tabs.AppendPage(settingsBox, gtk.NewLabel(po.Get("Settings")))
	settingsTab := tabs.Page(settingsBox)
	settingsTab.SetObjectProperty("tab-expand", true)

	ui.MainWindow.SetChild(tabs)
	ui.MainWindow.Show()
}
