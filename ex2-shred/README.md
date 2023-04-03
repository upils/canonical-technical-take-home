# Exercice 2 - Shred

## Instructions

Shred tool in Go
Implement a Shred(path) function that will overwrite the given file (e.g. “randomfile”) 3 times with
random data and delete the file afterwards. Note that the file may contain any type of data.
You are expected to give information about the possible test cases for your Shred function, including the ones that you don’t implement, and implementing the full test coverage is a bonus :)
In a few lines briefly discuss the possible use cases for such a helper function as well as advantages and drawbacks of addressing them with this approach.

## Assumptions

Given its name and descrption this function would probably be used to securely erase a file to prevent an attacker from recovering it if they were given access to the disk.

This function is probably supposed to replace the [shred](https://linux.die.net/man/1/shred) tool from coreutils, or to mimic the same behavior, so we apply the same limitations (overwrite in place, no filesystem specific strategy).

## Use cases and discussion

TODO