# Chapter 2. Modeling Your Code: Communicating Sequential Processes

## The Difference Between Concurrency and Parallelism
> The fact that concurrency is different from parallelism is often overlooked or misunderstood.

*The reason to differentiate goes well beyond pedantry.<br>
The difference between concurrency and parallelism turns out to be a very powerful abstraction when modeling your code,and Go takes full advantage of this.*

**Concurrency is a property of the code;parallelism is a property of the running program.**

That's kind of an interesting distinction.
1. Don't we usually think about these two things the same way?
2. We write our code so that it will execute in parallel.Right?

If I write my code with the intent that two chunks of the program will run in parallel, do I have any guarantee that will actually happen when the program is run? <br>
What happens if I run the code on a machine with only one core? <br>
Some of you may be thinking, It will run in parallel, but this isn’t true!

The chunks of our program may appear to be running in parallel, but really they’re executing in a sequential manner faster than is distinguishable.<br>
The CPU context switches to share time between different programs, and over a coarse enough granularity of time, the tasks appear to be running in parallel.<br> 
If we were to run the same binary on a machine with two cores, the program’s chunks might actually be running in parallel.

- The first is that we do not write parallel code,  only concurrent code that we hope will be run in parallel. 
>Once again, parallelism is a property of the runtime of our program, not the code。

- The second interesting thing is that we see it is possible — maybe even desirable — to be ignorant of whether our concurrent code is actually running in parallel.

- The third and final interesting thing is that parallelism is a function of time, or context.

## what Go do
Before Go was first revealed to the public, this was where the chain of abstraction ended for most of the popular programming languages. <br>
If you wanted to write concurrent code, you would model your program in terms of threads and synchronize the access to the memory between them. <br>
If you had a lot of things you had to model concurrently and your machine couldn’t handle that many threads, you created a thread pool and multiplexed your operations onto the thread pool.

1. Go has added another link in that chain: the goroutine.
2. In addition,Go has borrowed several concepts from the work of famed computer scientist Tony Hoare, and introduced new primitives for us to use, namely channels.

Threads are still there, of course, but we find that we rarely have to think about our problem space in terms of OS threads. <br>
Instead, we model things in goroutines and channels, and occasionally shared memory.

## What Is CSP?
> CSP stands for “Communicating Sequential Processes,” which is both a technique and the name of the
paper that introduced it. In 1978, Charles Antony Richard Hoare published the paper in the
Association for Computing Machinery (more popularly referred to as ACM).

*In this paper, Hoare suggests that input and output are two overlooked primitives of programming —
 particularly in concurrent code.*
 
 However, the shared memory model can be difficult to utilize correctly — especially in large or complicated programs. <br>
 It’s for this reason that concurrency is considered one of Go’s strengths: <br>
*it has been built from the start with principles from CSP in mind and therefore it is easy to read, write, and reason about.*

**Goroutines free us from having to think about our problem space in terms of parallelism and instead
  allow us to model problems closer to their natural level of concurrency.**
  
Channels
> for instance, are inherently composable with other channels. 
This makes writing large systems simpler because you can coordinate the input from multiple subsystems by easily composing the output together.

Select
> The select statement is the complement to Go's channels and is what enable all the difficult bits of composing channels.<br>
select statements allow you to efficiently wait for events, select a message from competing channels in a uniform random way, continue on if there are no messages waiting, and more.

## Go's Philosophy on Concurrency
>CSP was and is a large part of what Go was designed around;

Structs and methods in the sync and other packages allow you to perform locks, create pools of resources,preempt goroutine, and more.
    
    type Counter struct {
        mu sync.Mutex
        value int
    }
    func (c *Counter) Increment() {
        c.mu.Lock()
        defer c.mu.Unlock()
        c.value++
    }