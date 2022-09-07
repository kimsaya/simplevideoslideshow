apt update
apt -y upgrade
apt install -y firefox
apt install -y ntp

chmod 777 /home/orangepi/videoapp
cd /home/orangepi/videoapp

mkdir ./src
chmod 777 *

crontab ./initcron