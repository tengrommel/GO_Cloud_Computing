# Chapter 13 Remote Procedure Call

>Socket and HTTP programming use a message-passing paradigm.<br>
A client sends a message to a server which usually sends a message back. <br>
Both sides ae responsible for creating messages in a format understood by both sides, and in reading the data out of those messages.

However, most standalone applications do not make so much use of message of message passing techniques.
Generally the preferred mechanism is that of the function(or method or procedure) call. 


**In this style, a program will call a function with a list of parameters, and on completion of the function call will have a set of return values.
These values may be the function value, or if addresses have been passed as parameters then the contents of those addresses might have been changed.**

1. The remote procedure call is an attempt to bring this style of programming into the network world.
2. Thus a client will make what looks to it like a normal procedure call.
3. The client-side will package this into a procedure call on the server side.
4. The results of this call will be packaged up for return to the client.

## Go RPC

In Go, the restriction is that:

- the function must be public(begin with a capital letter)
- have exactly two arguments, the first is a pointer to be received by function from the client, and the second is a pointer to hold the answers to be returned to the client.
- have a return value of type error