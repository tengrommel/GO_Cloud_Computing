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
