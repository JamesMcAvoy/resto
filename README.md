# Resto
Partie "salle" du projet C#/.NET (en go, dans linux, sans .NET).  
Implémente la vue et la partie du contrôleur gérant la salle.

###Installation
Installer les dépendances de pixel et d'ui, puis:
```
go get github.com/faiface/pixel
go get github.com/andlabs/ui/...
go get gopkg.in/h2non/gock.v1
go get github.com/JamesMcAvoy/resto
resto
cd $GOPATH/github.com/JamesMcAvoy/resto
go test ./... # ou go test -v ./...
golint ./...
```
