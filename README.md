![kubernetes](https://raw.githubusercontent.com/kubernetes/kubernetes/master/logo/logo.svg)

# kPong

kPong is a chaos game for kubernetes. It uses your kubeconfig at its default path, at `$HOME/.kube/config`.

DO NOT PLAY THIS IN PRODUCTION (unless you're a bad ass or your services are resiliant AF).

## Installing

Tenatively, we'll place a binary on the [releases page](https://github.com/ryanhartje/kpong/releases)

By default, you'll have a single player game vs AI. Here are some flags that will alter your experience:

|Flag|Description|
|-|-|
|--kubeconfig|Specify location for kubeconfig|
|--namespace|Specify the namespace you want to select from, or leave empty for all namespaces|


## Building

Your platform is likely supported, but I wrote this on a Mac. If there's interest, I'll gladly expand install instructions to support your platform, please file an issue.

### MacOS

You will need xcode, which you can install by running:
```
xcode-select --install
```

Then, clone and fetch dependencies:
```
go get github.com/ryanhartje/kpong
cd ~/go/src/github.com/ryanhartje/kpong
go mod download
```


