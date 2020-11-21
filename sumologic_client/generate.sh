#!/usr/bin/env bash

wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.23/swagger-codegen-cli-3.0.23.jar -O ./swagger-codegen-cli.jar

java -jar ./swagger-codegen-cli.jar generate -i ./swagger.json -l go -c ./config.json