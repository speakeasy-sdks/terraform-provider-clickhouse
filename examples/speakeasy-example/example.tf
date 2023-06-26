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

variable "myIP" {
  type = string
}

provider "clickhouse" {
  username = var.keyID
  password = var.keySecret
}

resource "clickhouse_service" "test" {
  organization_id = var.orgID
  cloud_provider  = "gcp"
  idle_scaling    = true
  ip_access_list = [
    {
      source      = var.myIP
      description = "my IP"
    }
  ]
  idle_timeout_minutes = 10

  name   = "Speakeasy Test Instance"
  region = "us-central1"
  tier   = "development"
}


resource "clickhouse_user" "test_user" {
  organization_id = var.orgID
  email           = "bot@speakeasyapi.dev"
  role            = "admin"
}