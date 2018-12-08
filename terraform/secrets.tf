provider "aws" {
  profile = "card"
  region = "eu-west-1"
}

terraform {
  backend "s3" {
    bucket = "cardpayments-states"
    key = "common"
    region = "eu-west-1"
    profile = "card"
  }
}