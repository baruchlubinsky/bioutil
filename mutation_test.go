package bioutil

import(
	"testing"
)
 
func TestMutationEquality(t *testing.T) {
	f := MutationFactory{nil, nil}
	a := f.Parse("L100I")
	b := f.Parse("L100I")
	if a != b {
		t.Fatalf("\n%#v\n%#v", a, b)
	}
}

func TestParse(t *testing.T) {
	f := MutationFactory{nil, nil}
	a := f.Parse("100I")
	if a.Position() != 100 {
		t.Fatalf(a.String())
	}
	if a.Value() != "I" {
		t.Fatalf(a.String())	
	}
}