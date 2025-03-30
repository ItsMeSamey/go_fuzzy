package algorithms

import "fuzzy/common"

// Uses MultiSet, Calculates the Tversky index for the given strings.
// This may not follow triangle inequality, depending on the values of alpha and beta.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
//
// The Tversky index measures similarity between two strings using the following formula:
// Tversky Index = IntersectionCount(a, b) / (IntersectionCount(a, b) + alpha * (len(a) - IntersectionCount(a, b)) + beta * (len(b) - IntersectionCount(a, b)))
func TverskyIndexCharacter[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B, alpha F, beta F) F {
  intersection := common.IntersectionCharacterCount(a, b)
  return F(intersection) / (F(intersection) + alpha * F(uint(len(a)) - intersection) + beta * F(uint(len(b)) - intersection))
}

// Uses Bigram, Set Calculates the Tversky index for the given strings.
// This may not follow triangle inequality, depending on the values of alpha and beta.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 8kb
//
// The Tversky index measures similarity between two strings using the following formula:
// Tversky Index = IntersectionCount(a, b) / (IntersectionCount(a, b) + alpha * (len(a) - IntersectionCount(a, b)) + beta * (len(b) - IntersectionCount(a, b)))
func TverskyIndexBigram[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B, alpha F, beta F) F {
  intersection := common.IntersectionBigramOccurrence(a, b)
  return F(intersection) / (F(intersection) + alpha * F(uint(len(a)) - intersection) + beta * F(uint(len(b)) - intersection))
}

