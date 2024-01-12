package gzsave

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Tom5521/GZGoLauncher/internal/config"
	"github.com/Tom5521/GZGoLauncher/pkg/gzrun"
	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/ncruces/zenity"
)

var FilePath = config.CurrentPath + "/runner-config.json"

func errWin(txt ...any) {
	err := zenity.Error(fmt.Sprint(txt...))
	if err != nil {
		messages.FatalError(err)
	}
}

func Save(r gzrun.Pars) {
	bytedata, err := json.Marshal(r)
	if err != nil {
		errWin(err)
		return
	}
	err = os.WriteFile(FilePath, bytedata, os.ModePerm)
	if err != nil {
		errWin(err)
		return
	}
}

func Read() gzrun.Pars {
	var pars gzrun.Pars
	if _, err := os.Stat(FilePath); os.IsNotExist(err) {
		Save(pars)
		return pars
	}
	bytedata, err := os.ReadFile(FilePath)
	if err != nil {
		errWin(err)
		return pars
	}
	err = json.Unmarshal(bytedata, &pars)
	if err != nil {
		errWin(err)
		return pars
	}
	return pars
}
