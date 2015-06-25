resource "aws_iam_instance_profile" "ecs_profile" {
  name = "${aws_iam_role.ecs_role.name}"
  roles = ["${aws_iam_role.ecs_role.name}"]
}

resource "aws_iam_role" "ecs_role" {
  name = "ecs-role"
  path = "/"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {"AWS": "*"},
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "ecs_policy" {
  name = "ecs-policy"
  role = "${aws_iam_role.ecs_role.id}"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:Describe*",
        "elasticloadbalancing:*",
        "ecs:*",
        "iam:ListInstanceProfiles",
        "iam:ListRoles",
        "iam:PassRole"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

