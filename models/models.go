package models

type Args struct {
	File     string `help:"file to scan"`
	Dir      string `help:"directory with files to scan"`
	Category string `help:"category of data to scan for"`
	OutDir   string `help:"name of directory to save results to"`
	Workers  int    `help:"number of threads for scanning" default:"20"`
}
