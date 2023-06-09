#!/bin/sh

set -e
set -o pipefail

TOP=$PWD

echo "Download kernel source code"
curl https://cdn.kernel.org/pub/linux/kernel/v6.x/linux-6.2.10.tar.xz | tar xJf -
echo "Download busybox source code"
curl https://busybox.net/downloads/busybox-1.36.0.tar.bz2 | tar xjf - 

echo "Install needed tools"
sudo apt-get update && sudo apt-get install -y \
        bc \
        bison \
        build-essential \
        cpio \
        flex \
        libelf-dev \
        libncurses-dev \
        libssl-dev \
        qemu-utils \
        qemu-system-x86

echo "Configure busybox build"
cd busybox-1.36.0
mkdir -pv $TOP/obj/busybox-x86
CONFIG_EXTRA_CFLAGS="-static" make O=$TOP/obj/busybox-x86 defconfig

echo "Build busybox"
cd $TOP/obj/busybox-x86
make -j$(nproc)
make install

echo "Prepare initramfs"
cd $TOP/initramfs/x86-busybox
mkdir -pv {bin,sbin,etc,proc,sys,usr/{bin,sbin}}
cp -av $TOP/obj/busybox-x86/_install/* .
cd $TOP/initramfs/x86-busybox
find . -print0 \
  | cpio --null -ov --format=newc \
  | gzip -9 > $TOP/obj/initramfs-busybox-x86.cpio.gz

echo "Configure kernel"
cd $TOP/linux-6.2.10
make O=$TOP/obj/linux-x86-basic x86_64_defconfig
make O=$TOP/obj/linux-x86-basic kvm_guest.config
echo "Build kernel"
make O=$TOP/obj/linux-x86-basic -j$(nproc)

echo "Boot"
cd $TOP
sudo qemu-system-x86_64 \
    -kernel obj/linux-x86-basic/arch/x86_64/boot/bzImage \
    -initrd obj/initramfs-busybox-x86.cpio.gz \
    -nographic -append "console=ttyS0" -enable-kvm