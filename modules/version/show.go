package version

import "fmt"

var Version = "0.0.1"
var BuildDate string
var GitCommit string

func ShowVersion() {
	fmt.Printf("Commit: %s \nBuildDate: %s\n", GitCommit, BuildDate)
}
