#!/usr/bin/env node

const fs = require('fs');

const settingsPath = '.settings/config.json'

function getViewports() {
  const data = fs.readFileSync(settingsPath)
  const json = JSON.parse(data)
  return json.viewports
}

function getScenarios() {
    const data = fs.readFileSync(settingsPath)
    const json = JSON.parse(data)
    return json.scenarios
}

module.exports = {
    "id": "backstop_default",
    "viewports": getViewports(),
    "onBeforeScript": "puppet/onBefore.js",
    "onReadyScript": "puppet/onReady.js",
    "scenarios": getScenarios(),
    "paths": {
        "bitmaps_reference": "backstop_data/bitmaps_reference",
        "bitmaps_test": "backstop_data/bitmaps_test",
        "engine_scripts": "backstop_data/engine_scripts",
        "html_report": "backstop_data/html_report",
        "ci_report": "backstop_data/ci_report"
    },
    "report": ["browser", "json"],
    "engine": "puppeteer",
    "engineOptions": {
        "args": ["--no-sandbox"]
    },
    "asyncCaptureLimit": 5,
    "asyncCompareLimit": 50,
    "debug": false,
    "debugWindow": false
}


