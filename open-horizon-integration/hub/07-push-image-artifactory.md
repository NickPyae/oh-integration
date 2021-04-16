# Pushing Updated Image to Artifactory

If we want to push our updated docker image to our Artifactory, we need to do this either from local machine, HOP or Franklin lab.

For example, if we want to push our updated `app-init` image to Artifactory, we need to log into our Artifactory first using `docker login amaas-eos-mw1.cec.lab.emc.com:5070` by providing service account user name and API key as password.

``` bash
cd app-init
docker login amaas-eos-mw1.cec.lab.emc.com:5070
docker build -t amaas-eos-mw1.cec.lab.emc.com:5070/hellosally/app-init:0.0.1
docker push amaas-eos-mw1.cec.lab.emc.com:5070/hellosally/app-init:0.0.1
```