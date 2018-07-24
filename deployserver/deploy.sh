#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/newweb
git pull https://github.com/NodeBoy2/newweb.git
cd webserver/
./webserver &