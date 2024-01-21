package gui

import (
	"github.com/Tom5521/GZGoLauncher/internal/tools"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
)

func ErrWin(text ...any) {
	msg.Error(text...)
	tools.ErrWin(text...)
}
