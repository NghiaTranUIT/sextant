package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

const sh = `#!bin/sh
if [ -n "${BASH}" ]; then
	shell="bash"
elif [ -n "${ZSH_NAME}" ]; then
	shell="zsh"
else
	shell=$(echo "${SHELL}" | awk -F/ '{ print $NF }')
fi
if [ "${shell}" = "sh" ]; then
	return 0
fi
eval "$(sextant shell --type "$shell")"
`

const zsh = `__sextant_chpwd() {
	[[ "$(pwd)" == "$HOME" ]] && return
    sextant add "$(pwd)"
}

typeset -gaU chpwd_functions

chpwd_functions+=__sextant_chpwd
`

const bash = `__sextant_chpwd() {
	[[ "$(pwd)" == "$HOME" ]] && return
    sextant add "$(pwd)"
}
grep "sextant add" <<< "$PROMPT_COMMAND" >/dev/null || {
	PROMPT_COMMAND="$PROMPT_COMMAND"$'\n''(__sextant_chpwd 2>/dev/null &);'
}
`

func CmdShell(c *cli.Context) error {
	shell := c.String("type")
	fmt.Fprint(c.App.Writer, scriptForShell(shell))
	return nil
}

func scriptForShell(shell string) string {
	switch shell {
	case "sh":
		return sh
	case "bash":
		return bash
	case "zsh":
		return zsh
	default:
		return fmt.Sprintf("echo Sextant: We don't support %s shell yet :(", shell)
	}
}