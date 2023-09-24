# WcGo

My own version of the Unix command line tool `wc`.

## Code Challenge

Solution to the First Coding Challege in [John Crickett's Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc/).

## How to build

Execute `make build` to create the executable. It will be saved in the `bin/` directory.

## How to run

Execute the binary with a file as an argument.

The following are the possible arguments:

```bash
# Outputs the number of bytes

> ./bin/wcgo -b test.text
Bytes => 335041

# Outputs the number of lines

> ./bin/wcgo -l test.text
Lines => 7144

# Outputs the number of chars

> ./bin/wcgo -c test.txt
Chars => 332145

# Outputs the number of words

> ./bin/wcgo -w test.text
Words => 58164

# Outputs with -b, -l and -w flags
> ./bin/wcgo test.text
Bytes => 335041
Lines => 7144
Words => 58164
```

You can also read from standard input

```bash
> cat test.txt | ./bin/wcgo -l
Lines => 7144
```

## License

[MIT](LICENSE) © Pravesh Goyal
