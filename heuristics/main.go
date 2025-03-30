package heuristics

import (
  "fuzzy/common"
  "fuzzy/heuristics/algorithms"
)

// Uses MultiSet, Calculates the Dice-Sorensen coefficient for the given strings.
// This Does not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
//
// The Dice-Sorensen coefficient measures similarity between two strings using the following formula:
// Dice-Sorensen Coefficient = 2 * IntersectionCount(a, b) / (len(a) + len(b))
func DiceSorensenCoefficient[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.DiceSorensenCoefficientCharacter[F](a, b)
}

// Uses Bigram Set, Calculates the Dice-Sorensen coefficient for the given strings.
// This Does not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 8kb
//
// The Dice-Sorensen coefficient measures similarity between two strings using the following formula:
// Dice-Sorensen Coefficient = 2 * IntersectionCount(a, b) / (len(a) + len(b))
func DiceSorensenCoefficientBigram[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.DiceSorensenCoefficientBigram[F](a, b)
}

// A similarity measure that i made up
//
// Time complexity: O(n+m) = m + 2*n + 256*(log2(max(m, n)))
// Space complexity: O(n+m) = (m + n) * sizeof(uint32)
func FrequencySimilarity[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return 1 - algorithms.FrequencyDistance[F](a, b)
}

// Uses MultiSet, Calculates the Jaccard coefficient for the given strings.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
//
// The Jaccard coefficient measures similarity between two strings using the following formula:
// Jaccard Coefficient = IntersectionCount(a, b) / (len(a) + len(b) - IntersectionCount(a, b))
func JaccardCoefficient[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.JaccardCoefficientCharacter[F](a, b)
}

// Uses Bigram Set, Calculates the Jaccard coefficient for the given strings.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 8kb
//
// The Jaccard coefficient measures similarity between two strings using the following formula:
// Jaccard Coefficient = IntersectionCount(a, b) / (len(a) + len(b) - IntersectionCount(a, b))
func JaccardCoefficientBigram[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.JaccardCoefficientBigram[F](a, b)
}

// JaroSimilarity calculates the similarity between two strings using Jaro distance.
//
// Time Complexity: O(n*m)
// Space Complexity: O(n+m)
//
// jaro_distance = 1/3 * (m/|s1| + m/|s2| + (m - t)/m)
func JaroSimilarity[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.JaroDistance[F](a, b)
}

// Returns a function to Calculates the Jaro-Winkler distance between two strings by
// giving more favorable ratings to strings that match from the beginning and the end.
//
// Time Complexity: O(n*m)
// Space Complexity: O(n+m)
//
// `limit` of -1 means no limit
func GenJaroWinklerSimilarity[F common.FloatType](prefix_l F, prefix_limit int) func(a, b []byte) F {
  return func(a, b []byte) F {
    return algorithms.JaroWinklerDistance(a, b, prefix_l, prefix_limit)
  }
}

// Returns a function that Calculates JaroWinklerDistance, except this one matches from the end as well.
//
// Time Complexity: O(n*m)
// Space Complexity: O(n+m)
//
// `limit` of -1 means no limit
func GenJaroWinklerSimilarityBidirectional[F common.FloatType](prefix_l F, prefix_limit int, suffix_l F, suffix_limit int) func(a, b []byte) F {
  return func(a, b []byte) F {
    return algorithms.JaroWinklerDistanceBidirectional(a, b, prefix_l, prefix_limit, suffix_l, suffix_limit)
  }
}

// Returns a number between 0 and 1 that represents the percentage of the length of the longest common subsequence.
//
// Time Complexity: O(n*m)
// Space Complexity: O(2 * min(n,m))
//
// LCSPercentage = LCSLength(a, b) / min(len(a), len(b))
func LCSPercentage[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return F(algorithms.LCSLength(a, b)) / F(min(len(a), len(b)))
}


// Calculates the Damerau-Levenshtein distance as a similarity measure
//
// Time Complexity: O(n*m)
// Space Complexity: O(n*m)
//
// DamerauLevenshteinDistancePercentage = 1 - DamerauLevenshteinDistance(a, b) / max(len(a), len(b))
func LevenshteinDamerauSimilarityPercentage[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return 1 - F(algorithms.DamerauLevenshteinDistance(a, b)) / F(max(len(a), len(b)))
}

// Calculates the Optimal String Alignment distance as a similarity measure
//
// Time Complexity: O(n*m)
// Space Complexity: O(3 * min(n,m))
//
// OptimalStringAlignmentDistancePercentage = 1 - OptimalStringAlignmentDistance(a, b) / max(len(a), len(b))
func LevenshteinOSASimilarityPercentage[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return 1 - F(algorithms.LevenshteinOSADistance(a, b)) / F(max(len(a), len(b)))
}

