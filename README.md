# customerCSV
Dedicated app to parse big CSV files. Implemented in GO

Application splitted into two services in order to apply the micro services pattern. Each service is dedicated to a single responsability. It also allows for tuning each service according its needs in terms of performance.

CSV Reader Service 

Service in charge to load and read the CSV file, save the data in the database and make a request to the CRM integrator service along with the data.


CRM Integrator

Service which has an endpoint to receive the data from the CSV Reader

Run the application

Execute the following commands (you must have the port 80 free on the machine):

./services.sh build && ./services.sh up

Run the CSV processor

if you are running the app locally 

curl localhost 

otherwise

curl ip/domain name

To restart the application run:

./services.sh restart

Test

docker exec csv_reader go test -run Test_reader
docker exec csv_reader go test -run Test_save

