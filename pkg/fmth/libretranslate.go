package fmth

import tr "github.com/snakesel/libretranslate"

func LibretranslateToFactorioLangSwitch(input *string) {
	switch *input {
	case "es":
		input = new("es-ES")
	case "zh-Hans":
		input = new("zh-CN")
	case "zh-Hant":
		input = new("zh-TW")
	case "pt":
		input = new("pt-PT")
	case "ga":
		input = new("ga-IE")
	case "sv":
		input = new("sv-SE")
	}
}
func LibretranslateWrapper(tr *tr.Translation, sourceLang string) func(a, b string) (string, error) {
	return func(value, language string) (string, error) {
		return tr.Translate(value, sourceLang, language)
	}
}
