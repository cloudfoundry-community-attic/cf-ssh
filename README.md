cf-ssh
======

**DEPRECATED:** Cloud Foundry CLI now has `cf ssh` built-in.

SSH into a running container for your Cloud Foundry application, run one-off tasks, debug your app, and more.

Initial implementation requires the application to have a `manifest.yml`.

Also, `cf-ssh` requires that you run the command from within the project source folder. It performs a `cf push` to create a new application based on the same source code/path, buildpack, and variables. Once CF Runtime supports copy app bits [#78847148](https://www.pivotaltracker.com/story/show/78847148), then `cf-ssh` will be upgraded to use app bit copying, and not require local access to project app bits.

It is desired that `cf-ssh` works correctly from all platforms that support the `cf` CLI.

Windows is a target platform but has not yet been tested. Please give feedback in the [Issues](https://github.com/cloudfoundry-community/cf-ssh/issues).

Requirements
------------

This tool requires the following CLIs to be installed

-	`cf` ([download](https://github.com/cloudfoundry/cli/releases))
-	`ssh` (pre-installed on all *nix; [download](http://www.mls-software.com/opensshd.html) for Windows)

It is assumed that in using `cf-ssh` you have already successfully targeted a Cloud Foundry API, and have pushed an application (successfully or not).

This tool also currently requires outbound internet access to the http://tmate.io/ proxies. In future, to avoid the requirement of public Internet access, it would be great to package up the tmate server as a BOSH release and deploy it into the same infrastructure as the Cloud Foundry deployment.

### Why require `ssh` CLI?

This project is written in the Go programming language, and there is a candidate library [go.crypto](https://godoc.org/code.google.com/p/go.crypto/ssh#Session.RequestPty) that could have natively supported an interactive SSH session. Unfortunately, the SSL supports a subset of ciphers that don't seem to work with tmate.io proxies [[stackoverflow](http://stackoverflow.com/questions/18998473/failed-to-dial-handshake-failed-ssh-no-common-algorithms-error-in-ssh-client/19002265#19002265)]

Using the `go.crypto` library I was getting the following error. In future, perhaps either tmate.io or go.crypto will change to support each other.

```
unable to connect: ssh: handshake failed: ssh: no common algorithms
```

Installation
------------

Download a [pre-compiled release](https://github.com/cloudfoundry-community/cf-ssh/releases) for your platform. Place it in your `$PATH` or `%PATH%` and rename to `cf-ssh` (or `cf-ssh.exe` for Windows).

Alternately, if you have Go setup you can build it from source:

```
go get github.com/cloudfoundry-community/cf-ssh
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
