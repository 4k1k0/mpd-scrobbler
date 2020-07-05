#!/bin/bash

echo ""
echo -n "Username: "
read username
echo -n "Password: "
read -s password
echo ""
echo -n "MPD Server: (default localhost) "
read server
echo -n "MPD Port: (default 6600) "
read port
echo ""

touch .env

echo "USERNAME=$username" >> .env
echo "PASSWORD=$password" >> .env
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
