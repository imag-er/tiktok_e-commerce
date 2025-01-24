while ! curl http://0.0.0.0:2379/version; do
  echo 'Waiting for etcd to become available...'
  sleep 1
done

