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
| Trees | **Construct Binary Tree from Inorder & Preorder** | $O(n)$ | $O(n^2)$ | Use preorder to find inorder root idx, use inorder to partition preorder, recursive. |
| Heaps | **Kth Largest Element in Stream** | $O(m \log k)$ | $O(k)$ | min-heap, take nlargest and heapify in constructor. push always, pop if exceed k els (python) | 
| Heaps | **Last Stone Weight** | $O(n \log n)$ | $O(n)$ | python heapq defaults minheap; use heapify_max, heappop_max etc. (python) |
| Heaps | **K Closest Points to Origin** | $O(n \log k)$ | $O(k)$ | Use min-heap-reverse to store k closest points, pop max (most neg) when size >k. (python) |
| Heaps | **Kth Largest Element in List** | $O(n \log k)$ | $O(k)$ | min-heap / quick-select with partition. (python) | 
| Heaps | **Task Scheduler** | $O(m)$ | $O(1)$ | max heap and cooldown queue, int time to keep track of total time and cooldown time. (python) |
| Backtracking | **Subsets** | $O(n*2^n)$ | $O(2^n)$ | inner helper function, recursively build two paths, one with & one w/o current value. |
| Backtracking | **Combination Sum** | $O(2^(t/m))$ | $O(t/m)$ | at each step, take or skip el, stop when equal or >, use ... to append two lists. |
| Backtracking | **Combination Sum II** | $O(n*2^n)$ | $O(n)$ | sort first, two branches, loop to skip duplicates. |
| Backtracking | **Permutations** | $O(n * n!)$ | $O(n)$ | bool array to store progress, reset bool for backtracking after recursion. |
| Backtracking | **Subsets II** | $O(n * 2^n)$ | $O(n)$ | at each step, take or skip el, iterate for skip cond until next value reached. |
| Backtracking | **Generate Parentheses** | $O(4^n / \sqrt(n))$ | $O(n)$ | condition on start < n, start > end, then append string with + "". |
| Backtracking | **Word Search** | $O(m * 4^n)$ | $O(n)$ | mutate input board to mark path, then undo when backtracking. |
| Backtracking | **Palindrome Partitioning** | $O(n*2^n)$ | $O(n)$ | backtrack, at each start idx, try every cut pt end, check if palindrome. |
| Trie | **Implement Trie** | $O(n)$ | $O(t)$ | store children as map[byte]*PrefixTree, recurse along children node until end of word. |
| Trie | **Design Add Search Words** | $O(n)$ | $O(t+n)$ | helper function for parsing, recurse for wildcard '.' and return true when match. |
| Graphs | **Number of Islands** | $O(m*n)$ | $O(m*n)$ | iterate across all cells, apply dfs at val=1 and reset cells val->0 once parsed. |
| Graphs | **Max Area of Island** | $O(m*n)$ | $O(m*n)$ | dfs helper returns area as int. |
| Graphs | **Clone Graph** | $O(m+n)$ | $O(n)$ | call dfs on node, hmap stores old->new node, return clone if exist, or clone all neighbors. |
| Graphs | **Walls and Gates** | $O(m*n)$ | $O(m*n)$ | multi-source BFS, append neighbors to queue and add dist by 1. |
| Graphs | **Rotting Fruit** | $O(m*n)$ | $O(m*n)$ | per-layer BFS, iterate within each layer to process only cells rotten at given time. |
| Graphs | **Pacific Atlantic Water Flow** | $O(m*n)$ | $O(m*n)$ | multi-source BFS from all coasts, two [][]bool arrays for tracking reachable cells. | 
| Graphs | **Surrounded Regions** | $O(m*n)$ | $O(m*n)$ | fill all 'O' reachable from edges with temp 'V', then traverse twice to convert remaining to 'X'. |
| Graphs | **Course Schedule** | $O(m+n)$ | $O(m+n)$ | adj list, parse and mark visited nodes, true if no cycle. | 

> *Total Solved: 65 / 150*

## ⚙️ Systems Design

| Problem | Notes/Key Concepts |
| :--- | :--- |
| **Design Leetcode** | Code Execution (Serverless, API, VM, Container), RT Web Communication (Websockets vs SSE vs Polling), Redis sorted set for in-memory frequent updates with ZSET and skip lists |
| **Design URL Shortener** | Redis with cache libraries, using MachineID+Sequence as unique ID, base62 for conversion, sharding with machine ID |
| **Design Webhook** | Use message queue for buffering, retain idempotency keys for 30 days to preserve uniqueness, HMAC signature, IP whitelisting and rate limiting for security. | 
| **Design Yelp** | Efficient geo-search with Elastic Search / PostGreSQL + PostGIS; Dual write vs CDC (with costs) for search-store consistency; UX read from primary + optimistic display. | 
| **Design Rate Limiter** | Gateway/Edge with state sharing, extract per API/IP/tenant, token bucket algo, same counter in Redis. |
| **Design Comment System** | Comment database (NoSQL) + user database (DBMS), sharding, Redis for cache, Time-to-Live indexes for deleting expired posts. |
| **Design Dropbox** | Upload file chunks via multi-chunk, download large files directly via S3. Blob storage for files, DynamoDB for metadata. **DELTA SYNCING**: store chunk hashes in DB to handle network interruptions, download changed chunks and stitch on client side. |

> *Total Solved: 7 / 19*

