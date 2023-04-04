# Exercice 2 - Shred

## Instructions

Shred tool in Go
Implement a Shred(path) function that will overwrite the given file (e.g. “randomfile”) 3 times with
random data and delete the file afterwards. Note that the file may contain any type of data.
You are expected to give information about the possible test cases for your Shred function, including the ones that you don’t implement, and implementing the full test coverage is a bonus :)
In a few lines briefly discuss the possible use cases for such a helper function as well as advantages and drawbacks of addressing them with this approach.

## Assumptions

Given its name and description this function would probably be used to securely erase a file to prevent an attacker from recovering it if they were given access to the disk.

I took the liberty to add a second parameter to the function to select an iterations count. A wrappeer function could be added to call the current `shred()` one with a reasonable default iterations count.

## Tests

Currently I only tested basic use cases and made sure the output file do not contain its original content. But to make sure this function serves its purpose, I would need to test if the shredding is secure. It may not be a trivial problem. Some ideas:

- check how/if the original `shred` binary is tested and if the quality of the shredding is evaluated
- before automating it, use an empiric method to test the shredding, with a tool such as Photorec, to make sure the basic implemtation is heading in the right direction
- check if existing forensic tools or libraries can be integrated to a test suite to get a "shredding quality score"

Since this function may face arbitrary big files, for a function destined to production I would write some benchmark with various file sizes and various iterations count.

## Use cases and discussion

This function is probably supposed to replace the [shred](https://linux.die.net/man/1/shred) tool from coreutils, or to mimic the same behavior. My functoin have the same limitation: overwrite in place, no filesystem specific strategy, no support for distributed filesystem, etc.

I also suppose we run this function on a HDD. This method of shredding on a SSD would risk reducing its lifespan and a better strategy could be to use the TRIM command of the SSD.
