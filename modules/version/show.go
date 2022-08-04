package version

import "fmt"

var gitCommit string
var buildDate string

func ShowVersion() {
	fmt.Printf("Commit: %s \nBuildDate: %s\n", gitCommit, buildDate)
}
