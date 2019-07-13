rem @echo off
set bin=cron.exe
set binpath=%~dp0%bin%
set servicename=CronService
sc create %servicename% binPath= "%binpath%" start= auto displayName= "%servicename% displayName"
sc description %servicename% "%servicename% description"
net start %servicename%
:: sc delete %servicename%
pause