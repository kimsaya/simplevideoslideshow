# simplevideoslideshow


Turn off firefox security
New tap
about:config
security.fileuri.strict_origin_policy


# Install Os
sudo nand-sata-install
Boot from eMMC
format as ext4

# Check OS
hostnamectl
ps -a
ps -aux

# Cron
sudo apt update
sudo apt install cron
sudo systemctl enable cron
/var/spool/cron/crontabs/
30 17 * * 2 curl http://www.google.com

crontab -e
sudo nano /etc/crontab

# Install firefox
sudo apt install firefox

# Open firefox
firefox --new-tab http://localhost:8080