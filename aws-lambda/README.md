go build main.go 

zip function.zip main

trust-policy.json

{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json


aws lambda create-function --function-name hello-sample --handler main --zip-file ./function.zip --runtime go1.x --role arn:aws:iam:<ac id>:/role/lambda-ex


aws lambda invoke --function-name hello-sample --cli-binary-format raw-in-base64-out --playload '{"name":  "jim", "age": 33}' output.txt
