# DAVDiscover - a simple, configurable autodiscover service for distributed and self-hosted groupware services.
## What is DAVDiscover?
DAVDiscover is a self-hosted service which provides IMAP, SMTP, CalDAV, CardDav, etc
URLs to authenticated clients for ease in set up on mobile devices and also desktop/laptop computers.  

## What ~~does~~ will DAVDiscover do?
DAVDiscover is currently in _very_ early stages, with a _very_ limited set of features.
My hope for DAVDiscover is to mimic the AutoDiscover service found in Microsoft Exchange services,
but with the intent of providing a set of URLS for each service which may be self-hosted and/or distributed,
primary IMAP, SMTP, CalDAV and CardDAV URLS, but hopefully more services can be added in the future.  

__For example:__  
Company X may have many services that are selfhosted i.e. Dovecot for email (IMAP/SMTP) and Nextcloud for Calendar and Contacts.
Getting a set up like this configured on a mobile device is fairly involved for a standard user and tiresome if the company has many employees.
There are many points where set up configuration mistakes can happen, leading to service outage for a user, and the difficult job of
the IT consultant trying to talk the user through setting the device up over the phone.

DAVDiscover intends to patch this problem by providing the URLS and information (Port numbers, SSL/TLS type, domain name, etc.) for each service
in JSON format allowing for the connecting device to automatically set up this information on the device.

## What problems do I expect?
OK, the most glaringly obvious problem is that the JSON response information is no good if the device itself doesn't know what to do with it.  
I hope to develop a mobile app with mail, contact and calendar features in it, but we're in _very_ early stages of the base service at the moment.
Once the basic JSON response service is operational, I'll look into the mobile app which will then be when things will start to become useful.  

In the perfect world, this service starts to look so fantastic that mobile device Operating System and/or mail app developers start implementing the option
in their own code as an Account Type (i.e. ActiveSync, Office365, iCloud, IMAP, etc. are all already there), but I'm not sure if I see that happening just yet.  

### Another problem is authentication.
I'd like all clients to authenticate to the DAVDiscover service, but where do we get that authentication from?  
We could have local accounts on the DAVDiscover server obviously, but I don't think this feels "fluent" enough. Maybe, using the primary IMAP server address, we could do an authentication request and if that succeeds the login is accepted and DAVDiscover information is sent.   

### More problems regarding authentication.
As we're intended to be used for self-hosted, distributed services, each service may have different usernames and passwords.
This means there will have to be some sort of manual credential entry for each service.  
I am less concerned with this issue as it currently isn't really in the scope of DAVDiscover to handle this - the goal is to provide the core information (URL, Ports, etc.) not the credentials to log in. I of course want to make the experience as helpful as possible though, so I'll deal with any features surrounding this when I can.

## When will it be ready for production?
Well, not yet.  
I'm working on this ultimately for my own use for my own small business. I'm hoping once it's good enough I could deploy the set up to customers
and ultimately get them away from a Microsoft Exchange based environment. There's a long way to go for that right now though.  

If you feel you may be able to help, or ideas on features and their implementation, notice any bugs, or just want to say hi. Please do so and submit a pull request if required.

Thanks for reading!