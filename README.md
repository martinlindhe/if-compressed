# About

Command line tool to show resulting sizes of input file compressed with gzip, xz and brotli


# Installation

    brew install brotli xz

    go get -u github.com/martinlindhe/if-compressed


# Usage

```
if-compressed public/js/app.d75add8a.js --human

in   :  225 KiB

gz  -5:  51 KiB
xz  -5:  44 KiB
bro -5:  46 KiB

gz  -9:  50 KiB
xz  -9:  44 KiB
bro -9:  45 KiB
```


# License

Under [MIT](LICENSE)
