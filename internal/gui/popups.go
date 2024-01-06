package gui

import (
	"fmt"

	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/ncruces/zenity"
)

func ErrWin(text ...any) {
	txt := fmt.Sprint(text...)
	messages.Error(txt)
	err := zenity.Error(txt)
	if err != nil {
		fmt.Println(err)
	}
}
