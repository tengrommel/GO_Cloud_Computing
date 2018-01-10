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