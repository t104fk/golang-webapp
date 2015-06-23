# golang-webapp
Golang Webapp sample

### Prerequirement
- VirtualBox
- Vagrant

### Steps
1. provision
```sh
$ vagrant provision
```
1. migration
```sh
$ goose up
```
1. setup cookie  
set cookie named `test_uid`, and that's value is `users.id` in DB
1. create user  
insert db.
1. create article  
insert db.

#### run locally
```sh
$ make runlocal
```
