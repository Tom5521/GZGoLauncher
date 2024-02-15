package tools

import (
	"errors"
	"fmt"
	"os"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotkit/components/errpopup"
	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"
)

var (
	ParentWin        *gtk.Window
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
	if ParentWin == nil {
		ZenityErrWin(txt...)
		msg.Error(txt...)
		return
	}
	var errs []error
	for _, t := range txt {
		errs = append(errs, errors.New(fmt.Sprint(t)))
	}
	errpopup.Show(ParentWin, errs, func() {
		msg.Error(txt...)
	})
}

func FatalErrWin(txt ...any) {
	ErrWin(txt...)
	os.Exit(FatalErrExitCode)
}
