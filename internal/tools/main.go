package tools

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

func ErrWin(txt ...any) {
	msg.Error(txt...)
	app := fyne.CurrentApp()
	if app == nil || po == nil {
		text := fmt.Sprint(txt...)
		err := zenity.Error(text)
		if err != nil {
			msg.Error(err)
		}
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
	content := container.NewBorder(nil, buttonBox, nil, nil, label)
	w.SetContent(content)
	w.Show()
	w.RequestFocus()
}

func FatalErrWin(txt ...any) {
	ErrWin(txt...)
	os.Exit(FatalErrExitCode)
}
