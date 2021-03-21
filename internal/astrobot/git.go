package astrobot

import (
	"fmt"
	"log"
	"strings"

	"github.com/drpaneas/dudenetes/pkg/run"
)

// ChangeBranch changes to another branch
func ChangeBranch(branch string) {
	cmd := fmt.Sprintf("git checkout -b %s", branch)
	log.Println("Running: ", cmd)
	timeout := 5
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		CheckoutMaster()
		log.Fatal(err)
	}
	log.Println(output)
	output, _ = run.SlowCmdDir("git branch --show-current", timeout, directory)
	log.Println(output)
}

// IsItUpToDate checks if the repo is clear
func IsItUpToDate() (answer bool) {
	log.Println("Checking if repo is up to date ...")
	answer = false
	cmd := "git pull origin master"
	timeout := 5
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		CheckoutMaster()
		log.Fatal(err)
	}
	if strings.Contains(output, "Already up to date.") {
		answer = true
	}
	return answer
}

// GitAdd runs git add
func GitAdd() {
	log.Println("Running git add .")
	cmd := "git add ."
	timeout := 5
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		branch, _ := run.SlowCmdDir("git --no-pager log -1", timeout, directory)
		revertCleanState(branch)
		log.Fatal(err)
	}
	log.Println(output)
	output, _ = run.SlowCmdDir("git status", timeout, directory)
	log.Println(output)
}

// GitCommit commits the stuff
func GitCommit() {
	cmd := "git commit -a --allow-empty-message -m ''"
	log.Println("Running: ", cmd)
	timeout := 5
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		branch, _ := run.SlowCmdDir("git --no-pager log -1", timeout, directory)
		revertCleanState(branch)
		log.Fatal(err)
	}
	log.Println(output)
	output, _ = run.SlowCmdDir("git --no-pager log -1", timeout, directory)
	log.Println(output)
}

// GitPush on a given branch
func GitPush(branch string) {
	cmd := fmt.Sprintf("git push origin %s", branch)
	log.Println("Running: ", cmd)
	timeout := 10
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		branch, _ := run.SlowCmdDir("git --no-pager log -1", timeout, directory)
		revertCleanState(branch)
		log.Fatal(err)
	}
	log.Println(output)
}

// CheckoutMaster checks out at master
func CheckoutMaster() {
	log.Println("Checking out at master branch")
	cmd := "git checkout master"
	timeout := 5
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	fmt.Println("Running the command: " + cmd + " at the directory: " + directory)
	fmt.Println("The output is: ", output)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(output)
}

// DeleteBranch deletes a branch with force
func DeleteBranch(branch string) {
	cmd := fmt.Sprintf("git branch -D %s", branch)
	log.Println("Running: ", cmd)
	timeout := 5
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(output)
}

// gitResetHardHead deletes any changes to be commited
func gitResetHardHead() {
	cmd := "git reset --hard HEAD"
	log.Println("Running: ", cmd)
	timeout := 5
	output, err := run.SlowCmdDir(cmd, timeout, directory)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(output)
}

func revertCleanState(branch string) {
	CheckoutMaster()
	gitResetHardHead()
	DeleteBranch(branch)
}
