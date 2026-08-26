func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)   // adjacency list
    indeg := make([]int, numCourses)  // count of prereqs remaining

    // Step 1: build the graph
    // for each pair, figure out which course is "before" and which is "after"
	for _, pr := range prerequisites {
		graph[pr[1]] = append(graph[pr[1]], pr[0])
		indeg[pr[0]] += 1 
	}

    // Step 2: seed the queue with in-degree 0 courses
    q := []int{}
	for node, deg := range indeg {
		if deg == 0 {
			q = append(q, node)
		}
	}

    // Step 3: process the queue, build the result order
	res := []int{}
	for len(q) > 0 {
		head := q[0]
		res = append(res, head)
		q = q[1:]
		for _, node := range graph[head] {
			indeg[node] -= 1
			if indeg[node] == 0 {
				q = append(q, node)
			}
		}
	}

    // Step 4: compare result length to numCourses
	if len(res) == numCourses {
		return res
	}
	return []int{}
}
