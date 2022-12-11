# shantropy

![example workflow](https://github.com/oceansofvectors/shantropy/actions/workflows/main.yml/badge.svg)

This project provides a command-line utility for finding substrings with high Shannon entropy in the contents of files in a directory.

## Installation

To install the utility, clone the repository and run go install:
 
```
$ git clone https://github.com/shantropy
$ cd shantropy
$ go build
```

## Usage

To use the utility, run the shantropy command and provide the directory path and minimum entropy as command-line arguments:

```
$ shantropy /path/to/directory 0.5
```

This will search the given directory for files containing substrings with Shannon entropy greater than or equal to the minimum entropy, and print the path of any such files to the standard output.

## Example

Suppose you have the following files in the /tmp directory:

```
/tmp/file1.txt: hello, world!
/tmp/file2.txt: abcdefghijklmnopqrstuvwxyz
/tmp/file3.txt: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
```

Running the shantropy command with a minimum entropy of 1.0 will produce the following output:

```
$ shantropy /tmp 1.0
/tmp/file2.txt
```

## Limitations

Currently, this utility only works with files containing ASCII text. It does not support binary files or files with other encodings.

## Contributing

To contribute to this project, please submit a pull request on GitHub.

## Todo

* Support files with encodings other than ASCII
* Optimize search algorithm for finding substrings with high Shannon entropy
* Make search algorithm asynchronous
* Provide more detailed output from the utility, such as specific substrings found and their entropy values


## License

This project is licensed under the MIT License. See the LICENSE file for more details.
