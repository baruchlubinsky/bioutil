package bioutil

import(
	"sort"
	"strings"
)

const AMINO_ACIDS string = "ARNDCEQGHILKMFPSTWYV"

type Nucleotide byte
type Codon [3]Nucleotide
// May contain multiple characters, eg. iKL
type AminoAcid string
type AminoAcids []AminoAcid

func (codon *Codon) AminoAcid() AminoAcid {
	switch codon[0] {
	case 't', 'T':
		switch codon[1] {
		case 't', 'T':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return AminoAcid("F") // TTT TTC
			case 'a', 'A', 'g', 'G':
				return AminoAcid("L") // TTA TTG
			}
		case 'c', 'C':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return AminoAcid("S") // TCT TCC TCA TCG
			}
		case 'a', 'A':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return AminoAcid("Y") // TAT TAC
			case 'a', 'A', 'g', 'G':
				return AminoAcid("*") // TAA TAG
			}
		case 'g', 'G':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return AminoAcid("C") // TGT TGC
			case 'a', 'A':
				return AminoAcid("*") // TGA
			case 'g', 'G':
				return AminoAcid("W") // TGG
			}
		}
	case 'c', 'C':
		switch codon[1] {
		case 't', 'T':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return AminoAcid("L") // CTT CTC CTA CTG
			}
		case 'c', 'C':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return AminoAcid("P") // CCT CCC CCA CCG
			}
		case 'a', 'A':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return AminoAcid("H") // CAT CAC
			case 'a', 'A', 'g', 'G':
				return AminoAcid("Q") // CAA CAG
			}
		case 'g', 'G':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return AminoAcid("R") // CGT CGC CGA CGG
			}
		}
	case 'a', 'A':
		switch codon[1] {
		case 't', 'T':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A':
				return AminoAcid("I") //ATT ATC ATA
			case 'g', 'G':
				return AminoAcid("M") // ATG
			}
		case 'c', 'C':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return AminoAcid("T") // ACT ACC ACA ACG
			}
		case 'a', 'A':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return AminoAcid("N") // AAT AAC
			case 'a', 'A', 'g', 'G':
				return AminoAcid("K") // AAA AAG
			}
		case 'g', 'G':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return AminoAcid("S") // AGT AGC
			case 'a', 'A', 'g', 'G':
				return AminoAcid("R") // AGA AGG
			}
		}
	case 'g', 'G':
		switch codon[1] {
		case 't', 'T':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return AminoAcid("V") // GTT GTC GTA GTG
			}
		case 'c', 'C':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return AminoAcid("A") // GCT GCC GCA GCG
			}
		case 'a', 'A':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return AminoAcid("D") // GAT GAC
			case 'a', 'A', 'g', 'G':
				return AminoAcid("E") // GAA GAG
			}
		case 'g', 'G':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return AminoAcid("G") // GGT GGC GGA GGG
			}
		}
	}
	return AminoAcid("X")
}

func CodonsToAminoAcids(sequence []Nucleotide) AminoAcids {
	result := make(AminoAcids, len(sequence)/3)
	for i, _ := range result {
		codon := Codon{sequence[i*3], sequence[i*3+1], sequence[i*3+2]}
		result[i] = codon.AminoAcid()
	}
	return result
}

// Returns the code for an ambigious amino acid.
func AminoAcidCode(aa string) string {
	aa = strings.ToUpper(aa)
	if len(aa) == 1 {
		return aa
	}
	letters := sort.StringSlice{aa[:]}
	letters.Sort()
	sorted := strings.Join([]string(letters), "")
	switch sorted {
	case "AG":
		return "R"
	case "CT":
		return "Y"
	case "AC":
		return "M"
	case "GT":
		return "K"
	case "CG":
		return "S"
	case "AT":
		return "W"
	case "ACT":
		return "H"
	case "CGT":
		return "B"
	case "ACG":
		return "V"
	case "AGT":
		return "D"
	case "ACGT":
		return "N"
	}
	return aa
}
