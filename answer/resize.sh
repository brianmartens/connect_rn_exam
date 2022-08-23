#!/bin/sh

convert "$1.jpg" -resize 256x256^ -gravity center -extent 256x256 "$1.png"