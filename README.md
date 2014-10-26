cf-ssh
======

SSH into a running container for your Cloud Foundry application, run one-off tasks, debug your app, and more.

Initial implementation requires:

-	run `cf-ssh -f path/to/manifest.yml` to obtain information about the application

Installation
------------

```
go get -u github.com/cloudfoundry-community/cf-ssh
```

Usage
-----

```
cd path/to/app
cf-ssh -f manifest.yml
```
