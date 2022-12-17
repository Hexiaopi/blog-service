package main

import (
	"fmt"
	"runtime"
	"encoding/json"

	"github.com/spf13/cobra"
)

var (
	gitVersion   = "v0.0.0-master+$Format:%H$"
	gitCommit    = "$Format:%H$" // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState = ""            // state of git tree, either "clean" or "dirty"
	gitBranch    = ""

	buildDate = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')

	version = &Version{
		Version:      gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		GitBranch:    gitBranch,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

type Version struct {
	Version      string `json:"version"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	GitBranch    string `json:"gitBranch"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Platform     string `json:"platform"`
}

func (v Version) String() string {
	data, _ := json.MarshalIndent(v, "", "    ")
	return string(data)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of blog-service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
