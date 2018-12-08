variable "name" {
  default = "common"
}

resource "aws_codecommit_repository" "lambda" {
  repository_name = "${var.name}"
}
