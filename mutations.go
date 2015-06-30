package bioutil

import (
	"fmt"
	"strings"
)

// Represents a single mutation within Pol. Instances are immutable.
type Mutation struct {
	sequence *AminoAcids
	// The position in the reference sequence
	position int
	geneRegion         *GeneRegion
	// The position in the specific gene region
	geneRegionPosition int
	value        AminoAcid
	wildType     AminoAcid
}

type Mutations []Mutation

// Get the position within its gene region.
func (m Mutation) GeneRegionPosition() int {
	return m.geneRegionPosition
}

// Get the position, starting from 1.
func (m Mutation) Position() int {
	return m.position
}

// The wild type for this mutation.
// If it has not been explicitly set, the wildtype from the reference sequence is returned.
// If position is 0, throws a runtime index out-of-range error.
func (m *Mutation) WildType() AminoAcid {
	if m.wildType == "" && m.sequence != nil {
		index := m.Position() - 1
		m.wildType = AminoAcid((*m.sequence)[index])
	}
	return m.wildType
}

func (m Mutation) Value() AminoAcid {
	return m.value
}

// Returns a new Mutation object, having the same position, wild type and the specified value.
func (m Mutation) WithValue(v AminoAcid) Mutation {
	return Mutation{
		sequence: m.sequence,
		geneRegion: m.geneRegion,
		geneRegionPosition: m.geneRegionPosition,
		wildType: m.wildType,
		position: m.position,
		value:    v,
	}
}

func (m Mutation) Region() GeneRegion {
	return *m.geneRegion
}

// Mutations are equal if the have the same position and value.
func (m Mutation) Equals(b Mutation) bool {
	return m.Position() == b.Position() && m.value == b.value
}

// String representation of a Mutation with its position in the gene region.
func (m Mutation) RegionString() string {
	return fmt.Sprintf("%v%v%v", m.WildType(), m.GeneRegionPosition(), m.Value())
}

// String representation of a Mutation.
func (m Mutation) String() string {
	return fmt.Sprintf("%v%v%v", m.WildType(), m.Position(), m.Value())
}

func (a Mutations) RegionString() string {
	s := make([]string, len(a))
	for i, m := range a {
		s[i] = m.RegionString()
	}
	return strings.Join(s, ", ")
}


func (a Mutations) String() string {
	s := make([]string, len(a))
	for i, m := range a {
		s[i] = m.String()
	}
	return strings.Join(s, ", ")
}

// Sort interface
func (list Mutations) Len() int { return len(list) }

// Sort interface
func (list Mutations) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

// Sort interface. Order of position then value.
func (list Mutations) Less(i, j int) bool {
	if list[i].Position() != list[j].Position() {
		return list[i].Position() < list[j].Position()
	}
	return list[i].value < list[j].value
}
