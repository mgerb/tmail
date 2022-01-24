# TMail
A throwaway smtp mail server with an API to retrieve emails.

[Download the latest release here.](https://github.com/mgerb/tmail/releases)

- set up an MX record on your domain.
- point it at your server
- start tmail
- all emails sent to *@<your domain> are stored in `mail.db`
- hit these end points to check the email

```
All mail:
http://host:8090/api/mail

Specific mail:
http://host:8090/api/mail?to=<address>
```
  
## Server Configuration

[See this comment for additional information on configuring a server for tmail](https://github.com/mgerb/tmail/issues/2#issue-1112156904) - thanks @mrbluecoat
  
