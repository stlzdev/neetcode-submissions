class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        maxArea = 0
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == 1:
                    currArea = self.dfs(grid, r, c, 1)
                    maxArea = max(maxArea, currArea)
        return maxArea

    def dfs(self, grid: List[List[int]], r, c, area) -> int:
        grid[r][c] = 0
        if r+1 < len(grid) and grid[r+1][c] == 1:
            area = self.dfs(grid, r+1, c, area+1)
        if r-1 >= 0 and grid[r-1][c] == 1:
            area = self.dfs(grid, r-1, c, area+1)
        if c+1 < len(grid[0]) and grid[r][c+1] == 1:
            area = self.dfs(grid, r, c+1, area+1)
        if c-1 >= 0 and grid[r][c-1] == 1:
            area = self.dfs(grid, r, c-1, area+1)
        return area

#
# [ 1 0 0 ]
# [ 1 1 0 ]
# [ 0 1 0 ]
# [ 0 1 1 ]
# [ 1 0 0 ]