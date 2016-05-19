##Rest api using Goji

###Build docker container
docker build -t sv-mtest .

###Start docker container
docker run -it --rm -p 10100:8000  sv-mtest

##Dependencies

##Install dependencies from vendor.yml
govend

###Update dependencies
govend -v -u -l

Scan project, download all dependencies,
and create a vendor.yml file to lock dependency versions

govend -v -l
