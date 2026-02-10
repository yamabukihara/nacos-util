package utils

import (
	"log"

	"github.com/alecthomas/chroma/quick"
)

var plainLogger = log.New(log.Writer(), "", 0)

func PrettyYAML(content string) {
	err := quick.Highlight(log.Writer(), content, "yaml", "terminal256", "monokai")
	if err != nil {
		log.Println(content)
	}
}
