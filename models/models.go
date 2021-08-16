package models

type Args struct {
	File    string `help:"file or directory to scan" arg:"required"`
	OutFile string `help:"name of directory to save results to"`
	Workers int    `help:"number of threads for scanning" default:"20"`
	NoColor bool   `help:"turn off color output"`
}
