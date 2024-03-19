# StoriChallenge

Backend Coding Challenge made by Patricio Yegros

- [Challenge](#challenge)
- [Assumptions](#assumptions)
- [Execution](#execution)
- [EmailSending](#emailSending)
- [Money](#money)

## Challenge

For this challenge you must create a system that processes a file from a mounted directory. The file
    will contain a list of debit and credit transactions on an account. Your function should process the file
    and send summary information to a user in the form of an email.

An example file is shown below; but create your own file for the challenge. Credit transactions are
    indicated with a plus sign like +60.5. Debit transactions are indicated by a minus sign like -20.46

CSV:
    Id,Date,Transaction
    0,7/15,+60.5
    1,7/28,-10.3
    2,8/2,-20.46
    3,8/13,+10

We prefer that you code in Python or Golang; but other languages are ok too. Package your code in
one or more Docker images. Include any build or run scripts, Dockerfiles or docker-compose files
needed to build and execute your code.

Bonus points
1. Save transaction and account info to a database
2. Style the email and include Stori’s logo
3. Package and run code on a cloud platform like AWS. Use AWS Lambda and S3 in lieu of Docker.

Delivery and code requirements

Your project must meet these requirements:

1. The summary email contains information on the total balance in the account, the number of
transactions grouped by month, and the average credit and average debit amounts grouped by 
month. 

Using the transactions in the image above as an example, the summary info would be

Total balance is 39.74
Number of transactions in July: 2
Number of transactions in August: 2
Average debit amount: -15.38
Average credit amount: 35.25

2. Include the file you create in CSV format.
3. Code is versioned in a git repository. The README.md file should describe the code interface and
instructions on how to execute the code.

## Assumptions

1. Transactions are always in the folder "csv" in csv format
2. Id is a unrepeatable integer number that increments by 1.
3. The inserted email is considered correct only if it haves a "@" in it. (The app dont check if the domain or the email exist)
4. The data in the transactions file always has this order: Id,Date,Transaction

## Execution

1. Open your favourite IDE (Mine is VSCode)
2. In the terminal, go to the folder StoriChallenge
3. Execute "go run main.go"
4. Follow the instructions of the app!

## EmailSending

1. This app use a specific email created by me (patricioyegrosstori@gmail.com), the user can insert an email for destiny but not for origin.
2. The hardcoded password is a app-password created by gmail that only works with this app.

## Money
For all the money-related things, im using the library "shopspring/decimal", so we can forget about losing information in the parsing of the money.
(Every penny maters!)


