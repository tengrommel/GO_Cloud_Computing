# Chapter 4 Data serialisation
> Communication between a client and a service requires the exchange of data.<br>
This data may be highly structured, but has to be serialised for transport.<br>

## Introduction
>A client and server need to exchange information via messages. TCP and UDP provide the transport mechanisms to do this. <br>
The two processes also have to have a protocol in place so that message exchange can take place meaningfully.

*Messages are sent across the network as a sequence of bytes, which has no structure except for a linear stream of bytes.*

A program will typically build complex data structures to hold the current program state.

Programming languages use structured data such as

- records/structures
- variant records
- array - fixed size or varying
- string - fixed size or varying
- tables - e.g.arrays of records
- non-linear structures such as 
    1. circular linked list
    2. binary tree
    3. objects with references to other objects
    
## Mutual agreement
> The previous section gave an overview of the issue of data serialisation.

*The above serialisation is opaque or implicit. If data is marshalled using the above format, then there is nothing in the serialised data to say how it should be unmarshalled.*

## ASN.1

>Abstract Syntax Notation One (ASN.1) was originally designed in 1984 for the telecommunications industry.

