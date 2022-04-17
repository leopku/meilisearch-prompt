prog := "meilisearch-prompt"

build:
  go build -ldflags "-w -s" -o build/{{prog}} .
