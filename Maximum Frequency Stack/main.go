package main

import "fmt"

func main() {

	impl := Constructor()
	impl.Push(5)
	val := impl.Pop()
	fmt.Println("popped value:", val)

}

type FreqStack struct {
	freq    map[int]int
	group   map[int][]int
	maxfreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:  make(map[int]int),
		group: make(map[int][]int),
	}

}

func (fs *FreqStack) Push(val int) {
	f := fs.freq[val] + 1
	fs.freq[val] = f

	if f > fs.maxfreq {
		fs.maxfreq = f
	}

	fs.group[f] = append(fs.group[f], val)

}

func (fs *FreqStack) Pop() int {
	slice := fs.group[fs.maxfreq]
	val := slice[len(slice)-1]
	fs.group[fs.maxfreq] = slice[:len(slice)-1]

	fs.freq[val]--
	if len(fs.group[fs.maxfreq]) == 0 {
		fs.maxfreq--
	}

	return val

}
