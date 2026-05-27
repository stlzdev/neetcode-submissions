# NeetCode Tracker — @stlzdev

> My personal repository for tracking progress and solutions for NeetCode problems.

## 🚀 Progress Tracker

> Data Structures and Algorithms (in Go)

| Category | Problem | Time | Space | Notes / Key Concepts |
| :--- | :--- | :---: | :---: | :--- |
| Arrays & Hashing | **Contains Duplicate** | $O(n)$ | $O(n)$ | Used a `map[int]struct` hash set for O(1) lookups. |
| Arrays & Hashing | **Valid Anagram** | $O(n+m)$ | $O(n)$ | Use `map[byte]int` hash map, validate with ending zeroes. |
| Arrays & Hashing | **Two Sum** | $O(n)$ | $O(n)$ |  One-pass hash map with `map[int] int`, check complement exists. |
| Arrays & Hashing | **Group Anagrams** | $O(mnlogn)$ | $O(mn)$ | `map[string][]string` with sorted string as key. |
| Arrays & Hashing | **Top K Frequent Elements** | $O(n)$ | $O(n)$ | bucket sort by frequency count. |
| Arrays & Hashing | **Encode and Decode String** | $O(n)$ | $O(n)$ | Add length and delimiter to decoded array. |
| Arrays & Hashing | **Product of Array Except Self** | $O(n)$ | $O(n)$ | Two-pass Prefix & Suffix running products. |
| Two Pointers | **Is Valid Palindrome** | $O(n)$ | $O(1)$ | Alphanumeric check with `unicode.isLetter/Digit` (force type to rune). |
| Two Pointers | **Two Sum II** | $O(n)$ | $O(1)$ | Two pointer more space-efficient than hash map for sorted array. |
| Two Pointers | **Three Sum** | $O(n)$ | $O(1)$ | Nested for loop implementing Two Sum II for each element after sorting. |


> *Total Solved: 10 / 150*

> Systems Design

| Problem | Notes/Key Concepts |
| :--- | :--- |
| **Design Leetcode** | Code Execution (Serverless, API, VM, Container), RT Web Communication (Websockets vs SSE vs Polling), Redis sorted set for in-memory frequent updates with ZSET and skip lists |

