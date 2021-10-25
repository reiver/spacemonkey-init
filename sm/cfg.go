package sm

import (
	"github.com/reiver/spacemonkey-init/srv/log"

	"fmt"
	"os/exec"
	"strings"
)

// spacemonkey-cfg
func Cfg(name string) (string, error) {

	var output     strings.Builder
	var complaints strings.Builder

	const programname = "spacemonkey-cfg"

	var cmd *exec.Cmd
	{
		cmd = exec.Command(programname, name)
		if nil == cmd {
			return "", errInternalError
		}

		cmd.Stdout = &output
		cmd.Stderr = &complaints
	}

	{
		err := cmd.Run()
		if nil != err {
			return "", fmt.Errorf("program running `%s %s`: %s; %s", programname, name, err, complaints.String())
		}
	}

	var value string
	{
		value = output.String()
	}
	logsrv.Logf("value: %q", value)
	if "" == strings.TrimSpace(value) {
		return "", fmt.Errorf("bad value %q for name %q", value, name)
	}

	return value, nil
}
