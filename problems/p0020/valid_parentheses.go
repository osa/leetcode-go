/*
https://leetcode.com/problems/valid-parentheses/
20. Valid Parentheses
Easy
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package p0020

/*
space : O(n)
time: O(n)
*/
func isValid(s string) bool {
	var stack []rune = []rune{}
	var matches map[rune]rune = map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, r := range s {
		_, exists := matches[r]

		if exists {
			stack = append(stack, r)
		} else if len(stack) == 0 || matches[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1] //pop
		}
	}

	return len(stack) == 0
}
