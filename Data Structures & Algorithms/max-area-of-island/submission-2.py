class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        maxArea = 0
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == 1:
                    currArea = self.dfs(grid, r, c)
                    maxArea = max(maxArea, currArea)
        return maxArea

    def dfs(self, grid: List[List[int]], r, c) -> int:
        if r >= len(grid) or r < 0 or c >= len(grid[0]) or c < 0 or grid[r][c] != 1:
            return 0
        grid[r][c] = 0
        return 1 + self.dfs(grid, r+1, c) + self.dfs(grid, r-1, c) + self.dfs(grid, r, c+1) + self.dfs(grid, r, c-1)

#
# [ 1 0 0 ]
# [ 1 1 0 ]
# [ 0 1 0 ]
# [ 0 1 1 ]
# [ 1 0 0 ]