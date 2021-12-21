# plugs

## Importing the module: 

`go get github.com/matthewboyd/plugs`

```imports (

    "github.com/matthewboyd/plugs"
```

## How to use the module: 

There are two functions within the module: 

* Display
* DoINeedAConverter


## Display: 

The display will allow you to see the following properties: 

* The plug type
* The volts used in the country
* The hertz used in the country

Example use: 

plugs.Display("Serbia")

## DoINeedAConverter

This will allow you to compare two countries and see if you require a converter.

plugs.DoINeedAConverter("USA", "UK")
