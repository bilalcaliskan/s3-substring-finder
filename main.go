/*
Copyright © 2022 bilalcaliskan bilalcaliskan@protonmail.com

*/
package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/bilalcaliskan/s3-substring-finder/cmd"
	"github.com/dimiro1/banner"
)

func main() {
	if _, err := os.Stat("build/ci/banner.txt"); err == nil {
		bannerBytes, _ := ioutil.ReadFile("build/ci/banner.txt")
		banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
	}

	cmd.Execute()
}
