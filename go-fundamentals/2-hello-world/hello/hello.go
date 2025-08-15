package hello

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if len(name) == 0 {
		return "Hello, World"
	}
	return englishHelloPrefix + name
}
