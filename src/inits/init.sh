sudo apt update
sudo apt -y upgrade
sudo apt install -y firefox
sudo apt install -y ntp
sudo apt install -y unclutter

sudo chmod 777 ../

touch ../app.logs
touch ../firefox.logs

sudo mkdir ../resource
sudo chmod 777 ../*

sudo crontab ../inits/crons/rootcron
crontab ../inits/crons/usercron
