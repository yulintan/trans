# Description
A command line that can help you to translate words between languages.

# Manual
<img src="https://raw.githubusercontent.com/yulintan/translate/master/manual.png"  width="800"/>

# Installation
```
go get github.com/yulintan/translate
```

or
```
brew install yulintan/tap/translate
```

# Usage

for easy usage, in my bash or zsh, I

`alias t=translate`


Translate other language to English
```
$ translate 你好
```

translate other language to specified target language, Polish here
```
$ translate -t=pl hello
```

A list of language codes can be found here:

https://cloud.google.com/translate/docs/languages


