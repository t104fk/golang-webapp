#variable "region" {}
#
#resource "aws_elb" "main" {
#  name = "golang-webapp-elb"
#  availability_zones = ["${var.region}a", "${var.region}c"]
#
#  listener {
#    instance_port = 80
#    instance_protocol = "http"
#    lb_port = 80
#    lb_protocol = "http"
#  }
#
#  health_check {
#    healthy_threshold = 2
#    unhealthy_threshold = 2
#    timeout = 3
#    target = "HTTP:8000/article/1"
#    interval = 30
#  }
#
#  instances = ["${aws_instance.ecs_instance.id}"]
#  cross_zone_load_balancing = true
#  idle_timeout = 400
#  connection_draining = true
#  connection_draining_timeout = 400
#}
