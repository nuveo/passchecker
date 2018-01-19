# Run js

## Install by NPM

```
npm install --save passchecker
```

## Build

```
gopherjs build
```

## Rum example

```
node example.js 
```

## How to use it

```
require('./passchecker.js');
var main = global.passchecker;

console.log('Creating a new passchecker');
var pChecker = main.New();

console.log('Calling method to have the password strength!');
console.log(pChecker.GetStrength('test'));
console.log(pChecker.GetStrength('test pa'));
console.log(pChecker.GetStrength('test passwd'));

// the response is a list with two floats.
// ex: [ 0.0, 0.0 ] - the first is the password strength from 6.0 to 10 and second one is a percent from 0 to 100.
```