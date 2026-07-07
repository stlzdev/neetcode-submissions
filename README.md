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
| Stack | **Daily Temperatures** | $O(n)$ | $O(n)$ | Stack of idxs to keep track of days since highest temp. |
| Binary Search | **Binary Search** | $O(log n)$ | $O(1)$ | prevent int overflow when calculating mid. |
| Binary Search | **Search a 2D Matrix**| $O( log (mn))$ | $O(1)$ | use div-mod within function, no need helper func. |
| Binary Search | **Koko Eating Bananas** | $O(nlog m)$ | $O(1)$ | $$\left\lceil \frac{val}{m} \right\rceil = \left\lfloor \frac{val + m - 1}{m} \right\rfloor$$, manual max of array. |
| Binary Search | **Finding Minimum in Rotated Sorted Array** | $O(log n)$ | $O(1)$ | Update lpt to m+1 instead of m. |
| Binary Search | **Search in Rotated Sorted Array** | $O(log n)$ | $O(1)$ | At any time, one half of array is sorted. |
| Linked List | **Reverse Linked List** | $O(n)$ | $O(1)$ | Declare nil with type, e.g. `(*ListNode)(nil)`. Use next (temp), prev=curr, curr=next pointers. |
| Linked List | **Merge Two Sorted Lists** | $O(n)$ | $O(1)$ | Dummy variable `&ListNode{}`, assign as curr. Iterate pointer as min of two array add to Next. |
| Linked List | **Linked List Cycle** | $O(n)$ | $O(1)$ | Fast-slow pointer, gap between pointers always decrease by 1, max at length so time $O(n)$ bounded. |
| Linked List | **Reorder Linked List** | $O(n)$ | $O(1)$ | FS pointer + reverse 2nd half + dual merge, second half starts at `slow.Next`. |
| Linked List | **Remove Nth Node from End of List** | $O(n)$ | $O(1)$ | Two pointers diff by n, then use `second.Next = second.Next.Next` to remove node. |
| Linked List | **Copy Linked List with Random Pointer** | $O(n)$ | $O(n)$ | Two rounds of hmap, first for storing values, second for getting next & random pointers with old node.next/random as idx. |
| Linked List | **Add Two Numbers** | $O(n)$ | $O(1)$ | Use dummy and curr pointer, addition carry over logic and final digit edge case. |
| Linked List | **LRU Cache** | $O(n)$ | $O(n)$ | Doubly-linked list (`head`, `tail` pointers), delete + insert helpers for handling linked list re-pointing. |
| Trees | **Invert Binary Tree** | $O(n)$ | $O(n)$ | Recursive, DFS explore and reverse left and right, base case return nil. |
| Trees | **Maximum Depth of Binary Tree** | $O(n)$ | $O(n)$ | Can use BFS, iterative/recursive DFS, keep max depth count. |
| Trees | **Diameter of Binary Tree** | $O(n)$ | $O(n)$ | Recursive inner function that computes height, update max diameter of subtrees within function. | 
| Trees | **Balanced Binary Tree** | $O(n)$ | $O(n)$ | height=-1 to flag violation, recursive balance check, increase height bottom-up. | 
| Trees | **Same Binary Tree** | $O(n)$ | $O(n)$ | no helper function needed! |
| Trees | **Subtree of Another Tree** | $O(mn)$ | $O(m+n) | recurse call isSameTree & isSubtree, compare each equal val root with subtree until base. |
| Trees | **Lowest Common Ancestor** | $O(h)$ | $O(1)$ | simple logic. |
| Trees | **Binary Tree Level Order Traversal** | $O(n)$ | $O(n)$ | define `out=[][]int`, `queue.Len()` to iterate through # nodes in level. |
| Trees | **Binary Tree Right Side View** | $O(n)$ | $O(n)$ | BFS, store last node as var for each level. |
| Trees | **Count Good Nodes in Binary Tree** | $O(n)$ | $O(n)$ | DFS recursion, store max value so far on path. |
| Trees | **Valid Binary Search Tree** | $O(n)$ | $O(n)$ | Use `math.MaxInt64/MinInt64` for infinite bounds. |
| Trees | **Kth Smallest Integer in BST** | $O(n)$ | $O(n)$ | Inorder traversal (left -> node -> right), `count++` per node. |

> *Total Solved: 43 / 150*

## ⚙️ Systems Design

| Problem | Notes/Key Concepts |
| :--- | :--- |
| **Design Leetcode** | Code Execution (Serverless, API, VM, Container), RT Web Communication (Websockets vs SSE vs Polling), Redis sorted set for in-memory frequent updates with ZSET and skip lists |
| **Design URL Shortener** | Redis with cache libraries, using MachineID+Sequence as unique ID, base62 for conversion, sharding with machine ID |
| **Design Webhook** | Use message queue for buffering, retain idempotency keys for 30 days to preserve uniqueness, HMAC signature, IP whitelisting and rate limiting for security. | 
| **Design Yelp** | Efficient geo-search with Elastic Search / PostGreSQL + PostGIS; Dual write vs CDC (with costs) for search-store consistency; UX read from primary + optimistic display. | 
| **Design Rate Limiter** | Gateway/Edge with state sharing, extract per API/IP/tenant, token bucket algo, same counter in Redis. |

> *Total Solved: 5 / 19*

