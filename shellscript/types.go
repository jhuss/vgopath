package shellscript

type(
	InitOpts struct {
		VenvName string
		Gopath string
	}

	LinuxScript struct {
		Shebang string
		ScriptDir string
		ProjectPatch string
		GopathEnvSys string
		Ps1EnvSys string
		UnsetFunction string
	}
)
