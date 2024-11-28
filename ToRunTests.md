TestGetRisks: 
Tests the GET /v1/risks endpoint to ensure it returns a list of risks.


TestCreateRisk: 
Tests the POST /v1/risks endpoint to ensure it creates a new risk and returns the correct status code and response body.


TestGetRisk: 
Tests the GET /v1/risks/<id> endpoint to ensure it retrieves the correct risk by its ID.


You can run these tests using the following command:

go test -v