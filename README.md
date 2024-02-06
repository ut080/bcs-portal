# BCS Portal

BCS Portal is a suite of tools for the admin/personnel side of running the Blackhawk Cadet Squadron.

## Installation
- run `make install`
- A config file will be created at `${UserConfigDir}/bcs-portal/config/config.yaml`. You will want to update that config 
  file with the correct values for your instance of the application. See the Go docs for 
  [os.UserConfigDir](https://pkg.go.dev/os#UserConfigDir) to figure out where UserConfigDir is located for the relevant 
  operating system.

### Prerequisites
- [Go 1.21+](https://go.dev/)
- [TeX Live](https://www.tug.org/texlive/)
- [Graphviz](https://graphviz.org/)

### Set Up Local Environment
- [Clone this repo](https://github.com/ut080/bcs-portal)
- run `make install_dev`
- Update the config values in `${ProjectPath}/testdata/config/config.paml`

## Usage

## Contributing

## License

[MIT](https://choosealicense.com/licenses/mit/)