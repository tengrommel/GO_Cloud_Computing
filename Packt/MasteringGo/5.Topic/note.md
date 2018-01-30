# Syncs and locks
- Lockers and mutexes
- Once
- WaitGroup

# Timers and tickers
- Durations
    - Number of nanoseconds between two time instances
    - Go's way to identify durations
    - Examples:
        - 3*time.Second => time duration of 3 seconds
        - 3*time.Minute => time duration of 3 minutes
- Timers
- Tickers

# Channel generators

## The Pipeline Pattern
- A concurrent pattern
- A series of 'stages' connected by channels
- Each stage is a group of goroutines running the same function

Hydra Chat System Pipeline

1. One client sends a message to the chat server for a specific chat room
2. The chat server sends the same message to all clients in the same chat room


# fan-in fan-out

1. Stage 1 (Producer): StartClient() will send data to the msgCh
2. Stage 2: broadcastMsg() will fan-out the message to multiple client channels
3. Use writemonitor to sed the message to client

# Contexts



    
# Operation Hydra: A chat application




# Go laws of reflection

- What is Reflection
    - Advanced feature that allows meta programming
    - Ability of a program to inspect itself
    - Identify types, and fields

- Refresher on interfaces
    - Specifies a behavior of an object
    - Any user defined type can implement an interface
    - An interface value stores the concrete value assigned to the variable, and the value type descriptor
    - The interface value only allows access to methods that interface implements
    - The empty interface is the parent of all types in Go

- First law of reflection
    - Reflection goes from interface value to reflection object
    - At a basic level, reflection examines the type and value pair stored inside an interface variable
    - reflect.TypeOf to get the type,reflect.ValueOf to get the value struct

- Second law of reflection
    - Reflection goes form reflection object to interface value
    - value.Interface() to get the interface value from the reflect object
    - Interface() is the inverse of reflect.ValueOf()

- Third law of reflection
    - To modify a reflection object,the value must be settable
    - value.CanSet() to inspect if the reflect object value is settable


# Operation Hydra: Create new file format


