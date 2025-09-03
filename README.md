# nexus-bulk-downloader

Designed as a generic bulk downloader/extractor even though config.json.example is currently set up for MechWarrior 5: Mercenaries. 

To download mods from other games, add the game domain as a block under `games` in the config and your mods *by ID* into the *keys* underneath. Currently the values of the key-value pairs function as comments for clarity, since json doesn't support comments.

There are vague, distant plans for an extension to make generating these blocks easier. Emphasis on distant.

## Directions

Copy config.json.example to config.json

Add your [Nexus Mods API key](https://next.nexusmods.com/settings/api-keys) to the `apikey` value. Use the "Personal API Key" at the bottom. Save this to a password manager for ease of use.

Note the `autoextract` and `downloaddir` values. These can be overriden either in the config or on the command line.

> [!IMPORTANT]
> Currently only supports API keys from premium accounts. Support for free accounts is forthcoming, but due to Nexus' policies, only premium accounts can download from the API directly, free accounts must go through the website.

## Usage
```
nexus-bulk-downloader download [flags]

Flags:
  -x, --autoextract          Override autoextract setting
  -d, --downloaddir string   Override download directory
  -h, --help                 help for download
```

Reads the config.json file and downloads all the mod files for each entry to the current directory. 

If a mod offers multiple file options (i.e. MW5 Advanced Zoom), you will be prompted which option to download. Multiple file options are not supported at this time.

