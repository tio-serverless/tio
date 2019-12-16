apiVersion: v1
kind: ConfigMap
metadata:
  name: build
  namespace: {{ .Release.Namespace }}
data:
  config: |-
    {{ .Values.k8s.build.config}}
  tio-build.toml: |-
    log="{{ .Values.tio.log }}"
    port=80
    buildImage="tioserverless/build:{{ .Values.tio.version }}-{{ .Values.tio.branch }}"
    baseImage="{{ .Values.build.base }}"
    control="control.{{ .Release.Namespace }}.svc.cluster.local:80"
    [k8s]
      namespace="{{ .Values.k8s.build.namespace }}"
      config="/conf/k8s"
    [docker]
      user="{{ .Values.docker.user }}"
      passwd="{{ .Values.docker.passwd }}"

---

apiVersion: v1
kind: ConfigMap
metadata:
  name: control
  namespace: {{ .Release.Namespace }}
data:
  control.toml: |-
    log="{{ .Values.tio.log }}"
    rest_port=80
    build_agent_address="build.{{ .Release.Namespace }}.svc.cluster.local:80"
    deploy_agent_address="deploy.{{ .Release.Namespace }}.svc.cluster.local:80"
    rpc_port=8000
    [db]
      engine="postgres"
      connect="{{ .Values.db.connect }}"

---

apiVersion: v1
kind: ConfigMap
metadata:
  name: deploy
  namespace: tio
data:
  deploy-agent.toml: |-
    log="{{ .Values.tio.log }}"
    port=80
    [k8s]
      config="/k8s/config"
      consul="{{ .Values.k8s.deploy.consul }}"
      sidecar="{{ .Values.k8s.deploy.sidecar}}:{{ .Values.tio.version }}-{{ .Values.tio.branch }}"
  k8s-config: |-
    {{ .Values.k8s.deploy.config}}

---

apiVersion: v1
kind: ConfigMap
metadata:
  name: bootstarp
  namespace: {{ .Release.Namespace }}
data:
  envoy.yaml: |-
    admin:
      access_log_path: /tmp/admin_access.log
      address:
        socket_address: { address: 0.0.0.0, port_value: 9901 }

    dynamic_resources:
      cds_config:
        api_config_source:
          api_type: GRPC
          grpc_services:
            envoy_grpc:
              cluster_name: xds_cluster

    static_resources:
      listeners:
        - name: listener_0
          address:
            socket_address: { address: 0.0.0.0, port_value: 80 }
          filter_chains:
            - filters:
                - name: envoy.http_connection_manager
                  config:
                    stat_prefix: ingress_http
                    codec_type: AUTO
                    access_log:
                      name: envoy.file_access_log
                      config:
                        path: /dev/stdout
                    rds:
                      route_config_name: tio
                      config_source:
                        api_config_source:
                          api_type: GRPC
                          grpc_services:
                            envoy_grpc:
                              cluster_name: xds_cluster
                    http_filters:
                      - name: envoy.grpc_web
                      - name: envoy.cors
                      - name: envoy.router
      clusters:
        - name: xds_cluster
          connect_timeout: 60s
          type: STATIC
          lb_policy: ROUND_ROBIN
          http2_protocol_options: {}
          hosts: [{ socket_address: { address: 0.0.0.0, port_value: 8000 }}]
