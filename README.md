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

```
azreg init --profile <PROFILE_NAME> --subcription <SUBSCRIPTION_ID> --output <ABSOLUTE_FILE_PATH>
```

```
Usage:
  azreg init [flags]

Flags:
  -h, --help                  help for init
  -s, --subscription string   Subscription id of the

Global Flags:
      --profile string   profile to be used (default "default")
```