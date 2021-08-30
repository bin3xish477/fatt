#! /usr/bin/env bash

_fatt_completions() {
  COMPREPLY+=("--file")
  COMPREPLY+=("--url")
  COMPREPLY+=("--outfile")
  COMPREPLY+=("--workers")
  COMPREPLY+=("--nocolor")
  COMPREPLY+=("--quiet")
  COMPREPLY+=("--list")
  COMPREPLY+=("--exclude")
}

complete -f -F _fatt_completions -o filenames fatt
