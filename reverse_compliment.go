package bioutil

// Reverse compliment the sequence of nucleotides.
// All characters are also made caps.
func ReverseCompliment(sequence []byte) []byte {
	n := len(sequence)
	result := make([]byte, n)
	a, A, t, T, c, C, g, G := byte('a'), byte('A'), byte('t'), byte('T'), byte('c'), byte('C'), byte('g'), byte('G')
	for i, base := range sequence {
		switch base {
		case a, A:
			result[n-i-1] = T
		case t, T:
			result[n-i-1] = A
		case g, G:
			result[n-i-1] = C
		case c, C:
			result[n-i-1] = G
		default:
			result[n-i-1] = base
		}
	}
	return result
}

func Reverse(sequence []byte) []byte {
	n := len(sequence)
	result := make([]byte, n)
	for i, b := range sequence {
		result[n-i-1] = b
	}
	return result
}
