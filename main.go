// Change so that use ask command to prompt, and enter others to edit field

// commands will be help, ask, info, model, key, prompt

package main

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "io"
	// "net/http"
	"os"
	// "unicode/utf8"
)

func main() {

	process(os.Args[1:])

	// 	// old code

	// 	args := os.Args[1:]

	// 	if noArgs(args) {
	// 		// --prompt=... enter custom prompt to use during API calls, set to default to return to original
	// 		fmt.Println(`
	// Enter a query or edit options:

	// --model=...  choose the model, currently only supports gpt4
	// --key=...    enter private API key to access model

	// All other input will be interpreted as a query to the selected model.`)

	// 	} else if isQuery(args) {
	// 		processQuery(args)

	// 	} else {
	// 		updateConfig(args)

	// 	}
}

// func isQuery(args []string) bool {
// 	char0, _ := utf8.DecodeRuneInString(args[0])
// 	return char0 != '-'
// }

// func noArgs(args []string) bool {
// 	return len(args) == 0
// }

// type Body struct {
// 	Model    string              `json:"model"`
// 	Messages []map[string]string `json:"messages"`
// }

// func processQuery(query []string) {

// 	message := ""
// 	for _, word := range query {
// 		message += word + " "
// 	}

// 	config := Config{}
// 	json.Unmarshal(readFile(), &config)

// 	body := Body{Model: config.Model, Messages: []map[string]string{

// 		{"role": "system", "message": config.Prompt},
// 		{"role": "user", "message": message},
// 	}}

// 	endpoint := "https://api.openai.com/v1/chat/completions"

// 	jsonReq, err := json.Marshal(body)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonReq))

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+config.Key)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer resp.Body.Close()

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	fmt.Println(string(respBody))
// }
