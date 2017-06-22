package colorize

import "strings"
import "strconv"

const (
	Black     int = 0
	Red       int = 1
	Green     int = 2
	Yellow    int = 3
	Blue      int = 4
	Magenta   int = 5
	Cyan      int = 6
	White     int = 7
	Bright    int = 16
	Bold      int = 32
	Blink     int = 64
	Underline int = 128
	Italic    int = 256
	Negative  int = 512
)

func Paint(str string, options ...int) string {
	return "" + start + parseColor(options...) + str + finish
}

const (
	start     = "\033["
	finish    = "\033[0m"
	fgNormal  = 30
	bgNormal  = 40
	bright    = 60
	bold      = "1"
	italic    = "3"
	underline = "4"
	blink     = "5"
	negative  = "7"
)

func parseColor(options ...int) string {
	var params []string
	var color int
	
	if options[0] != 0 {
		color = getColor(options[0]) + getIntense(options[0]) + fgNormal
		params = append(params, strconv.Itoa(color))
		params = append(params, getEffects(options[0])...)
	}
	
	if options[1] != 0 {
		color = getColor(options[1]) + getIntense(options[1]) + bgNormal
		params = append(params, strconv.Itoa(color))
	}
	
	return strings.Join(params, ";") + "m"
}

func getColor(param int) int {
	return param & 7
}

func getIntense(param int) int {
	var intense int = 0
	
	if param&Bright != 0 {
		intense = bright
	}
	return intense
}

func getEffects(param int) []string {
	var effects []string
	
	if param&Bold != 0 {
		effects = append(effects, bold)
	}
	
	if param&Blink != 0 {
		effects = append(effects, blink)
	}
	
	if param&Italic != 0 {
		effects = append(effects, italic)
	}
	
	if param&Underline != 0 {
		effects = append(effects, underline)
	}
	
	if param&Negative != 0 {
		effects = append(effects, negative)
	}
	
	return effects
}
