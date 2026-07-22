class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        heapq.heapify(nums)
        count = len(nums) - k
        while count > 0:
            heapq.heappop(nums)
            count -= 1
        return heapq.heappop(nums)
