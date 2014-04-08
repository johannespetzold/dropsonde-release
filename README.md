# Dropsonde Release [![Build Status](https://travis-ci.org/cf-lamb/hurricanehunter-release.svg?branch=master)](https://travis-ci.org/cf-lamb/hurricanehunter-release)

BOSH release for example components using [dropsonde](https://github.com/cloudfoundry-incubator/dropsonde) metrics emission.

## Usage

HurricaneHunter is an instrumented proxy on port 8080; pass it a form-encoded URL and it will fetch it for you. Once it's running, issue a command like `curl --data "url=http://www.github.com" hurricanehunter.somewhere:8080`

The application will use the dropsonde instrumentation to record the transaction, and will return the result of a GET request against the provided URL.

If you wish to see the results of the instrumentation, provide a local UDP listener on the same instance, port 42420. (Please keep in mind that the dropsonde library emits Protocol Buffer messages, so you will also want to decode those.) Metrics about the sent metrics (aka "heartbeats") are emitted (also in Protocol Buffer format) over TCP to 42421.

## Boshlite deployment

1. Check out the repository and navigate to it
1. `bin/update`
1. `bosh create release`
1. `bosh upload release`
1. `./make_manifest_boshlite`
1. `bosh deploy`
