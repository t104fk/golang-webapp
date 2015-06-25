#variable "region" {}
#
#resource "aws_vpc" "main" {
#  cidr_block = "10.0.0.0/16"
#  tags {
#    Name = "main"
#  }
#}
#
#resource "aws_subnet" "subnet_a" {
#  vpc_id = "${aws_vpc.main.id}"
#  cidr_block = "10.0.0.0/24"
#  availability_zone = "${var.region}a"
#  tags {
#    Name = "main"
#  }
#}
#
##resource "aws_subnet" "subnet_c" {
##  vpc_id = "${aws_vpc.main.id}"
##  cidr_block = "10.0.1.0/24"
##  availability_zone = "${var.region}c"
##  tags {
##    Name = "main"
##  }
##}
#
#resource "aws_internet_gateway" "gw" {
#  vpc_id = "${aws_vpc.main.id}"
#  tags {
#    Name = "main"
#  }
#}
#
## TODO: [Route Tables]を選択後、今回作成したVPCへのルートをグローバル0.0.0.0/0からの設定をする。入力後、[SAVE]して終わり。
## http://qiita.com/hiroshik1985/items/9de2dd02c9c2f6911f3b
#resource "aws_route_table" "r" {
#  vpc_id = "${aws_vpc.main.id}"
#  route {
#    cidr_block = "10.0.1.0/24"
#    gateway_id = "${aws_internet_gateway.gw.id}"
#  }
#  tags {
#    Name = "main"
#  }
#}
