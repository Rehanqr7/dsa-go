# DSA in Go (Golang)

A comprehensive, modular repository for Data Structures and Algorithms interview preparation using Go. Organized by topic and difficulty, with clean, idiomatic solutions, problem statements, examples, and tests.

## Getting Started

- Prerequisites: Go 1.21+

## Repository Structure

```
dsa-go/
├── README.md
├── go.mod
├── arrays/
├── strings/
├── linkedlist/
├── stacks/
├── queues/
├── trees/
├── graphs/
├── heaps/
├── dp/
├── backtracking/
├── math/
├── searching/
├── sorting/
├── patterns/
└── utils/
```

- Each file solves one problem: `problem-name.go`
- Each solution includes:
  - Problem statement (comment)
  - Example input/output
  - Time and space complexity
  - Clean, idiomatic Go implementation
  - A corresponding `*_test.go` with example tests

## Coding Standards

- One problem per file, named `kebab-case.go` (e.g., `two-sum.go`).
- Prefer pure functions that can be tested easily.
- Keep helpers reusable inside `utils/`.
- Document complexities and edge cases.

## Contribution Guidelines

1. Create a new file in the appropriate topic directory using `kebab-case.go`.
2. Start with a top-of-file comment: statement, examples, complexity.
3. Add unit tests in `problem-name_test.go`.
4. Ensure `go test ./...` passes locally.
5. Open a PR with a clear description and references (if any).

## Roadmap (Easy → Medium → Hard)

- Arrays: Two Sum → Maximum Subarray → Trapping Rain Water
- Strings: Palindrome → Longest Substring Without Repeating → KMP
- Linked List: Reverse List → Detect Cycle → LRU Cache
- Stacks/Queues: Valid Parentheses → Min Stack → Next Greater Element
- Trees: Traversals → LCA → Diameter → Serialize/Deserialize
- BST: Search/Insert/Delete → Kth Smallest/Largest → Validate BST
- Graphs: BFS/DFS → Topological Sort → Dijkstra/Bellman-Ford → MST
- Heaps: Kth Largest → Median from Stream → Heap Sort
- DP: Fibonacci → Coin Change → LCS/Knapsack → LIS → MCM
- Backtracking: N-Queens → Rat in a Maze → Sudoku → Word Search
- Math: GCD/LCM → Sieve → Fast Power → Mod Exp → Factorization
- Searching: Binary Search variants → Rotated Array
- Sorting: QuickSort → MergeSort → Counting/Radix Sort

## Patterns

- Sliding Window
- Two Pointers
- Fast & Slow Pointers
- Binary Search on Answer
- Greedy
- Divide & Conquer

## Utilities

- `utils/` provides reusable types and helpers: `ListNode`, `TreeNode`, graph builders, traversal printers, and small helpers for tests.

