package algorithms

import "fuzzy/common"

// Uses MultiSet, Calculates the Overlap Coefficient for the given strings.
// This Does not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
//
// The Overlap Coefficient measures similarity between two strings using the following formula:
// Overlap Coefficient = IntersectionCount(a, b) / min(len(a), len(b))
func OverlapCoefficientCharacter[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  return F(common.IntersectionCharacterCount(a, b)) / F(min(len(a), len(b)))
}

// Uses Bigram, Set Calculates the Overlap Coefficient for the given strings.
// This Does not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 8kb
//
// The Overlap Coefficient measures similarity between two strings using the following formula:
// Overlap Coefficient = IntersectionCount(a, b) / min(len(a), len(b))
func OverlapCoefficientBigram[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  return F(common.IntersectionBigramOccurrence(a, b)) / F(min(len(a), len(b)))
}

