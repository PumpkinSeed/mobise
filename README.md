# Modified binary search

Based on the [article](https://confusedcoders.com/data-structures/find-minimum-element-in-rotated-sorted-array-modified-binary-search).

## Usage

```
var arr = []int{7, 8, 9, 10, 11, 13, 24, 1, 2, 3, 4, 5, 6}

value := Search(arr, 0, len(arr)-1)
// value == 7
// arr[value] == 1
```

### Benchmark

```
Number of elements: 13055
BenchmarkSearch-4
20000000	        70.0 ns/op
```