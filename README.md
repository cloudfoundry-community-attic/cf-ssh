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

Download a [pre-compiled release](https://github.com/cloudfoundry-community/cf-ssh/releases) for your platform. Place it in your `$PATH` or `%PATH%` and rename to `cf-ssh` (or `cf-ssh.exe` for Windows).

Alternately, if you have Go setup you can build it from source:

```
go get -u github.com/cloudfoundry-community/cf-ssh
```

Usage
-----

```
cd path/to/app
cf-ssh -f manifest.yml
```

Publish releases
----------------

To generate the pre-compiled executables for the target platforms, using [gox](https://github.com/mitchellh/gox):

```
gox -output "out/{{.Dir}}_{{.OS}}_{{.Arch}}" -osarch "darwin/amd64 linux/amd64 windows/amd64 windows/386" ./...
```

They are now in the `out` folder:

```
-rwxr-xr-x  1 drnic  staff   4.0M Oct 25 23:05 cf-ssh_darwin_amd64
-rwxr-xr-x  1 drnic  staff   4.0M Oct 25 23:05 cf-ssh_linux_amd64
-rwxr-xr-x  1 drnic  staff   3.4M Oct 25 23:05 cf-ssh_windows_386.exe
-rwxr-xr-x  1 drnic  staff   4.2M Oct 25 23:05 cf-ssh_windows_amd64.exe
```

```bash
VERSION=v0.1.0
github-release release -u cloudfoundry-community -r cf-ssh -t $VERSION --name "cf-ssh $VERSION" --description 'SSH into a running container for your Cloud Foundry application, run one-off tasks, debug your app, and more.'

for arch in darwin_amd64 linux_amd64 windows_amd64 windows_386; do
  github-release upload -u cloudfoundry-community -r cf-ssh -t $VERSION --name cf-ssh_$arch --file out/cf-ssh_$arch*
done
```
