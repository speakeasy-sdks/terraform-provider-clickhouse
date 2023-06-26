terraform {
  required_providers {
    clickhouse = {
      source  = "speakeasy/clickhouse"
      version = "0.0.1"
    }
  }
}

variable "keyID" {
  type = string
}

variable "keySecret" {
  type = string
}

variable "orgID" {
  type = string
}

provider "clickhouse" {
  username = var.keyID
  password = var.keySecret
}

resource "clickhouse_service" "test" {
  organization_id = var.orgID
}