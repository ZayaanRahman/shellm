// Change so that use ask command to prompt, and enter others to edit field

// commands will be help, ask, info, model, key, prompt

package main

import (
	"os"
)

func main() {

	process(os.Args[1:])
}
