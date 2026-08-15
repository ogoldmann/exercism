package rnatranscription

var nucleotidePairs = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

func ToRNA(dna string) string {
	rnai := ""
	for _, n := range dna {
		rnai += nucleotidePairs[string(n)]
	}
	return rnai
}
