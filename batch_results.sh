#!/bin/bash

set -x

echo "Reviewing whole project ..."
cd example;
lingo review -o ../results/all_issues.json;

echo "Reviewing server ..."
cd server;
lingo review -o ../../results/server_issues.json;

echo "Reviewing client ..."
cd ../client;
lingo review -o ../../results/client_issues.json;

echo "Done! all results written to the results directory."