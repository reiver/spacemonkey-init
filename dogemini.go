package main

import (
	"github.com/reiver/spacemonkey-init/sm"
	"github.com/reiver/spacemonkey-init/srv/log"

	"fmt"
	"os"
	"path/filepath"
)

const (
	geminiREADME =
		`⸿ Gemini Protocol`                                                                                  +"\r\n"+
		``                                                                                                   +"\r\n"+
		`This directory is meant to store the user's user-data for the Gemini Protocol`                      +"\r\n"+
		``                                                                                                   +"\r\n"+
		`If the user is inclined to backup their data then the user may want to backup this directory too.`  +"\r\n"+
		``                                                                                                   +"\r\n"+
		`§ SPACEMONKEY`                                                                                      +"\r\n"+
		``                                                                                                   +"\r\n"+
		`This directory was created by a program named: spacemonkey`                                         +"\r\n"+
		``                                                                                                   +"\r\n"+
		`https://github.com/reiver/spacemonkey`                                                              +"\r\n"+
		``                                                                                                   +"\r\n"+
		`But although spacemonkey created this directory it does not mean other programs cannot use it too.` +"\r\n"+
		`In fact, it is hoped that other programs will use use this directory too.`                          +"\r\n"+
		`Just stick to the same conventions so that all programs using this directory can interoperate.`     +"\r\n"+
		``                                                                                                   +"\r\n"+
		`♜ shâh mât`                                                                                         +"\r\n"
)

func doGemini() error {

	var geminipath string
	{
		var err error

		const name = "protocol/gemini/data-path"
		geminipath, err = sm.Cfg(name)

		if nil != err {
			return fmt.Errorf("problem figuring the path to store the user's user-data for the Gemini Protocol: %s", name, err)
		}
	}

	{
		const permissions = 0700

		err := os.Mkdir(geminipath, permissions)
		if nil != err {
			return fmt.Errorf("problem: could not create %q: %s", geminipath, err)
		}
	}
	logsrv.Log("success — created hidden directory for Gemini Protocol at %q", geminipath)

	var readmepath string
	{
		const readmename string = "README.txt"

		readmepath = filepath.Join(geminipath, readmename)
	}
	logsrv.Log("readme-path:", readmepath)

	{
		const permissions = 0644

		os.WriteFile(readmepath,
		[]byte(geminiREADME),
		permissions)
	}
	logsrv.Logf("success — created README.txt file for Gemini Protocol at %q", readmepath)

	{
		fmt.Println("Initialized empty directory for Gemini Protocol at", geminipath)
	}

	return nil
}
