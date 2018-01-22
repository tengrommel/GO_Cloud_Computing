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

Confinement can also allow for a lighter cognitive load on the developer and smaller critical sections. 

Confinement is the simple yet powerful idea of ensuring information is only ever available from one concurrent process.

There are two kinds of confinement possible: ad hoc and lexical.

This is why I prefer lexical confinement: it wields the compiler to enforce the confinement.

Lexical: it wields the compiler to enforce the confinement.

Lexical confinement involves using lexical scope to expose only the correct data and concurrency primitives
for multiple concurrent processes to use.

## The for-select Loop

    for { // Either loop infinitely or range over something
        select {
        // Do some work with channels
        }
    }

*Sending iteration variables out on a channel*
> Oftentimes you'll want to convert something that can be iterated over into values on a channel.

    for _, s := range []string{"a", "b", "c"} {
        select {
        case <-done:
            return
        case stringStream <- s:    
        }
    }

*Looping infinitely waiting to be stopped*
>It's very common to create goroutines that loop infinitely until they're stopped.<br>
There are a couple variations of this one.

    for {
        select {
        case <-done:
            return
        default:    
        }
        // Do non-preemptable work
    }
    
If the done channel isn't closed, we'll exit the select statement and continue on to the rest of out for loop's body.

*The second variation embeds the work in a default clause of the select statement*

    for {
        select {
        case <-done:
            return
        default:
            // Do non-preemptable work    
        }
    }

## Preventing Goroutine Leaks

*As we covered in the section “Goroutines”, we know goroutines are cheap and easy to create; 
it’s one of the things that makes Go such a productive language.*

The runtime handles multiplexing the goroutines onto any number of operating system threads so that we don't often have to worry about that level of abstraction.

## The or-channel
>At times you may find yourself wanting to combine one or more done channels into a single done channel that closes if any of its component channels close.

1. It is perfectly acceptable, albeit verbose, to write a select statement that performs this coupling; 
2. however, sometimes you can’t know the number of done channels you’re working with at runtime.
