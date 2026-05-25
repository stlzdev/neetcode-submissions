class DynamicArray:
    
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.size = 0
        self.a = [0] * capacity

    def get(self, i: int) -> int:
        return self.a[i]

    def set(self, i: int, n: int) -> None:
        self.a[i] = n

    def pushback(self, n: int) -> None:
        if self.size == self.capacity:
            self.resize()
        self.a[self.size] = n
        self.size += 1

    def popback(self) -> int:
        self.size -= 1
        return self.a[self.size]

    def resize(self) -> None:
        self.capacity *= 2
        newArr = [0] * self.capacity
        for i in range(self.size):
            newArr[i] = self.a[i]
        self.a = newArr

    def getSize(self) -> int:
        return self.size
        
    
    def getCapacity(self) -> int:
        return self.capacity
