# Chapter 3.Go's Concurrency Building Blocks
>Go's rich tapestry of features that support its concurrency story.

## Goroutines
> Goroutines are one of the most basic unints of organization in a Go program,so it's important 
we understand what they are and how they work.

*In fact, every Go program has at least one goroutine: the main goroutine,which is automatically created and
started when the process begins.*

deep integration with Go's runtime.

- Goroutines don't define their own suspension or reentry points;
- Go's runtime observes the runtime behavior of goroutines and automatically suspends them when they block 
and then resumes them when they become unblocked.
- Go's mechanism for hosting goroutines is an implementation of what's called an M:N scheduler, which means it maps
M green threads to N OS threads.

