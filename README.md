# QUIK GAMING GO BACKEND DEVELOPER INTERVIEW
## Configuration
###  Setup
You need to have Golang installed at the very least `Go 1.16`. 


The easiest way to get started with the rest api in your local environment is to clone it using git:

```
git clone https://github.com/Yaska1706/quik-gaming-interview.git
```

**Environment Variables**

You need to configure environment variables for the rest api to work as expected.

Copy the content on .env_example to .env file
`cat .env_example > .env`

The .env file should look like this:

```
# database configurations
DB_DRIVER = "mysql"
DB_HOST = "localhost"
DB_PORT = "5432"
DB_USER = "root"
DB_NAME = "wallet_player"
DB_PASSWORD = ""

# port for app to listen and serve
LISTEN_ADDRESS = "127.0.0.1"
LISTEN_PORT = "8080"

# Logging file
LOG_FILE = "wallet.log"

# redis configuration
REDIS_ADDRESS = "127.0.0.1"
REDIS_PORT = "6379"

```
You can change the values to your own preference.
## Run
To run the REST api you need to run this on the terminal so as to download the required dependancies:

`make deps`

Run the application

`make run`

## Task 
The goal is to write a JSON API in Golang to get the balance and manage credit/debit  operations on the players wallets. For example, you might receive calls on your API to get  the balance of the wallet with id 123, or to credit the wallet with id 456 by 10.00 €. The  storage mechanism to use will be MySQL.

Below are the 3 endpoints to implement, as well as the business rules. Endpoints 
- balance : retrieves the balance of a given wallet id 

```GET /api/v1/wallets/{wallet_id}/balance ```
- credit : credits money on a given wallet id 

```POST / api/v1/wallets/{wallet_id}/credit ```

- debit : debits money from a given wallet id 

```POST / api/v1/wallets/{wallet_id}/debit ```
#### Business rules 
- A wallet balance cannot go below 0. 
- Amounts sent in the credit and debit operations cannot be negative. 

## Results

#### DEBIT

```json
{
    "data": "debit created",
    "status": "success",
    "amount": "4000"
}
```
 #### CREDIT

```json
{
    
    "data": "credit successful",
    "status": "success",
    "amount": "3000"

}

```
 
#### Balance

```json
{
    "balance": "1000"
}

```
