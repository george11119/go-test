echo "creating tasks test:"
curl localhost:3001/task/ \
  -H "Content-Type: application/json" \
  -d '{"text":"first task", "tags":["tag1"], "due":"2006-01-02T15:04:05Z"}' \
  -iL -X POST
printf "\n\n"

curl localhost:3001/task/ \
  -H "Content-Type: application/json" \
  -d '{"text":"second task", "tags":["tag2"], "due":"2006-01-02T15:04:05Z"}' \
  -iL -X POST
printf "\n\n"

curl localhost:3001/task/ \
  -H "Content-Type: application/json" \
  -d '{"text":"third task", "tags":["tag1", "tag2", "tag3"], "due":"2006-01-02T15:04:05Z"}' \
  -iL -X POST
printf "\n\n"

echo "getting all tasks:"
curl localhost:3001/task/ | jq

echo "deleting all tasks:"
curl localhost:3001/task/ \
  -iL -X DELETE

echo "getting all tasks:"
curl localhost:3001/task/ | jq
