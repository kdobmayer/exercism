package dna

import (
	"errors"
	"strings"
)

const validNucleotides = "ACGT"

var ErrInvalidNucleotide = errors.New("invalid nucleotide")

type Histogram map[rune]int

type DNA string

func (d DNA) Count(nucleotide byte) (int, error) {
	if !isValid(rune(nucleotide)) {
		return 0, ErrInvalidNucleotide
	}
	return strings.Count(string(d), string(nucleotide)), nil
}

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, r := range d {
		if !isValid(r) {
			return nil, ErrInvalidNucleotide
		}
		h[r]++
	}
	return h, nil
}

func isValid(nucleotide rune) bool {
	return strings.ContainsRune(validNucleotides, nucleotide)
}