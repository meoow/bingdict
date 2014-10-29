package main

import "os"

//import "regexp"
import "strings"
import "net/url"

func popclip() string {
	var popcliptext string
	if i := os.ExpandEnv("$POPCLIP_TEXT"); i != "" {
		i = strings.Trim(i, " ,.")
		i = strings.Replace(i, "- ","",-1)
		popcliptext = url.QueryEscape(i)
	}
	return popcliptext
}

func commandline() string {
	return url.QueryEscape(strings.Join(os.Args[1:], " "))
}

func getText() string {
	var text string
	if len(os.Args) > 1 {
		text = commandline()
	} else {
		text = popclip()
	}
	return text
}
