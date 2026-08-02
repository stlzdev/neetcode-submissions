type PrefixTree struct {
	children map[byte]*PrefixTree
	isWord bool
}

func Constructor() *PrefixTree {
    return &PrefixTree {
		children: make(map[byte]*PrefixTree),
		isWord: false,
	}
}

func (this *PrefixTree) Insert(word string) {
	for _, c := range word {
		if _, ok := this.children[byte(c)]; !ok {
			this.children[byte(c)] = &PrefixTree{
				children: make(map[byte]*PrefixTree),
				isWord: false,
			}
		}
		this = this.children[byte(c)]
	}
	this.isWord = true
}

func (this *PrefixTree) Search(word string) bool {
	for _, c := range word {
		if _, ok := this.children[byte(c)]; !ok {
			return false
		}
		this = this.children[byte(c)]
	}
	return this.isWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if _, ok := this.children[byte(c)]; !ok {
			return false
		}
		this = this.children[byte(c)]
	}
	return true
}
