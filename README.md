# ASCII-Art

Ascii-art is a program written in Go programming language which consists in receiving a string as an argument and displaying the string in a graphic representation in ASCII.

## Features

- Accepts text passed as an argument and displays it in an ASCII representation

## Getting Started

To get started with the program, the user needs to have Go installed in their local workstation. Click on the link below to download and install Go.

https://go.dev/doc/install

### Installation

1. Clone the repository

```
$ git clone https://github.com/johneliud/ASCII-Art.git
```

2. Change path to the cloned repository

```
$ cd ASCII-Art
```

## Usage

The project can be run in two ways. Use the commands below to run the program

```
$ go run . "Hello World!"
```

or

```
$ go run . "Hello World!" "shadow.txt"
```

The user has to specify the banner file to be accessed when running the second command from either of the below files:

- standard.txt

- shadow.txt

- thinkertoy.txt

## Examples

- Display the ASCII representation of the text "Hello World!"

```
$ go run . "Hello World" | cat -e
```

Output

```
 _    _          _   _                __          __                 _       _   _  $
| |  | |        | | | |               \ \        / /                | |     | | | | $
| |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | | | $
|  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | | | $
| |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | |_| $
|_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| (_) $
                                                                                    $
                                                                                    $
```

- Display the ASCII representation of the text "Hello World!" while specifying the banner file to access

```
$ go run . "Hello World!" "shadow.txt" | cat -e
```

Output

```
                                                                                       $
_|    _|          _| _|                _|          _|                   _|       _| _| $
_|    _|   _|_|   _| _|   _|_|         _|          _|   _|_|   _|  _|_| _|   _|_|_| _| $
_|_|_|_| _|_|_|_| _| _| _|    _|       _|    _|    _| _|    _| _|_|     _| _|    _| _| $
_|    _| _|       _| _| _|    _|         _|  _|  _|   _|    _| _|       _| _|    _|    $
_|    _|   _|_|_| _| _|   _|_|             _|  _|       _|_|   _|       _|   _|_|_| _| $
                                                                                       $
                                                                                       $
```

## Contact

Incase of further enquiries, please reach out via the email address johneliud4@gmail.com

## Credits

- [Denil](https://github.com/denilany)

- [Joab Owala](https://github.com/JoabOwala)
