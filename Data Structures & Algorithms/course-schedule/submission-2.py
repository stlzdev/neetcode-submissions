class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph = [[] for _ in range(numCourses)]
        for course, prereq in prerequisites:
            graph[prereq].append(course)
        state = [0] * numCourses
        def hasCycle(course: int) -> bool:
            if state[course] == 2:
                return False
            if state[course] == 1:
                return True
            state[course] = 1
            for neighbor in graph[course]:
                if hasCycle(neighbor):
                    return True
            state[course] = 2
            return False
        return not any(hasCycle(course) for course in range(numCourses))