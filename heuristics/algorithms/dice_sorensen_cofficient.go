package algorithms

import "fuzzy/common"

// Uses MultiSet, Calculates the Dice-Sorensen coefficient for the given strings.
// This Does not follow triangle inequality
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
//
// The Dice-Sorensen coefficient measures similarity between two strings using the following formula:
// Dice-Sorensen Coefficient = 2 * IntersectionCount(a, b) / (len(a) + len(b))
func DiceSorensenCoefficientCharacter[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  return F(2 * common.IntersectionCharacterCount(a, b)) / F(uint(len(a)) + uint(len(b)))
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
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  return F(2 * common.IntersectionBigramOccurrence(a, b)) / F(uint(len(a)) + uint(len(b)))
}

