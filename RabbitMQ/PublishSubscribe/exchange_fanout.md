# emit_log
> The assumption behind a work queue is that each task is delivered to exactly one worker.

To illustrate the pattern, we're going to build a simple logging system.
> In our logging system every running copy of the receiver program wil get the messages.

That way we'll be able to run one receiver and direct the logs to disk; 
and at the same time we'll be able to run another receiver and see the logs on the screen.
 
Essentially, published log messages are going to be broadcast to all the receivers.

## Exchanges

> Queue
1. A producer is a user application that sends messages.
2. A queue is a buffer that stores messages.
3. A consumer is a user application that receives messages.

> The core idea in the messaging model in RabbitMQ is that the producer never sends any messages directly to a queue.
Actually, quite often the producer doesn't even know if a message will be delivered to any queue at all.

**Instead,the producer can only send messages to an exchange.
An exchange is a very simple thing.On one side it receives thing.
On one side it receives messages from producters and the other side it pushes them to queues.
The exchange must know exactly what to do with a message it receives.**


There are a few exchange types available: direct, topic, headers and fanout.

*The fanout exchange is very simple.As you can probably guess from the name,it just broadcasts 
all the messages it receives to all the queues it knows.And that's exactly what we need for our logger.*

