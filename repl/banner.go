package repl

import (
	"fmt"
	"io"
)

// Gold gradient, top -> bottom: warm amber down to deep gold.
var berriesGold = [6][3]int{
	{255, 224, 130}, {252, 204, 92}, {244, 184, 60},
	{230, 162, 40}, {212, 143, 26}, {194, 126, 18},
}

var berriesArt = [6]string{
	` _                     _           `,
	`| |                   (_)          `,
	`| |__   ___ _ __ _ __  _  ___  ___ `,
	`| '_ \ / _ \ '__| '__|| |/ _ \/ __|`,
	`| |_) |  __/ |  | |   | |  __/\__ \`,
	`|_.__/ \___|_|  |_|   |_|\___||___/`,
}

const (
	ansiReset = "\x1b[0m"
	subtitle  = "\x1b[38;2;170;138;92m"   // muted bronze
	hintDim   = "\x1b[2;37m"              // dim gray
	promptCol = "\x1b[1;38;2;244;184;60m" // bold gold
)

func fg(r, g, b int) string { return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b) }

// PrintBanner writes the colored berries REPL header to w.
func PrintBanner(w io.Writer) {
	fmt.Fprintln(w)
	for i, line := range berriesArt {
		c := berriesGold[i]
		fmt.Fprintf(w, "%s%s%s\n", fg(c[0], c[1], c[2]), line, ansiReset)
	}
	fmt.Fprintf(w, "%sFinancial Runtime REPL | 2026%s\n", subtitle, ansiReset)
	fmt.Fprintf(w, "%sType 'help' for docs, 'exit' to quit.%s\n\n", hintDim, ansiReset)
}

// Prompt returns the colored prompt for the read loop.
func Prompt() string { return promptCol + "b> " + ansiReset }
