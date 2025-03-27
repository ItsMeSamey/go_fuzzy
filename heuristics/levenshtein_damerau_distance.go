package heuristics

// Calculates the Damerau-Levenshtein distance between two strings.
// Implementation from https://wikipedia.org/wiki/Damerau-Levenshtein_distance
func DamerauLevenshteinDistance(a, b string) int {
  var da [256]int // Ascii character set

	// length of the input strings
	la := len(a)
	lb := len(b)

  if (lb < 3) { return LevenshteinOSADistance(a, b) }

	// the 2D matrix
	d := make([]int, (la+2)*(lb+2))

	// maximum possible distance, used for initialization.
	maxdist := la + lb

	// Initialize the distance matrix.
	d[0] = maxdist // d[-1, -1]
	for i := 1; i <= la+1; i++ { // i from 0 to la
		d[i*(lb+2)] = maxdist // d[i, -1]
		d[i*(lb+2)+1] = i - 1 // d[i, 0]
	}
	for j := 1; j <= lb+1; j++ { // j from 0 to lb
		d[j] = maxdist // d[-1, j]
		d[j+lb+2] = j - 1// d[0, j]
	}

	// Calculate the Damerau-Levenshtein distance.
	for i := 1; i <= la; i++ {
		db := 0
		for j := 1; j <= lb; j++ {
			k := da[b[j-1]] // Go is 0-indexed, pseudocode 1-indexed
			l := db
			cost := 0
			if a[i-1] == b[j-1] { // Go is 0-indexed
				cost = 0
				db = j
			} else {
				cost = 1
			}
			substitutionCost := d[(i-1+1)*(lb+2)+(j-1+1)] + cost
			insertionCost := d[(i+1)*(lb+2)+(j-1+1)] + 1
			deletionCost := d[(i-1+1)*(lb+2)+(j+1)] + 1
			transpositionCost := 0
			if k > 0 && l > 0 {
				transpositionCost = d[(k-1+1)*(lb+2)+(l-1+1)] + (i - k - 1) + 1 + (j - l - 1)
			} else {
				transpositionCost = maxdist // Or any value > substitutionCost, insertionCost, deletionCost
			}

			d[(i+1)*(lb+2)+(j+1)] = min(substitutionCost, insertionCost, deletionCost, transpositionCost)
		}
		da[a[i-1]] = i
	}

	return d[(la+1)*(lb+2)+(lb+1)] // d[la, lb]
}

