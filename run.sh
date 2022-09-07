nohup /home/orangepi/videoapp/app > /home/orangepi/videoapp/app.logs &
export DISPLAY=:0 && nohup firefox --kiosk http://localhost:8080/ > /home/orangepi/videoapp/firefox.logs &