package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
)

type args struct {
	help        bool
	interactive bool
	xtermNumber bool
	args        []string
}

const HELP = `Find the closest xterm 256 color of a hex value.
SYNTAX:
    hexto256 [OPTIONS] hex ...

OPTIONS:
    -h/--help               help
    -i/--interactive        interactive mode
    -x/--xterm-number       return the xterm-number instead of the hex value`

var colors = map[Color]uint8{
	*parseColor("000000"): 16, *parseColor("00005f"): 17,
	*parseColor("000087"): 18, *parseColor("0000af"): 19,
	*parseColor("0000d7"): 20, *parseColor("0000ff"): 21,
	*parseColor("005f00"): 22, *parseColor("005f5f"): 23,
	*parseColor("005f87"): 24, *parseColor("005faf"): 25,
	*parseColor("005fd7"): 26, *parseColor("005fff"): 27,
	*parseColor("008700"): 28, *parseColor("00875f"): 29,
	*parseColor("008787"): 30, *parseColor("0087af"): 31,
	*parseColor("0087d7"): 32, *parseColor("0087ff"): 33,
	*parseColor("00af00"): 34, *parseColor("00af5f"): 35,
	*parseColor("00af87"): 36, *parseColor("00afaf"): 37,
	*parseColor("00afd7"): 38, *parseColor("00afff"): 39,
	*parseColor("00d700"): 40, *parseColor("00d75f"): 41,
	*parseColor("00d787"): 42, *parseColor("00d7af"): 43,
	*parseColor("00d7d7"): 44, *parseColor("00d7ff"): 45,
	*parseColor("00ff00"): 46, *parseColor("00ff5f"): 47,
	*parseColor("00ff87"): 48, *parseColor("00ffaf"): 49,
	*parseColor("00ffd7"): 50, *parseColor("00ffff"): 51,
	*parseColor("5f0000"): 52, *parseColor("5f005f"): 53,
	*parseColor("5f0087"): 54, *parseColor("5f00af"): 55,
	*parseColor("5f00d7"): 56, *parseColor("5f00ff"): 57,
	*parseColor("5f5f00"): 58, *parseColor("5f5f5f"): 59,
	*parseColor("5f5f87"): 60, *parseColor("5f5faf"): 61,
	*parseColor("5f5fd7"): 62, *parseColor("5f5fff"): 63,
	*parseColor("5f8700"): 64, *parseColor("5f875f"): 65,
	*parseColor("5f8787"): 66, *parseColor("5f87af"): 67,
	*parseColor("5f87d7"): 68, *parseColor("5f87ff"): 69,
	*parseColor("5faf00"): 70, *parseColor("5faf5f"): 71,
	*parseColor("5faf87"): 72, *parseColor("5fafaf"): 73,
	*parseColor("5fafd7"): 74, *parseColor("5fafff"): 75,
	*parseColor("5fd700"): 76, *parseColor("5fd75f"): 77,
	*parseColor("5fd787"): 78, *parseColor("5fd7af"): 79,
	*parseColor("5fd7d7"): 80, *parseColor("5fd7ff"): 81,
	*parseColor("5fff00"): 82, *parseColor("5fff5f"): 83,
	*parseColor("5fff87"): 84, *parseColor("5fffaf"): 85,
	*parseColor("5fffd7"): 86, *parseColor("5fffff"): 87,
	*parseColor("870000"): 88, *parseColor("87005f"): 89,
	*parseColor("870087"): 90, *parseColor("8700af"): 91,
	*parseColor("8700d7"): 92, *parseColor("8700ff"): 93,
	*parseColor("875f00"): 94, *parseColor("875f5f"): 95,
	*parseColor("875f87"): 96, *parseColor("875faf"): 97,
	*parseColor("875fd7"): 98, *parseColor("875fff"): 99,
	*parseColor("878700"): 100, *parseColor("87875f"): 101,
	*parseColor("878787"): 102, *parseColor("8787af"): 103,
	*parseColor("8787d7"): 104, *parseColor("8787ff"): 105,
	*parseColor("87af00"): 106, *parseColor("87af5f"): 107,
	*parseColor("87af87"): 108, *parseColor("87afaf"): 109,
	*parseColor("87afd7"): 110, *parseColor("87afff"): 111,
	*parseColor("87d700"): 112, *parseColor("87d75f"): 113,
	*parseColor("87d787"): 114, *parseColor("87d7af"): 115,
	*parseColor("87d7d7"): 116, *parseColor("87d7ff"): 117,
	*parseColor("87ff00"): 118, *parseColor("87ff5f"): 119,
	*parseColor("87ff87"): 120, *parseColor("87ffaf"): 121,
	*parseColor("87ffd7"): 122, *parseColor("87ffff"): 123,
	*parseColor("af0000"): 124, *parseColor("af005f"): 125,
	*parseColor("af0087"): 126, *parseColor("af00af"): 127,
	*parseColor("af00d7"): 128, *parseColor("af00ff"): 129,
	*parseColor("af5f00"): 130, *parseColor("af5f5f"): 131,
	*parseColor("af5f87"): 132, *parseColor("af5faf"): 133,
	*parseColor("af5fd7"): 134, *parseColor("af5fff"): 135,
	*parseColor("af8700"): 136, *parseColor("af875f"): 137,
	*parseColor("af8787"): 138, *parseColor("af87af"): 139,
	*parseColor("af87d7"): 140, *parseColor("af87ff"): 141,
	*parseColor("afaf00"): 142, *parseColor("afaf5f"): 143,
	*parseColor("afaf87"): 144, *parseColor("afafaf"): 145,
	*parseColor("afafd7"): 146, *parseColor("afafff"): 147,
	*parseColor("afd700"): 148, *parseColor("afd75f"): 149,
	*parseColor("afd787"): 150, *parseColor("afd7af"): 151,
	*parseColor("afd7d7"): 152, *parseColor("afd7ff"): 153,
	*parseColor("afff00"): 154, *parseColor("afff5f"): 155,
	*parseColor("afff87"): 156, *parseColor("afffaf"): 157,
	*parseColor("afffd7"): 158, *parseColor("afffff"): 159,
	*parseColor("d70000"): 160, *parseColor("d7005f"): 161,
	*parseColor("d70087"): 162, *parseColor("d700af"): 163,
	*parseColor("d700d7"): 164, *parseColor("d700ff"): 165,
	*parseColor("d75f00"): 166, *parseColor("d75f5f"): 167,
	*parseColor("d75f87"): 168, *parseColor("d75faf"): 169,
	*parseColor("d75fd7"): 170, *parseColor("d75fff"): 171,
	*parseColor("d78700"): 172, *parseColor("d7875f"): 173,
	*parseColor("d78787"): 174, *parseColor("d787af"): 175,
	*parseColor("d787d7"): 176, *parseColor("d787ff"): 177,
	*parseColor("d7af00"): 178, *parseColor("d7af5f"): 179,
	*parseColor("d7af87"): 180, *parseColor("d7afaf"): 181,
	*parseColor("d7afd7"): 182, *parseColor("d7afff"): 183,
	*parseColor("d7d700"): 184, *parseColor("d7d75f"): 185,
	*parseColor("d7d787"): 186, *parseColor("d7d7af"): 187,
	*parseColor("d7d7d7"): 188, *parseColor("d7d7ff"): 189,
	*parseColor("d7ff00"): 190, *parseColor("d7ff5f"): 191,
	*parseColor("d7ff87"): 192, *parseColor("d7ffaf"): 193,
	*parseColor("d7ffd7"): 194, *parseColor("d7ffff"): 195,
	*parseColor("ff0000"): 196, *parseColor("ff005f"): 197,
	*parseColor("ff0087"): 198, *parseColor("ff00af"): 199,
	*parseColor("ff00d7"): 200, *parseColor("ff00ff"): 201,
	*parseColor("ff5f00"): 202, *parseColor("ff5f5f"): 203,
	*parseColor("ff5f87"): 204, *parseColor("ff5faf"): 205,
	*parseColor("ff5fd7"): 206, *parseColor("ff5fff"): 207,
	*parseColor("ff8700"): 208, *parseColor("ff875f"): 209,
	*parseColor("ff8787"): 210, *parseColor("ff87af"): 211,
	*parseColor("ff87d7"): 212, *parseColor("ff87ff"): 213,
	*parseColor("ffaf00"): 214, *parseColor("ffaf5f"): 215,
	*parseColor("ffaf87"): 216, *parseColor("ffafaf"): 217,
	*parseColor("ffafd7"): 218, *parseColor("ffafff"): 219,
	*parseColor("ffd700"): 220, *parseColor("ffd75f"): 221,
	*parseColor("ffd787"): 222, *parseColor("ffd7af"): 223,
	*parseColor("ffd7d7"): 224, *parseColor("ffd7ff"): 225,
	*parseColor("ffff00"): 226, *parseColor("ffff5f"): 227,
	*parseColor("ffff87"): 228, *parseColor("ffffaf"): 229,
	*parseColor("ffffd7"): 230, *parseColor("ffffff"): 231,
	*parseColor("080808"): 232, *parseColor("121212"): 233,
	*parseColor("1c1c1c"): 234, *parseColor("262626"): 235,
	*parseColor("303030"): 236, *parseColor("3a3a3a"): 237,
	*parseColor("444444"): 238, *parseColor("4e4e4e"): 239,
	*parseColor("585858"): 240, *parseColor("626262"): 241,
	*parseColor("6c6c6c"): 242, *parseColor("767676"): 243,
	*parseColor("808080"): 244, *parseColor("8a8a8a"): 245,
	*parseColor("949494"): 246, *parseColor("9e9e9e"): 247,
	*parseColor("a8a8a8"): 248, *parseColor("b2b2b2"): 249,
	*parseColor("bcbcbc"): 250, *parseColor("c6c6c6"): 251,
	*parseColor("d0d0d0"): 252, *parseColor("dadada"): 253,
	*parseColor("e4e4e4"): 254, *parseColor("eeeeee"): 255,
}


