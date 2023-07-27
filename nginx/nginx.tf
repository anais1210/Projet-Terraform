terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.13.0"
    }
  }
}

provider "docker" {}

resource "docker_image" "nginx" {
  name         = var.nginximage
  keep_locally = false
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.latest
  name  = var.nginxapp
  ports {
    internal = 80
    external = 8080
  }
 provisioner "local-exec" {
    command = "curl http://localhost:8080 > myIndex.html"
  }
}
