# slackmessenger

Integrate Monit with Slack - allow simple and small cli tool to send monit alerts to slack

### Installation

```bash
go get github.com/maded2/slackmessenger/...
```

### Typical usage 

your slack Incoming WebHooks URL string should be stored in a plain text file somewhere secured on your server.
 
slackmsg <slack-url-file>


#### Monit config

```
check filesystem rootfs with path /sw
        if space usage > 85% then exec "/usr/local/bin/slackmsg /usr/local/etc/slack-url"
check process nginx with pidfile /var/run/nginx.pid
        if does not exist then exec "/usr/local/bin/slackmsg /usr/local/etc/slack-url"
```

### Example Alert in Slack


