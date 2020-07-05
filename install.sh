#!/bin/bash

echo ""
echo -n "Username: "
read username
echo -n "Password: "
read -s password
echo ""
echo -n "Key: "
read key
echo -n "Secret Key: "
read -s secret
echo ""
echo -n "MPD Server: (default localhost) "
read server
echo -n "MPD Port: (default 6600) "
read port
echo ""

touch .env mpd-scrobbler.service

echo "USERNAME=$username" >> .env
echo "PASSWORD=$password" >> .env
echo "KEY=$key" >> .env
echo "SECRET=$secret" >> .env
mkdir $HOME/.config/mpd-scrobbler
touch $HOME/.config/mpd-scrobbler/database.db
echo "DATABASE=$HOME/.config/mpd-scrobbler/database.db" >> .env
if [[ "$server" != "" ]]
then
    echo "SERVER=$server" >> .env
fi
if [[ "$port" != "" ]]
then
    echo "PORT=$port" >> .env
fi
echo "" >> .env

go build

echo "" >> mpd-scrobbler.service
echo "[Unit]" >> mpd-scrobbler.service
echo "Description=MPD Scrobbler for Last.fm" >> mpd-scrobbler.service
echo "After=network.target" >> mpd-scrobbler.service
echo "" >> mpd-scrobbler.service
echo "[Service]" >> mpd-scrobbler.service
echo "PORT=1409" >> mpd-scrobbler.service
echo "Type=simple" >> mpd-scrobbler.service
echo "User=$USER" >> mpd-scrobbler.service
echo "WorkingDirectory=$GOPATH/xxx" >> mpd-scrobbler.service
echo "ExecStart=$GOPATH/xxx" >> mpd-scrobbler.service
echo "Restart=on-failure" >> mpd-scrobbler.service
echo "" >> mpd-scrobbler.service
echo "[Install]" >> mpd-scrobbler.service
echo "WantedBy=multi-user.target" >> mpd-scrobbler.service
echo "" >> mpd-scrobbler.service