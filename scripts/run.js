#! /usr/bin/env node
"use strict";

var exec = require("child_process").spawnSync;
var path = require("path");

var homedir = path.join(__dirname, "../bin/csla");

exec(homedir, process.argv.slice(2), {
  stdio: "inherit"
});
