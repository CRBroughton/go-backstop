# Go, Backstop!

Go, Backstop! is a TUI wrapper application for [BackstopJS](https://github.com/garris/BackstopJS), a automated visual regression testing application.

Go, Backstop! currently can:

- Create new custom viewports
- Create reference images
- Run existing tests
- Visualise results of JSON report in the terminal

## Dependencies & Getting Started

Go, Backstop! requires the following:

- The latest version of backstopJS docker image: ``docker pull backstopjs/backstopjs``


With the lastest `backstopjs` image downloaded, starting Go, Backstop! with `go run main.go` with initialise the project and `backstopjs`. Once initialised, you can create
a new test scenario in the main menu, and then test your scenario by running your existing test suite.

Go, Backstop! will create a `.settings` folder, where you can find your JSON configuration file.


If you wish to change the size of the table that shows your test results, simply change
the `resultstable` object in your configuration file; These values must be numbers.


---

To do list:

- ~~Add more robust dependency checker~~
- Creation of cookies - partially done
- Allow cookies to be used in tests
- ~~Creation of custom scenarios~~
- Improved TUI styling
- Dynamic loading screens - partially done
- ~~Visualise output of BackstopJS reporting into terminal~~
- Better seperate concerns of each view - mostly done
- ~~Handle failed tests (application currently `log.fatal`'s if it gets back a failed test)~~