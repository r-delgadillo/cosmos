# cosmos
 eval $(minikube -p minikube docker-env)
 minikube image ls --format table
 

# kubectl run kafka-producer -ti --image=bitnami/kafka --rm=true --restart=Never --command -- /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic test --partitions 1 --replication-factor 1 --zookeeper # // # # kafka-zookeeper:2181


 kubectl run kafka-producer -ti --image=bitnami/kafka --rm=true --restart=Never --command -- /opt/bitnami/kafka/bin/kafka-console-producer.sh --topic test --broker-list kafka:9092

kubectl run kafka-consumer -ti --image=bitnami/kafka --rm=true --restart=Never --command -- /opt/bitnami/kafka/bin/kafka-console-consumer.sh --topic test --bootstrap-server kafka:9092 --from-beginning
