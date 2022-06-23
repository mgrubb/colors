# colors

A simple utility to convert color formats

Licensed under MIT license.

## Installation

`go install github.com/mgrubb/colors@latest`

```
Usage: colors [-l] <color spec>
  -h	Provide help
  -l	Output in LaTeX Color spec format
Where <color spec> is one of:
	#hexcode       - Typical HTML Hex color code
	<R> <G> <B>    - Decimal RGB values
	0.00,0.00,0.00 - LaTeX color spec

If -l is not given, then hexcodes and LaTeX codes give decimal RGB values, and vice versa.
```

## Examples

This converts a hex color code to decimal RGB values
```shell
$ colors '#183C54'
24 60 84
```

The `#` can also be omitted
```shell
$ colors 183C54
24 60 84
```

This converts RGB values to Hex code:
```shell
$ colors 24 60 84
#183C54
```

Get a color in LaTeX format:
```shell
$ colors -l '#183C54'
0.09,0.24,0.33
```

Convert a LaTeX color to RGB:
```shell
$ colors 0.09,0.24,0.33
22 61 84
```

Convert A LaTeX color to Hex:
```shell
$ colors `colors 0.09,0.24,0.33`
#163D54
```
