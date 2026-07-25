class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        counts = Counter(tasks)
        heap = [c for c in counts.values()]
        heapq.heapify_max(heap)
        cooldown = []
        time = 0
        while heap or cooldown:
            time += 1
            if heap:
                head = heapq.heappop_max(heap) - 1
                if head > 0:
                    cooldown.append((time + n, head))
            if cooldown and cooldown[0][0] == time:
                _, cnt = cooldown.pop(0)
                heapq.heappush_max(heap, cnt)
        return time

        