#! /usr/bin/env node
"use strict";

var exec = require("child_process").spawnSync;
var path = require("path");
var os = require("os");
var fs = require("fs");

var homedir = path.join(__dirname, "../bin/csla");
var platform = os.platform();
var arch = process.arch;

console.log(homedir);
exec(homedir, [process.argv.slice(2).join(' ')], {
  stdio: "inherit"
});
