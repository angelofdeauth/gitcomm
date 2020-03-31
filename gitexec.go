// @File:     gitexec.go
// @Created:  2020-03-31 00:57:45
// @Modified: 2020-03-31 01:02:47
// @Author:   aescaler
// @Commiter: aescaler
// @Mail:     aescaler@redhat.com
// @Copy:     Copyright © 2020 aescaler <aescaler@redhat.com>

package gitcomm

import (
	"os/exec"
	"strings"
)

// GitExec function performs git workflow
func GitExec(addAll, show bool, msg string) {
	if addAll {
		gitColorCmd("add", "-A")
	}
	gitColorCmd("commit", "-m", msg)
	if show {
		gitColorCmd("show", "-s")
	}
}

// CheckForUncommited function checks if there are changes that need commit
func CheckForUncommited() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.CombinedOutput()
	CheckIfError(err)
	return len(out) != 0
}

// CheckIsGitDir function checks is dir inside git worktree
func CheckIsGitDir() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	out, err := cmd.Output()
	isGitDir := strings.TrimSpace(string(out))
	if err == nil && isGitDir == "true" {
		return true
	}
	return false
}
