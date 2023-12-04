
# EduTrack
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
![CodeFactor Grade](https://img.shields.io/codefactor/grade/github/Tom5521/EduTrack)

This is a program that manages a database of students and grades.


## Features

- Light/dark mode
- Fullscreen mode
- Cross platform
- User profile pictures and detailed fields to fill out
- SQLite managed database
- Embedded SQLite (no additional installations required)

## Installation

Install my-project with go or by executing the binary in its own folder

Install with go:
```bash
go install github.com/Tom5521/EduTrack
```

You can compile it, you need a C compiler for your system, git and the go compiler.

```bash
git clone https://github.com/Tom5521/EduTrack
git checkout <latest version>
go build -o EduTrack cmd/EduTrack/main.go
```

You can also download the binaries and unzip them, to run them in a portable way.



    
## Lessons Learned

- More advanced pointer handling
- More knowledge with the fyne framework and improve my knowledge regarding containers and layouts.
- Handling raw SQL in go.
- Management and adaptation to GORM.
- Improve the file structure of projects.
- Good practices in the Go language.
- Basic use of github actions.

## License

[MIT](https://choosealicense.com/licenses/mit/)


## Optimizations

- Change the use of raw SQL to GORM (I reduced ~1000 lines of code).
- Change the use of buttons to toolbars
- Improve data handling and data uploading
- Switching from storing data in yml files to SQLite files
- Improve packaging scripts for distribution.
- Move the data library that can be imported from other projects to modify databases and make scripts in general!


## Roadmap

- Add password protection for modifying the database

- Improve the graphic management system of grades and records


## Screenshots

![App Screenshot](./screenshots/Screenshot1.png)



## Support

For support open an issue or message me through my [reddit](https://www.reddit.com/u/Sad-Technician3861)!



## Running Tests

To run tests, run the following command

```bash
  go test -v tests/data/
  go test -v -tags delete tests/data/
```

