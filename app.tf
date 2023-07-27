terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}

resource "docker_image" "vault-go" {
  name = var.image_name
}
# Create a container
resource "docker_container" "vault-cont" {
  image = docker_image.vault-go.image_id
  name  = var.container_name
  ports {
    internal = var.port_internal
    external = var.port_external
  }
}

