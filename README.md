# golang-webapp
Golang Webapp sample

### Prerequirement
- VirtualBox
- Vagrant

### Local
```sh
$ make runlocal
```

### Vagrant
#### vagrant setup
1. provision
```sh
$ vagrant provision
```
1. migration
```sh
$ go get github.com/tools/godep
$ goose up
```
1. setup cookie  
set cookie named `test_uid`, and that's value is `users.id` in DB
1. create user  
insert db.
1. create article  
insert db.

### ECS
First, setup ECS, second, setup EC2 because of prepare container instance cluster.

1. setup ECS
setup ECS cluster, task, and service
```sh
$ cd ecs
$ ./prepare.sh
```
1. setup EC2
```sh
$ cd terraform
$ terraform plan -var-file=aws_access_token.tfvars -var-file=variables.tfvars
$ terraform apply -var-file=aws_access_token.tfvars -var-file=variables.tfvars
```
1. migration
```sh
DB_HOST=<your instance hostname> goose -env release up
```
1. setup cookie, create user, article.
1. request!
