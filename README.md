# Create Serverless App

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme) [![Build Status](https://travis-ci.com/jazztong/create-serverless-app.svg?branch=master)](https://travis-ci.com/jazztong/create-serverless-app) [![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-0.1.2-red.svg)](https://conventionalcommits.org)

> This is the generator to generate serverless app project template

TODO: Fill out this long description.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Issue](#issue)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Install

```
npm install -g csla
```

If you are experiencing `EACCES: permission denied` errors during installation
using NPM then you can try:

```
npm i -g csla --unsafe-perm=true --allow-root
```

## Usage

```
csla -name my-lambda
```

## Issue

If you encounter permission issue please check your node install correctly, or use

```
npm config set unsafe-perm=true
```

## Maintainers

[@JazzTong](https://github.com/jazztong)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2019 Jazz Tong
