package app

type jvmOptions struct {
	verbose     bool
	cpOption    string
	class       string
	args        []string
}

func newOptions() *jvmOptions {
	return &jvmOptions{}
}
