package longestsubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		Input  string
		Expect int
	}{
		{
			Input:  "abcabcbb",
			Expect: 3,
		}, {
			Input:  "bbbbb",
			Expect: 1,
		}, {
			Input:  "pwwkew",
			Expect: 3,
		},
	}

	for _, c := range cases {
		actual := lengthOfLongestSubstring(c.Input)
		if actual != c.Expect {
			t.Errorf("Expect %d, But got %d", c.Expect, actual)
		}
	}
}
