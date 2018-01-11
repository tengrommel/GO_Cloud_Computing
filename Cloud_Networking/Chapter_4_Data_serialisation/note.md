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

## JSON
> JSON stands for JavaScript Object Notation.<br>
 It was designed to be a lightweight means of passing data between JavaScript systems.<br>
 It uses a text-based format and is sufficiently general that it has become used as a general purpose serialisation method for many programing languages.

From the Go JSON package specification, marshalling uses the following type-dependent default encodings:

- Boolean values encode as JSON booleans.
- Floating point and integer values encode as JSON numbers.
- String values encode as JSON strings, with each invalid UTF-8 sequence replaced by the encoding of the Unicode replacement character U+FFFD.
- Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string.
- Struct values encode as JSON objects. Each struct field becomes a member of the object. By default the object's key name is the struct field name converted to lower case. <br>
If the struct field has a tag, that tag will be used as the name instead.
- Map values encode as JSON objects. The map's key type must be string; the object keys are used directly as map keys.
- Pointer values encode as the value pointed to. (Note: this allows trees, but not graphs!). A nil pointer encodes as the null JSON object.
- Interface values encode as the value contained in the interface. A nil interface value encodes as the null JSON object.
- Channel, complex, and function values cannot be encoded in JSON. Attempting to encode such a value cause Marshal to return an InvalidTypeError.
- JSON cannot represent cyclic data structures and Marshal does not handle them. Passing cyclic structures to Marshal will result in an infinite recursion.