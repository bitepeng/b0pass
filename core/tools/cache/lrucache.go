package tools

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type elem struct {
	key        interface{}
	data       interface{}
	expireTime int64
	next       *elem
	pre        *elem
}

type lrucache struct {
	maxSize   int
	elemCount int
	elemList  map[interface{}]*elem
	first     *elem
	last      *elem
	mu        sync.Mutex
}

// New create a new lrucache
// size: max number of element
func New(size int) (*lrucache, error) {
	newCache := new(lrucache)
	newCache.maxSize = size
	newCache.elemCount = 0
	newCache.elemList = make(map[interface{}]*elem)
	return newCache, nil
}

// Set create or update an element using key
//      key:    The identity of an element
//      value:  new value of the element
//      ttl:    expire time, unit: second
func (c *lrucache) Set(key interface{}, value interface{}, ttl ...int) error {

	// Ensure ttl are correct
	if len(ttl) > 1 {
		return errors.New("wrong para number, 2 or 3 expected but more than 3 received")
	}
	var elemTTL int64
	if len(ttl) == 1 {
		elemTTL = int64(ttl[0])
	} else {
		elemTTL = -1
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.elemList[key]; ok {
		e.data = value
		if elemTTL == -1 {
			e.expireTime = elemTTL
		} else {
			e.expireTime = time.Now().Unix() + elemTTL
		}
		c.mvKeyToFirst(key)
	} else {
		if c.elemCount+1 > c.maxSize {
			if c.checkExpired() <= 0 {
				c.eliminationOldest()
			}
		}
		newElem := &elem{
			key:        key,
			data:       value,
			expireTime: -1,
			pre:        nil,
			next:       c.first,
		}
		if elemTTL != -1 {
			newElem.expireTime = time.Now().Unix() + elemTTL
		}
		if c.first != nil {
			c.first.pre = newElem
		}
		c.first = newElem
		c.elemList[key] = newElem

		c.elemCount++
	}
	return nil
}

// updateKeyPtr 更新对应key的指针，放到链表的第一个
func (c *lrucache) mvKeyToFirst(key interface{}) {
	elem := c.elemList[key]
	if elem.pre == nil {
		// 当key是第一个元素时，不做动作
		return
	} else if elem.next == nil {
		// 当key不是第一个元素，但是是最后一个元素时，提到第一个元素去
		elem.pre.next = nil

		c.last = elem.pre

		elem.pre = nil
		elem.next = c.first
		c.first = elem

	} else {
		elem.pre.next = elem.next
		elem.next.pre = elem.pre

		elem.next = c.first
		elem.pre = nil
		c.first = elem
	}
}

func (c *lrucache) eliminationOldest() {
	if c.last == nil {
		return
	}
	if c.last.pre != nil {
		c.last.pre.next = nil
	}
	key := c.last.key
	c.last = c.last.pre
	delete(c.elemList, key)
}

func (c *lrucache) deleteByKey(key interface{}) {
	if v, ok := c.elemList[key]; ok {
		if v.pre == nil && v.next == nil {
			// 当key是第一个元素时，清空元素列表，充值指针和元素计数
			c.elemList = make(map[interface{}]*elem)
			c.elemCount = 0
			c.last = nil
			c.first = nil
			return
		} else if v.next == nil {
			// 当key不是第一个元素，但是是最后一个元素时,修改前一个元素的next指针并修改c.last指针
			v.pre.next = v.next
			c.last = v.pre
		} else if v.pre == nil {
			c.first = v.next
			c.first.pre = nil
		} else {
			// 中间元素，修改前后指针
			v.pre.next = v.next
			v.next.pre = v.pre
		}
		delete(c.elemList, key)
		c.elemCount--
	}
}

// 遍历链表，检查并删除已经过期的元素
func (c *lrucache) checkExpired() int {
	now := time.Now().Unix()
	tmp := c.first
	count := 0
	for tmp != nil {
		if tmp.expireTime != -1 && now > tmp.expireTime {
			c.deleteByKey(tmp.key)
			count++
		}
		tmp = tmp.next
	}
	return count
}

// Get Get the value of a cached element by key. If key do not exist, this function will return nil and a error msg
//      key:    The identity of an element
//      return:
//          value:  the cached value, nil if key do not exist
//          err:    error info, nil if value is not nil
func (c *lrucache) Get(key interface{}) (value interface{}, err error) {
	if v, ok := c.elemList[key]; ok {
		if v.expireTime != -1 && time.Now().Unix() > v.expireTime {
			// 如果过期了
			c.deleteByKey(key)
			return nil, errors.New("the key was expired")
		}
		c.mvKeyToFirst(key)
		return v.data, nil
	}
	return nil, errors.New("no value found")
}

// Delete delete an element
func (c *lrucache) Delete(key interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.elemList[key]; !ok {
		return fmt.Errorf("key %T do not exist", key)
	}
	c.deleteByKey(key)
	return nil
}
