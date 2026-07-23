# DSync

A local directory syncer that doesn't require large I/O operations.

## Installation

### GitHub Releases

Download latest release from [GitHub Releases](https://github.com/MarkIvory2973/dsync/releases/latest).

### Build from source

#### Requirements

- Go 1.26+
- nFPM
- UPX
- GNU Make
- Git

Clone the repository:

```bash
git clone https://github.com/MarkIvory2973/dsync.git
cd dsync
```

Install dependencies:

```bash
make install
```

Build binaries:

```bash
make build
```

Build packages (optional):

```bash
make package
```

Clean files:

```bash
make clean
```

## Usage

Run the following command:

```bash
./dsync path/to/src path/to/dst
```