func closest256Color(c *Color) (r Color) {
	var min int64 = math.MaxInt64
	for color := range colors {
		d := c.distance(&color)
		if d < min {
			min = d
			r = color
		}
	}
	return
}


func resetBoolFlag(s string, flag *bool) {
	s = strings.ToLower(s)
	if (s == "1" || s == "t" || s == "true") && !*flag {
		*flag = true
	}
}


func parseArgs() (a *args) {
	a = &args{false, false, false, []string{}}
	flag.BoolVar(&a.help, "h", false, "print help")
	flag.BoolFunc("help", "print help", func(s string) error {
		resetBoolFlag(s, &a.help)
		return nil
	})
	flag.BoolVar(&a.interactive, "i", false, "interactive mode")
	flag.BoolFunc("interactive", "interactive mode", func(s string) error {
		resetBoolFlag(s, &a.interactive)
		return nil
	})
	flag.BoolVar(&a.xtermNumber, "x", false,
		"return the xterm-number instead of the hex value")
	flag.BoolFunc("xterm-number", "return the xterm-number instead of the hex value",
		func(s string) error {
			resetBoolFlag(s, &a.xtermNumber)
			return nil
		})
	flag.Parse()
	a.args = flag.Args()
	return
}


func formatXTerm(c *Color) (s string) {
    return fmt.Sprintf("%d \033[38;2;%d;%d;%dm \033[m", colors[*c],
        c.red(), c.green(), c.blue())
}


func formatHex(c *Color) (s string) {
    return fmt.Sprintf("%s \033[38;2;%d;%d;%dm \033[m", c.toString(), 
            c.red(), c.green(), c.blue())
}


func interactive(a *args) {
    for {
        var input string
        fmt.Print("Insert a hex value (empty to exit): ")
        fmt.Scanf("%s", &input)
        if input == "" {
            break
        }
        closest := closest256Color(parseColor(input))
        if a.xtermNumber {
            fmt.Println(formatXTerm(&closest))
        } else {
            fmt.Println(formatHex(&closest))
        }
    }
}


func nonInteractive(a *args) {
	for _, hex := range a.args {
		closest := closest256Color(parseColor(hex))
		if a.xtermNumber {
			fmt.Println(formatXTerm(&closest))
		} else {
			fmt.Println(formatHex(&closest))
		}
	}
}


func main() {
	a := parseArgs()

	if a.help {
		fmt.Println(HELP)
		return
	}

	if a.interactive {
        interactive(a)
        return
	}

    nonInteractive(a)
	return
}
