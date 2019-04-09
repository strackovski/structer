provider "aws" {
  region  = "${var.aws_region}"
  profile = "${var.aws_profile}"
}

terraform {
  backend "s3" {}
  required_version = ">= 0.11.7"
}

## Variables ##
variable "aws_profile" {}
variable "aws_region" {}
variable "acm_region" {
  default = ""
}

variable "organization" {}
variable "environment" {}

variable "project" {
  default = ""
}

variable "owner" {}
variable "terraform_state_aws_region" {}
variable "terraform_state_s3_bucket" {}
variable "acm_domain_name" {}
variable "acm_name" {
  default = ""
}

variable "acm_aliases" {
  type    = "list"
  default = []
}

## Data sources ##
data "terraform_remote_state" "hosted_zone" {
  backend = "s3"

  config {
    region  = "${var.terraform_state_aws_region}"
    bucket  = "${var.terraform_state_s3_bucket}"
    key     = "${var.aws_region}/${var.environment}/route53/terraform.tfstate"
    profile = "${var.aws_profile}"
  }
}

## Modules ##
module "acm" {
  source             = "source"
  organization       = "${var.organization}"
  environment        = "${var.environment}"
  owner              = "${var.owner}"
  acm_name           = "${var.acm_name != "" ? var.acm_name : var.project}"
  acm_region         = "${var.acm_region != "" ? var.acm_region :var.aws_region}"
  acm_profile        = "${var.aws_profile}"
  acm_domain_name    = "${var.acm_domain_name}"
  acm_aliases        = ["${var.acm_aliases}"]
  acm_hosted_zone_id = "${data.terraform_remote_state.hosted_zone.zone_id}"
}

## outputs ##
output "arn" {
  value = "${module.acm.arn}"
}
