# Smorgasbord

[![Go Documentation](https://img.shields.io/badge/go-doc-blue.svg?style=flat)](https://pkg.go.dev/github.com/kubism/smorgasbord/pkg)
[![Build Status](https://travis-ci.org/kubism/smorgasbord.svg?branch=master)](https://travis-ci.org/kubism/smorgasbord)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubism/smorgasbord)](https://goreportcard.com/report/github.com/kubism/smorgasbord)
[![Coverage Status](https://coveralls.io/repos/github/kubism/smorgasbord/badge.svg?branch=master)](https://coveralls.io/github/kubism/smorgasbord?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/b6fbe93e1b95f6b7f5e3/maintainability)](https://codeclimate.com/github/kubism/smorgasbord/maintainability)

> a range of open sandwiches and delicacies served as hors d'oeuvres or a buffet

Smorgasbord purpose is to ease up the administration of a wireguard-based VPN.  
It creates, stores and distributes client configurations for its users and can
derive server configuration using the provided agent.
Users can self-service their public keys after authenticating via OpenID Connect.
Rather than using a database the public keys and metadata are commited to a
git repository, which is used as storage endpoint.

Smorgasbord primary goal is to provide a minimalistic environment to manage
users across multiple wireguard servers applicable to embedded systems as well
as more complex installments.

![Concept of Smorgasbord](./docs/concept.svg)

## Backlog topics

The backlog contains some bigger topics, which we might implement in the future.
However feel free to implement them yourself you need them.

### Automatic removal

Currently it is required to manage the removal/deactivation of users manually,
e.g. admin removing entries from git repository.

However if the information about the deactivation is available via OIDC, e.g.
refresh token failing. It would be possible to deactivate users automatically.

### Configure wireguard directly

Rather than provide the configuration and issuing a command (e.g. `wg syncconf`)
the agent could configure wireguard interface directly using the [go library](https://github.com/WireGuard/wgctrl-go).

## About the name

This project started a late night project and the name was essentially what
came up first after googling "synonym self-service".
It might therefore be subject to change.


