# Распаковка строки

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* `a4bc2d5e` => `aaaabccddddde`
* `abcd` => `abcd`
* `45` => "" (некорректная строка)
* `qwe\4\5` => `qwe45`
* `qwe\45` => `qwe44444`
* `qwe\\5` => `qwe\\\\\`
