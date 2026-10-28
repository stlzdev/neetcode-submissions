class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort()
        result = [intervals[0]]
        for start, end in intervals[1:]:
            last = result[-1]
            if last[1] >= start:
                result[-1] = [last[0], max(last[1], end)]
            else:
                result.append([start, end])
        return result


# queue -> pop top two, then merge if needed?
# [1 3] [2 4] -> [1 4] (i[0][1] >= i[1][0] -> push([i[0][0], i[1][1]]))
# [1 3] [1 5] -> [1 5] (i[0][0] == i[1][0] and i[0][1] <= i[1][1] -> push([i[0][0], i[1][1]]))