# Factorio Mod Translation Helper

## Usage

### Prepare a local libretranslate instance with docker

```bash
git clone https://github.com/LibreTranslate/LibreTranslate
cd LibreTranslate
docker compose up -d # starts a local libretranslate api on port 5000
```

### Install

`go install github.com/ichbestimmtnicht/factorio-mod-translation-helper@main`

### Execution

Within the Factorio mod folder:

```bash
factorio-mod-translation-helper
```

Output of `-h`

```txt
  -absoluteInputPath
        Will the input path reader ignore adding the fileName?
  -debug
        Use debug output? (default true)
  -fileName string
        The filename to read from and write to. (default "/locale.cfg")
  -l value
        Generate for this language. If you want to generate for many add a flag for every language you want to generate for.
        Default languages: [ar bg ca cs da de el eo es et eu fa fi fr ga he hu id it ja ko lv lt nl pl pt pt-BR ro ru sk sl sq sv th tr uk zh-Hans zh-Hant]
  -out string
        The base path to create files in. (default "locale/")
  -sourceLang string
        The original language to translate from. (default "en")
  -sourcePath string
        Where to find the source file. (default "locale/en")
  -url string
        The libretranslate URL. (default "http://localhost:5000")
  -url-key string
        The secret for libretranslate.
```
