package anonymousfunctions

func MakeGreeting(greet string) func(name string) string {
	return func(name string) string {
		return greet + " " + name + "!"
	}
}
