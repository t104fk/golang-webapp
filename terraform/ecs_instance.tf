variable "region" {}

resource "template_file" "user_data" {
  filename = "userdata.sh"
}

resource "aws_instance" "ecs_instance" {
  ami = "ami-9cc1119c"
  #availability_zone = "${var.region}a"
  availability_zone = "ap-northeast-1a"
  instance_type = "t2.micro"
  key_name = "toyo-tokyo"
  tags {
    Name = "api01"
  }
  iam_instance_profile = "${aws_iam_instance_profile.ecs_profile.name}"
  security_groups = ["${aws_security_group.api.name}"]
  #subnet_id = "${aws_subnet.subnet_a}"
  user_data = "${template_file.user_data.rendered}"
}
