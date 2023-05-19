// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfileCloudsqlExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDatabaseMigrationServiceConnectionProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfileCloudsqlExample(context),
			},
			{
				ResourceName:            "google_database_migration_service_connection_profile.cloudsqlprofile",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "mysql.0.password", "mysql.0.ssl.0.ca_certificate", "mysql.0.ssl.0.client_certificate", "mysql.0.ssl.0.client_key"},
			},
		},
	})
}

func testAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfileCloudsqlExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

resource "google_sql_database_instance" "cloudsqldb" {
  name             = "tf-test-my-database%{random_suffix}"
  database_version = "MYSQL_5_7"
  settings {
    tier = "db-n1-standard-1"
    deletion_protection_enabled = false
  }
  deletion_protection = false
}

resource "google_sql_ssl_cert" "sql_client_cert" {
  common_name = "tf-test-my-cert%{random_suffix}"
  instance = google_sql_database_instance.cloudsqldb.name
  
  depends_on = [google_sql_database_instance.cloudsqldb]
}

resource "google_sql_user" "sqldb_user" {
  name     = "tf-test-my-username%{random_suffix}"
  instance = google_sql_database_instance.cloudsqldb.name
  password = "tf-test-my-password%{random_suffix}"
  
  depends_on = [google_sql_ssl_cert.sql_client_cert]
}



resource "google_database_migration_service_connection_profile" "cloudsqlprofile" {
  location = "us-central1"
  connection_profile_id = "tf-test-my-fromprofileid%{random_suffix}"
  display_name = "tf-test-my-fromprofileid%{random_suffix}_display"
  labels = { 
    foo = "bar"
  }
  mysql {
    host = google_sql_database_instance.cloudsqldb.ip_address.0.ip_address
    port = 3306
    username = google_sql_user.sqldb_user.name
    password = google_sql_user.sqldb_user.password
    ssl {
      client_key = google_sql_ssl_cert.sql_client_cert.private_key
      client_certificate = google_sql_ssl_cert.sql_client_cert.cert
      ca_certificate = google_sql_ssl_cert.sql_client_cert.server_ca_cert
    }
    cloud_sql_id = "tf-test-my-database%{random_suffix}"
  }

  depends_on = [google_sql_user.sqldb_user]
}


resource "google_database_migration_service_connection_profile" "cloudsqlprofile_destination" {
  location = "us-central1"
  connection_profile_id = "tf-test-my-toprofileid%{random_suffix}"
  display_name = "tf-test-my-toprofileid%{random_suffix}_displayname"
  labels = { 
    foo = "bar"
  }
  cloudsql {
    settings {
      database_version = "MYSQL_5_7"
      user_labels = { 
        cloudfoo = "cloudbar"
      }
    tier = "db-n1-standard-1"
    storage_auto_resize_limit = "0"
    activation_policy = "ALWAYS"
    ip_config {
      enable_ipv4 = true
      require_ssl = "true"
    }
    auto_storage_increase = true
    data_disk_type = "PD_HDD"
    data_disk_size_gb = "11"
    zone = "us-central1-b"
    source_id = "projects/${data.google_project.project.project_id}/locations/us-central1/connectionProfiles/tf-test-my-fromprofileid%{random_suffix}"
    root_password = "testpasscloudsql"
    }


  }
  depends_on = [google_database_migration_service_connection_profile.cloudsqlprofile]
}
`, context)
}

func TestAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfilePostgresExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDatabaseMigrationServiceConnectionProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfilePostgresExample(context),
			},
			{
				ResourceName:            "google_database_migration_service_connection_profile.postgresprofile",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "postgresql.0.password", "postgresql.0.ssl.0.ca_certificate", "postgresql.0.ssl.0.client_certificate", "postgresql.0.ssl.0.client_key"},
			},
		},
	})
}

func testAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfilePostgresExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "postgresqldb" {
  name             = "tf-test-my-database%{random_suffix}"
  database_version = "POSTGRES_12"
  settings {
    tier = "db-custom-2-13312"
  }
  deletion_protection = false
}

resource "google_sql_ssl_cert" "sql_client_cert" {
  common_name = "tf-test-my-cert%{random_suffix}"
  instance    = google_sql_database_instance.postgresqldb.name

  depends_on = [google_sql_database_instance.postgresqldb]
}

resource "google_sql_user" "sqldb_user" {
  name     = "tf-test-my-username%{random_suffix}"
  instance = google_sql_database_instance.postgresqldb.name
  password = "tf-test-my-password%{random_suffix}"


  depends_on = [google_sql_ssl_cert.sql_client_cert]
}

resource "google_database_migration_service_connection_profile" "postgresprofile" {
  location = "us-central1"
  connection_profile_id = "tf-test-my-profileid%{random_suffix}"
  display_name = "tf-test-my-profileid%{random_suffix}_display"
  labels = { 
    foo = "bar" 
  }
  postgresql {
    host = google_sql_database_instance.postgresqldb.ip_address.0.ip_address
    port = 5432
    username = google_sql_user.sqldb_user.name
    password = google_sql_user.sqldb_user.password
    ssl {
      client_key = google_sql_ssl_cert.sql_client_cert.private_key
      client_certificate = google_sql_ssl_cert.sql_client_cert.cert
      ca_certificate = google_sql_ssl_cert.sql_client_cert.server_ca_cert
    }
    cloud_sql_id = "tf-test-my-database%{random_suffix}"
  }
  depends_on = [google_sql_user.sqldb_user]
}
`, context)
}

func TestAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfileAlloydbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDatabaseMigrationServiceConnectionProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfileAlloydbExample(context),
			},
			{
				ResourceName:            "google_database_migration_service_connection_profile.alloydbprofile",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "alloydb.0.settings.0.initial_user.0.password"},
			},
		},
	})
}

func testAccDatabaseMigrationServiceConnectionProfile_databaseMigrationServiceConnectionProfileAlloydbExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

resource "google_compute_network" "default" {
  name = "tf-test-alloydb-cp%{random_suffix}"
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          =  "private-ip-alloc%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = google_compute_network.default.id

  depends_on    = [google_compute_network.default]
}

resource "google_service_networking_connection" "vpc_connection" {
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]

  depends_on = [google_compute_global_address.private_ip_alloc]
}


resource "google_database_migration_service_connection_profile" "alloydbprofile" {
  location = "us-central1"
  connection_profile_id = "tf-test-my-profileid%{random_suffix}"
  display_name = "tf-test-my-profileid%{random_suffix}_display"
  labels = { 
    foo = "bar" 
  }
  alloydb {
    cluster_id = "dbmsalloycluster%{random_suffix}"
    settings {
      initial_user {
        user = "alloyuser%{random_suffix}"
        password = "alloypass%{random_suffix}"
      }
      vpc_network = "projects/${data.google_project.project.number}/global/networks/${google_compute_network.default.name}"
      labels  = { 
        alloyfoo = "alloybar" 
      }
      primary_instance_settings {
        id = "priminstid"
        machine_config {
          cpu_count = 2
        }
        database_flags = { 
        }
        labels = { 
          alloysinstfoo = "allowinstbar" 
        }
      }
    }
  }

  depends_on = [google_service_networking_connection.vpc_connection]
}
`, context)
}

func testAccCheckDatabaseMigrationServiceConnectionProfileDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_database_migration_service_connection_profile" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DatabaseMigrationServiceBasePath}}projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DatabaseMigrationServiceConnectionProfile still exists at %s", url)
			}
		}

		return nil
	}
}