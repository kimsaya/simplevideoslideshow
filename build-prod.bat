set VERSION=SDISPLAY_V_0_0_1_Linux
set DIRECTORY=.\builds\%VERSION%\

go env -w GOOS=linux GOARCH=arm
go build .\src\application\main.go
go env -w GOOS=windows GOARCH=amd64

if not exist .\builds mkdir .\builds
if not exist %DIRECTORY% mkdir %DIRECTORY%
if not exist %DIRECTORY%webs mkdir %DIRECTORY%webs
if not exist %DIRECTORY%inits mkdir %DIRECTORY%inits
if not exist %DIRECTORY%resource mkdir %DIRECTORY%resource
@REM if not exist %~dp0builds mkdir %~dp0builds
@REM if not exist %~dp0builds\apps mkdir %~dp0builds\apps
@REM if not exist %~dp0builds\pre-config mkdir %~dp0builds\pre-config
@REM if not exist %~dp0builds\pre-config\default mkdir %~dp0builds\pre-config\default
@REM if not exist %~dp0builds\release mkdir %~dp0builds\release

@REM del /s /q %~dp0builds\release\*

move /y .\main %DIRECTORY%app
copy /y .\src\application\webs\ %DIRECTORY%webs\
xcopy /y .\src\inits\* %DIRECTORY%inits\ /e
copy /y .\src\application\application.env %DIRECTORY%application.env
copy /y .\src\application\env-readme.md %DIRECTORY%env-readme.md
@REM move /y %~dp0poscar_face_agent builds\apps\poscar_face_agent_linux_arm_%VERSION%
@REM copy /y %~dp0init.sh-vscode\vscode\* builds\release\*
@REM copy /y %~dp0builds\pre-config\default\new.conf builds\release\new.conf
@REM copy /y %~dp0builds\readme.md builds\release\readme(%VERSION%).md

@REM explorer %~dp0builds\release\