# Bitcoin transaction builder

## Project Description

A script that generates a redeem script, derives a p2sh address from it and constructs a bitcoin transaction using Golang.

## Installation

To run this project, you need to have Go installed on your machine. If you haven't installed Go yet, you can download and install it from 
the [official Go website](https://golang.org/doc/install).

### Cloning the Repository

First, clone this repository to your local machine using Git:

```bash
git clone https://github.com/theedtron/btctxbuilder.git
```

## Running the Project

Navigate to the project directory:

```bash
cd btctxbuilder
```
Download necessary dependancies

```bash
go mod tidy
```

Then, run the project using the following command:

```bash
go run main.go
```
