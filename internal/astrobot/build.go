package astrobot

import (
	"log"

	"github.com/drpaneas/dudenetes/pkg/run"
)

// BuildFails checks if Hugo build succeeds
func BuildFails() bool {
	cmd := "hugo --gc --themesDir themes"
	log.Println("Running: ", cmd)
	timeout := 10
	directory := "/Users/drpaneas/github/starlordgr"
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		log.Println("git resett hard ...")
		gitResetHardHead()
		return true
	}
	log.Println(output)
	return false
}
