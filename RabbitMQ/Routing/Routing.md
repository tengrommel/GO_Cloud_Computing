# Routing
>In this tutorial we're going to add a feature to it we're going to make it possible to subscribe
only to a subset of the messages.

For example, we will be able to direct only critical error messages to log file(to save disk space), 
while still being able to print all of the log messages on the console.

## Bindings
    err = ch.QueueBind(
        q.Name,
        "",
        "logs",
        false,
        nil)

*A binding is a relationship between an exchange and a queue. This can be simply read as: the queue is interested in messages from this exchange.*

Bindings can take an extra routing_key parameter. 
To avoid the confusion with a Channel.
Publish parameter we're going to call it a binding key. 
    
    err = ch.QueueBind(
        q.Name,    // queue name
        "black",   // routing key
        "logs",    // exchange
        false,
        ni)

**The meaning of a binding key depends on the exchange type. The fanout exchanges, which we used previously, simply ignored its value.**

## Direct exchange
>Our logging system from the previous tutorial broadcasts all messages to all consumers.

We want to extend that to allow filtering messages based on their severity. <br>
For example we may want the script which is writing log messages to the disk to only receive critical errors,<br>
and not waste disk space on warning or info log messages.

We will use a direct exchange instead.The routing algorithm behind a direct exchange is simple - a message goes to
the queues whose binding key exactly matches the routing key of the message.

It is perfectly legal to bind multiple queues with the same binding key. <br>
In our example we could add a binding between X and Q1 with binding key black. <br>
In that case, the direct exchange will behave like fanout and will broadcast the message to all the matching queues. <br>
A message with routing key black will be delivered to both Q1 and Q2.

