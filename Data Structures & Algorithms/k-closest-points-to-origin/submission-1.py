class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        heap = []
        for pair in points:
            dist = -(pair[0]**2 + pair[1]**2)
            heapq.heappush(heap, (dist, pair))
            if len(heap) > k:
                heapq.heappop(heap)
        return [pair[1] for pair in heap]