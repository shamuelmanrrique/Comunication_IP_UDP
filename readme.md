
<!-- generar ssh -->
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_rsa

sudo apt-get install xclip
xclip -sel clip < ~/.ssh/id_rsa.pub

IMPORTANTE 

// buscar elementos o que te retornes los elementos
// de la session

// change
// chsh /bin/bash