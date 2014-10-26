cf-ssh
======

SSH into a running container for your Cloud Foundry application, run one-off tasks, debug your app, and more.

Initial implementation requires the application to have a `manifest.yml`.

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
