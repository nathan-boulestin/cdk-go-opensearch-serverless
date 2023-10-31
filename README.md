# Disclamer

I made this code to share a simple setup of an AWS Serverless Opensearch collecion with CDK.

**Make sure to adapt the security policies** to your need with [the documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-security.html).

This code does not handle index creation but this can be done with a custom resource.

# Usage

- `cdk deploy -c account=YOUR_ACCOUNT_ID -c region=YOUR_REGION -c collection-name=YOUR_COLLECTION_NAME -c principal-arn=YOUR_USER_ARN`
- Make sure to use a region supported by Opensearch Serverless.
- `principal-arn` can be a user arn or a role arn. This stack will grant permission to this principal to manage indicies in the collection.

![infra](https://github.com/nathan-boulestin/cdk-go-opensearch-serverless/blob/main/infra.drawio.png)