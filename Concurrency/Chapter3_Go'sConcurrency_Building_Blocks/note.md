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

This is an interesting side note about how Go manages memory. 
The Go runtime is observant enough　to know that a reference to the salutation variable is still being held,
 and therefore will transfer the memory to the heap so that the goroutines can continue to access it.
 
 **A few kilobytes per goroutine**
 
*Something that might dampen out spirits is context switching,which is when something hosting a concurrent process must save its state to switch to running a different concurrent process.*
>If we have too many concurrent processes, we can spend all of our CPU time context switching between them and never get any real work done.<br>
At the OS level, with threads, this can be quite costly. <br>
The OS thread must save things like register values, lookup tables, and memory maps to successfully be able to switch back to the current thread when it is time.<br>
Then it has to load the same information for the incoming thread.

Context switching in software is comparatively much, much cheaper.<br>
Under a software-defined scheduler, the runtime can be more selective in what is persisted for retrieval, how it is persisted, and when the persisting need occur.<br>
Let’s take a look at the relative performance of context switching on my laptop between OS threads and goroutines.

## The sync Package
>The sync package contains the concurrency primitives that are most useful for low-level memory access synchronization.

### WaitGroup
>WaitGroup is a great way to wait for a set of concurrent operations to complete when you either
 don’t care about the result of the concurrent operation, or you have other means of collecting their results.
 
