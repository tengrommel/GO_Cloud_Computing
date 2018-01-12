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

## Version control
> A protocol used in a client/server system will evolve over time, changing as the system expands.<br>
This raises compatibility problems: a version 2 client will make requests that a version 1 server doesn't understand,<br>
whereas a version 2 server will send replies that a version 1 client won't understand.

**Ech side should ideally be able to understand messages for its own version and all earlier ones.<br>
It should be able to write replies to old style queries in old style response format**
> The ability to talk earlier version formats may be lost if the protocol changes too much.<br>
In this case, you need to be able to ensure that no copies of the earlier version still exist - and that is generally impossible.

## The Web
>The Web is a good example of a system that is messed up by different versions. The protocol has been through three versions, and most servers/browsers now use the latest version. 

## Message Format

1. The client and server will exchange messages with different meanings,e.g.
    - login request
    - get record request
    - login reply
    - record data reply
2. The client will prepare a request which must be understood by the server.
3. The server will prepare a reply which must be understood by the client.

Commonly, the first part of the message will be a message type.

- Client to server
    
    
    LOGIN name passwd
    GET cpe4001 grade

- Server to client

    
    LOGIN succeeded
    GRADE cpe4001 D
    
## Data Format
> There are two main format choices for messages: byte encoded or character encoded.

1. Byte format

- the first part of the message is typically a byte to distinguish between message types.
- The message handler would examine this first byte to distinguish message type and then perform a switch to select the appropriate handler for that type.
- Further bytes in the message would contain message content according to a pre-defined format.

*The advantages are compactness and hence speed. <br>
The disadvantages are caused by the opaqueness of the data: it may be harder to spot errors, harder to debug, require special purpose decoding functions.<br>
There are many examples of byte-encoded formats, including major protocols such as DNS and NFS, upto recent ones such as Skype. <br>
Of course, if your protocol is not publicly specified, then a byte format can also make it harder for others to reverse-engineer it!*

**Pseudocode for a byte-format server is**

    handleClient(conn) {
        while (true) {
            byte b = conn.readByte()
            switch (b) {
                case MSG_1: ...
                case MSG_2: ...
                ...
            }
        }
    }

Go has basic support for managing byte streams.The interface Conn has methods

    func (c Conn) Read(b []byte) (n int, err os.Error)
    func (c Conn) Write(b []byte) (n int, err os.Error)
**and these methods are implemented by TCPConn and UDPConn.**    

## Character Format
> In this mode, everything is sent as characters if possible.

In character format:
- A message is a sequence of one or more lines the start of the first line of the message is typically a word that represents the message type.
- String handling functions may be used to decode the message type and data.
- The rest of the first line and successive lines contain the data.
- Line-oriented functions and line-oriented conventions are used to manage this.

Pseudocode is 

    handleClient() {
        line = conn.readLine()
        if (line.startsWith(...) {
            ...
        } else if (line.startsWith(...) {
            ...
        }
    }
    
*If we just pretend everything is ASCII, like it was once upon a time, then character formats are quite straightforward to deal with. The principal complication at this level is the varying status of "newline" across different operating systems. Unix uses the single character "\n". Windows and others (more correctly) use the pair "\r\n". On the internet, the pair "\r\n" is most common - Unix systems just need to take care that they don't assume "\n".*