package main

import (
	"fmt"
	"bufio"
	//"log"
	"os"
	"encoding/json"
	"strings"
)

// Get user data from stdio
func collectDataFromUser(input string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(input + ": ")
	data, _ := reader.ReadString('\n')
	return strings.Trim(data, "\n")
}

// Write data in config
func cacheUsersDataInFile(data map[string]string) {
		pathToConfig := os.Getenv("HOME") + "/.deb-build-config"
		file, err := os.Open(pathToConfig)
		defer file.Close()


		if err != nil {
			newFile, _ := os.Create(pathToConfig)
			userDataToJson, _ := json.Marshal(data)
			newFile.WriteString(string(userDataToJson))
		}
}

func main() {
	userData := make(map[string]string)

	userData["host"] = collectDataFromUser("Host to intall")
	userData["emails"] = collectDataFromUser("Emails to notify, (string splited spaces, for example - khvostov@yamoney.ru vladkhvo@gmial.com")
	cacheUsersDataInFile(userData)
}
