package astrobot

import (
	"log"

	"github.com/drpaneas/dudenetes/pkg/run"
)

// BuildFails checks if Hugo build fails
func BuildFails() bool {
	cmd := "hugo --gc --themesDir themes"
	log.Println("Running: ", cmd)
	timeout := 10
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		log.Println(err)
		// gitResetHardHead()
		return true
	}
	log.Println(output)
	return false
}
