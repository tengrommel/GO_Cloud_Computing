# Chapter 7 Security

## Data integrity
1. Ensuring data integrity means supplying a means of testing that the data has not been tampered with.
2. Usually this is done by forming a simple number out of the bytes in the data.
3. This process is called hashing and the resulting numbers is called a hash or hash value.

A naive hashing algorithm is just to sum up all the bytes in the data.

**Go has support for several hashing algorithms, including MD4, MD5, RIPEMD-160, SHA1, SHA224, SHA256, SHA384 and SHA512. 
<br>They all follow the same pattern as far as the Go programmer is concerned: a function New (or similar) in the appropriate package returns a Hash object from the hash package.**