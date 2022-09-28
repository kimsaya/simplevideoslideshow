# simplevideoslideshow

# Fire Fox
- Turn off firefox security
  - New tap
  - about:config
  - security.fileuri.strict_origin_policy

- Install firefox
  - sudo apt install firefox

- Open firefox
  - firefox --new-tab http://localhost:8080


# Install Os
- Generic
  - sudo nand-sata-install
  - Boot from eMMC
  - format as ext4
- For Armbian
  - sudo armbian-config
  - System
  - Install
  - ext4

# Check OS
- hostnamectl
- ps -a
- ps -aux
- ps aux | grep appName

# Cron
- sudo apt update
- sudo apt install cron
- sudo systemctl enable cron
- File: /var/spool/cron/crontabs/
```cron
30 17 * * 2 curl http://www.google.com
```

- crontab -e
- sudo nano /etc/crontab

# Armbian 
- Autologin

https://github.com/armbian/build/blob/master/packages/bsp/common/usr/lib/armbian/armbian-firstlogin#L297-L304

- Autologin
  - Remove password event better
File: /etc/lightdm/lightdm.conf.d/11-armbian.conf
```sh
[Seat:*]
autologin-user=pi
autologin-user-timeout=0
user-session=xfce
greeter-show-manual-login=false
greeter-hide-users=false
allow-guest=false
```
Note that also removed the password for my user (in my case 'pi')

# System
- Monitor always on
https://superuser.com/questions/404012/how-to-prevent-the-monitor-from-turning-off-in-linux

# Users
- Remove password
```sh
$ sudo passwd -d sweta
Removing password for user sweta.
passwd: Success
```
