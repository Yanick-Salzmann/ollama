package llm

import (
	"embed"
)

//go:embed llama.cpp/build/linux/.so*
var libEmbed embed.FS
