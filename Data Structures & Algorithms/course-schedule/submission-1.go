func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	parse := make([]int, numCourses) 
	for _, req := range prerequisites {
		adj[req[1]] = append(adj[req[1]], req[0])
	}
	var hasCycle func(int) bool
	hasCycle = func(course int) bool {
		if parse[course] == 2 {
			return false
		}
		if parse[course] == 1 {
			return true
		}
		parse[course] = 1
		for _, next := range adj[course] {
			if hasCycle(next) {
				return true
			}
		}
		parse[course] = 2
		return false
	}
	for i := 0; i < numCourses; i++ {
		if parse[i] == 0 && hasCycle(i) {
			return false
		}
	}
	return true
}
