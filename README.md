# sp
Stepfinance price, or sp, returns step price in terminal. That's the only price you need fam :)

# Installation

Depending on your OS, run the following commands in your cmd / terminal

## Windows

```bash
curl.exe -Lo sp-windows-amd64.exe https://github.com/forkchili/sp/releases/download/v1.0.0/sp-windows-amd64.exe
Move-Item .\sp-windows-amd64.exe c:\a-dir-in-your-PATH\sp.exe
```

## Linux

```bash
curl -Lo ./sp https://github.com/forkchili/sp/releases/download/v1.0.0/sp-linux-amd64
chmod +x ./sp
mv sp /usr/local/bin/sp
```

## Mac (using bash)

```bash
curl -Lo ./sp https://github.com/forkchili/sp/releases/download/v1.0.0/sp-darwin-amd64
chmod +x ./sp
mv sp /a-dir-in-your-PATH/sp
```

# Usage
In terminal :

```bash
sp
```

Price by default is displayed in USD.

You can change the displayed currency with parameter --c.

```bash
sp --c=CURRENCY
```

CURRENCY value must be a valid 3 letter fiat currencies code (e.g. EUR, JPY, AUD, GBP ect...)

# Special Thanks 

Just wanted to thank the steppers fam.
Enjoy a little bit of fun on your terminal ! LFG ! 🤝