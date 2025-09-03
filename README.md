# nexus-bulk-downloader

## Directions

Copy config.json.example to config.json

Add your [Nexus Mods API key](https://next.nexusmods.com/settings/api-keys) to the `apikey` value. Use the "Personal API Key" at the bottom. Save this to a password manager for ease of use.

Note the `autoextract` and `downloaddir` values. These can be overriden either in the config or on the command line.

## Usage


`nexus-bulk-downloader download [flags]`

Flags:
  -x, --autoextract          Override autoextract setting
  -d, --downloaddir string   Override download directory
  -h, --help                 help for download

Reads the config.json file and downloads all the mod files for each entry to the current directory. 

If a mod offers multiple file options (i.e. MW5 Advanced Zoom), you will be prompted which option to download. Multiple file options are not supported at this time.

