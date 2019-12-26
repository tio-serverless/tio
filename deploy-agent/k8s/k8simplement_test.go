package k8s
//
//import (
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	apiv1 "k8s.io/api/core/v1"
//)
//
//func TestEnvMerge(t *testing.T) {
//
//	cases := []struct {
//		env1   []apiv1.EnvVar
//		env2   map[string]string
//		expect []apiv1.EnvVar
//	}{
//		{
//			env1: []apiv1.EnvVar{
//				apiv1.EnvVar{
//					Name:  "k1",
//					Value: "123",
//				},
//				apiv1.EnvVar{
//					Name:  "k2",
//					Value: "124",
//				},
//				apiv1.EnvVar{
//					Name:  "k3",
//					Value: "125",
//				},
//				apiv1.EnvVar{
//					Name:  "k4",
//					Value: "126",
//				},
//			},
//			env2: map[string]string{
//				"k1": "321",
//				"v1": "abc",
//				"v2": "cba",
//			},
//			expect: []apiv1.EnvVar{
//				apiv1.EnvVar{
//					Name:  "k1",
//					Value: "321",
//				},
//				apiv1.EnvVar{
//					Name:  "k2",
//					Value: "124",
//				},
//				apiv1.EnvVar{
//					Name:  "k3",
//					Value: "125",
//				},
//				apiv1.EnvVar{
//					Name:  "k4",
//					Value: "126",
//				},
//				apiv1.EnvVar{
//					Name:  "v1",
//					Value: "abc",
//				},
//				apiv1.EnvVar{
//					Name:  "v2",
//					Value: "cba",
//				},
//			},
//		},
//		{
//			env1: []apiv1.EnvVar{
//				apiv1.EnvVar{
//					Name:  "k1",
//					Value: "123",
//				},
//				apiv1.EnvVar{
//					Name:  "k2",
//					Value: "124",
//				},
//				apiv1.EnvVar{
//					Name:  "k3",
//					Value: "125",
//				},
//				apiv1.EnvVar{
//					Name:  "k4",
//					Value: "126",
//				},
//			},
//			env2: map[string]string{
//			},
//			expect: []apiv1.EnvVar{
//				apiv1.EnvVar{
//					Name:  "k1",
//					Value: "123",
//				},
//				apiv1.EnvVar{
//					Name:  "k2",
//					Value: "124",
//				},
//				apiv1.EnvVar{
//					Name:  "k3",
//					Value: "125",
//				},
//				apiv1.EnvVar{
//					Name:  "k4",
//					Value: "126",
//				},
//			},
//		},
//	}
//
//	for _, c := range cases {
//		actual := envMerge(c.env1, c.env2)
//		for i, e := range actual {
//			assert.EqualValues(t, c.expect[i].Name, e.Name, e.Name)
//			assert.EqualValues(t, c.expect[i].Value, e.Value, e.Value)
//		}
//	}
//}
