package main

import (
	"fmt"
)

func toKey(i, j int) string {
	return fmt.Sprintf("%d,%d\n", i, j)
}

func getLongestPalindromicIndeces(s string, si, e int) (int, int) {
	for true {
    if s[si] != s[e] {
      break
    } else if si-1 < 0 || e +1 >= len(s) {
      break
    } else if s[si-1] != s[e+1] {
      break
    } else {
      si--
      e++
    }
	}

	return si, e
}

func longestPalindrome(s string) string {

  if len(s) == 0 {
    return ""
  }

  var longestI int 
  var longestE int

  for i := 0; i < len(s)-1; i++ {
    tI, tE := getLongestPalindromicIndeces(s, i, i)

    if s[i] == s[i+1] {
      sI, sE := getLongestPalindromicIndeces(s, i, i+1)
      if sE - sI > tE - tI {
        tI = sI
        tE = sE
      }
    }

    if longestE - longestI < tE - tI {
      longestE = tE
      longestI = tI
    }
  }
  return s[longestI: longestE+1]
}

func main() {
	fmt.Printf("longestPalindrome %s\n", longestPalindrome(""))
}
