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

## Error Handling
>In concurrent programs, error handling can be difficult to get right.

**With concurrent processes, this question becomes a little more complex.Because a concurrent
  process is operating independently of its parent or siblings, it can be difficult for it to reason about
  what the right thing to do with the error is.**

When Go eschewed the popular exception model of errors, it made a statement that error handling was important, 
and that as we develop our programs, we should give our error paths the same attention we give our algorithms. 

*The key thing to note here is how we're coupled the potential result with the potential error.*

1. This represents the complete set of possible outcomes created from the goroutine checkStatus, 
and allows our main goroutine to make decisions about what to do when errors occur. 

2. In broader terms, we’ve successfully separated the concerns of error handling from our producer goroutine.

>This is desirable because the goroutine that spawned the producer goroutine — in this case our main
goroutine — has more context about the running program, and can make more intelligent decisions about what to do with errors.

## Pipelines
>When you write a program, you probably don't sit down and write one long function -- at least I hope you don't.

Partly to abstract away details that don't matter to the greater flow, and partly so that we can work
on one area of code without affecting other areas.

1. A pipeline is just another tool you can use to form an abstraction in your system.
2. By using a pipeline, you separate the concerns of each stage, which provides numberous benefits.

As mentioned previously, a stage is just something that takes data in, performs a transformation on it, and sends the data back out.

What are the properties of a pipeline stage:

1. A stage consumes and return the same type.
2. A stage must be reified by the language so that it may be passed around.

*Those of you familiar with functional programming may be nodding your head and thinking of terms like higher order functions and monads.*

**Indeed, pipeline stages are very closely related to functional programming and can be considered a subset of monads**

>Notice how each stage is taking a slice of data and returning a slice of data?

1. These stages are performing what we call batch processing.
2. This just means that they operate on chunks of data all at once instead of one discrete value at a time.
3. There is another type of pipeline stage that performs stream processing.
4. This means that the stage receives and emits one element at a time.

## Best Practices for Constructing Pipelines
> Channels are uniquely suited to constructing pipelines in Go because they fulfill all of our basic requirements.

## Some Handy Generators
> As a reminder, a generator for a pipeline is any function that converts a set of discrete values into a stream of values on channel.
 
When you need to deal in specific types, you can place a stage that performs the type assertion for you.
> The performance overhead of having an extra pipeline stage (and thus goroutine) and the type assertion are negligible, as we'll see in just a bit.

## Fan-Out, Fan-In

Sometimes, stages in your pipeline can be particularly computationally expensive.
When this happens, upstream stages in your pipeline can become blocked while waiting for your expensive stages to complete.

1. One of the interesting properties of pipelines is the ability they give you to operate on the stream of data using a combination of separate, often reorderable stages. 
2. You can even reuse stages of the pipeline multiple times. 

You might consider fanning out one of your stages if both the following apply:
- It doesn't rely on values that the stage had calculated before.
- It takes a long time to run.

The property of order-independence is important because you have no guarantee in what order<br>
concurrent copies of your stage will run, nor in what order they will return.


