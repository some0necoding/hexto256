package main

import (
	"fmt"
	. "regexp"
	"strconv"
	"strings"
)

const HEX_PATTERN = "^(#|0x)?[0-9a-fA-F]{6}$"

/* RGB Color */
type Color [3]uint8

/* Parse a string representing a hex value into a Color object */
func parseColor(hex string) (c *Color) {
	if match, err := MatchString(HEX_PATTERN, hex); err != nil || !match {
		return nil
	}

	if strings.HasPrefix(hex, "#") {
		hex = strings.Replace(hex, "#", "", 1)
	}
	if strings.HasPrefix(hex, "0x") {
		hex = strings.Replace(hex, "0x", "", 1)
	}

	primary, _ := strconv.ParseUint(hex[:2], 16, 8)
	red := uint8(primary)
	primary, _ = strconv.ParseUint(hex[2:4], 16, 8)
	green := uint8(primary)
	primary, _ = strconv.ParseUint(hex[4:], 16, 8)
	blue := uint8(primary)
	return &Color{red, green, blue}
}

/* Get red component */
func (c *Color) red() uint8 {
	return c[0]
}

/* Get green component */
func (c *Color) green() uint8 {
	return c[1]
}

/* Get blue component */
func (c *Color) blue() uint8 {
	return c[2]
}

/* Compute distance between two colors */
func (c1 *Color) distance(c2 *Color) (d int64) {
	diffRed := int64(c1.red()) - int64(c2.red())
	diffGreen := int64(c1.green()) - int64(c2.green())
	diffBlue := int64(c1.blue()) - int64(c2.blue())
	return (diffRed * diffRed) + (diffGreen * diffGreen) +
		(diffBlue * diffBlue)
}

func (c *Color) toString() string {
	red := strconv.FormatUint(uint64(c.red()), 16)
	green := strconv.FormatUint(uint64(c.green()), 16)
	blue := strconv.FormatUint(uint64(c.blue()), 16)
	return fmt.Sprintf("#%02s%02s%02s", red, green, blue)
}
