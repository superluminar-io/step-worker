# DynamoDB Step Worker

> Example how to use a [State Machine](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) with [Step Functions](https://aws.amazon.com/step-functions/) to process objects stored in DynamoDB.

## Usage

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
    "OutputKey": "StateMachine",
    "OutputValue": "arn:aws:states:eu-west-1:1234567890:stateMachine:step-worker-machine-stable",
    "Description": "ARN for State Machine"
  }
]

# Execute State Machine and configure BatchSize

$ > aws stepfunctions start-execution \
    --state-machine-arn arn:aws:states:eu-west-1:1234567890:stateMachine:step-worker-machine-stable \
    --input '
    {
      "Comment": "Run State Machine with BatchSize 25",
      "Configuration": {
        "BatchSize": 250
      }
    }'
```

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.