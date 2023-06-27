# Go, Backstop!

Go, Backstop! is a TUI wrapper application for [BackstopJS](https://github.com/garris/BackstopJS), a automated visual regression testing application.

Go, Backstop! currently can:

- Create new custom viewports
- Create reference images
- Run existing tests
- Visualise results of JSON report in the terminal

## Dependencies & Getting Started

Go, Backstop! requires the following:

- Node + NPM
- Docker
- The latest version of backstopJS docker image: ``docker pull backstopjs/backstopjs``

---

To do list:

- Add more robust dependency checker
- Creation of cookies - partially done
- ~~Creation of custom scenarios~~
- Improved TUI styling
- Dynamic loading screens
- ~~Visualise output of BackstopJS reporting into terminal~~
- Better seperate concerns of each view
- ~~Handle failed tests (application currently `log.fatal`'s if it gets back a failed test)~~