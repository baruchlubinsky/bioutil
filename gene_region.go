package bioutil

import "fmt"

// A named region with a gene. Positions are inclusive.
type GeneRegion struct {
	startPosition int
	endPosition int
	shortName string
	fullName string
}

func NewGeneRegion(startPosition, endPosition int, fullName, shortName string) GeneRegion {
	return GeneRegion{startPosition, endPosition, shortName, fullName}
}

// The name of this region in the format <fullNamee> (<shortName>)
func (region GeneRegion) String() string {
	return fmt.Sprintf("%v (%v)", region.fullName, region.shortName)
}