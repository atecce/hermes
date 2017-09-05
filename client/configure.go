package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/kr/pretty"
)

var userAgent = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>www</string>
	<key>Program</key>
	<string>` + filepath.Join(buildDir, "main") + `</string>
	<key>ProgramArguments</key>
	<array>
		<string>` + filepath.Join(buildDir, "main") + `</string>
		<string>-root</string>
		<string>` + buildDir + `</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
</dict>
</plist>`

func configure() error {
	pretty.Logln("[INFO] configuring...")
	pretty.Logln("[INFO] writing launchd User Agent", userAgent)
	if err := ioutil.WriteFile(userAgentPath, []byte(userAgent), 0644); err != nil {
		pretty.Logln("[ERROR]", err)
		return errors.New("failed to write launchd User Agent")
	}
	return nil
}
