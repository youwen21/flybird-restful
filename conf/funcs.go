package conf

import (
	"fmt"
	"gofly/conf/flag_vars"
	"os"
	"runtime"
	"strings"
)

func GetEnv(key, dft string) string {
	s := os.Getenv(key)
	if s == "" {
		return dft
	}
	return s
}

func GetAddress() string {
	return flag_vars.GetAddress()
}

func IsRunInProduction() bool {
	return BuildGo != ""
}

func GetBuildInfo() map[string]interface{} {
	res := map[string]interface{}{
		"buildTime":      BuildTime,
		"buildVersion":   BuildVersion,
		"go":             BuildGo,
		"branch":         BuildGitBranch,
		"lastCommitUser": BuildLastCommitUser,
		"lastCommitTime": BuildLastCommitTime,
		"lastCommitId":   BuildLastCommitId,
		"lastCommitMsg":  BuildLastCommitMsg,
		"isRunInBuild":   IsRunInProduction(),
		"version":        Version,
	}
	return res
}

func getProjROOT() string {
	if IsRunInProduction() {
		return GetEnv("ROOT", "./")
	}
	_, fpath, _, _ := runtime.Caller(0)
	ProjROOT := strings.Replace(fpath, "/conf/env.go", "", 1)
	return ProjROOT
}

func GetConfigNameByRunMode() string {
	configName := "config"
	if flag_vars.GetRunMode() != "" {
		configName = fmt.Sprintf("config_%s", flag_vars.GetRunMode())
	}

	return configName
}

func GetVersion() string {
	if Version != "" {
		return Version
	}

	if BuildTime != "" {
		return BuildTime
	}

	return ""
}
