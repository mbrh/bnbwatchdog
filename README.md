# bnbwatchdog

Use an API key to check your BNB balance on the Binance exchange. If
there is more available than the specified treshold, the program is
silent.

Place this in a cronjob to warn you when you run out of currency to
cover trading fees.


## Example
```shell
% export BINANCE_API_KEY=...
% export BINANCE_API_SECRET=...
% bnbwatchdog -treshold 0.5
BNB balance 0.209811 below or equal to treshold 0.500000
```


## Usage
```shell
  -symbol string
        coin symbol (default "BNB")
  -treshold float
        treshold below which to complain
  -verbose
        verbosity
```
