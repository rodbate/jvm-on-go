package app

type jvmOptions struct {
	helpFlag    bool
	versionFlag bool
	verbose     bool
	cpOption    string
	class       string
	args        []string
}

func newOptions() *jvmOptions {
	return &jvmOptions{}
}
