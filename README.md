# About

Command line tool to show resulting sizes of input file compressed with gzip and xz


# Installation

    go get -u github.com/martinlindhe/if-compressed


# Usage

```
if-compressed public/js/app.d75add8a.js --human

in   :  218 KiB
gz -5:  50 KiB
gz -9:  49 KiB
xz -5:  43 KiB
xz -9:  43 KiB
```


# License

Under [MIT](LICENSE)
