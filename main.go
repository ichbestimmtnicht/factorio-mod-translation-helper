package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/ichbestimmtnicht/factorio-mod-translation-helper/pkg/fmth"
	"github.com/snakesel/libretranslate"
)

type languages []string

func (i *languages) String() string {
	return fmt.Sprintf("%v", *i)
}
func (i *languages) Set(value string) error {
	*i = append(*i, value)
	return nil
}
func main() {
	var (
		outputPath, sourceLang, sourcePath, fileName, url, key string

		langs           languages
		absolute, debug bool
	)
	flag.StringVar(&outputPath, "out", "locale/", "The base path to create files in.")
	flag.StringVar(&fileName, "fileName", "/locale.cfg", "The filename to read from and write to.")
	flag.StringVar(&sourceLang, "sourceLang", "en", "The original language to translate from.")
	flag.StringVar(&sourcePath, "sourcePath", "locale/en", "Where to find the source file.")
	flag.StringVar(&url, "url", "http://localhost:5000/", "The libretranslate URL.")
	flag.StringVar(&key, "url-key", "", "The secret for libretranslate.")
	flag.Var(&langs, "l", fmt.Sprintf("Generate for this language. If you want to generate for many add a flag for every language you want to generate for.\nDefault languages: %s", fmth.DefaultLangs))
	flag.BoolVar(&absolute, "absoluteInputPath", false, "Will the input path reader ignore adding the fileName?")
	flag.BoolVar(&debug, "debug", true, "Use debug output?")
	flag.Parse()
	if len(langs) == 0 {
		langs = append(langs, fmth.DefaultLangs...)
	}
	translater := libretranslate.New(libretranslate.Config{
		Url: url,
		Key: key,
	})
	if debug {
		translater.Debug = os.Stdout
	}
	if !absolute {
		sourcePath = path.Join(sourcePath, fileName)
	}
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		panic(err)
	}
	err = fmth.TranslateFactorioLines(outputPath, fileName, langs, strings.Split(string(data), "\n"), fmth.LibretranslateWrapper(translater, sourceLang))
	if err != nil {
		fmt.Print(err)
	}
}
