Dans cet examen pratique, vous allez démontrer votre capacité à appliquer les connaissances acquises durant le cours pour construire une infrastructure sécurisée en utilisant Terraform, ainsi que l’outil Vault.
Vous allez créer une application en Golang et la conteneuriser à l’aide Docker. L'application interagira avec une instance de HashiCorp Vault déployée à l'aide de Docker et Terraform.
Vous allez mettre en place des méthodes d'authentification, créer des politiques de sécurité (polycies) et gérer des secrets dans Vault.
L'application Golang utilisera les identités définies dans Vault pour accéder et interagir avec les secrets.
De surcroît, vous allez déployer un conteneur Nginx dans Docker à l’aide de terraform.
Après déploiement du conteneur, votre host devra automatiquement effectuer un CURL de la page index de nginx et stocker son contenu dans un fichier (la démo sera demandée).
