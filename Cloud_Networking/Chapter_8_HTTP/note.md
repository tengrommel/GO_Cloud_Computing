# Chapter 8 HTTP

- The World Wide Web is a major distributed system, with millions of users. 
    A site may become a Web host by running an HTTP server.
 While Web clients are typically users with a browser, there are many other "user agents" such as web spiders, web application clients and so on.
- The Web is built on top of the HTTP (Hyper-Text Transport Protocol) which is layered on top of TCP.
 HTTP has been through three publicly available versions, but the latest - version 1.1 - is now the most commonly used.

## URLs and resources
>URLs specify the location of a resource.

## Configuring HTTP requests
> Go also supplies a lower-level interface for user agents to communicate with HTTP servers.<br>
As you might expect, not only does it give you more control over the client requests, but require you to spend more effort in building the requests. <br>
However, there is only a small increase.

-  The simplest way to create a request with default values is by for example
    
    
    request, err := http.NewRequest("GET", url.String(), nil)
    

-  To specify that you only wish to receive UTF-8, add an "Accept-Charset" field to a request by


    request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
    
    
## The client object
> To send a request to a server and get a reply, the convenience object Client is the easiest way.

*This object can manage multiple requests and will look after issues such as whether the server keeps the TCP connection alive and so on.*

## Proxy handling

### Simple proxy
>HTTP 1.1 laid out how HTTP should work through a proxy. A "GET" request should be made to a proxy.<br>
However, the URL requested should be the full URL of the destination.<br>
In addition the HTTP header should contain a "Host" field, set to the proxy.<br>
As long as the proxy is configured to pass such requests through, then that is all that needs to be done.

Go considers this to be part of the HTTP transport layer.To manage this it has a class Transport.<br>
This contains a field which can be set to a function that returns a URL for a proxy.<br>
If we have a URL as a string for the proxy, the appropriate transport object is created and then given to a client object by

    proxyURL, err := url.Parse(proxyString)
    transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
    client := &http.Client{Transport: transport}