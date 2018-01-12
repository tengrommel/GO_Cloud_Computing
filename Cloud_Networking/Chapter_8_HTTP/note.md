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
    
### Authenticating proxy
> Some proxies will require authentication, by a user name and password in order to pass requests.<br>
A common scheme is "basic authentication" in which the user name and password are concatenated into a string "user:password" and then BASE64 encoded.<br> 
This is then given to the proxy by the HTTP request header "Proxy-Authorisation" with the flag that it is the basic authentication
                                                                 
## HTTPS connections by clients
>For secure, encrypted connections, HTTP uses TLS which is described in the chapter on security.<br>
The protocol of HTTP_TLS is called HTTPS and uses https:// urls instead of http:// urls.

*Servers are required to return valid X.509 certificates before a client will accept data from them.<br>
If the certificate is valid, then Go handles everything under the hood and the clients given previously run okay with https URLs.*

*Many sites have invalid certificates. <br>
They may have expired, they may be self-signed instead of by a recognised Certificate Authority or they may just have errors (such as having an incorrect server name). <br>
Browsers such as Firefox put a big warning notice with a "Get me out of here!" button, but you can carry on at your risk - which many people do.*

*Go presently bails out when it encounters certificate errors. There is cautious support for carrying on but I haven't got it working yet. <br>
So there is no current example for "carrying on in the face of adversity :-)". Maybe later.*

## Servers   
>The other side to building a client is a Web server handling HTTP requests. <br>
The simplest - and earliest - servers just returned copies of files.<br>
However, any URL can now trigger an arbitrary computation in current servers.
   
### File server
   
   We start with a basic file server. 
   >Go supplies a multi-plexer, that is, an object that will read and interpret requests. <br>
   It hands out requests to handlers which run in their own thread. 
   <br>Thus much of the work of reading HTTP requests, decoding them and branching to suitable functions in their own thread is done for us.
   
   *For a file server, Go also gives a FileServer object which knows how to deliver files from the local file system.<br>
    It takes a "root" directory which is the top of a file tree in the local system, and a pattern to match URLs against.<br> 
    The simplest pattern is "/" which is the top of any URL. This will match all URLs.*
   
   An HTTP server delivering files from the local file system is almost embarrassingly trivial given these objects.
