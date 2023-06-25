#!/usr/bin/env node

const fs = require('fs')

const settingsPath = '.settings/config.json'

function getViewports() {
  const data = fs.readFileSync(settingsPath)
  const json = JSON.parse(data)
  console.log(json.viewports)
  return json.viewports
}

module.exports = {
    "id": "backstop_default",
    "viewports": getViewports(),
    "onBeforeScript": "puppet/onBefore.js",
    "onReadyScript": "puppet/onReady.js",
    "scenarios": [
        {
            "label": "BackstopJS Homepage",
            "cookiePath": "backstop_data/engine_scripts/cookies.json",
            "url": "https://garris.github.io/BackstopJS/",
            "referenceUrl": "",
            "readyEvent": "",
            "readySelector": "",
            "delay": 0,
            "hideSelectors": [],
            "removeSelectors": [],
            "hoverSelector": "",
            "clickSelector": "",
            "postInteractionWait": 0,
            "selectors": [],
            "selectorExpansion": true,
            "expect": 0,
            "misMatchThreshold": 0.1,
            "requireSameDimensions": true
        }
    ],
    "paths": {
        "bitmaps_reference": "backstop_data/bitmaps_reference",
        "bitmaps_test": "backstop_data/bitmaps_test",
        "engine_scripts": "backstop_data/engine_scripts",
        "html_report": "backstop_data/html_report",
        "ci_report": "backstop_data/ci_report"
    },
    "report": ["browser"],
    "engine": "puppeteer",
    "engineOptions": {
        "args": ["--no-sandbox"]
    },
    "asyncCaptureLimit": 5,
    "asyncCompareLimit": 50,
    "debug": false,
    "debugWindow": false
}


