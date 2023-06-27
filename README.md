# Ceifador Converter


A Ceifador tem como objetivo encontrar caminhos alternativos no sistema de anti fraud da Mobills e gerar conversões para os parceiros afiliados.


## Deploy

```shell
docker build -t ceifador:latest .
```

```shell
docker tag ceifador us-west1-docker.pkg.dev/conversion-toolkit/toolkitrepo/ceifador:latest
```

```shell
docker push us-west1-docker.pkg.dev/conversion-toolkit/toolkitrepo/ceifador:latest     
```


