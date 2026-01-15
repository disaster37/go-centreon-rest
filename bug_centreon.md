# All bugs Centreon

## Host

### Create

The response forget the following fields:
- snmp_community


### Get

Method not allowed. No way to get all fields for a given host from configuration/host API

### Get from real time api

"last_hard_state": 1 is integer not datetime string

---

## Host Template

### Get

Method not allowed. No way to get all fields for a given host template from configuration/hosttemplate API

---

## Host Category

### List

List with filter not working. Always return empty list.

---

## Media

### Delete

Method not allowed. No way to delete media from configuration/media API

---

## Command

### Get

Method not allowed. No way to get all fields for a given command from configuration/command API

### Delete

Method not allowed. No way to delete command from configuration/command API

### Update

Method not allowed. No way to update command from configuration/command API

---

## Timzeone

How to get the timezone ID from API. We need to provide it when creattin host and host template but no way to get it from API.

---

## Service

### Get

Method not allowed. No way to get all fields for a given service from configuration/service API
issue: https://github.com/centreon/centreon/issues/9277