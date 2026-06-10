# NeetCode Tracker — @stlzdev

> My personal repository for tracking progress and solutions for NeetCode problems.

## 🧩 Data Structures and Algorithms (in Go)

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
| Sliding Window | **Best Time to Buy and Sell Stock** | $O(n)$ | $O(1)$ | Use pointer to indicate min value, max diff is global. |
| Sliding Window | **Longest Substring without Repeating Characters** | $O(n)$ | $O(m)$ | Move lpt until no repeating in hashmap |
| Sliding Window | **Longest Repeating Character Replacement** | $O(n)$ | $O(m)$ | Most freq el in window only updates new el. For loop increment lpt until replacements < k |
| Sliding Window | **Permutation in String** | $O(n)$ | $O(1)$ | Compare s1 alphabet freq array to each sliding window freq array from s2. |
| Stack | **Valid Parentheses** | $O(n)$ | $O(n)$ | Use map[rune]rune to encode matching parentheses. Pop from stack when match. |
| Stack | **Min Stack** | $O(1)$ | $O(n)$ | Create constructor with custom type, use min stack to track min at any point. |
| Stack | **Evaluate Reverse Polish Notation** | $O(n)$ | $O(n)$ | Use strconv.Atoi to convert from str to int. |
| Binary Search | **Binary Search** | $O(log n)$ | $O(1)$ | prevent int overflow when calculating mid. |
| Binary Search | **Search a 2D Matrix**| $O( log (mn))$ | $O(1)$ | use div-mod within function, no need helper func. |
| Binary Search | **Koko Eating Bananas** | $O(nlog m)$ | $O(1)$ | $$\left\lceil \frac{val}{m} \right\rceil = \left\lfloor \frac{val + m - 1}{m} \right\rfloor$$, manual max of array. |
| Binary Search | **Finding Minimum in Rotated Sorted Array** | $O(log n)$ | $O(1)$ | Update lpt to m+1 instead of m. |
| Binary Search | **Search in Rotated Sorted Array** | $O(log n)$ | $O(1)$ | At any time, one half of array is sorted. |
| Linked List | **Reverse Linked List** | $O(n)$ | $O(1)$ | Declare nil with type, e.g. `(*ListNode)(nil)`. Use curr, next (temp), prev pointers. |

> *Total Solved: 23 / 150*

## ⚙️ Systems Design

| Problem | Notes/Key Concepts |
| :--- | :--- |
| **Design Leetcode** | Code Execution (Serverless, API, VM, Container), RT Web Communication (Websockets vs SSE vs Polling), Redis sorted set for in-memory frequent updates with ZSET and skip lists |
| **Design URL Shortener** | Redis with cache libraries, using MachineID+Sequence as unique ID, base62 for conversion, sharding with machine ID |
| **Design Webhook** | Use message queue for buffering, retain idempotency keys for 30 days to preserve uniqueness, HMAC signature, IP whitelisting and rate limiting for security. | 

> *Total Solved: 3 / 55*

