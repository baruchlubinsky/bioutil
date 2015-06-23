package bioutil

import(
	"regexp"
	"strconv"
)

var MutationRegex, MutationListRegex regexp.Regexp

func init() {
	MutationRegex = *regexp.MustCompile("([" + AMINO_ACIDS + "]?)([0-9]+)([i" + AMINO_ACIDS + "]+|[" + AMINO_ACIDS + "]+|d)")
	MutationListRegex = *regexp.MustCompile("([" + AMINO_ACIDS + "]?)([0-9]+)([i" + AMINO_ACIDS + "]+|[" + AMINO_ACIDS + "]+|d)")
}

type MutationFactory struct {
	Sequence AminoAcids
	Regions []*GeneRegion
}

func (factory *MutationFactory) RegionFor(position int) *GeneRegion {
	for i, region := range factory.Regions {
		if position >= region.startPosition && position <= region.endPosition {
			return factory.Regions[i]
		}
	}
	return nil
}

// Create a Mutation at the specified position. If the position is within a region for this gene, that is assigned to the Mutation.
func (factory *MutationFactory) Create(position int) Mutation {
	var res Mutation
	if factory.Sequence == nil {
		res = Mutation{position: position}
	} else {
		res = Mutation {
			sequence: &factory.Sequence, 
			position: position,
		}
	}
	region := factory.RegionFor(position)
	if region != nil {
		res.geneRegionPosition = position - region.startPosition
		res.geneRegion = region
	} 
	res.WildType()
	return res
}

// Create a Mutation at the specified position. If the position is within a region for this gene, that is assigned to the Mutation.
func (factory *MutationFactory) CreateWithValue(position int, value AminoAcid) Mutation {
	res := Mutation {
			sequence: &factory.Sequence, 
			position: position,
			value: value,
		}
	region := factory.RegionFor(position)
	if region != nil {
		res.geneRegionPosition = position - region.startPosition
		res.geneRegion = region
	} 
	res.WildType()
	return res
}

func (factory *MutationFactory) CreateWithValueInRegion(position int, value AminoAcid, region *GeneRegion) Mutation {
	res := factory.CreateInRegion(position, region)
	res.WildType()
	return res.WithValue(value)
}

func (factory *MutationFactory) CreateInRegion(position int, region *GeneRegion) Mutation {
	res := Mutation{
		sequence: &factory.Sequence,
		geneRegionPosition: position,
		position: position + region.startPosition,
		geneRegion: region,
	}
	res.WildType()
	return res
}

// Convert a string (such as A287K) to a Mutation. This is the inverse of Mutation.String().
func (factory *MutationFactory) Parse(s string) Mutation {
	matches := MutationRegex.FindAllSubmatch([]byte(s), 1)
	var res Mutation
	if matches != nil {
		fields := len(matches[0])
		position, _ := strconv.Atoi(string(matches[0][fields-2]))
		res = factory.Create(position)
		res.value = AminoAcid(matches[0][fields-1])
		if fields == 4 {
			res.wildType = AminoAcid(matches[0][fields-3])
		} else {
			res.WildType()
		}
	}
	return res
}

// Convert a string to a slice of Mutations.
// A100LMN becomes {A100L, A100M, A100N}
// A100AiLP becomes {A100A, A100iLP}
// For comma separated data use dnaio.FastmAlignment
func (factory *MutationFactory) ParseList(s string) Mutations {
	matches := MutationListRegex.FindAllSubmatch([]byte(s), -1)
	if matches != nil {
		res := make(Mutations, 0, len(matches))
		for _, match := range matches {
			wildType := AminoAcid(match[1])
			position, _ := strconv.Atoi(string(match[2]))
			for i, m := range match[3] {
				if m == 'i' {
					newMutation := factory.Create(position)
					newMutation.value = AminoAcid(match[3][i:])  
					newMutation.wildType = wildType
					res = append(res, newMutation)
					break
				}
				if m != 'X' {
					newMutation := factory.Create(position)
					newMutation.value = AminoAcid(m)
					newMutation.wildType = wildType
					res = append(res, newMutation)
				}
			}
		}
		return res
	}
	return nil
}

// Create a Mutation within a specified gene region. "K65R", RT
func (factory *MutationFactory) ParseInRegion(s string, region *GeneRegion) Mutation {
	matches := MutationRegex.FindAllSubmatch([]byte(s), 1)
	var res Mutation
	if matches != nil {
		fields := len(matches[0])
		regionPos, _ := strconv.Atoi(string(matches[0][fields-2]))
		res = factory.CreateInRegion(regionPos, region)
		res.value = AminoAcid(matches[0][fields-1])
		if fields == 4 {
			res.wildType = AminoAcid(matches[0][fields-3])
		} else {
			res.WildType()
		}
	}
	return res
}

// NewMutautionList in the specified region.
func (factory *MutationFactory) ParseListInRegion(s string, region *GeneRegion) Mutations {
	matches := MutationListRegex.FindAllSubmatch([]byte(s), -1)
	if matches != nil {
		res := make(Mutations, 0, len(matches))
		for _, match := range matches {
			wildType := AminoAcid(match[1])
			position, _ := strconv.Atoi(string(match[2]))
			for i, m := range match[3] {
				if m == 'i' {
					newMutation := factory.CreateInRegion(position, region)
					newMutation.wildType = wildType
					newMutation.value = AminoAcid(match[3][i:])
					newMutation.WildType()
					res = append(res, newMutation)
					break
				}
				if m != 'X' {
					newMutation := factory.CreateInRegion(position, region)
					newMutation.wildType = wildType
					newMutation.value = AminoAcid(m)
					newMutation.WildType()
					res = append(res, newMutation)
				}
			}
		}
		return res
	}
	return nil
}