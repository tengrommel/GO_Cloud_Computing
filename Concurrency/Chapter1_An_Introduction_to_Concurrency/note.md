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

*If that sounds grim, it’s because it is!<br>
 The Go runtime attempts to do its part and will detect some deadlocks (all goroutines must be blocked, or “asleep”), but this doesn’t do much to help you prevent deadlocks.*
 
    type value struct{
        mu sync.Mutex
        value int
    }
    
    var wg sync.WaitGroup
    printSum := func(v1, v2 *value) {
        defer wg.Done()
        v1.mu.Lock()
        defer v1.mu.Unlock()
        
        time.Sleep(2*time.Second)
        v2.mu.Lock()
        defer v2.mu.Unlock()
        
        fmt.Printf("sum=%v\n", v1.value + v2.value)
    }
    
    var a, b value
    wg.Add(2)
    go printSum(&a, &b)
    go printSum(&b, &a)
    wg.Wait()

## Livelock
> Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward.

    cadence := sync.NewCond(&sync.Mutex{})
    go func() {
        for range time.Tick(1*time.Millisecond) {
            cadence.Broadcast()
        }
    }
    takeStep := func() {
        cadence.L.Lock()
        cadence.Wait()
        cadence.L.Unlock()
    }
    tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
        fmt.Fprintf(out, " %v", dirName)
        atomic.AddInt32(dir, 1)
        takeStep()
        if atomic.LoadInt32(dir) == 1 {
            fmt.Fprint(out, ". Success!")
            return true
        }
        takeStep()
        atomic.AddInt32(dir, -1)
        return false
    }
    var left, right int32
    tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
    tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }
    
## Starvation
> Starvation is any situation where a concurrent process cannot get all the resources it needs to work.
