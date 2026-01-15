# Contribute

## Tests

Run local centreon before start test with:

```bash
docker run -d --name centreon --privileged -v /projects/centreon/centreon/src:/fix -p 8080:80 webcenter/centreon:25.10-configured

docker exec -ti centreon bash


rsync -a /fix/ /usr/share/centreon/src
chmod -R 777 /var/cache/centreon/symfony
su centreon -c "/usr/share/centreon/bin/console cache:clear"
chmod -R 777 /var/cache/centreon/symfony
```

http://linuxworkgroup-hotmail-com-golang-web.che.home.webcenter.fr