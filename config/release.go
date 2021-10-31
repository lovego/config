package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/lovego/fs"
)

// DetectReleaseConfigDirOf return the release config dir of a major environment,
// which contains clusters.yml, deploy.yml.
func DetectReleaseConfigDirOf(majorEnv string) string {
	releaseDir, hasDirectContent := DetectReleaseDir()
	if releaseDir == "" {
		return ""
	}
	if majorEnv == "" || hasDirectContent {
		return releaseDir
	}
	if majorEnv == "" {
		majorEnv = "default"
	}
	return filepath.Join(releaseDir, majorEnv)
}

// DetectReleaseDir return the top release dir of a project.
func DetectReleaseDir() (dir string, hasDirectContent bool) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	if releaseDir := detectReleaseDir(cwd, true); releaseDir != "" {
		return releaseDir, true
	}
	if releaseDir := detectReleaseDir(cwd, false); releaseDir != "" {
		return releaseDir, false
	}
	return "", false
}

func detectReleaseDir(cwd string, directContent bool) string {
	var dir = "release/"
	if !directContent {
		dir += "*/"
	}
	if result := fs.DetectDir(
		cwd, dir+"clusters.yml", dir+"deploy.yml", dir+"img-*",
	); result != "" {
		return filepath.Join(result, "release")
	}
	return ""
}
