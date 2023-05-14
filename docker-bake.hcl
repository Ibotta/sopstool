variable "VERSION" {
  default = "$VERSION"
}

target "sopstool" {
  dockerfile = "Dockerfile"
  tags = [
    "latest",
    ""
  ]
  pull = true
}
