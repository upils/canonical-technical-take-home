# Exercice 1 - Bootable Linux image via QEMU

## Instructions

In this exercise you are expected to create a shell script that will run in a Linux environment (will be tested on Ubuntu 20.04 LTS or 22.04 LTS). This shell script should create and run an AMD64 Linux filesystem image using QEMU that will print “hello world” after successful startup. Bonus points for creating a fully bootable filesystem image (but not mandatory). The system shouldn’t contain any user/session management or prompt for login information to access the filesystem.

You can use any version/flavor of the Linux kernel. The script can either download and build the kernel from source on the host environment or download a publicly available pre-built kernel.

The script shouldn’t ask for any user input unless superuser privileges are necessary for some functionality, therefore any additional information that you require for the script should be available in your repository.

The script should run within the working directory and not consume any other locations on the host file system.

## Notes

In the current state this is not working without asking user input. In the time frame I was unable to properly configure busybox to be built statically.