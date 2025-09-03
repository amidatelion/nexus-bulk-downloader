# nexus-bulk-downloader

## Directions

Copy config.json.example to config.json

Add your [Nexus Mods API key](https://next.nexusmods.com/settings/api-keys) to the `apikey` value. Use the "Personal API Key" at the bottom. Save this to a password manager for ease of use.

Note the `autoextract` and `downloaddir` values. These can be overriden either in the config or on the command line.

> [!IMPORTANT]
> Currently only supports API keys from premium accounts. Support for free accounts is forthcoming, but due to Nexus' policies, only premium accounts can download from the API directly, free accounts must go through the website.

## Usage


`nexus-bulk-downloader download [flags]`

Flags:
  -x, --autoextract          Override autoextract setting
  -d, --downloaddir string   Override download directory
  -h, --help                 help for download

Reads the config.json file and downloads all the mod files for each entry to the current directory. 

If a mod offers multiple file options (i.e. MW5 Advanced Zoom), you will be prompted which option to download. Multiple file options are not supported at this time.

