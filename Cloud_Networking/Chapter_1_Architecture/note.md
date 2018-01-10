# Chapter 1 Architecture
> This chapter covers the high level architectural aspects of distributed systems.<br>
There are many ways of looking at such systems, and many of these are dealt with.

## Protocol Layers
> Distributed systems are hard. There are multiple computers involved, which have to be connected in some way.<br>
 *Programs have to be written to run on each computer in the system and they all have to co-operate to get a distributed task done.*

 - The common way to deal with complexity is to break it down into smaller and simpler parts.
 > These parts have their own structure,but they also have defined means of communicating with other related parts.

 ## TCP/IP Protocol
> There are many protocols occupying significant niches, such as<br> 
<tab>*Firewire*<br>
<tab>*USB*<br>
<tab>*Bluetooth*<br>
<tab>*WiFi*

## Networking
> A network is a communications system for connecting end systems called hosts.

## Gateways
> A gateway is a generic term for an entity used to connect two or more networks.<br>

## Packet encapsulation
> Each layer has administrative information that it has to keep about its own layer.

## Connection Models
> In order for two computers to communicate, they must set up a path whereby they can send at least one message in a session.

- Connection oriented
> A single connection is established for the session.<br>
 Two-way communications flow along the connection.<br>
 When the session is over, the connection is broken.<br> 
 The analogy is to a phone conversation.<br>
 An example is TCP

- Connectionless
>In a connectionless system, messages are sent independent of each other. <br>Ordinary mail is the analogy. Connectionless messages may arrive out of order. <br>An example is the IP protocol.


*Connection oriented transports may be established on top of connectionless ones - TCP over IP. Connectionless transports may be established on top of connection oriented ones - HTTP over TCP.*

**There can be variations on these. For example, a session might enforce messages arriving, but might not guarantee that they arrive in the order sent. However, these two are the most common.**


# Communications Models

## Message passing
>Message passing is a primitive mechanism for distributed systems.<br>
 Set up a connection and pump some data down it. <br>
 At the other end, figure out what the message was and respond to it, possibly sending messages back.

 *Higher level event driven systems assume that this decoding has been done by the underlying system and the event is then dispatched to an appropriate object such as a ButtonPress handler.*

 *This can also be done in distributed message passing systems, whereby a message received across the network is partly decoded and dispatched to an appropriate handler.*

## Remote procedure call
>In any system,there is a transfer of information and flow control from one part of the system to another. <br>
In procedural languages this may consist of the procedure call, where information is placed on a call stack and then control flow is transferred to another part of the program.

## Distributed Computing Models
>At the highest level,we could consider the equivalence of the non-equivalence of components of a distributed system.

*The most common occurrence is an asymmetric one: a client sends requests to a server, and the server responds. This is a client-server system.*
- If both components are equivalent, both able to initiate and to respond to messages, then we have a peer-to-peer system.
- A third model is the so-called filter. Here one component passes information to another which modifies it before passing it to a third. 

## Server Distribution
> the master receives requests and instead of handling them one at a time itself, passes them off to other servers to handle. <br>
This is a common model when concurrent clients are possible.

## Component Distribution

A simple but effective way of decomposing many applications is to consider them as made up of three parts:

- Presentation component
> The presentation component is responsible for interactions with the user, both displaying data and gathering input.<br>
 It may be a modern GUI interface with buttons, lists, menus, etc., or an older command-line style interface, asking questions and getting answers.<br> 
 The details are not important at this level.
- Application logic
> The application logic is responsible for interpreting the users' response,for applying business rules,for preparing queries and managing responses from the their component.
- Data access
> The data access component is responsible for storing and retrieving data. This will often be through a database, but not necessarily.

## Gartner Classification

- distributed data
- remote data
- distribute transaction
- remote presentation
- distribute presentation

## Middleware

Components of middleware include
- The network services include things like TCP/IP
- The middleware layer is application-independent s/w using the network services.
- Examples of middleware are: DCE,RPC,Corba
- Middleware many only perform one function (such as RPC) or many (such as DCE)

Examples of middleware include
- Primitive services such as terminal emulators, file transfer, email
- Basic services such as RPC
- Integrated services such as DCE,Network O/S
- Distributed object services such as CORBA,OLE/ActiveX
- Mobile object services such as RMI,Jini
- World Wide Web

The functions of middleware include
- Initiation of processes at different computers
- Session management
- Directory services to allow clients to locate servers
- Remote data access
- Concurrency control to allow servers to handle multiple clients
- Security and integrity
- Monitoring
- Termination of processes both local and remote

## Points of Failure
> Distributed applications run in a complex environment.<br>
This makes them much more prone to failure than standalone applications on a single computer. <br>
The points of failure include:
- The client side of the application could crash
- The client system may have h/w problems
- The client's network card could fail
- Network contention could cause timeouts
- There may be network address conflicts
- Network elements such as routers could fail
- Transmission errors may lose messages
- The client and server versions may be incompatible
- The server's network card could fail
- The server system may have h/w problems
- The server s/w may crash
- The server's database may become corrupted

## Acceptance Factors
- Reliability
- Performance
- Responsiveness
- Scalability
- Capacity
- Security

## Transparency
- access transparency
- location transparency
- migration transparency
- replication transparency
- concurrency transparency
- scalability transparency
- performance transparency
- failure transparency

## Eight fallacies of distributed computing

- The network is reliable.
- Latency is zero.
- Bandwidth is infinite.
- The network is secure.
- Topology don't change.
- There is one administrator.
- Transport cost is zero.
- The network is homogeneous.