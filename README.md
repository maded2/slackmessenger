# slackmessenger

Integrate Monit with Slack - allow simple and small cli tool to send monit alerts to slack

Installation

```bash
go 
```

Typical usage - Monit config

```
check filesystem rootfs with path /sw
        if space usage > 85% then exec "/usr/local/bin/slackmsg /usr/local/etc/slack-url"
check process nginx with pidfile /var/run/nginx.pid
        if does not exist then exec "/usr/local/bin/slackmsg /usr/local/etc/slack-url"
```