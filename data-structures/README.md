# Data Structures

The available data structures are:

* Bitset
* Heap
* LinkedList
* Matrix
* PriorityQueue
* Queue
* Stack
* Tree
* Tuple
* Vector


## Containers

The available containers are
* Heap
* LinkedList
* Matrix
* PriorityQueue
* Queue
* Stack
* Tree
* Vector

All of these containers implement the `Container` interface.

All the containers that can be iterated upon sequentially without being mutated are also implement standard iterator methods, such as:
* All
* Keys
* Values
* Collect

so you can do
```go
for v := range vec.Values() {
    fmt.Println(v)
}
```

