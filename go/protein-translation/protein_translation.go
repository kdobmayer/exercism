package protein

func FromCodon(codon string) string {
	switch codon {
	case "AUG":
		return "Methionine"
	case "UUU", "UUC":
		return "Phenylalanine"
	case "UUA", "UUG":
		return "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine"
	case "UAU", "UAC":
		return "Tyrosine"
	case "UGU", "UGC":
		return "Cysteine"
	case "UGG":
		return "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "STOP"
	default:
		return "invalid codon"
	}
}

func FromRNA(rna string) []string {
	var proteins []string
	for i := 0; i < len(rna); i += 3 {
		protein := FromCodon(rna[i : i+3])
		if protein == "STOP" {
			break
		}
		proteins = append(proteins, protein)
	}
	return proteins
}