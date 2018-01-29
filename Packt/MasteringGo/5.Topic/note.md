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



# Operation Hydra: Create new file format


