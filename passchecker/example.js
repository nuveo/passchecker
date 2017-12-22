// Imports and uses the 'passchecker.js' generated through gopherjs
'use strict';
console.log('Starting');

console.log('Loading passchecker.js');
require('./passchecker.js');
var main = global.passchecker;

console.log('Creating a new passchecker');
var p = main.New();
console.log(p.GetStrength('test'));
console.log(p.GetStrength('test pa'));
console.log(p.GetStrength('test passwd'));

console.log('Done!');