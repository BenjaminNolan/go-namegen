# Team Name Generator for Go

This package generates random names for teams and similar things such as Xenial Giraffe, Happy Rabbit, etc.

## Why? There's, like, thirty already!

Yeah, but they keep generating things like `fat-cow` and `obese-cat`, which we didn't really want when generating random team names for [Comnoco](https://comnoco.com)!

## Install

```sh
go get github.com/BenjaminNolan/go-namegen
```

## Usage

```go
package main

import namegen "github.com/BenjaminNolan/namegen"

func main() {
	generator := namegen.New()

	// Set to false to generate lower-case names
	generator.Capitalize = false

	// If you want to generate alliterative names like Kind Kiwi, Amazing Aardvark, Lucky Lemur, etc, set this to true
	generator.Alliterate = true

	// If you want dashes or underscores instead of spaces, set Separator to your desired character(s)
	generator.Separator = "-"

	println(generator.Generate())
}
```

## Adding words

If you want to add words or functionality here, feel free to whang over a PR. :)

## License

This is MIT licensed, see [LICENSE](LICENSE) for more info.
