// Package configuration implements methods for get config variables.
//
// Config main project.
// jenkins/gdo/compiled = go main.
// jenkins/gdo/.env.
// jenkins/gdo/config/soa.php.
package configuration

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

const (
	// PathBin expose compiled dir
	PathBin = "compiled"

	// PathScript expose script dir
	PathScript = "script"

	// PathFile expose file dir
	PathFile = "file"

	// PathLib expose lib dir
	PathLib = "lib"

	// PathApp expose app dir
	PathApp = "app"

	// PathConfig expose config dir
	PathConfig = "config"

	// PathStorage expose storage dir
	PathStorage = "storage"

	// PathLogs expose logs dir
	PathLogs = "logs"
)

var (
	// DirCwd current working dir
	DirCwd string

	// DirProject - php framework
	DirProject string

	// DirProjectApp - php app dir
	DirProjectApp string

	// DirProjectConfig - php config dir
	DirProjectConfig string

	// DirProjectStorage - php storage dir
	DirProjectStorage string

	// DirScript - python framework
	DirScript string

	// DirScriptApp - python app dir
	DirScriptApp string

	// DirScriptConfig -python config dir
	DirScriptConfig string

	// DirScriptStorage -python storage dir
	DirScriptStorage string

	// DirBinStorage - go storage dir
	DirBinStorage string
)

func init() {
	// get current working directory
	cwd, _ := os.Getwd()
	cwd, _ = filepath.EvalSymlinks(cwd)

	relativeDir := ".."

	switch baseDir := filepath.Base(cwd); baseDir {
	case PathBin:
		// pass
	case PathConfig, PathFile, PathLib:
		relativeDir = filepath.Join(relativeDir, "..")
	default:
		relativeDir = filepath.Join(relativeDir, "..", "..")
	}

	projectDir, _ := filepath.Abs(filepath.Join(cwd, relativeDir))

	DirCwd = cwd            // /
	DirProject = projectDir // gdo/
	BuildProject()
}

// BuildProject (or rebuild) dir structure
func BuildProject() {
	// DirProjectApp - gdo/app
	DirProjectApp = filepath.Join(DirProject, PathApp)

	// DirProjectConfig - gdo/config
	DirProjectConfig = filepath.Join(DirProject, PathConfig)

	// DirProjectStorage - gdo/storage
	DirProjectStorage = filepath.Join(DirProject, PathStorage)

	// DirScript - gdo/script
	DirScript = filepath.Join(DirProject, PathScript)

	// DirScriptApp - gdo/script/app
	DirScriptApp = filepath.Join(DirScript, PathApp)

	// DirScriptConfig - gdo/script/app/config
	DirScriptConfig = filepath.Join(DirScriptApp, PathConfig)

	// DirScriptStorage - gdo/script/storage
	DirScriptStorage = filepath.Join(DirScript, PathStorage)

	// DirBinStorage - compiled/script/storage
	DirBinStorage = filepath.Join(DirProject, PathBin, PathStorage)
}

// Read a file, from a path and extract configuration from regexp.
func cfg(filePath string, regexpString string) (object map[string][]byte, err error) {
	// content of file
	var bytesContent []byte

	isAbs := path.IsAbs(filePath)

	// absolute path
	if isAbs {
		bytesContent, err = ioutil.ReadFile(filePath)
	}

	// directory path
	if !isAbs && err == nil {
		bytesContent, err = ioutil.ReadFile(filepath.Join(DirProject, filePath))
	}

	// relative path
	if err != nil {
		bytesContent, err = ioutil.ReadFile(filePath)
	}

	// open and read an absolute file
	if err != nil {
		return
	}

	// compiled regex
	re := regexp.MustCompile(regexpString)
	names := re.SubexpNames()
	subexpressions := re.NumSubexp()
	matches := re.FindAllSubmatch(bytesContent, -1)
	if matchesLen := len(matches); matchesLen < subexpressions {
		subexpressions = matchesLen
	}

	// temporary dict of results
	object = map[string][]byte{}
	for i, key := 0, 1; i < subexpressions; i, key = i+1, key+1 {
		match := matches[i]
		name := names[key]
		// for key, match := range matches[i] {

		// if len(match) > 0 && len(name) > 0 {
		object[name] = match[key]
		// }

		// }
	}

	return
}
