package Problems

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(isPalindrome(10)).To(BeFalse(), "isPalindrome() of 10 should be false")
	g.Expect(isPalindrome(11)).To(BeTrue(), "isPalindrome() of 11 should be true")
	g.Expect(isPalindrome(12)).To(BeFalse(), "isPalindrome() of 12 should be false")

	g.Expect(isPalindrome(221)).To(BeFalse(), "isPalindrome() of 221 should be false")
	g.Expect(isPalindrome(222)).To(BeTrue(), "isPalindrome() of 222 should be true")
	g.Expect(isPalindrome(223)).To(BeFalse(), "isPalindrome() of 223 should be false")

	g.Expect(isPalindrome(322)).To(BeFalse(), "isPalindrome() of 322 should be false")
	g.Expect(isPalindrome(323)).To(BeTrue(), "isPalindrome() of 323 should be true")
	g.Expect(isPalindrome(324)).To(BeFalse(), "isPalindrome() of 324 should be false")

	g.Expect(isPalindrome(2441)).To(BeFalse(), "isPalindrome() of 2441 should be false")
	g.Expect(isPalindrome(2442)).To(BeTrue(), "isPalindrome() of 2442 should be true")
	g.Expect(isPalindrome(2443)).To(BeFalse(), "isPalindrome() of 2443 should be false")
}
