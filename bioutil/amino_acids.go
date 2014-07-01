package bioutil

type Nucleotide byte
type Codon [3]Nucleotide
type AminoAcid byte

func (codon *Codon) AminoAcid() AminoAcid {
	switch codon[0] {
	case 't', 'T':
		switch codon[1] {
		case 't', 'T':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return 'F' // TTT TTC
			case 'a', 'A', 'g', 'G':
				return 'L' // TTA TTG
			}
		case 'c', 'C':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return 'S' // TCT TCC TCA TCG
			}
		case 'a', 'A':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return 'Y' // TAT TAC
			case 'a', 'A', 'g', 'G':
				return '*' // TAA TAG
			}
		case 'g', 'G':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return 'C' // TGT TGC
			case 'a', 'A':
				return '*' // TGA
			case 'g', 'G':
				return 'W' // TGG
			}
		}
	case 'c', 'C':
		switch codon[1] {
		case 't', 'T':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return 'L' // CTT CTC CTA CTG
			}
		case 'c', 'C':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return 'P' // CCT CCC CCA CCG
			}
		case 'a', 'A':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return 'H' // CAT CAC
			case 'a', 'A', 'g', 'G':
				return 'Q' // CAA CAG
			}
		case 'g', 'G':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return 'R' // CGT CGC CGA CGG
			}
		}
	case 'a', 'A':
		switch codon[1] {
		case 't', 'T':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A':
				return 'I' //ATT ATC ATA
			case 'g', 'G':
				return 'M' // ATG
			}
		case 'c', 'C':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return 'T' // ACT ACC ACA ACG
			}
		case 'a', 'A':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return 'N' // AAT AAC
			case 'a', 'A', 'g', 'G':
				return 'K' // AAA AAG
			}
		case 'g', 'G':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return 'S' // AGT AGC
			case 'a', 'A', 'g', 'G':
				return 'R' // AGA AGG
			}
		}
	case 'g', 'G':
		switch codon[1] {
		case 't', 'T':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return 'V' // GTT GTC GTA GTG
			}
		case 'c', 'C':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return 'A' // GCT GCC GCA GCG
			}
		case 'a', 'A':
			switch codon[2] {
			case 't', 'T', 'c', 'C':
				return 'D' // GAT GAC
			case 'a', 'A', 'g', 'G':
				return 'E' // GAA GAG
			}
		case 'g', 'G':
			switch codon[2] {
			case 't', 'T', 'c', 'C', 'a', 'A', 'g', 'G':
				return 'G' // GGT GGC GGA GGG
			}
		}
	}
	return 'X'
}

func CodonsToAminoAcids(sequence []Nucleotide) []AminoAcid {
	result := make([]AminoAcid, len(sequence)/3)
	for i, _ := range result {
		codon := Codon{sequence[i*3], sequence[i*3+1], sequence[i*3+2]}
		result[i] = codon.AminoAcid()
	}
	return result
}
