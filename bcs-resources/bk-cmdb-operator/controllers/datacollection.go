/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package controllers

import (
	"fmt"

	bkcmdbv1 "github.com/Tencent/bk-bcs/bcs-resources/bk-cmdb-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// reconcileDataCollection reconciles bk-cmdb datacollection
func (r *BkcmdbReconciler) reconcileDataCollection(instance *bkcmdbv1.Bkcmdb) error {
	deploy := makeDataCollectionDeploy(instance)
	if err := controllerutil.SetControllerReference(instance, deploy, r.Scheme); err != nil {
		return fmt.Errorf("failed to set datacollection deployment owner reference: %s", err.Error())
	}

	err := r.Client.CreateOrUpdateDeploy(deploy)
	if err != nil {
		return fmt.Errorf("failed to create or update datacollection deploy: %s", err.Error())
	}

	service := makeDataCollectionService(instance)
	if err := controllerutil.SetControllerReference(instance, service, r.Scheme); err != nil {
		return fmt.Errorf("failed to set datacollection service owner reference: %s", err.Error())
	}

	err = r.Client.CreateOrUpdateService(service)
	if err != nil {
		return fmt.Errorf("failed to create or update datacollection service: %s", err.Error())
	}

	return nil
}

// makeDataCollectionDeploy build datacollection deployment object
func makeDataCollectionDeploy(z *bkcmdbv1.Bkcmdb) *appsv1.Deployment {
	replicas := int32(z.Spec.DataCollection.Replicas)
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      z.GetName() + "-datacollection",
			Namespace: z.Namespace,
			Labels: map[string]string{
				"app":       "bk-cmdb",
				"release":   z.GetName(),
				"component": "datacollection",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":       "bk-cmdb",
					"release":   z.GetName(),
					"component": "datacollection",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":       "bk-cmdb",
						"release":   z.GetName(),
						"component": "datacollection",
					},
				},
				Spec: v1.PodSpec{
					Containers: makeDataCollectionContainers(z),
				},
			},
		},
	}
}

// makeDataCollectionContainers build datacollection containers object
func makeDataCollectionContainers(z *bkcmdbv1.Bkcmdb) []v1.Container {
	zkHost := z.GetName() + "-zookeeper"
	zkPort := defaultZkClientPort
	if z.Spec.Zookeeper != nil {
		zkHost = z.Spec.Zookeeper.Host
		zkPort = int(z.Spec.Zookeeper.Port)
	}
	rediscv := fmt.Sprintf("--regdiscv=%s:%d", zkHost, zkPort)
	return []v1.Container{
		{
			Name:            "datacollection",
			Image:           z.Spec.Image,
			ImagePullPolicy: "IfNotPresent",
			WorkingDir:      "/data/bin/bk-cmdb/cmdb_datacollection/",
			Command:         []string{"./cmdb_datacollection", "--addrport=$(POD_IP):80", rediscv, "--log-dir", "./logs", "--v", "3", "--enable-auth", "false"},
			LivenessProbe: &v1.Probe{
				Handler: v1.Handler{
					HTTPGet: &v1.HTTPGetAction{
						Path: "/healthz",
						Port: intstr.IntOrString{
							Type:   intstr.Int,
							IntVal: 80,
						},
					},
				},
				InitialDelaySeconds: 30,
				PeriodSeconds:       10,
			},
			ReadinessProbe: &v1.Probe{
				Handler: v1.Handler{
					HTTPGet: &v1.HTTPGetAction{
						Path: "/healthz",
						Port: intstr.IntOrString{
							Type:   intstr.Int,
							IntVal: 80,
						},
					},
				},
				InitialDelaySeconds: 30,
				PeriodSeconds:       10,
			},
			Env: []v1.EnvVar{
				{
					Name: "POD_IP",
					ValueFrom: &v1.EnvVarSource{
						FieldRef: &v1.ObjectFieldSelector{
							FieldPath: "status.podIP",
						},
					},
				},
			},
			Ports: []v1.ContainerPort{
				{
					ContainerPort: 80,
				},
			},
			Resources: z.Spec.DataCollection.Resources,
		},
	}
}

func makeDataCollectionService(z *bkcmdbv1.Bkcmdb) *v1.Service {
	return &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      z.GetName() + "-datacollection",
			Namespace: z.Namespace,
			Labels: map[string]string{
				"app":     "bk-cmdb",
				"release": z.GetName(),
			},
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Port: 80,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
					},
				},
			},
			Selector: map[string]string{
				"app":       "bk-cmdb",
				"release":   z.GetName(),
				"component": "datacollection",
			},
		},
	}
}
