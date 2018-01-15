# Longest Palindromic Substring
> Given a string s, find the longest palindromic substring in s.You may assume that the maximum length of s is 1000.

- Example:
    
    
    Input: "babad"
    Output: "bab"
    Note: "aba" is also a valid answer.
    
- Example:

    
    Input: "cbbd"
    Output: "bb"
    
## 解题思路
> 题目要求寻找字符串中最长的回文。当然我们可以使用下面的程序判断字符串s[i:j+1]是否是回文。

    func isPalindromic(s *string, i, j int) bool {
        for i < j {
            if (*s)[i] != (*s)[j]{
                return false
            }
            i++
            j--
        }
        return true
    }
    
但是，这样就没有利用回文的特点，假定回文的长度为l，x为任意字符

1. 当l为奇数时，回文的正中间段只会是，"x"，或"xxx"，或"xxxxx",或...
2. 当l为偶数时，回文的正中间段只会是，"xx"，或"xxxx"，或"xxxxxx"，或...
3. 同一个字符串的任意两个回文substring的正中间段，不会重叠。

## 总结
**充分利用查找对象的特点，可以加快查找速度。**