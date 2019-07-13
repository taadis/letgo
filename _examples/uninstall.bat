rem @echo off
set servicename=CronService
net stop %servicename%
sc delete %servicename%
pause