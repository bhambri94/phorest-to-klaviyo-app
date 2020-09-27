### phorest-to-klaviyo-app

Setup

```
git clone https://github.com/bhambri94/phorest-to-klaviyo-app.git

cd phorest-to-klaviyo-app/

vi config.json 
//Change the configs with spreadhseet id and DB creds

docker build -t phorest-to-klaviyo-app:v1.0 .

docker images

docker run -it --name phorest-to-klaviyo-app -v $PWD/src:/go/src/phorest-to-klaviyo-app phorest-to-klaviyo-app:v1.0

```

#### Cron job

To setup a Daily Cron job, please follow following steps:
 
```
cd phorest-to-klaviyo-app/

Vi bash.sh

```
```
#!/bin/bash
sudo /usr/bin/docker restart phorest-to-klaviyo-app
```

Save the sheet script and run command 

```
chmod 777 bash.sh

Crontab -e

0 */12 * * * /path_to_phorest-to-klaviyo-app/bash.sh > /path_to_phorest-to-klaviyo-app_repo/phorest-to-klaviyo-app.logs

```
This above command written in crontab will run the script daily twice.
