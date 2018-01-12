# Chapter 6 Managing character sets and encodings
> There are many languages in use throughout the world, and they use many different character sets.<br>
There are also many ways of encoding character sets into binary formats of bytes.

## Transport encoding
> A character encoding will suffice for handling characters within a single application.<br>
However, once you start sending text between applications, then there is the further issue of how the bytes, shorts or words are put on the wire.<br>
An encoding can be based on space-and hence bandwidth-saving techniques such as zip ping the text.<br>
Or it could be reduced to a 7-bit format to allow a parity checking bit, such as base64.