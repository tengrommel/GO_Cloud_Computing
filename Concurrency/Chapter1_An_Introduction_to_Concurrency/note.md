# An Introduction to Concurrency
> When most people use the word “concurrent,” they’re usually referring to a process that occurs simultaneously with one or more processes. 

## Why is Concurrency Hard?
It usually takes a few iterations to get it working as expected,and even then it's not uncommon for bugs to exist in code for years before some change in timing (heavier disk utilization, more users logged into the system, etc.) causes a previously undiscovered bug to rear its head.

## Race Conditions
> A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.

Most of the time, data races are introduced because the developers are thinking about the problem sequentially. They assume that because a line of code falls before another that it will run first. 
<br>They assume the goroutine above will be scheduled and execute before the data variable is read in the if statement.

**When writing concurrent code, you have to meticulously iterate through the possible scenarios.**

Indeed, some developers fall into the trap of sprinkling sleeps throughout their code exactly because it seems to solve their concurrency problems. Let’s try that in the preceding program.

    var data int 
    go func() {data ++}()
    time.Sleep(1*time.Second) // This is bad!
    if data == 0{
        fmt.Printf("the value is %v.\n" data)
    }


## Atomicity
>When something is considered atomic, or to have the property of atomicity, this means that within the context that it is operating, it is indivisible, or uninterruptible.

1. The first thing that’s very important is the word “context.” Something may be atomic in one context, but not another. 
2. Operations that are atomic within the context of your process may not be atomic in the context of the operating system; 
3. operations that are atomic within the context of the operating system may not be atomic within the context of your machine;
 and operations that are atomic within the context of your machine may not be atomic within the context of your application. 
 
 *In other words, the atomicity of an operation can change depending on the currently defined scope.*
 This fact can work both for and against you!
 
 ## Memory Access Synchronization
 
    var data int
    go func() {data++}()
    if data == 0 {
        fmt.Println("the value is 0.")
    } else {
        fmt.Printf("The value is %v.\n", data)
    }
    
- Our goroutine, which is incrementing the data variables.
- Our if statement, which checks whether the value of data is 0.
- Out fmt.Printf statement, which retrieves the value of data for output.


    var memoryAccess sync.Mutex
    var value int
    go func() {
        memoryAccess.Lock()
        value++
        memoryAccess.Unlock()
    }()
    memoryAccess.Lock()
    if value == 0{
        fmt.Printf("the value is %v.\n", value)
    } else {
        fmt.Printf("the value is %v.\n", value)
    }
    memoryAccess.Unlock()

## Deadlock
> A deadlocked program is one in which all concurrent processes are waiting on one another.In this state, the program will never recover without outside intervention

