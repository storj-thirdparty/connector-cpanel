# Run

The following flags can be used with the `store` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase.
* `share` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Once you have built the project you can run the following:

## Get help

```
$ ./connector-cpanel --help
```

## Check version

```
$ ./connector-cpanel --version
```

## Create back-up from cPanel and upload them to Storj

```
$ ./connector-cpanel store --cpanel <path_to_cpanel_config_file> --storj <path_to_storj_config_file>
```

## Create back-up from cPanel and upload it to Storj bucket using Access Key

```
$ ./connector-cpanel store --accesskey
```

## Create back-up from cPanel and upload it to Storj and generate a Shareable Access Key based on restrictions in `storj_config.json`

```
$ ./connector-cpanel store --share
```
