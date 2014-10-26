cf-ssh
======

SSH into a running container for your Cloud Foundry application, run one-off tasks, debug your app, and more.

Initial implementation requires the application to have a `manifest.yml`.

It is desired that `cf-ssh` works correctly from all platforms that support the `cf` CLI.

Windows is a target platform but has not yet been tested. Please give feedback in the [Issues](https://github.com/cloudfoundry-community/cf-ssh/issues).

Requirements
------------

This tool requires the following CLIs to be installed

-	`cf`
-	`ssh`

It is assumed that in using `cf-ssh` you have already successfully targeted a Cloud Foundry API, and have pushed an application (successfully or not).

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
