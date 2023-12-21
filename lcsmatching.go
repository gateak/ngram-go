package main

import ("fmt"; "os")

// longestCommonSubstrings finds all longest common substrings between two strings
func longestCommonSubstrings(str1, str2 string) []string {
    len1, len2 := len(str1), len(str2)
    maxLength := 0
    dp := make([][]int, len1+1)
    var result []string
    var substringSet = make(map[string]bool)

    for i := range dp {
        dp[i] = make([]int, len2+1)
    }

    for i := 1; i <= len1; i++ {
        for j := 1; j <= len2; j++ {
            if str1[i-1] == str2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
                if dp[i][j] > maxLength {
                    maxLength = dp[i][j]
                    result = result[:0]
                    substringSet = make(map[string]bool)
                }
                if dp[i][j] == maxLength {
                    substr := str1[i-maxLength : i]
                    if _, exists := substringSet[substr]; !exists {
                        result = append(result, substr)
                        substringSet[substr] = true
                    }
                }
            }
        }
    }

    return result
}

func main() {
    args := os.Args
    if len(args) < 3 {
        fmt.Println("Usage: go run program.go <string1> <string2>")
        return
    }

    str1 := args[1]
    str2 := args[2]

    fmt.Println("Longest Common Substrings:", longestCommonSubstrings(str1, str2))
}
