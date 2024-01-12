package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Run mg.Namespace

func (Run) Current() error {
	return sh.RunV("go", "run", "-v", MainFile)
}

func (Run) Windows() error {
	return sh.RunWithV(WindowsEnv, "go", "run", "-v", "-tags", "windows", MainFile)
}

func (Run) Linux() error {
	return sh.RunV("go", "run", "-v", "-tags", "linux", MainFile)
}
