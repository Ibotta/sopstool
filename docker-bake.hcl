variable "MAJOR" {
  default = "$MAJOR"
}
variable "MINOR" {
  default = "$MINOR"
}
variable "PATCH" {
  default = "$PATCH"
}

target "default" {
  dockerfile = "Dockerfile"
  tags = [
    "latest",
    "${MAJOR}.${MINOR}.${PATCH}",
    "v${MAJOR}.${MINOR}.${PATCH}",
    "v${MAJOR}.${MINOR}",
    "v${MAJOR}"
  ]
  labels = {
    "org.label-schema.schema-version"="1.0",
    "org.label-schema.version"="${MAJOR}.${MINOR}.${PATCH}",
    "org.label-schema.name"="sopstool",
  }
  platforms = ["linux/amd64", "linux/arm64"]
}

function "repo_tags" {
  params = [tags]
  result = concat(
    formatlist("ghcr.io/ibotta/sopstool:%s", tags),
    formatlist("ibotta/sopstool:%s", tags)
  )
}
