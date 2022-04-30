# SBIT DApp Server

A Web Server for running third-party DApps.

All requests made by a DApp would block until you approve it in the authorization UI.

# Install sbit-portal

If you have golang installed, you can install the latest version:

```
go get -u github.com/SBit-Project/sbit-portal/cli/sbitportal
```

# Running A SBIT DApp

First, we'll need to make sure that sbitd is running.

For testing/development purposes, let's start sbitd in regtest mode:

```
sbitd -regtest -rpcuser=user -rpcpassword=pass
```

Then use the env variable `SBIT_RPC` to specify the URL of your local sbitd RPC node:

```
export SBIT_RPC=http://user:pass@localhost:22302
```
