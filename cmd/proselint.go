package cmd

import "github.com/skybet/cali"
import "os/user"
import log "github.com/Sirupsen/logrus"

func init() {

	command := cli.NewCommand("proselint [command]")
	command.SetShort("Run proselint")
	command.SetLong(`Starts a container and runs rake proselint.
Examples:
  To render the site locally.
  # staticli proselint foo.md
`)

	task := command.Task("staticli/proselint")
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to find uid for user: %s", err)
	}
	task.AddEnv("HOST_USER_ID", u.Uid)
	task.AddEnv("HOST_GROUP_ID", u.Gid)
	task.SetInitFunc(func(t *cali.Task, args []string) {

	})
}
