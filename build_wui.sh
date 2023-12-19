#!/bin/bash

cd wui
npm run astro build
cp -R dist/* ../public
