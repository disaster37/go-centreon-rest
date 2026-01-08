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

### Host Template

### Get

Method not allowed. No way to get all fields for a given host template from configuration/hosttemplate API

### Host Category

### List

List with filter not working. Always return empty list.