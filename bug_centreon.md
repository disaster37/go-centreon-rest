# All bugs Centreon

## Host

### Create

The response forget the following fields:
- snmp_community

> We need PR https://github.com/centreon/centreon/pull/9335


### Get

Method not allowed. No way to get all fields for a given host from configuration/host API

> We need PR https://github.com/centreon/centreon/pull/9335

### Get from real time api

"last_hard_state": 1 is the hard state, not a date like say a doc.

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

### Get

Method not allowed. No way to get all fields for a given media from configuration/media API
Need PR https://github.com/centreon/centreon/pull/9329

### Delete

Method not allowed. No way to delete media from configuration/media API
Need PR https://github.com/centreon/centreon/pull/9329

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