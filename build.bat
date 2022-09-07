set VERSION=V_0_8_1

go env -w GOOS=linux GOARCH=arm
go build main.go
go env -w GOOS=windows GOARCH=amd64

@REM if not exist %~dp0builds mkdir %~dp0builds
@REM if not exist %~dp0builds\apps mkdir %~dp0builds\apps
@REM if not exist %~dp0builds\pre-config mkdir %~dp0builds\pre-config
@REM if not exist %~dp0builds\pre-config\default mkdir %~dp0builds\pre-config\default
@REM if not exist %~dp0builds\release mkdir %~dp0builds\release

@REM del /s /q %~dp0builds\release\*

@REM copy /y %~dp0poscar_face_agent builds\release\poscar_face_agent
@REM move /y %~dp0poscar_face_agent builds\apps\poscar_face_agent_linux_arm_%VERSION%
@REM copy /y %~dp0init.sh-vscode\vscode\* builds\release\*
@REM copy /y %~dp0builds\pre-config\default\new.conf builds\release\new.conf
@REM copy /y %~dp0builds\readme.md builds\release\readme(%VERSION%).md

@REM explorer %~dp0builds\release\