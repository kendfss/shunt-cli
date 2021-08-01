package pathtools

import (
	"os"
	"path/filepath"
	"io/fs"
	
	// et"tildegit.org/eli2and40/rube/errortools"
	et"tildegit.org/eli2and40/rube/cli/errortools"
)

func getDirPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		exe, err := os.Executable()
		
		et.Assert("Couldn't find user's home directory or the executable")
		
		home = filepath.Dir(exe)
	}
	return filepath.Join(home, ".config", "gonzo", "rube")
}

func GameFile() string {
	return filepath.Join(getDirPath(), "game.json")
}
func BindingsFile() string {
	return filepath.Join(getDirPath(), "bindings.json")
}
func AssureFile(pth string) {
	// If a file does not exist, all necessary directories will be created for it
	if !Exists(pth) {
		dir, _ := filepath.Dir(pth)
		AssureDir(dir)
	}
}
func AssureDir(pth string) {
	// If a directory does not exist it will be created
	if !Exists(pth) {
		os.MkdirAll(pth, fs.ModeDir)
	}
}
func Exists(pth string) bool {
    _, err := os.Lstat(pth)
    return !et.Bool(err)
}
