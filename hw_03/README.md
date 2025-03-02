# Решение 3-ей домашней работы по курсу Microservice Architecture

## Последовательность шагов для запуска
1. Клонируем репозиторий
    ```bash
    git clone https://github.com/Vasilui/MicroserviceArchitecture.git
    ```
2. Переходим в директорию с манифестами
    ```bash
    cd hw_03/k8s
    ```
3. Запускаем манифесты на исполнение
    ```bash
    minikube start && kubectl apply -f .
    ```
4. Делаем запись в файле /etc/hosts(для macos), добавляем строчку со следующим содержимым
   ```bash
   127.0.0.1 arch.homework
   ```
5. Теперь можно запускать запросы из приложенной [postman коллекции](./microservice_architecture.postman_collection.json)