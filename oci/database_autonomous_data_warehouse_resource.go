// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDataWarehouseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createAutonomousDataWarehouse,
		Read:     readAutonomousDataWarehouse,
		Update:   updateAutonomousDataWarehouse,
		Delete:   deleteAutonomousDataWarehouse,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"high": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"low": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"medium": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_console_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAutonomousDataWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return CreateResource(d, sync)
}

func readAutonomousDataWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func updateAutonomousDataWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return UpdateResource(d, sync)
}

func deleteAutonomousDataWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type AutonomousDataWarehouseResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDataWarehouse
	DisableNotFoundRetries bool
}

func (s *AutonomousDataWarehouseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AutonomousDataWarehouseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseLifecycleStateProvisioning),
		string(oci_database.AutonomousDataWarehouseLifecycleStateStarting),
	}
}

func (s *AutonomousDataWarehouseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseLifecycleStateAvailable),
	}
}

func (s *AutonomousDataWarehouseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseLifecycleStateTerminating),
		string(oci_database.AutonomousDataWarehouseLifecycleStateUnavailable),
	}
}

func (s *AutonomousDataWarehouseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseLifecycleStateTerminated),
	}
}

func (s *AutonomousDataWarehouseResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseLifecycleStateProvisioning),
		string(oci_database.AutonomousDataWarehouseLifecycleStateUnavailable),
		string(oci_database.AutonomousDataWarehouseLifecycleStateScaleInProgress),
	}
}

func (s *AutonomousDataWarehouseResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseLifecycleStateAvailable),
	}
}

func (s *AutonomousDataWarehouseResourceCrud) Create() error {
	request := oci_database.CreateAutonomousDataWarehouseRequest{}

	if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
		tmp := adminPassword.(string)
		request.AdminPassword = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
		tmp := dataStorageSizeInTBs.(int)
		request.DataStorageSizeInTBs = &tmp
	}

	if dbName, ok := s.D.GetOkExists("db_name"); ok {
		tmp := dbName.(string)
		request.DbName = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.CreateAutonomousDataWarehouseDetailsLicenseModelEnum(licenseModel.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDataWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDataWarehouse
	return nil
}

func (s *AutonomousDataWarehouseResourceCrud) Get() error {
	request := oci_database.GetAutonomousDataWarehouseRequest{}

	tmp := s.D.Id()
	request.AutonomousDataWarehouseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDataWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDataWarehouse
	return nil
}

func (s *AutonomousDataWarehouseResourceCrud) Update() error {
	request := oci_database.UpdateAutonomousDataWarehouseRequest{}

	// @CODEGEN 09/2018: Cannot update the password and scale the ATP/ADW in the same request, only include changed properties in request
	if adminPassword, ok := s.D.GetOkExists("admin_password"); ok && s.D.HasChange("admin_password") {
		tmp := adminPassword.(string)
		request.AdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.AutonomousDataWarehouseId = &tmp
	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok && s.D.HasChange("cpu_core_count") {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}
	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok && s.D.HasChange("data_storage_size_in_tbs") {
		tmp := dataStorageSizeInTBs.(int)
		request.DataStorageSizeInTBs = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDataWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDataWarehouse
	return nil
}

func (s *AutonomousDataWarehouseResourceCrud) Delete() error {
	request := oci_database.DeleteAutonomousDataWarehouseRequest{}

	tmp := s.D.Id()
	request.AutonomousDataWarehouseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteAutonomousDataWarehouse(context.Background(), request)
	return err
}

func (s *AutonomousDataWarehouseResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{AutonomousDataWarehouseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func AutonomousDataWarehouseConnectionStringsToMap(obj *oci_database.AutonomousDataWarehouseConnectionStrings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.High != nil {
		result["high"] = string(*obj.High)
	}

	if obj.Low != nil {
		result["low"] = string(*obj.Low)
	}

	if obj.Medium != nil {
		result["medium"] = string(*obj.Medium)
	}

	return result
}
