package algorithms

import "github.com/ItsMeSamey/go_fuzzy/common"

// Uses MultiSet, Calculates the Jaccard coefficient for the given strings.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
//
// The Jaccard coefficient measures similarity between two strings using the following formula:
// Jaccard Coefficient = IntersectionCount(a, b) / (len(a) + len(b) - IntersectionCount(a, b))
func JaccardCoefficientCharacter[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  intersection := common.IntersectionCharacterCount(a, b)
  return F(intersection) / F(uint(len(a)) + uint(len(b)) - intersection)
}

// Uses Bigram Set, Calculates the Jaccard coefficient for the given strings.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 8kb
//
// The Jaccard coefficient measures similarity between two strings using the following formula:
// Jaccard Coefficient = IntersectionCount(a, b) / (len(a) + len(b) - IntersectionCount(a, b))
func JaccardCoefficientBigram[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  intersection := common.IntersectionBigramOccurrence(a, b)
  return F(intersection) / F(uint(len(a)) + uint(len(b)) - intersection)
}

