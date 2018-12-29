# PULUMI DEMO
## Description
This example creates a file from template in GCS displaying the `GREETING` variable

## Running the example

Execute the following bash command in order to setup the enviroment.

```bash
export CREDENTIALS_FILE="/Users/omar/Downloads/EXAMPLE-1bf0619764f8.json"
export CREDENTIALS=$(cat $CREDENTIALS_FILE)
export PULUMI_ACCESS_TOKEN="zzzz"
export GOOGLE_PROJECT="XXX"
export GREETING="Hello World"
```

Pulumi requires access token in order to save state, you can get one at https://app.pulumi.com/

Create the container with the pulumi code binary

```bash
make build-gcp
```

Preview GCP resources to create

```bash
make pulumi-preview
```

Create/Update GCP resources

```bash
make pulumi-update
```
