#!/usr/bin/env bash
#install nginx and postgresql
sudo add-apt-repository 'deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get install nginx libpq-dev postgresql-common postgresql-contrib postgresql-10 -y

sudo mv /home/ubuntu/configs/default /etc/nginx/sites-available/default
sudo service nginx restart

#Setup DB
sudo su - postgres
createdb wizeweb
createuser wizeweb
psql -c "alter user wizeweb with encrypted password 'secret';"
psql -c "grant all privileges on database wizeweb to wizeweb ;"

#clone web server repo
cd $GOPATH/src
ssh-agent bash -c 'ssh-add ~/wize_web; git clone git@bitbucket.org:udt/wizeweb.git'
cd wizeweb/frontend
npm run build
mv /home/ubuntu/configs/db.conf /home/ubuntu/go/src/wizeweb/backend/conf/db.conf
cd ../backend
go build
./wizeweb &

