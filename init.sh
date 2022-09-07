apt update
apt -y upgrade
apt install -y firefox
apt install -y ntp

cd /home/orangepi/videoapp

# chmod 777 ./run.sh
chmod 777 ./app

mkdir ./src

crontab ./initcron