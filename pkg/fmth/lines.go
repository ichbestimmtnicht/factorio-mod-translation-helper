package fmth

import (
	"fmt"
	"os"
	"os/exec"
)

// example: path: "locale/", name: "/example.cfg", Lang: ["en"] -> {working_dir}/locale/en/example.cfg
//
// createDir tries silently to create path+lang with "mkdir -p {working_dir}/locale/en/""
func TranslateFactorioLines(path, name string, langs, lines []string, tr func(value, language string) (string, error)) (err error) {
	for _, l := range langs {
		ll := l
		LibretranslateToFactorioLangSwitch(&ll)
		fmt.Printf("Language: %s\n", ll)
		if createDir {
			cmd := exec.Command("mkdir", "-p", path+ll)
			cmd.Run()
		}
		f, err := os.Create(path + ll + name)
		if err != nil {
			return err
		}
		defer f.Close()
		for _, line := range lines {
			newLine, err := translateFactorioLine(l, line, tr)
			if err != nil {
				if err.Error() == "Skip" {
					continue
				}
				return err
			}
			_, err = f.WriteString(newLine + "\n")
			if err != nil {
				return err
			}
		} // finish range lines.
	} // finish range languages.
	return
}
