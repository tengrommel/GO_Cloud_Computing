# Remote procedure call(RPC)
>We need to run a function on a remote computer and wait for the result?

In this tutorial we're going to use RabbitMQ to build an RPC system:
> a client and a scalable RPC server.

As we don't have any time-consuming tasks that are worth distributing,we're going to create a dummy RPC service that returns Fibonacci numbers.

## Callback queue
> A client sends a request message and a server replies with a response message.

In order to receive a response we need to send a 'callback' queue address with the request.

## Correlation Id
>In the method presented above we suggest creating a callback queue for every RPC request.

That raises a new issue, having received a response in that queue it's not clear to which request the response belongs. 
1. That's when the correlation_id property is used. 
We're going to set it to a unique value for every request.

2. Later, when we receive a message in the callback queue we'll look at this property, 
and based on that we'll be able to match a response with a request. 

3. If we see an unknown correlation_id value,
we may safely discard the message - it doesn't belong to our requests.

*It's due to a possibility of a race condition on the server side. 
Although unlikely, it is possible that the RPC server will die just after sending us the answer, 
but before sending an acknowledgment message for the request. 
If that happens, the restarted RPC server will process the request again. 
That's why on the client we must handle the duplicate responses gracefully, 
and the RPC should ideally be idempotent.*

Our RPC will work like this:

1. When the Client starts up,it creates an anonymous exclusive callback queue.
2. For an RPC request,the Client sends a message with two properties:
>reply_to,which is set to the callback queue and correlation_id,which is set to a unique value for every request.
3. The RPC worker (aka: server) is waiting for requests on that queue.
> When a request appears, it does the job and sends a message with the result back to the Client, 
using the queue from the reply_to field.
4. The client waits for data on the callback queue.
> When a message appears,it checks the correlation_id property.<br>
If it matches the value from the request it returns the response to the application.

