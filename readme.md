# Description
A command line that can help you to translate words between languages.

# Manual

```
➜  ~ trans --help
NAME:
   trans - translation words/sentencse in your terminal

USAGE:
   trans [global options] command [command options] [arguments...]

VERSION:
   v0.1.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --target value, -t value    target language which will be translated to, list of supported language codes at https://cloud.google.com/translate/docs/languages (default: "en")
   --provider value, -p value  translator providers, e.g. google, youdao (default: "youdao")
   --brief value, -b value     return brief defination of the given word(s) (default: "true")
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```

# Installation
```
go get github.com/yulintan/trans
```

or
```
brew install yulintan/tap/trans
```

# Usage

Translate other language to English
```
$ tran 你好
```

Sentense with whitespace should be quoted
```
$ trans 'hello, world'
```

translate other language to specified target language, Polish here
```
$ trans --target=pl --provider=google hello
```

A list of language codes can be found here:

https://cloud.google.com/translate/docs/languages


