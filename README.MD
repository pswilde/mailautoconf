# MailAutoConf - a simple, configurable autodiscover/autoconfig service for distributed and self-hosted services.  

## New GoLang version - please make sure you update your ini files to yaml!
Github : https://github.com/pswilde/mailautoconf  
  
<a href="https://www.buymeacoffee.com/pswilde" target="_blank">
  <img src="https://cdn.buymeacoffee.com/buttons/v2/default-blue.png" height="60" width="217" alt="Buy Me A Coffee" style="height: 30px !important;width: 106px !important;" >
</a>  
  
## What is MailAutoConf?
MailAutoConf is autodiscover/autoconfig web server for self-hosted mail services
which do not have their own autodiscover service.

## What does MailAutoConf do?
MailAutoConf is currently in _very_ early stages, but it does generate valid
autoconfig XML files (/mail/config-v1.1.xml) compatible with many mail clients
i.e. Thunderbird, Evolution, etc.  
Theoretically, anything that can read the standard autoconfig XML file -
https://wiki.mozilla.org/Thunderbird:Autoconfiguration:ConfigFileFormat, should
be able to auto-configure using this service.

## Installation
MailAutoConf runs its own webserver on port 8010.
You will need to supply a volume for the configuration file and port forwarding.
```
podman run -dt \
  --name mailautoconf \
  -v ./config:/mailautoconf/config \
  -p 8010:8010 \
  pswilde/mailautoconf
```
You will need a reverse proxy server to publish to the outside world and handle SSL encryption.  
For example, in nginx:
```
server {
  listen                443 ssl;
  server_name           autoconfig.example.com;
  ssl_certificate       /path/to/certificate.file;
  ssl_certificate_key   /path/to/certificate.key
  location / {
    proxy_set_header    X-Forwarded-Host $host:$server_port;
    proxy_set_header    X-Forwarded-Server $host;
    proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass          http://localhost:8010/;
    proxy_redirect      http://localhsot:8010/ /;
    proxy_read_timeout  60s;
  }
}
```  
First run will create sample.yaml files in the config directory. Copy these to config.yaml and services.yaml and configure them to your needs.  

MailAutoConf will handle all the URLs it's able to deal with, i.e. /mail/config-v1.1.xml, /Autodiscover/Autodiscover.xml automatically.  

## DNS config
You will need to configure DNS records for your domain.  
I would recommend setting up an A record for autoconfig.your.domain to point
directly to your MailAutoConf's instance which will work for Thunderbird and
Evolution clients.
Then, if Outlook Autodiscover is required (currently not implemented), configure
a SRV record to redirect to the autoconfig.your.domain A record. i.e.
```         
Type                                TTL  P  M  Port Target
SRV _autodiscover._tcp.your.domain  3600 10 10 443  autoconfig.your.domain
```

## Compatibility  
MailAutoConf has been tested and confirmed working (for IMAP and SMTP) with the following software packages  
- [x] Thunderbird (v78 and probably earlier versions too)
- [x] Evolution Mail (v3.40.3 and probably earlier versions too)
- [x] Nextcloud Mail app
- [ ] Other Mail Clients are likely supported if they support /mail/config-v1.1.xml



## Future plans
MailAutoConf is currently in _very_ early stages, with a _very_ limited set of features.
My hope for MailAutoConf is to run as an autoconfig service which provides Calendar and AddressBook locations as well.
Calendar and AddressBook is in the autoconfig XML documentation, but currently not implemented. MailAutoConf can generate the correct information for Calendar and AddressBook into the autoconfig XML file, though no clients currently use it.

## TODO:

- [x] /mail/config-v1.1.xml file for Thunderbird and Evolution mail clients.
- [ ] Add more code commenting, it's pretty terrible right now.
- [ ] Refactor the autoconfig and autodiscover XML files, they've been copied and bodged together and could look so much better.
- [ ] Autodiscover.xml for Outlook clients - currently does generate, but untested.
- [ ] Autodiscover.json for more modern Outlook clients - this is currently undocumented. I have implemented some features of it but it is currently not working.
- [ ] Calendars and Address Book - does generate into mail/config-v1.1.xml, but no actual clients support it yet.
- [ ] Additional, custom services. For example, Nextcloud - instead of having to type in your Nextcloud instance URL, you can just enter your email address associated with you Nextcloud account and it looks up the URL from autoconfig - wouldn't that be nice?
- [ ] Probably loads of stuff I haven't thought of yet.


## When will it be ready for production?
It works for non-Microsoft email clients now (see Compatibility above).  
Outlook's autodiscover is a troublesome little blighter, MailAutoConf does generate a valid Autodiscover.xml, but modern Outlook clients use an Autodiscover.json file now which isn't documented anywhere. I'm working on this and hope to get Outlook Compatibility as soon as possible.  

Then it's down to Autoconfiguration of Calendars and AddressBooks... but that's down to the email client developers really...  


If you feel you may be able to help, or ideas on features and their implementation, notice any bugs, or just want to say hi. Please do so and submit a pull request if required.

Thanks for reading!
