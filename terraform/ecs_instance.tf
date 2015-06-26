variable "region" {}

resource "template_file" "user_data" {
  filename = "userdata.sh"
}

resource "aws_instance" "ecs_instance01" {
  ami = "ami-9cc1119c"
  #availability_zone = "${var.region}a"
  availability_zone = "ap-northeast-1a"
  instance_type = "t2.small"
  key_name = "toyo-tokyo"
  tags {
    Name = "api01"
  }
  iam_instance_profile = "${aws_iam_instance_profile.ecs_profile.name}"
  security_groups = ["${aws_security_group.api.name}"]
  user_data = "${template_file.user_data.rendered}"
}

resource "aws_instance" "ecs_instance02" {
  ami = "ami-9cc1119c"
  #availability_zone = "${var.region}a"
  availability_zone = "ap-northeast-1a"
  instance_type = "t2.small"
  key_name = "toyo-tokyo"
  tags {
    Name = "api02"
  }
  iam_instance_profile = "${aws_iam_instance_profile.ecs_profile.name}"
  security_groups = ["${aws_security_group.api.name}"]
  user_data = "${template_file.user_data.rendered}"
}
