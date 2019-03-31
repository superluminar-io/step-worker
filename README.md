# DynamoDB Step Worker

> Example project for AWS [Step Functions](https://aws.amazon.com/step-functions/) and [State Machines](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) to process data objects stored in DynamoDB with parallel DynamoDB Scans using Segments.

The CloudFormation Stack creates two State Machines: **Puppeteer** and **Worker**. You can configure the **Puppeteer** to fan out the desired number of **Worker** executions to process the DynamoDB objects with the following configuration:

```json
{
  "Comment": "Run State Machine Workers",
  "Configuration": {
    "TableName": "step-worker-url-stable",
    "TableHash": "url",
    "Workers": 4,
    "BatchSize": 250
  }
}
```


![State Machine](/machine.png)

## Usage



### Dependencies

```bash
# Install Go dependencies

$ > make install
```

### Deployment

```bash
# Compile Lambda w/ Go and deploy CloudFormation Stack

$ > make build package deploy
```

### Execute State Machine

```bash
# Get ARN of State Machine 

$ > make outputs

[
  {
    "OutputKey": "Puppeteer",
    "OutputValue": "arn:aws:states:eu-west-1:1234567890:stateMachine:step-worker-puppeteer-stable",
    "Description": "ARN for State Machine"
  }
]

# Execute State Machine and configure BatchSize

$ > aws stepfunctions start-execution \
    --state-machine-arn arn:aws:states:eu-west-1:1234567890:stateMachine:step-worker-puppeteer-stable \
    --input '
    {
      "Comment": "Run State Machine with BatchSize 25",
      "Configuration": {
        "TableName": "step-worker-url-stable",
        "TableHash": "url",
        "Workers": 4,
        "BatchSize": 250
      }
    }'
```

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.
