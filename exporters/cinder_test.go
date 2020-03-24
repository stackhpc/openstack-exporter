package exporters

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

type CinderTestSuite struct {
	BaseOpenStackTestSuite
}

var cinderExpectedUp = `
# HELP openstack_cinder_agent_state agent_state
# TYPE openstack_cinder_agent_state counter
openstack_cinder_agent_state{adminState="enabled",disabledReason="",hostname="devstack@lvmdriver-1",service="cinder-volume",uuid="` + DEFAULT_UUID + `",zone="nova"} 1
openstack_cinder_agent_state{adminState="enabled",disabledReason="Test1",hostname="devstack",service="cinder-scheduler",uuid="` + DEFAULT_UUID + `",zone="nova"} 1
openstack_cinder_agent_state{adminState="enabled",disabledReason="Test2",hostname="devstack",service="cinder-backup",uuid="` + DEFAULT_UUID + `",zone="nova"} 1
# HELP openstack_cinder_limits_volume_max_gb limits_volume_max_gb
# TYPE openstack_cinder_limits_volume_max_gb gauge
openstack_cinder_limits_volume_max_gb{project_id="0c4e939acacf4376bdcd1129f1a054ad",tenant="admin"} 1000
openstack_cinder_limits_volume_max_gb{project_id="fdb8424c4e4f4c0ba32c52e2de3bd80e",tenant="alt_demo"} 1000
openstack_cinder_limits_volume_max_gb{project_id="0cbd49cbf76d405d9c86562e1d579bd3",tenant="demo"} 1000
openstack_cinder_limits_volume_max_gb{project_id="5961c443439d4fcebe42643723755e9d",tenant="invisible_to_admin"} 1000
openstack_cinder_limits_volume_max_gb{project_id="3d594eb0f04741069dbbb521635b21c7",tenant="service"} 1000
openstack_cinder_limits_volume_max_gb{project_id="43ebde53fc314b1c9ea2b8c5dc744927",tenant="swifttenanttest1"} 1000
openstack_cinder_limits_volume_max_gb{project_id="2db68fed84324f29bb73130c6c2094fb",tenant="swifttenanttest2"} 1000
openstack_cinder_limits_volume_max_gb{project_id="4b1eb781a47440acb8af9850103e537f",tenant="swifttenanttest4"} 1000
# HELP openstack_cinder_limits_volume_used_gb limits_volume_used_gb
# TYPE openstack_cinder_limits_volume_used_gb gauge
openstack_cinder_limits_volume_used_gb{project_id="0c4e939acacf4376bdcd1129f1a054ad",tenant="admin"} 0
openstack_cinder_limits_volume_used_gb{project_id="fdb8424c4e4f4c0ba32c52e2de3bd80e",tenant="alt_demo"} 0
openstack_cinder_limits_volume_used_gb{project_id="0cbd49cbf76d405d9c86562e1d579bd3",tenant="demo"} 0
openstack_cinder_limits_volume_used_gb{project_id="5961c443439d4fcebe42643723755e9d",tenant="invisible_to_admin"} 0
openstack_cinder_limits_volume_used_gb{project_id="3d594eb0f04741069dbbb521635b21c7",tenant="service"} 0
openstack_cinder_limits_volume_used_gb{project_id="43ebde53fc314b1c9ea2b8c5dc744927",tenant="swifttenanttest1"} 0
openstack_cinder_limits_volume_used_gb{project_id="2db68fed84324f29bb73130c6c2094fb",tenant="swifttenanttest2"} 0
openstack_cinder_limits_volume_used_gb{project_id="4b1eb781a47440acb8af9850103e537f",tenant="swifttenanttest4"} 0
# HELP openstack_cinder_pool_capacity_free_gb pool_capacity_free_gb
# TYPE openstack_cinder_pool_capacity_free_gb gauge
openstack_cinder_pool_capacity_free_gb{name="i666testhost@FastPool01",vendor_name="EMC",volume_backend_name="VNX_Pool"} 636.316
# HELP openstack_cinder_pool_capacity_total_gb pool_capacity_total_gb
# TYPE openstack_cinder_pool_capacity_total_gb gauge
openstack_cinder_pool_capacity_total_gb{name="i666testhost@FastPool01",vendor_name="EMC",volume_backend_name="VNX_Pool"} 1692.429
# HELP openstack_cinder_snapshots snapshots
# TYPE openstack_cinder_snapshots gauge
openstack_cinder_snapshots 1
# HELP openstack_cinder_up up
# TYPE openstack_cinder_up gauge
openstack_cinder_up 1
# HELP openstack_cinder_volume_size volume_size
# TYPE openstack_cinder_volume_size gauge
openstack_cinder_volume_size{bootable="false",name="test-volume-attachments",project_id="bab7d5c60cd041a0a36f7c4b6e1dd978",resource_id="6edbc2f4-1507-44f8-ac0d-eed1d2608d38",status="in-use",user_id="32779452fcd34ae1a53a797ac8a1e064",volume_type="lvmdriver-1"} 2
openstack_cinder_volume_size{bootable="true",name="test-volume",project_id="bab7d5c60cd041a0a36f7c4b6e1dd978",resource_id="173f7b48-c4c1-4e70-9acc-086b39073506",status="available",user_id="32779452fcd34ae1a53a797ac8a1e064",volume_type="lvmdriver-1"} 1
# HELP openstack_cinder_volume_status volume_status
# TYPE openstack_cinder_volume_status gauge
openstack_cinder_volume_status{bootable="false",id="6edbc2f4-1507-44f8-ac0d-eed1d2608d38",name="test-volume-attachments",project_id="bab7d5c60cd041a0a36f7c4b6e1dd978",size="2",status="in-use",volume_type="lvmdriver-1"} 5
openstack_cinder_volume_status{bootable="true",id="173f7b48-c4c1-4e70-9acc-086b39073506",name="test-volume",project_id="bab7d5c60cd041a0a36f7c4b6e1dd978",size="1",status="available",volume_type="lvmdriver-1"} 1
# HELP openstack_cinder_volume_status_counter volume_status_counter
# TYPE openstack_cinder_volume_status_counter gauge
openstack_cinder_volume_status_counter{status="attaching"} 0
openstack_cinder_volume_status_counter{status="available"} 1
openstack_cinder_volume_status_counter{status="awaiting-transfer"} 0
openstack_cinder_volume_status_counter{status="backing-up"} 0
openstack_cinder_volume_status_counter{status="creating"} 0
openstack_cinder_volume_status_counter{status="deleting"} 0
openstack_cinder_volume_status_counter{status="detaching"} 0
openstack_cinder_volume_status_counter{status="downloading"} 0
openstack_cinder_volume_status_counter{status="error"} 0
openstack_cinder_volume_status_counter{status="error_backing-up"} 0
openstack_cinder_volume_status_counter{status="error_deleting"} 0
openstack_cinder_volume_status_counter{status="error_extending"} 0
openstack_cinder_volume_status_counter{status="error_restoring"} 0
openstack_cinder_volume_status_counter{status="extending"} 0
openstack_cinder_volume_status_counter{status="in-use"} 1
openstack_cinder_volume_status_counter{status="maintenance"} 0
openstack_cinder_volume_status_counter{status="reserved"} 0
openstack_cinder_volume_status_counter{status="restoring-backup"} 0
openstack_cinder_volume_status_counter{status="retyping"} 0
openstack_cinder_volume_status_counter{status="uploading"} 0
# HELP openstack_cinder_volumes volumes
# TYPE openstack_cinder_volumes gauge
openstack_cinder_volumes 2
`

func (suite *CinderTestSuite) TestCinderExporter() {
	err := testutil.CollectAndCompare(*suite.Exporter, strings.NewReader(cinderExpectedUp))
	assert.NoError(suite.T(), err)
}
