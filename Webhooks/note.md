## Webhooks

- User-defined HTTP callback
- Trigger remote endpoint behaviour (e.g.web application)
- Web technology
> Generally REST API<br>
Methods: POST,GET;Encodings(JSON, form encoding)
- Ingredients
> Event specification(sender side)<br>
URL to 'hook'(recipient side)<br>
Payload format specification

## API vs. Webhooks

- API
    > Interation Pattern: Polling<br>
    Server provides functionality<br>
    Client polls repeatedly

    > Problems:<br>
        Rate limits<br>
        Connection limits<br>
        Timeouts

- Webhooks    
    > Interaction Pattern: Publish-subscribe<br>
    Client lodges callback(destination, event, ...)<br>
    Server only contacts client upon event condition
    