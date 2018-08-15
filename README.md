# githubls

Simple command line utility to list repos

## Prerequisites

Create a github token with minimal permissions:

Go to https://github.com/settings/tokens and create a new Personal access token

Only the following permissions are necessary:

- repo:status
- public_repo

## Usage

set GITHUB_TOKEN environment variable

```
export GITHUB_TOKEN=98cfzfd8f5f2fz835275z292679f2ed6e3zqdf13
```

### Get all repos in an organization

```
$ ./githubls -org shokunin  |head
ops_scripts
ec2secviz
zero_ipstream
```

### Get all repos in an organization with the name matching a regex

```
$ ./githubls -org shokunin -regex "nagios$"
puppet-nagios
```

### Get all public organizations for a user

```
$ ./githubls -list-public-orgs myfriend
choria-io
google
ruby-foo
```

## Building

```
$ make
```
