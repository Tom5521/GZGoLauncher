package tools

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"
)

var (
	po               *gotext.Po
	FatalErrExitCode int
)

func ReceivePo(p *gotext.Po) {
	po = p
}

func ZenityErrWin(txt ...any) {
	errtext := fmt.Sprint(txt...)
	err := zenity.Error(errtext)
	if err != nil {
		msg.Error(err)
	}
}

func ErrWin(txt ...any) {
	msg.Error(txt...)
	app := fyne.CurrentApp()
	if app == nil || po == nil {
		ZenityErrWin(txt...)
		return
	}
	text := fmt.Sprint(txt...)
	w := app.NewWindow(po.Get("Error"))
	w.Resize(fyne.NewSize(300, 150))
	w.SetIcon(theme.ErrorIcon())
	label := &widget.Label{
		Alignment: fyne.TextAlignCenter,
		Text:      text,
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}
	button := &widget.Button{
		Text: po.Get("Accept"),
		OnTapped: func() {
			w.Close()
		},
		Importance: widget.DangerImportance,
	}

	buttonBox := boxes.NewCenter(button)
	content := boxes.NewBorder(nil, buttonBox, nil, nil, label)
	w.SetContent(content)
	w.Show()
	w.RequestFocus()
}

func FatalErrWin(txt ...any) {
	ErrWin(txt...)
	os.Exit(FatalErrExitCode)
}
