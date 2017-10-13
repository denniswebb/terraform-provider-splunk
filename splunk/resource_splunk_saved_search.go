package splunk

import (
	"fmt"
	"log"

	"github.com/denniswebb/go-splunk/splunk"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSplunkSavedSearch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSplunkSavedSearchCreate,
		Read:   resourceSplunkSavedSearchRead,
		Update: resourceSplunkSavedSearchUpdate,
		Delete: resourceSplunkSavedSearchDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				ForceNew: true,
				Type:     schema.TypeString,
				Required: true,
			},
			"search": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action_email": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"action_email_format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "plain",
			},
			"action_email_from": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_subject": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cron_schedule": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dispatch_earliest_time": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "-30m",
			},
			"dispatch_latest_time": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "now",
			},
			"dispatch_max_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_scheduled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_visible": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"max_concurrent": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}

func resourceSplunkSavedSearchCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*splunk.Client)

	savedSearch := savedSearchFromResourceData(d)

	log.Printf("[DEBUG] Splunk Saved Search create configuration: %#v", savedSearch)

	r, err := client.SavedSearchCreate(savedSearch)
	if err != nil {
		return fmt.Errorf("Failed to create saved search: %s", err)
	}

	d.SetId(r.Name)

	log.Printf("[INFO] Splunk Saved Search ID: %s", d.Id())

	return resourceSplunkSavedSearchRead(d, meta)
}

func resourceSplunkSavedSearchRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*splunk.Client)
	savedSearch, err := client.SavedSearchRead(d.Id())
	if err != nil {
		return err
	}

	d.SetId(savedSearch.Name)
	d.Set("name", savedSearch.Name)
	d.Set("search", savedSearch.Search)
	d.Set("action_emal", savedSearch.ActionEmail)
	d.Set("action_email_format", savedSearch.ActionEmailFormat)
	d.Set("action_email_from", savedSearch.ActionEmailFrom)
	d.Set("action_email_subject", savedSearch.ActionEmailSubject)
	d.Set("action_email_to", savedSearch.ActionEmailTo)
	d.Set("cron	_schedule", savedSearch.CronSchedule)
	d.Set("description", savedSearch.Description)
	d.Set("disabled", savedSearch.Disabled)
	d.Set("dispatch_earliest_time", savedSearch.DispatchEarliestTime)
	d.Set("dispatch_latest_time", savedSearch.DispatchLatestTime)
	d.Set("dispatch_max_count", savedSearch.DispatchMaxCount)
	d.Set("is_scheduled", savedSearch.IsScheduled)
	d.Set("is_visible", savedSearch.IsVisible)

	return nil
}

func resourceSplunkSavedSearchUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*splunk.Client)

	updateSavedSearch := savedSearchFromResourceData(d)

	log.Printf("[DEBUG] Splunk Saved Search update configuration: %#v", updateSavedSearch)
	_, err := client.SavedSearchUpdate(updateSavedSearch)
	if err != nil {
		return fmt.Errorf("Failed to update Splunk Saved Search: %s", err)
	}

	return resourceSplunkSavedSearchRead(d, meta)
}

func resourceSplunkSavedSearchDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*splunk.Client)

	log.Printf("[INFO] Deleting Splunk Saved Search: %s", d.Id())

	err := client.SavedSearchDelete(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Splunk Saved Search: %s", err)
	}

	return nil
}

func savedSearchFromResourceData(d *schema.ResourceData) *splunk.SavedSearch {
	savedSearch := &splunk.SavedSearch{
		Name:                 d.Get("name").(string),
		Search:               d.Get("search").(string),
		ActionEmailFormat:    d.Get("action_email_format").(string),
		ActionEmailFrom:      d.Get("action_email_from").(string),
		ActionEmailSubject:   d.Get("action_email_subject").(string),
		ActionEmailTo:        d.Get("action_email_to").(string),
		Description:          d.Get("description").(string),
		Disabled:             d.Get("disabled").(bool),
		DispatchEarliestTime: d.Get("dispatch_earliest_time").(string),
		DispatchLatestTime:   d.Get("dispatch_latest_time").(string),
		IsScheduled:          d.Get("is_scheduled").(bool),
		IsVisible:            d.Get("is_visible").(bool),
	}

	if cronSchedule, ok := d.GetOk("cron_schedule"); ok {
		savedSearch.CronSchedule = cronSchedule.(string)
	}

	if dispatchMaxCount, ok := d.GetOk("dispatch_max_count"); ok {
		savedSearch.DispatchMaxCount = dispatchMaxCount.(int)
	}

	return savedSearch
}
