package shellscript

import "fmt"

func createScript(options InitOpts) func() string {
	var script LinuxScript

	script.Shebang = "#!/usr/bin/env bash"
	script.ScriptDir = fmt.Sprintf(`"%s"`, options.Gopath)
	script.ProjectPatch = `"$( dirname "${SCRIPT_DIR}" )"`
	script.GopathEnvSys = `"${GOPATH}"`
	script.Ps1EnvSys = `"${PS1}"`
	script.UnsetFunction = `ungopath() {
    unset GOPATH
    unset PS1

    export GOPATH="${GOPATH_ENV_SYS}"
    export PS1="${PS1_ENV_SYS}"

    unset GOPATH_ENV_SYS
    unset PS1_ENV_SYS
    unset -f ungopath
}`

	return func() string {
		return fmt.Sprintf(
			`%s

# vars
SCRIPT_DIR=%s
PROJECT_PATH=%s
GOPATH_ENV_SYS=%s
PS1_ENV_SYS=%s

# unset
%s

# set project as GOPATH
export GOPATH="${SCRIPT_DIR}"
export PS1="${PS1}(%s) "
`,
			script.Shebang,
			script.ScriptDir,
			script.ProjectPatch,
			script.GopathEnvSys,
			script.Ps1EnvSys,
			script.UnsetFunction,
			options.VenvName,
		)
	}
}
