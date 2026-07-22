class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        heap = []
        for pair in points:
            dist = pair[0]**2 + pair[1]**2
            heapq.heappush(heap, (dist, pair))
        count = k
        out = []
        while count > 0:
            val = heapq.heappop(heap)
            out.append(val[1])
            count -= 1
        return out