// Calculates Levenshtein distance as a similarity measure
//
// Time Complexity: O(n*m)
// Space Complexity: O(2 * min(n,m))
//
// LevenshteinDistancePercentage = 1 - LevenshteinDistance(a, b) / max(len(a), len(b))
func LevenshteinSimilarityPercentage[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return 1 - F(algorithms.LevenshteinDistance(a, b)) / F(max(len(a), len(b)))
}

// Calculates the Morisitas Overlap Coefficient for the given strings using MultiSet.
// This May? not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
func MorisitasOverlapCoefficient[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.MorisitasOverlapCoefficient[F](a, b)
}

// Calculates the Horns modification of the Morisitas Overlap Coefficient for the given strings using MultiSet.
// This May? not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
func HornsMorisitasOverlapCoefficient[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.HornsMorisitasOverlapCoefficient[F](a, b)
}


// Uses MultiSet, Calculates the Overlap Coefficient for the given strings.
// This Does not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
func OverlapCoefficient[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.OverlapCoefficientCharacter[F](a, b)
}

// Uses Bigram, Set Calculates the Overlap Coefficient for the given strings.
// This Does not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 8kb
func OverlapCoefficientBigram[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  return algorithms.OverlapCoefficientBigram[F](a, b)
}

// Uses MultiSet, Calculates the Tversky index for the given strings.
// This may not follow triangle inequality, depending on the values of alpha and beta.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
func GenTverskyIndex[F common.FloatType](alpha F, beta F) func(a, b []byte) F {
  return func(a, b []byte) F {
    return algorithms.TverskyIndexCharacter(a, b, alpha, beta)
  }
}
// Uses Bigram, Set Calculates the Tversky index for the given strings.
// This may not follow triangle inequality, depending on the values of alpha and beta.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 8kb
func GenTverskyIndexBigram[F common.FloatType](alpha F, beta F) func(a, b []byte) F {
  return func(a, b []byte) F {
    return algorithms.TverskyIndexBigram(a, b, alpha, beta)
  }
}

// Every function that does not start with `Gen` must be wrepped befote being used (to make it non-generic)
func Wrap[F common.FloatType](f func(a, b []byte) F) func(a, b []byte) F { return f }

func WrapTrim[F common.FloatType](f func(a, b []byte) F, prefix_l F, prefix_limit int, suffix_l F, suffix_limit int) func(a, b []byte) F {
  if !(common.Abs(prefix_l) <= 1) { panic("prefix_l must be between -1 and 1") }
  if !(common.Abs(suffix_l) <= 1) { panic("suffix_l must be between -1 and 1") }
  return func(a, b []byte) F {
    min_len := min(len(a), len(b))
    pre := 0
    for pre != prefix_limit && pre < min(len(a), len(b)) && a[pre] == b[pre] { pre += 1 }
    a = a[pre:]
    b = b[pre:]
    suf := 0
    for suf != suffix_limit && suf < min(len(a), len(b)) && a[len(a)-1-suf] == b[len(b)-1-suf] { suf += 1 }

    out := f(a[:len(a)-suf], b[:len(b)-suf])
    return F(pre)*prefix_l/F(min_len) + out*F(min_len-(pre+suf))/F(min_len) + F(suf)*suffix_l/F(min_len)
  }
}

func WrapTrimStart[F common.FloatType](f func(a, b []byte) F, prefix_l F, prefix_limit int) func(a, b []byte) F {
  if !(common.Abs(prefix_l) <= 1) { panic("prefix_l must be between -1 and 1") }
  return func(a, b []byte) F {
    min_len := min(len(a), len(b))
    pre := 0
    for pre != prefix_limit && pre < min(len(a), len(b)) && a[pre] == b[pre] { pre += 1 }

    out := f(a[pre:], b[pre:])
    return F(pre)*prefix_l/F(min_len) + out*F(min_len-pre)/F(min_len)
  }
}

func WrapTrimEnd[F common.FloatType](f func(a, b []byte) F, suffix_l F, suffix_limit int) func(a, b []byte) F {
  if !(common.Abs(suffix_l) <= 1) { panic("suffix_l must be between -1 and 1") }
  return func(a, b []byte) F {
    min_len := min(len(a), len(b))
    suf := 0
    for suf != suffix_limit && suf < min(len(a), len(b)) && a[len(a)-1-suf] == b[len(b)-1-suf] { suf += 1 }

    out := f(a[:len(a)-suf], b[:len(b)-suf])
    return out*F(min_len-suf)/F(min_len) + F(suf)*suffix_l/F(min_len)
  }
}

