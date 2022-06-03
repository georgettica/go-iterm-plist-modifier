package main

import (
	"encoding/json"
	"fmt"
	"howett.net/plist"
	"os"
)

func main() {
	itermPlistModifierLocationName := "ITERM_PLIST_MODIFIER_LOCATION"
	itermPlistModifierLocation, hasItermPlistModifierLocation := os.LookupEnv(itermPlistModifierLocationName)
	if !hasItermPlistModifierLocation {
		fmt.Printf("could not find the envvar '%s', exiting\n",itermPlistModifierLocationName)
		os.Exit(1)
	}
	//plistLocation := "/Users/georgettica/Documents/internal/dekel-macos-installation/com.googlecode.iterm2.plist"
	fileContent, err := os.ReadFile(itermPlistModifierLocation)
	if err != nil {
		fmt.Println("could not read the file:", itermPlistModifierLocation, err)
	}
	//var response map[string]interface{}
	var response ItermPlist
	plist.Unmarshal(fileContent, &response)
	expectedNewBookmarksCount := 1

	newBookmarksCount := len(response.NewBookmarks)

	if newBookmarksCount != 1 {
		fmt.Printf("extracted bookmakrs in not expected (wanted '%d' got %d')", expectedNewBookmarksCount, newBookmarksCount)
	}
	newBookmark := &response.NewBookmarks[0]
	newBookmark.SilenceBell = true
	b, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(b))
}
