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

> We need PR https://github.com/centreon/centreon/pull/9436

---

## Host Category

### List

List with filter not working. Always return empty list.

> It's fix on develop branch, but not on 25.10.

---

## Media

### Get

Method not allowed. No way to get all fields for a given media from configuration/media API

> Need PR https://github.com/centreon/centreon/pull/9329

### Delete

Method not allowed. No way to delete media from configuration/media API

> Need PR https://github.com/centreon/centreon/pull/9329

---

## Command

### Get

Method not allowed. No way to get all fields for a given command from configuration/command API

> Need PR https://github.com/centreon/centreon/pull/9359

### Delete

Method not allowed. No way to delete command from configuration/command API
> Need PR https://github.com/centreon/centreon/pull/9359

### Update

Method not allowed. No way to update command from configuration/command API
> Need PR https://github.com/centreon/centreon/pull/9359

---

## Timzeone

How to get the timezone ID from API. We need to provide it when creattin host and host template but no way to get it from API.

---

## Service

### Get

Method not allowed. No way to get all fields for a given service from configuration/service API
> Need PR https://github.com/centreon/centreon/pull/9306

## Service Template

### Get
Method not allowed. No way to get all fields for a given service template from configuration/servicetemplate API
> Nedd PR https://github.com/centreon/centreon/pull/9451


## Service Severity

### Get

Method not allowed. No way to get all fields for a given service severity from configuration/serviceseverity API
> Need PR https://github.com/centreon/centreon/pull/9459

## Service group

### Get
Method not allowed. No way to get all fields for a given service group from configuration/servicegroup API
> Need PR https://github.com/centreon/centreon/pull/9463

### Update
Method not allowed. No way to update service group from configuration/servicegroup API
> Need PR https://github.com/centreon/centreon/pull/9463

## Service category

### Get

Method not allowed. No way to get all fields for a given service category from configuration/servicecategory API
> Need PR https://github.com/centreon/centreon/pull/9472

### Update 

Method not allowed. No way to update service category from configuration/servicecategory API
> Need PR https://github.com/centreon/centreon/pull/9472