sudo apt update
sudo apt -y upgrade
sudo apt install -y firefox
sudo apt install -y ntp
sudo apt install -y unclutter

sudo chmod 777 /home/orangepi/videoapp
cd /home/orangepi/videoapp

touch ./app.logs
touch ./firefox.logs

sudo mkdir ./src
sudo chmod 777 *

sudo crontab ./rootcron
crontab ./usercron
