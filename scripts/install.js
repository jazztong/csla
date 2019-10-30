#!/usr/bin/env node
"use strict";

const request = require("request"),
  path = require("path"),
  tmp = require("tmp"),
  tar = require("tar-fs"),
  zlib = require("zlib"),
  mkdirp = require("mkdirp"),
  fs = require("fs"),
  AdmZip = require("adm-zip"),
  exec = require("child_process").exec;
// Mapping from Node's `process.arch` to Golang `$GOARCH`
const ARCH_MAPPING = {
  ia32: "386",
  x64: "x86_64",
  arm: "arm"
};

// Mapping between Node's `process.platform` to Golang's
const PLATFORM_MAPPING = {
  darwin: "darwin",
  linux: "linux",
  win32: "windows",
  freebsd: "freebsd"
};

function getInstallationPath(callback) {
  // `npm bin` will output the path where binary files should be installed
  exec("npm bin", function(err, stdout, stderr) {
    let dir = null;
    if (err || stderr || !stdout || stdout.length === 0) {
      // We couldn't infer path from `npm bin`. Let's try to get it from
      // Environment variables set by NPM when it runs.
      // npm_config_prefix points to NPM's installation directory where `bin` folder is available
      // Ex: /Users/foo/.nvm/versions/node/v4.3.0
      let env = process.env;
      if (env && env.npm_config_prefix) {
        dir = path.join(env.npm_config_prefix, "bin");
      }
    } else {
      dir = stdout.trim();
    }

    mkdirp.sync(dir);

    callback(null, dir);
  });
}

function verifyAndPlaceBinary(binName, binPath, callback) {
  if (!fs.existsSync(path.join(binPath, binName)))
    return callback(
      `Downloaded binary does not contain the binary specified in configuration - ${binName}`
    );

  getInstallationPath(function(err, installationPath) {
    if (err)
      return callback("Error getting binary installation path from `npm bin`");

    console.log(
      `Install ${binName} successful into ${path.join(binPath, binName)}`
    );
    callback(null);
  });
}

function validateConfiguration(packageJson) {
  if (!packageJson.version) {
    return "'version' property must be specified";
  }

  if (!packageJson.install || typeof packageJson.install !== "object") {
    return "'install' property must be defined and be an object";
  }

  if (!packageJson.install.binaryName) {
    return "'binaryName' property is required";
  }

  if (!packageJson.install.url) {
    return "'url' property is required";
  }
}

function parsePackageJson() {
  if (!(process.arch in ARCH_MAPPING)) {
    console.error(
      "Installation is not supported for this architecture: " + process.arch
    );
    return;
  }

  if (!(process.platform in PLATFORM_MAPPING)) {
    console.error(
      "Installation is not supported for this platform: " + process.platform
    );
    return;
  }

  const packageJsonPath = path.join(".", "package.json");
  if (!fs.existsSync(packageJsonPath)) {
    console.error(
      "Unable to find package.json. " +
        "Please run this script at root of the package you want to be installed"
    );
    return;
  }

  let packageJson = JSON.parse(fs.readFileSync(packageJsonPath));
  let error = validateConfiguration(packageJson);
  if (error && error.length > 0) {
    console.error("Invalid package.json: " + error);
    return;
  }

  // We have validated the config. It exists in all its glory
  let binName = packageJson.install.binaryName;
  let binPath = "./bin";
  let url = packageJson.install.url;
  let version = packageJson.version;
  if (version[0] === "v") version = version.substr(1); // strip the 'v' if necessary v0.0.1 => 0.0.1

  // Binary name on Windows has .exe suffix
  if (process.platform === "win32") {
    binName += ".exe";
  }

  // Interpolate variables in URL, if necessary
  url = url.replace(/{{arch}}/g, ARCH_MAPPING[process.arch]);
  url = url.replace(/{{platform}}/g, PLATFORM_MAPPING[process.platform]);
  url = url.replace(/{{version}}/g, version);
  url = url.replace(/{{zipExt}}/, zipExt());

  return {
    binName: binName,
    binPath: binPath,
    url: url,
    version: version
  };
}

function zipExt() {
  if (isWindow()) {
    return "zip";
  }
  return "tar.gz";
}

function isWindow() {
  return PLATFORM_MAPPING[process.platform] === "windows";
}

const INVALID_INPUT = "Invalid inputs";
function install(callback) {
  let opts = parsePackageJson();
  if (!opts) return callback(INVALID_INPUT);

  mkdirp.sync(opts.binPath);
  let ungz = zlib.createGunzip();
  let untar = tar.extract(opts.binPath);

  ungz.on("error", callback);
  untar.on("error", callback);

  // First we will Un-GZip, then we will untar. So once untar is completed,
  // binary is downloaded into `binPath`. Verify the binary and call it good
  untar.on(
    "end",
    verifyAndPlaceBinary.bind(null, opts.binName, opts.binPath, callback)
  );

  console.log("Downloading from URL: " + opts.url);
  let req = request({ uri: opts.url });
  req.on(
    "error",
    callback.bind(null, "Error downloading from URL: " + opts.url)
  );
  req.on("response", function(res) {
    if (res.statusCode !== 200)
      return callback(
        "Error downloading binary. HTTP Status Code: " + res.statusCode
      );
    if (isWindow()) {
      // handler for window
      var tmpFile = tmp.fileSync();
      req.pipe(fs.createWriteStream(tmpFile.name)).on("finish", function() {
        var zip = new AdmZip(tmpFile.name);
        zip.extractAllTo(opts.binPath);
        console.log("install successful");
      });
    } else {
      req.pipe(ungz).pipe(untar);
    }
  });
}

function uninstall(callback) {
  let opts = parsePackageJson();
  getInstallationPath(function(err, installationPath) {
    if (err) callback("Error finding binary installation directory");

    try {
      fs.unlinkSync(path.join(installationPath, opts.binName));
    } catch (ex) {
      // Ignore errors when deleting the file.
    }
    console.log(`Uninstall ${opts.binName} successful`);
    return callback(null);
  });
}

// Parse command line arguments and call the right method
let actions = {
  install: install,
  uninstall: uninstall
};

let argv = process.argv;
if (argv && argv.length > 2) {
  let cmd = process.argv[2];
  if (!actions[cmd]) {
    console.log(
      "Invalid command to go-npm. `install` and `uninstall` are the only supported commands"
    );
    process.exit(1);
  }

  actions[cmd](function(err) {
    if (err) {
      console.error(err);
      process.exit(1);
    } else {
      process.exit(0);
    }
  });
} else {
  console.warn("Invalid command");
}
