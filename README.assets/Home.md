## Flow Diagram

![](https://github.com/storj-thirdparty/connector-cpanel/blob/master/README.assets/arch.drawio.png)

## Config Files

There are two config files that contain Storj network and cPanel connection information.  The tool is designed so you can specify a config file as part of your tooling/workflow.



##### `cpanel_property.json`

Inside the `./config` directory there is a  `cpanel_property.json` file, with following information about your cPanel instance:

* hostname :- Host Name connect to cPanel
* port	   :- Port Number connect to cPanel
* username :- User Name of cPanel
* password :- Password of cPanel

##### `storj_config.json`

Inside the `./config` directory a `storj_config.json` file, with Storj network configuration information in JSON format:

* `apiKey` - API Key created in Storj Satellite GUI (mandatory)
* `satelliteURL` - Storj Satellite URL (mandatory)
* `encryptionPassphrase` - Storj Encryption Passphrase (mandatory)
* `bucketName` - Name of the bucket to upload data into (mandatory)
* `uploadPath` - Path on Storj Bucket to store data (optional) or "" or "/" (mandatory)
* `serializedAccess` - Serialized access shared while uploading data used to access bucket without API Key (mandatory)
* `allowDownload` - Set *true* to create serialized access with restricted download (mandatory while using *share* flag)
* `allowUpload` - Set *true* to create serialized access with restricted upload (mandatory while using *share* flag)
* `allowList` - Set *true* to create serialized access with restricted list access
* `allowDelete` - Set *true* to create serialized access with restricted delete
* `notBefore` - Set time that is always before *notAfter*
* `notAfter` - Set time that is always after *notBefore*



## Run

Once you have built the project run the following commands as per your requirement:

##### Get help

```
$ ./connector-cpanel --help
```

##### Check version

```
$ ./connector-cpanel --version
```

##### Create back-up from cPanel and upload them to Storj

```
$ ./connector-cpanel store --cpanel <path_to_cpanel_config_file> --storj <path_to_storj_config_file>
```

##### Create back-up from cPanel and upload it to Storj bucket using Access Key

```
$ ./connector-cpanel store --accesskey
```

##### Create back-up from cPanel and upload it to Storj and generate a Shareable Access Key based on restrictions in `storj_config.json`.

```
$ ./connector-cpanel store --share
```



##  Testing

The project has been tested on the following operating systems:

```
	* ubuntu
		* Version: 16.04 LTS
		* Processor: AMD A6-7310 APU with AMD Radeon R4 Graphics × 4
```
