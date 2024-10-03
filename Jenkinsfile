pipeline {
    agent any
    environment {
        PORT = '8080'
        DB = credentials('CONNECTION_URL')
        DOCKER_TAG = 'template-api'
        EXTERNAL_PORT = '8083'
        DB_PASS = credentials('DB_PASS')
        DB_NAME = 'database_name'
    }
    stages {
        stage('Migrate') {
            steps {
                sh 'cat migrationConf.json'
                sh 'rm ./migration || echo "cleaned cache"'
                sh 'wget https://github.com/Juanma1223/Golang_migrations/releases/download/v1.0/migration \
&& chmod +x ./migration \
&& ./migration -h $DOMAIN -d $DB_NAME -u lila -p $DB_PASS'
            }
        }
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'docker build -t $DOCKER_TAG .'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Stopping previous version...'
                sh 'docker stop $DOCKER_TAG || echo No hay nada que detener'
                sh 'docker rm $DOCKER_TAG || echo No hay nada que eliminar'
                echo 'Deploying....'
                sh 'docker run -d -e DB -e PORT -p $EXTERNAL_PORT:8080 --name $DOCKER_TAG $DOCKER_TAG'
            }
        }
    }
}
