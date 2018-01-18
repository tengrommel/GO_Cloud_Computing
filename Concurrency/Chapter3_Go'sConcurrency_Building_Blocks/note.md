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
 
## Mutex and RWMutex
> If you're already familiar with languages that handle concurrency through memory access synchronization,
then you'll probably immediately recognize Mutex.

*Mutex stands for "mutual exclusion" and is a way to guard critical sections of your program.*

>One strategy for doing so is to reduce the cross-section of the critical section. 
There may be memory　that needs to be shared between multiple concurrent processes, 
but perhaps not all of these processes　will read and write to this memory.

*The sync.RWMutex is conceptually the same thing as a Mutex: it guards access to memory; 
however, RWMutex gives you a little bit more control over the memory.*
>You can request a lock for reading, in which case you will be granted access unless the lock is being held for writing.
This means that an arbitrary number of readers can hold a reader lock so long as nothing else is holding a writer lock.

## Cond
>Cond implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event.

*In that definition, an “event” is any arbitrary signal between two or more goroutines that carries no
 information other than the fact that it has occurred. Very often you’ll want to wait for one of these
 signals before continuing execution on a goroutine.*
 
 **Signal & Broadcast**
1. Signal is one of two methods that the Cond type provides for notifying goroutines blocked on a Wait call 
that the condition has been triggered.

2. The other is a method called Broadcast. 
Internally, the runtime maintains a FIFO list of goroutines waiting to be signaled.

**Signal finds the goroutine that's been waiting the longest and notifies that, whereas Broadcast sends 
a signal to all goroutines that are waiting.**

## Once
> As the name implies, sync.Once is a type that utilizes some sync primitives internally to ensure that
only one call to Do ever calls the function passed in --- even on different goroutines.

## Pool
>Pool is a concurrent-safe implementation of the object pool pattern.

A complete explanation of the object pool pattern is best left to literature on design patterns;
**however, since Pool resides in the sync package, we’ll briefly discuss why you might be interested in utilizing it.**

*It’s commonly used to constrain the creation of things that are expensive (e.g., database connections)<br>
so that only a fixed number of them are ever created, but an indeterminate number of operations can still request access to these things.<br>
In the case of Go’s sync.Pool, this data type can be safely used by multiple goroutines.*
 
>Pool’s primary interface is its Get method. 
When called, Get will first check whether there are any available instances within the pool to return to the caller,
and if not, call its New member variable to create a new one.

*When finished, callers call Put to place the instance they were working with back in the pool for use by other processes.*

*Another common situation where a Pool is useful is for warming a cache of pre-allocated objects for 
operations that must run as quickly as possible.*