func lengthOfLongestSubstring(s string) int {
    l := 0
    r := 0
    maxlen := 0
    set := make(map[byte]bool)
    for r < len(s) {
        if !set[s[r]] {
            set[s[r]] = true
            r++
            maxlen = max(maxlen, r - l)
        } else {
            delete(set, s[l])
            l++
        }
    }
    return maxlen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
