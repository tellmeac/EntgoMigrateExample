go generate ./ent/...;

atlas migrate apply \
  --dir "file://migrations" \
  --url "postgres://postgres:postgres@localhost:5432/Test?sslmode=disable"