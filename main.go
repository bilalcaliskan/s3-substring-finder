/*
Copyright © 2022 bilalcaliskan bilalcaliskan@protonmail.com

*/
package main

import (
	"io/ioutil"
	"os"
	"s3-substring-finder/cmd"
	"strings"

	"github.com/dimiro1/banner"
)

func main() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
	cmd.Execute()
}
