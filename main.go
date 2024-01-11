package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var timeout time.Duration

func main() {
  //change this line below to fit your own settings
	filePath := "C:/Users/Derp/AppData/Roaming/Code/User/settings.json"
	animeGifsPath := "F:/dank stuff/animu/cute gifs/vsc"
	flag.DurationVar(&timeout, "timeout", 60*time.Minute, "specify the timeout duration in minutes")
	repeatFlagPtr := flag.Bool("repeat", false, "a bool")
	stringFlagPtr := flag.String("c", "", "a string")
	flag.Parse()
	for {

		fileNames, err := getAllFileNames(animeGifsPath)
		randomIndex := rand.Intn(len(fileNames))
		randomFileName := fileNames[randomIndex]
		randomFilePath := animeGifsPath + "/" + randomFileName
		if len(*stringFlagPtr) > 1 {
			randomFilePath = animeGifsPath + "/" + *stringFlagPtr + ".gif"
		}
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		var modifiedLines []string

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "doki.sticker.path") {
				newSticker := `"doki.sticker.path": "` + randomFilePath + `",`
				line = newSticker
			}
			modifiedLines = append(modifiedLines, line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		file.Close()
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		for _, line := range modifiedLines {
			_, err = fmt.Fprintln(file, line)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		if !*repeatFlagPtr {
			return
		}
		time.Sleep(timeout)
	}
}
