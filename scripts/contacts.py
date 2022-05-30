#!/usr/bin/env python3
#
# This file needs "click" module, so run: pip3 install click
#
import click
import subprocess
import shutil


@click.group("contacts")
def contacts():
    pass


@contacts.command("configure")
def contacts_configure():
    cmd = "go get".split()
    proc = subprocess.Popen(cmd, stdout=subprocess.PIPE)
    __print(proc)


@contacts.command("build")
def contacts_build():
    cmd = "go build".split()
    proc = subprocess.Popen(cmd, stdout=subprocess.PIPE)
    __print(proc)


def __print(proc: subprocess.Popen):
    while True:
        output = proc.stdout.readline()
        if not output: break
        print(output.rstrip().decode('utf-8'))

contacts()
