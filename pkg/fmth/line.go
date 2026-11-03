package fmth

import (
	"errors"
	"strings"
)

func translateFactorioLine(lang, line string, tr func(value, language string) (string, error)) (newLine string, err error) {
	switch {
	default:
		kv := strings.Split(line, "=")
		v := strings.Join(kv[1:], "=")
		var newVal string
		if strings.Contains(v, "[") {
			var toTranslate string
			for _, c := range v {
				char := string(c)
				switch char {
				case "[":
					if toTranslate != "" {
						var n string
						n, err = tr(toTranslate, lang)
						if err != nil {
							return
						}
						newVal += n
						toTranslate = ""
					}
				case "]":
					newVal += " [" + toTranslate + "] "
					toTranslate = ""
				default:
					toTranslate += char
				}
			} // end loop over runes.
			if toTranslate != "" {
				var n string
				n, err = tr(toTranslate, lang)
				if err != nil {
					return
				}
				newVal += n
			}
		} else {
			newVal, err = tr(v, lang)
			if err != nil {
				return
			}
		}
		newLine = kv[0] + "=" + newVal
	case strings.HasPrefix(line, "["):
		newLine = line
	case line == "":
		return "", errors.New("Skip")
	} // end switch
	return
}
