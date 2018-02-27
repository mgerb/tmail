# TMail
A throwaway smtp mail server with an API to retrieve emails.

- set up an MX record on your domain.
- point it at your server
- start tmail
- all emails sent to *@<your domain> are stored in `mail.db`
- hit these end points to check the email

[Download the latest release here.](https://github.com/mgerb/tmail/releases)

```
All mail:
http://host:8090/api/mail

Specific mail:
http://host:8090/api/mail?to=<address>
```
