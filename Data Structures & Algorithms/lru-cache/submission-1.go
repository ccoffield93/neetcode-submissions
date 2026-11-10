type LRUCache struct {
    capacity int
	list     *list.List 
	datamap map[int]*list.Element
}

type cacheItem struct { 
	key int 
	value int
}

func Constructor(capacity int) LRUCache {
    cache := LRUCache{capacity: capacity, list: list.New(), datamap: make(map[int]*list.Element)}
	return cache
}

func (this *LRUCache) Get(key int) int {
    val, found := this.datamap[key]
	if !found {
		return -1
	}

	// move to the front of the line 
	// put at the front of the line
	this.list.MoveToFront(val)

	return val.Value.(*cacheItem).value
}

func (this *LRUCache) Put(key int, value int) {
	listItem, found := this.datamap[key]
	if found {
		// update value 
		listItem.Value.(*cacheItem).value = value

		// put at the front of the line
		this.list.MoveToFront(listItem)

	} else {
		// remove the last item if needed 
		for this.list.Len() >= this.capacity {
			toRemove := this.list.Back()
			if toRemove != nil {
				this.list.Remove(toRemove)
			}
			delete(this.datamap, toRemove.Value.(*cacheItem).key)
		}

		// add to front
		ent := &cacheItem{key: key, value: value}
		elem := this.list.PushFront(ent)
		this.datamap[key] = elem
	}
}
