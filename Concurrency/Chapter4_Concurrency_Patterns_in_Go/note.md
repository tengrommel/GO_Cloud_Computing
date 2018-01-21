# Concurrency Patterns in Go
>We've explored the fundamentals of Go's concurrency primitives and discussed how to properly use these primitives.

## Confinement
>When working with concurrent code, there are a few different options for safe operation.

- Synchronization primitives for sharing memory(e.g., sync.Mutex)
- Synchronization via communicating(e.g., channels)

However, there are a couple of other options that are implicitly safe within multiple concurrent processes:
- Immutable data
- Data protected by confinement

In some sense, immutable data is ideal because it is implicitly concurrent-safe.<br>
Each concurrent process may operate on the same data, but it may not modify it.<br>
If it wants to create new data, it must create a new copy of the data with the desired modifications.<br>
This allows not only a lighter cognitive load on the developer, but can also lead to faster programs if it leads to smaller critical sections(or eliminates them altogether)

In Go, you can achieve this by writing code that utilizes copies of values instead of pointers to values in memory.<br>
Some languages support utilization of pointers with explicitly immutable values; however, Go is not among these.