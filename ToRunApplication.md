go run main.go


Create a new risk:
curl -X POST -H "Content-Type: application/json" -d '{"state":"open","title":"Sample Risk","description":"This is a sample risk."}' http://localhost:8080/v1/risks


Get the list of risks:
curl http://localhost:8080/v1/risks


Get an individual risk:
curl http://localhost:8080/v1/risks/<id>