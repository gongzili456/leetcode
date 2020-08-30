package leetcode

/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// MyCircularDeque is
// @lc code=start
type MyCircularDeque struct {
	Cap   int
	Front int
	Rear  int
	Body  []int
}

// Constructor is
/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		Body: make([]int, k+1),
		Cap:  k + 1,
	}
}

// InsertFront is
/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (t *MyCircularDeque) InsertFront(value int) bool {
	if t.IsFull() {
		return false
	}
	t.Front = (t.Front - 1 + t.Cap) % t.Cap
	t.Body[t.Front] = value
	return true
}

// InsertLast is
/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (t *MyCircularDeque) InsertLast(value int) bool {
	if t.IsFull() {
		return false
	}
	t.Body[t.Rear] = value
	t.Rear = (t.Rear + 1) % t.Cap
	return true
}

// DeleteFront is
/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (t *MyCircularDeque) DeleteFront() bool {
	if t.IsEmpty() {
		return false
	}
	t.Front = (t.Front + 1) % t.Cap
	return true
}

// DeleteLast is
/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (t *MyCircularDeque) DeleteLast() bool {
	if t.IsEmpty() {
		return false
	}
	t.Rear = (t.Rear - 1 + t.Cap) % t.Cap
	return true
}

// GetFront is
/** Get the front item from the deque. */
func (t *MyCircularDeque) GetFront() int {
	if t.IsEmpty() {
		return -1
	}

	return t.Body[t.Front]
}

// GetRear is
/** Get the last item from the deque. */
func (t *MyCircularDeque) GetRear() int {
	if t.IsEmpty() {
		return -1
	}

	return t.Body[(t.Rear-1+t.Cap)%t.Cap]
}

// IsEmpty is
/** Checks whether the circular deque is empty or not. */
func (t *MyCircularDeque) IsEmpty() bool {
	return t.Front == t.Rear
}

// IsFull is
/** Checks whether the circular deque is full or not. */
func (t *MyCircularDeque) IsFull() bool {
	return (t.Rear+1)%t.Cap == t.Front
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end
