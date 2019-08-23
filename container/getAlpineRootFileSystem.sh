#!/bin/bash
# i did choose to make a basic FS with Alpine because it is very small
# choose your best ALPINE mirror here : http://nl.alpinelinux.org/alpine/MIRRORS.txt
AlpineMirror=http://pkg.adfinis-sygroup.ch/alpine
VERSION=2.10.4-r2
chroot_dir="$(pwd)/alpinelinux_${VERSION}/ROOTFS"
# based on the information in :
# https://wiki.alpinelinux.org/wiki/Alpine_Linux_in_a_chroot
# https://www.tldp.org/HOWTO/Bootdisk-HOWTO/buildroot.html
# another way would be to use debootstrap
# https://help.ubuntu.com/community/DebootstrapChroot
mkdir alpinelinux_${VERSION}
cd alpinelinux_${VERSION}
mkdir ROOTFS
wget ${AlpineMirror}/latest-stable/main/x86_64/apk-tools-static-${VERSION}.apk
wget ${AlpineMirror}/latest-stable/main/x86_64/readline-8.0.0-r0.apk
wget ${AlpineMirror}/latest-stable/main/x86_64/ncurses-libs-6.1_p20190518-r0.apk
wget ${AlpineMirror}/latest-stable/main/x86_64/bash-5.0.0-r0.apk
tar xvf ./apk-tools-static-${VERSION}.apk
sudo ./sbin/apk.static -X ${AlpineMirror}/latest-stable/main -U --allow-untrusted --root ROOTFS --initdb add alpine-base
sudo mknod -m 666 "${chroot_dir}/dev/full" c 1 7
sudo mknod -m 666 "${chroot_dir}/dev/ptmx" c 5 2
sudo mknod -m 644 "${chroot_dir}/dev/random" c 1 8
sudo mknod -m 644 "${chroot_dir}/dev/urandom" c 1 9
sudo mknod -m 666 "${chroot_dir}/dev/zero" c 1 5
sudo mknod -m 666 "${chroot_dir}/dev/tty" c 5 0
# uncomment if you need SCSI disc access:
#sudo mknod -m 666 "${chroot_dir}/dev/sda" b 8 0
#sudo mknod -m 666 "${chroot_dir}/dev/sda1" b 8 1
#sudo mknod -m 666 "${chroot_dir}/dev/sda2" b 8 2
#sudo mknod -m 666 "${chroot_dir}/dev/sda3" b 8 3
#sudo mknod -m 666 "${chroot_dir}/dev/sda4" b 8 4
#sudo mknod -m 666 "${chroot_dir}/dev/sda5" b 8 5
#sudo mknod -m 666 "${chroot_dir}/dev/sda6" b 8 6
#sudo mknod -m 666 "${chroot_dir}/dev/sdb" b 8 16
#sudo mknod -m 666 "${chroot_dir}/dev/sdb1" b 8 17
#sudo mknod -m 666 "${chroot_dir}/dev/sdb2" b 8 18
sudo tar -C ROOTFS -xzf ../readline-8.0.0-r0.apk
sudo tar -C ROOTFS -xzf ../ncurses-libs-6.1_p20190518-r0.apk
sudo tar -C ROOTFS -xzf ../bash-5.0.0-r0.apk
echo "now you are going to  chroot inside ALPINE FS type exit to come back"
sudo mount -t proc none "${chroot_dir}/proc"
sudo mount -o bind /sys "${chroot_dir}/sys"
sudo chroot ROOTFS /bin/bash
sudo umount "${chroot_dir}/proc"
sudo umount "${chroot_dir}/sys"


