# Topics
*主题*
>Instead of using a fanout exchange only capable of dummy broadcasting,we used a direct one,
and gained a possibility of selectively receiving the logs.

Although using the direct exchange improved our system,
it still has limitations - it can't do routing based on multiple criteria.

In our logging system we might want to subscribe to not only logs based on severity, 
but also based on the source which emitted the log. 
You might know this concept from the syslog unix tool, 
which routes logs based on both severity (info/warn/crit...) and facility (auth/cron/kern...).

That would give us a lot of flexibility - we may want to listen to just critical errors coming from 'cron' but also all logs from 'kern'.

To implement that in our logging system we need to learn about a more complex topic exchange.

## Topic exchange
>Messages sent to a topic exchange can't have an arbitrary routing_key - it must be a list of words,delimited by dots.

The words can be anything,but usually they specify some features connected to the message.
1. A few valid routing key examples:"stock.usd.nyse","nyse.vmw","quick.orange.rabbit".
2. There can be as many words in the routing key as you like,up to the limit of 255 bytes.

The binding key must also be in the same form. 
The logic behind the topic exchange is similar to a direct one - a message sent with a particular 
routing key will be delivered to all the queues that are bound with a matching binding key. 
However there are two important special cases for binding keys:
    - *(star)can substitute for exactly one word.
    - #(hash)can substitute for zero or more words.

