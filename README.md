# mamuro
aplicacion realizada con Go y Vue

# Instalaciones necesarias
<ul>
  <li><a href="https://zincsearch-docs.zinc.dev/">Zincsearch</a></li>
  <li><a href="https://go.dev/">go</a> version 1.21.3 o superior</li>
  <li>Data de enron_mail indexado en zincsearch</li>
  <li><a href="https://nodejs.org/en">Node</a></li>
</ul>

# Ejecutar el programa
<ol>
  <li>Ejecutar Zincsearch con los siguientes parametros:</li>
  <pre>
ZINC_FIRST_ADMIN_USER=admin ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123</pre>
  <li>Clonar el proyecto</li>
  <pre> git clone https://github.com/Moisesmp75/mamuro.git</pre>
  <li>Generar el build</li>
  <pre>
  cd view
  npm install
  npm run build
  cd ..</pre>
  <li>Ejecutar el programa y especificar el puerto</li>
  <pre>
  go run . -port 3000</pre>
  <li>Ingresar al enlace con el port indicado</li>
  <pre>http://localhost:3000/</pre>
</ol>

# Visualizar la aplicacion

