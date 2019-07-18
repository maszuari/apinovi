# apinovi
Go Rest API that uses Gorilla Mux and reads from a JSON file

This is an example of a cart checkout process that uses Go and Gorilla Mux. It uses a JSON file to get the details of a product. 
It uses Docker Compose to run multiple-container Docker applications. 

Prerequisite:-
- Docker
- Docker Compose

Installation:-

1. Install Docker. You can find instructions for the installation here; https://docs.docker.com/install/
2. Install Docker Compose. You can find instructions for the installation here; https://docs.docker.com/compose/install/
3. On terminal, git clone https://github.com/maszuari/apinovi.git
4. Run this command: docker-compose up

Use these commands:-

1. curl -X POST http://localhost:3000/checkout -H 'Content-Type: application/json' -d '{"items":["VOUCHER","TSHIRT","MUG"]}'
2. curl -X POST http://localhost:3000/checkout -H 'Content-Type: application/json' -d '{"items":["VOUCHER","TSHIRT","VOUCHER"]}'
3. curl -X POST http://localhost:3000/checkout -H 'Content-Type: application/json' -d '{"items":["TSHIRT", "TSHIRT", "TSHIRT", "VOUCHER", "TSHIRT"]}'
4. curl -X POST http://localhost:3000/checkout -H 'Content-Type: application/json' -d '{"items":["VOUCHER", "TSHIRT", "VOUCHER", "VOUCHER", "MUG", "TSHIRT", "TSHIRT"]}'
