# Gollections
###### an experiment with generics in Go ..._totally usable in production..._

### Installation
`go get github.com/michaelwomack/gollections`

## Structures

### Min Heap
A collection to remove elements in ascending order by implementing the `Lesser` interface. Adds and removes 
elements in O(logn) time:

```go
type Player struct {
	Username string
	Rank     int64
}

func (t *Player) Less(other Player) bool {
	return t.Rank < other.Rank
}

players := []*Player{
    {
        Username: "player1",
        Rank: 341,
    },
    {
        Username: "player2",
        Rank: 21,
    },
    {
        Username: "player3",
        Rank: 87,
    },
}

min := gollections.NewMinHeap(players...) // create min heap of size 3
p2, ok := min.Pop() // pop player2
p3, ok := min.Pop() // pop player3
p1, ok := min.Pop() // pop player1
```

### Max Heap
A collection to remove elements in descending order by implementing the `Lesser` interface. This structure is not 
really necessary since you can simply reverse the way `Lesser` is implemented, but it's there anyway. Also adds and 
removes elements in O(logn) time.

