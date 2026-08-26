import "slices"

func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)   // adjacency list
    state := make([]int, numCourses)     // 0=unvisited, 1=visiting, 2=visited

    // Step 1: build the graph (same as BFS)
	for _, el := range prerequisites {
		graph[el[1]] = append(graph[el[1]], el[0])
	}

    res := []int{}
    hasCycle := false

    // Step 2: define the dfs closure
    var dfs func(course int)
    dfs = func(course int) {
        // base case: visiting -> cycle detected
		if state[course] == 1 {
			hasCycle = true
			return
		}

        // base case: visited -> already done, return
		if state[course] == 2 {
			return
		}

        // mark visiting, recurse into neighbors
		state[course] = 1
		for _, node := range graph[course] {
			dfs(node)
			if hasCycle {
				return
			}
		}

        // mark visited, append to order
		state[course] = 2
		res = append(res, course)
    }

    // Step 3: drive dfs from every unvisited course
	for node, curr := range state {
		if curr == 0 {
			dfs(node)
			if hasCycle {
				break
			}
		}
	}

    // Step 4: if cycle, return empty
	if hasCycle {
		return []int{}
	}

    // Step 5: reverse order, return
	slices.Reverse(res)
	return res
}