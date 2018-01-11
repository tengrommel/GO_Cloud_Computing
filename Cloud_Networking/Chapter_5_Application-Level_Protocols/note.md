# Chapter 5 Application-Level Protocols
> A client and server need to exchange information via messages. TCP and UDP provide the transport mechanisms to do this. The two processes also need to have a protocol in place so that message exchange can take place meaningfully. A protocol defines what type of conversation can take place between two components of a distributed application, by specifying messages, data types, encoding formats and so on.

## Protocol Design
> There are many possibilities and issues to be decided on when designing a protocol.

- Is it to be broadcast or point to point?Broadcast must be UDP, local multicast or the more experimental MBONE.- In order to handle masking operations, there is the type.Point to point could be either TCP or UDP.
- Is it to be stateful vs stateless?Is it reasonable for one side to maintain state about the other side? It is often simple to do so, but what happens if something crashes?
- Is the transport protocol reliable or unreliable? Reliable is often slower, but then you don't have to worry so much about loss messages.
- Are replies needed? If a reply is needed, how do you handle a lost reply? Timeouts may be used.
- What data format do you want? Two common possibilities are MIME or byte encoding.
- Is your communication bursty or steady stream?
> Ethernet and the Internet are best at bursty traffic. Steady stream is needed for video streams and particularly for voice. If required, how do you manage Quality of Service (QoS)?
- Are there multiple streams with synchronisation required?
- Are you building a standalone application or a library to be used by others? The standards of documentation required might vary.