type WordDictionary struct {
	children map[byte]*WordDictionary
    isEnd bool
}

func Constructor() *WordDictionary {
    return &WordDictionary {
        children: make(map[byte]*WordDictionary),
        isEnd: false,
    }
}

func (this *WordDictionary) AddWord(word string)  {
    for _, c := range word {
        if _, exist := this.children[byte(c)]; !exist {
            this.children[byte(c)] = &WordDictionary {
                children: make(map[byte]*WordDictionary),
                isEnd: false,
            }
        }
        this = this.children[byte(c)]
    }
    this.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    return parse(this, word)
}

func parse(node *WordDictionary, word string) bool {
    if word == "" {
        return node.isEnd
    }
    c := word[0]
    if c == '.' {
        for _, child := range node.children {
            if parse(child, word[1:len(word)]) {
                return true
            }
        }
        return false
    }  
    if _, exist := node.children[byte(c)]; !exist {
        return false
    }
    return parse(node.children[byte(c)], word[1:len(word)])
}
