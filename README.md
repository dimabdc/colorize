# colorize
Golang. Coloring text for logging and cli application

## Installation
    $ go get github.com/dimabdc/colorize
    
    
## Usage example
    fmt.Println("print ", colorize.Paint("Blue", colorize.Blue|colorize.Bold|colorize.Underline, colorize.Black))