# S3 Substring Finder
[![CI](https://github.com/bilalcaliskan/s3-substring-finder/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/s3-substring-finder/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/s3-substring-finder)](https://goreportcard.com/report/github.com/bilalcaliskan/s3-substring-finder)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-substring-finder&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-substring-finder)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-substring-finder&metric=coverage)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-substring-finder)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-substring-finder&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-substring-finder)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-substring-finder&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-substring-finder)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-substring-finder&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-substring-finder)
[![Release](https://img.shields.io/github/release/bilalcaliskan/s3-substring-finder.svg)](https://github.com/bilalcaliskan/s3-substring-finder/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/s3-substring-finder)](https://github.com/bilalcaliskan/s3-substring-finder)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This tool gets the **AWS S3** credentials from user as input and also gets a specific substring to search across the files with specified extensions with **--fileExtensions** flag in a bucket.
Then prints the file names that contains provided substring.

You can provide access credentials of your AWS account with below environment variables or CLI flags. Keep in mind that command line flags
will override environment variables if you set both of them:
```
"--accessKey" CLI flag or "AWS_ACCESS_KEY" environment variable
"--secretKey" CLI flag or "AWS_SECRET_KEY" environment variable
"--region" CLI flag or "AWS_REGION" environment variable
"--bucketName" CLI flag or "AWS_BUCKET_NAME" environment variable
```

## Configuration
This tool provides below command line arguments:
```
This tool searches the specific substring in files on AWS S3 and returns the file names

Usage:
  s3-substring-finder [flags]

Flags:
      --accessKey string        access key credential to access S3 bucket
      --bucketName string       name of the target bucket on S3
      --fileExtensions string   comma separated list of file extensions to search on S3 bucket (default "txt")
  -h, --help                    help for s3-substring-finder
      --region string           region of the target bucket on S3
      --secretKey string        secret key credential to access S3 bucket
      --substring string        substring to find on txt files on target bucket
  -v, --verbose                 verbose output of the logging library (default false)
      --version                 version for s3-substring-finder
```

## Installation
### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/s3-substring-finder/releases) page.

After then, you can simply run binary by providing required command line arguments:
```shell
$ ./s3-substring-finder --accessKey asdasfasfasfasfasfas --secretKey asdasfasfasfasfasfas --bucketName demo-bucket --region us-east-2 --substring "catch me if you can"
```

### Homebrew
This project can also be installed with [Homebrew](https://brew.sh/):
```shell
$ brew tap bilalcaliskan/tap
$ brew install bilalcaliskan/tap/s3-substring-finder
```

Then similar to binary method, you can run it by calling below command:
```shell
$ s3-substring-finder --accessKey asdasfasfasfasfasfas --secretKey asdasfasfasfasfasfas --bucketName demo-bucket --region us-east-2 --substring "catch me if you can"
```

## Development
This project requires below tools while developing:
- [Golang 1.19](https://golang.org/doc/go1.19)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
- [gocyclo](https://github.com/fzipp/gocyclo) - required by [pre-commit](https://pre-commit.com/)

After you installed [pre-commit](https://pre-commit.com/), simply run below command to prepare your development environment:
```shell
$ pre-commit install -c build/ci/.pre-commit-config.yaml
```

## License
Apache License 2.0
