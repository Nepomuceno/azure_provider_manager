# AZReg

This is a command line to to help you administer the providers that are regiestered in onw subscription.

## Pre requisites

This tool need to run in a computer that has [az cli](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest).

The user authenticated in Azure CLI needs to haver permission to read register and unregister providers.

## Commands

There are 2 commands current available in this tool `init` and `sync`

## Init

This is the initializaiton of the profile
Some sample usages would be:

```bash
azreg init --profile <PROFILE_NAME> --subscription <SUBSCRIPTION_ID>
```

If you want to setup the output file just pass the file path to `output` parameter

```sh
azreg init --profile <PROFILE_NAME> --subcription <SUBSCRIPTION_ID> --output <ABSOLUTE_FILE_PATH>
```

```text
Usage:
  azreg init [flags]

Flags:
  -h, --help                  help for init
  -s, --subscription string   Subscription id of the

Global Flags:
      --profile string   profile to be used (default "default")
```

## Sync

This is the update of the subscription to comply with the profile
Some sample usages would be:

```sh
azreg sync --subscription <SUBSCRIPTION_ID> --input <INPUT_FILE>
```

If you want to setup the output file just pass the file path to 'output' parameter

```sh
azreg sync --subcription <SUBSCRIPTION_ID> --input <INPUT_FILE>  --output <OUPUT_FILE_PATH>
```

```text
Usage:
  azreg sync [flags]

Flags:
  -h, --help                  help for sync
  -i, --input string          input configuration file to be used in the sync
  -o, --output string         out configuration file tolet the generated file to get to
  -s, --subscription string   Subscription id of the

Global Flags:
      --profile string   profile to be used (default "default")
```