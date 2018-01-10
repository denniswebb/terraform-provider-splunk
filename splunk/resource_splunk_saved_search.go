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
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
			"action_email_auth_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_auth_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_from": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_bcc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_cc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_subject": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "plain",
			},
			"action_email_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_inline": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_email_mailserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_max_results": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"action_email_max_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_message_alert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_pdfview": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_preprocess_results": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_email_report_cid_font_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_report_include_splunk_logo": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_email_report_paper_orientation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_report_paper_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_report_server_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_email_report_server_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_send_pdf": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_email_send_results": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_email_track_alert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_email_ttl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_email_use_ssl": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_email_use_tls": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_email_width_sort_columns": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_populate_lookup": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"action_populate_lookup_command": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_populate_lookup_dest": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_populate_lookup_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_populate_lookup_max_results": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"action_populate_lookup_max_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_populate_lookup_track_alert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_populate_lookup_ttl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_rss": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"action_rss_command": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_rss_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_rss_max_results": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"action_rss_max_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_rss_track_alert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_rss_ttl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_script": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"action_script_command": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_script_filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_script_hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_script_max_results": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"action_script_max_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_script_track_alert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_script_ttl": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_slack": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"action_slack_channel": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_slack_message": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_summary_index": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"action_summary_index_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_summary_index_command": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_summary_index_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_summary_index_inline": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_summary_index_max_results": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"action_summary_index_max_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action_summary_index_track_alert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"action_summary_index_ttl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"actions": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_digest_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"alert_expires": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "24h",
			},
			"alert_severity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"alert_suppress": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"alert_suppress_fields": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_suppress_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_track": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
			"alert_comparator": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "greater than",
			},
			"alert_condition": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_threshold": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
			"alert_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "number of events",
			},
			"auto_summarize": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_summarize_command": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_summarize_cron_schedule": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_summarize_dispatch_earliest_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_summarize_dispatch_latest_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_summarize_dispatch_time_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_summarize_dispatch_ttl": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_summarize_max_disabled_buckets": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_summarize_max_summary_ratio": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"auto_summarize_max_summary_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_summarize_max_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_summarize_suspend_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_summarize_timespan": {
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
			"dispatch_buckets": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dispatch_indexed_realtime": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dispatch_lookups": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dispatch_max_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dispatch_reduce_freq": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dispatch_rt_backfill": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dispatch_spawn_process": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dispatch_time_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dispatch_ttl": {
				Type:     schema.TypeString,
				Optional: true,
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
				Default:  500000,
			},
			"displayview": {
				Type:     schema.TypeString,
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
			"realtime_schedule": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"request_ui_dispatch_app": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "search",
			},
			"request_ui_dispatch_view": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "search",
			},
			"restart_on_searchpeer_add": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"run_on_startup": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vsid": {
				Type:     schema.TypeString,
				Optional: true,
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
	d.Set("action_email_auth_username", savedSearch.ActionEmailAuthUsername)
	d.Set("action_email_auth_password", savedSearch.ActionEmailAuthPassword)
	d.Set("action_email_bcc", savedSearch.ActionEmailBCC)
	d.Set("action_email_cc", savedSearch.ActionEmailCC)
	d.Set("action_email_format", savedSearch.ActionEmailFormat)
	d.Set("action_email_hostname", savedSearch.ActionEmailHostname)
	d.Set("action_email_inline", savedSearch.ActionEmailInline)
	d.Set("action_email_mailserver", savedSearch.ActionEmailMailserver)
	d.Set("action_email_max_results", savedSearch.ActionEmailMaxResults)
	d.Set("action_email_max_time", savedSearch.ActionEmailMaxTime)
	d.Set("action_email_message_alert", savedSearch.ActionEmailMessageAlert)
	d.Set("action_email_pdfview", savedSearch.ActionEmailPDFView)
	d.Set("action_email_preprocess_results", savedSearch.ActionEmailPreprocessResults)
	d.Set("action_email_report_cid_font_list", savedSearch.ActionEmailReportCIDFontList)
	d.Set("action_email_report_include_splunk_logo", savedSearch.ActionEmailReportIncludeSplunkLogo)
	d.Set("action_email_report_paper_orientation", savedSearch.ActionEmailReportPaperOrientation)
	d.Set("action_email_report_paper_size", savedSearch.ActionEmailReportPaperSize)
	d.Set("action_email_report_server_enabled", savedSearch.ActionEmailReportServerEnabled)
	d.Set("action_email_report_server_url", savedSearch.ActionEmailReportServerURL)
	d.Set("action_email_send_pdf", savedSearch.ActionEmailSendPDF)
	d.Set("action_email_send_results", savedSearch.ActionEmailSendResults)
	d.Set("action_email_track_alert", savedSearch.ActionEmailTrackAlert)
	d.Set("action_email_ttl", savedSearch.ActionEmailTTL)
	d.Set("action_email_use_ssl", savedSearch.ActionEmailUseSSL)
	d.Set("action_email_use_tls", savedSearch.ActionEmailUseTLS)
	d.Set("action_email_width_sort_columns", savedSearch.ActionEmailWidthSortColumns)
	d.Set("action_populate_lookup", savedSearch.ActionPopulateLookup)
	d.Set("action_populate_lookup_command", savedSearch.ActionPopulateLookupCommand)
	d.Set("action_populate_lookup_dest", savedSearch.ActionPopulateLookupDest)
	d.Set("action_populate_lookup_hostname", savedSearch.ActionPopulateLookupHostname)
	d.Set("action_populate_lookup_max_results", savedSearch.ActionPopulateLookupMaxResults)
	d.Set("action_populate_lookup_max_time", savedSearch.ActionPopulateLookupMaxTime)
	d.Set("action_populate_lookup_track_alert", savedSearch.ActionPopulateLookupTrackAlert)
	d.Set("action_populate_lookup_ttl", savedSearch.ActionPopulateLookupTTL)
	d.Set("action_rss", savedSearch.ActionRSS)
	d.Set("action_rss_command", savedSearch.ActionRSSCommand)
	d.Set("action_rss_hostname", savedSearch.ActionRSSHostname)
	d.Set("action_rss_max_results", savedSearch.ActionRSSMaxResults)
	d.Set("action_rss_max_time", savedSearch.ActionRSSMaxTime)
	d.Set("action_rss_track_alert", savedSearch.ActionRSSTrackAlert)
	d.Set("action_rss_ttl", savedSearch.ActionRSSTTL)
	d.Set("action_slack", savedSearch.ActionSlack)
	d.Set("action_slack_channel", savedSearch.ActionSlackChannel)
	d.Set("action_slack_message", savedSearch.ActionSlackMessage)
	d.Set("action_script", savedSearch.ActionScript)
	d.Set("action_script_command", savedSearch.ActionScriptCommand)
	d.Set("action_script_filename", savedSearch.ActionScriptFilename)
	d.Set("action_script_hostname", savedSearch.ActionScriptHostname)
	d.Set("action_script_max_results", savedSearch.ActionScriptMaxResults)
	d.Set("action_script_max_time", savedSearch.ActionScriptMaxTime)
	d.Set("action_script_track_alert", savedSearch.ActionScriptTrackAlert)
	d.Set("action_script_ttl", savedSearch.ActionScriptTTL)
	d.Set("action_summary_index", savedSearch.ActionSummaryIndex)
	d.Set("action_summary_index_name", savedSearch.ActionSummaryIndexName)
	d.Set("action_summary_index_command", savedSearch.ActionSummaryIndexCommand)
	d.Set("action_summary_index_hostname", savedSearch.ActionSummaryIndexHostname)
	d.Set("action_summary_index_inline", savedSearch.ActionSummaryIndexInline)
	d.Set("action_summary_index_max_results", savedSearch.ActionSummaryIndexMaxResults)
	d.Set("action_summary_index_max_time", savedSearch.ActionSummaryIndexMaxTime)
	d.Set("action_summary_index_track_alert", savedSearch.ActionSummaryIndexTrackAlert)
	d.Set("action_summary_index_ttl", savedSearch.ActionSummaryIndexTTL)
	d.Set("actions", savedSearch.Actions)
	d.Set("alert_digest_mode", savedSearch.AlertDigestMode)
	d.Set("alert_expires", savedSearch.AlertExpires)
	d.Set("alert_severity", savedSearch.AlertSeverity)
	d.Set("alert_suppress", savedSearch.AlertSuppress)
	d.Set("alert_suppress_fields", savedSearch.AlertSuppressFields)
	d.Set("alert_suppress_period", savedSearch.AlertSuppressPeriod)
	d.Set("alert_track", savedSearch.AlertTrack)
	d.Set("alert_comparator", savedSearch.AlertComparator)
	d.Set("alert_condition", savedSearch.AlertCondition)
	d.Set("alert_threshold", savedSearch.AlertThreshold)
	d.Set("alert_type", savedSearch.AlertType)
	d.Set("auto_summarize", savedSearch.AutoSummarize)
	d.Set("auto_summarize_command", savedSearch.AutoSummarizeCommand)
	d.Set("auto_summarize_cron_schedule", savedSearch.AutoSummarizeCronSchedule)
	d.Set("auto_summarize_dispatch_earliest_time", savedSearch.AutoSummarizeDispatchEarliestTime)
	d.Set("auto_summarize_dispatch_latest_time", savedSearch.AutoSummarizeDispatchLatestTime)
	d.Set("auto_summarize_dispatch_time_format", savedSearch.AutoSummarizeDispatchTimeFormat)
	d.Set("auto_summarize_dispatch_ttl", savedSearch.AutoSummarizeDispatchTTL)
	d.Set("auto_summarize_max_disabled_buckets", savedSearch.AutoSummarizeMaxDisabledBuckets)
	d.Set("auto_summarize_max_summary_ratio", savedSearch.AutoSummarizeMaxSummaryRatio)
	d.Set("auto_summarize_max_summary_size", savedSearch.AutoSummarizeMaxSummarySize)
	d.Set("auto_summarize_max_time", savedSearch.AutoSummarizeMaxTime)
	d.Set("auto_summarize_suspend_period", savedSearch.AutoSummarizeSuspendPeriod)
	d.Set("auto_summarize_timespan", savedSearch.AutoSummarizeTimespan)
	d.Set("dispatch_buckets", savedSearch.DispatchBuckets)
	d.Set("dispatch_indexed_realtime", savedSearch.DispatchIndexedRealtime)
	d.Set("dispatch_lookups", savedSearch.DispatchLookups)
	d.Set("dispatch_max_time", savedSearch.DispatchMaxTime)
	d.Set("dispatch_reduce_freq", savedSearch.DispatchReduceFreq)
	d.Set("dispatch_rt_backfill", savedSearch.DispatchRtBackfill)
	d.Set("dispatch_spawn_process", savedSearch.DispatchSpawnProcess)
	d.Set("dispatch_time_format", savedSearch.DispatchTimeFormat)
	d.Set("dispatch_ttl", savedSearch.DispatchTTL)
	d.Set("dispatch_earliest_time", savedSearch.DispatchEarliestTime)
	d.Set("dispatch_latest_time", savedSearch.DispatchLatestTime)
	d.Set("dispatch_max_count", savedSearch.DispatchMaxCount)
	d.Set("displayview", savedSearch.DisplayView)
	d.Set("max_concurrent", savedSearch.MaxConcurrent)
	d.Set("realtime_schedule", savedSearch.RealtimeSchedule)
	d.Set("request_ui_dispatch_app", savedSearch.RequestUIDispatchApp)
	d.Set("request_ui_dispatch_view", savedSearch.RequestUIDispatchView)
	d.Set("restart_on_searchpeer_add", savedSearch.RestartOnSearchPeerAdd)
	d.Set("run_on_startup", savedSearch.RunOnStartup)
	d.Set("vsid", savedSearch.VSID)
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
		Name:                               d.Get("name").(string),
		Search:                             d.Get("search").(string),
		ActionEmailAuthPassword:            d.Get("action_email_auth_password").(string),
		ActionEmailAuthUsername:            d.Get("action_email_auth_username").(string),
		ActionEmailBCC:                     d.Get("action_email_bcc").(string),
		ActionEmailCC:                      d.Get("action_email_cc").(string),
		ActionEmailFormat:                  d.Get("action_email_format").(string),
		ActionEmailFrom:                    d.Get("action_email_from").(string),
		ActionEmailHostname:                d.Get("action_email_hostname").(string),
		ActionEmailInline:                  d.Get("action_email_inline").(bool),
		ActionEmailMailserver:              d.Get("action_email_mailserver").(string),
		ActionEmailMaxResults:              d.Get("action_email_max_results").(int),
		ActionEmailMaxTime:                 d.Get("action_email_max_time").(string),
		ActionEmailMessageAlert:            d.Get("action_email_message_alert").(string),
		ActionEmailPDFView:                 d.Get("action_email_pdfview").(string),
		ActionEmailPreprocessResults:       d.Get("action_email_preprocess_results").(string),
		ActionEmailReportCIDFontList:       d.Get("action_email_report_cid_font_list").(string),
		ActionEmailReportIncludeSplunkLogo: d.Get("action_email_report_include_splunk_logo").(bool),
		ActionEmailReportPaperOrientation:  d.Get("action_email_report_paper_orientation").(string),
		ActionEmailReportPaperSize:         d.Get("action_email_report_paper_size").(string),
		ActionEmailReportServerEnabled:     d.Get("action_email_report_server_enabled").(bool),
		ActionEmailReportServerURL:         d.Get("action_email_report_server_url").(string),
		ActionEmailSendPDF:                 d.Get("action_email_send_pdf").(bool),
		ActionEmailSendResults:             d.Get("action_email_send_results").(bool),
		ActionEmailSubject:                 d.Get("action_email_subject").(string),
		ActionEmailTo:                      d.Get("action_email_to").(string),
		ActionEmailTrackAlert:              d.Get("action_email_track_alert").(bool),
		ActionEmailTTL:                     d.Get("action_email_ttl").(string),
		ActionEmailUseSSL:                  d.Get("action_email_use_ssl").(bool),
		ActionEmailUseTLS:                  d.Get("action_email_use_tls").(bool),
		ActionEmailWidthSortColumns:        d.Get("action_email_width_sort_columns").(bool),
		ActionPopulateLookupCommand:        d.Get("action_populate_lookup_command").(string),
		ActionPopulateLookupDest:           d.Get("action_populate_lookup_dest").(string),
		ActionPopulateLookupHostname:       d.Get("action_populate_lookup_hostname").(string),
		ActionPopulateLookupMaxResults:     d.Get("action_populate_lookup_max_results").(int),
		ActionPopulateLookupMaxTime:        d.Get("action_populate_lookup_max_time").(string),
		ActionPopulateLookupTrackAlert:     d.Get("action_populate_lookup_track_alert").(bool),
		ActionPopulateLookupTTL:            d.Get("action_populate_lookup_ttl").(string),
		ActionRSSCommand:                   d.Get("action_rss_command").(string),
		ActionRSSHostname:                  d.Get("action_rss_hostname").(string),
		ActionRSSMaxResults:                d.Get("action_rss_max_results").(int),
		ActionRSSMaxTime:                   d.Get("action_rss_max_time").(string),
		ActionRSSTrackAlert:                d.Get("action_rss_track_alert").(bool),
		ActionRSSTTL:                       d.Get("action_rss_ttl").(string),
		Actions:                            d.Get("actions").(string),
		ActionScriptCommand:                d.Get("action_script_command").(string),
		ActionScriptFilename:               d.Get("action_script_filename").(string),
		ActionScriptHostname:               d.Get("action_script_hostname").(string),
		ActionScriptMaxResults:             d.Get("action_script_max_results").(int),
		ActionScriptMaxTime:                d.Get("action_script_max_time").(string),
		ActionScriptTrackAlert:             d.Get("action_script_track_alert").(bool),
		ActionScriptTTL:                    d.Get("action_script_ttl").(string),
		ActionSlack:                        d.Get("action_slack").(bool),
		ActionSlackChannel:                 d.Get("action_slack_channel").(string),
		ActionSlackMessage:                 d.Get("action_slack_message").(string),
		ActionSummaryIndexCommand:          d.Get("action_summary_index_command").(string),
		ActionSummaryIndexHostname:         d.Get("action_summary_index_hostname").(string),
		ActionSummaryIndexInline:           d.Get("action_summary_index_inline").(bool),
		ActionSummaryIndexMaxResults:       d.Get("action_summary_index_max_results").(int),
		ActionSummaryIndexMaxTime:          d.Get("action_summary_index_max_time").(string),
		ActionSummaryIndexName:             d.Get("action_summary_index_name").(string),
		ActionSummaryIndexTrackAlert:       d.Get("action_summary_index_track_alert").(bool),
		ActionSummaryIndexTTL:              d.Get("action_summary_index_ttl").(string),
		AlertComparator:                    d.Get("alert_comparator").(string),
		AlertCondition:                     d.Get("alert_condition").(string),
		AlertDigestMode:                    d.Get("alert_digest_mode").(bool),
		AlertExpires:                       d.Get("alert_expires").(string),
		AlertSeverity:                      d.Get("alert_severity").(int),
		AlertSuppress:                      d.Get("alert_suppress").(bool),
		AlertSuppressFields:                d.Get("alert_suppress_fields").(string),
		AlertSuppressPeriod:                d.Get("alert_suppress_period").(string),
		AlertThreshold:                     d.Get("alert_threshold").(string),
		AlertTrack:                         d.Get("alert_track").(string),
		AlertType:                          d.Get("alert_type").(string),
		AutoSummarize:                      d.Get("auto_summarize").(bool),
		AutoSummarizeCommand:               d.Get("auto_summarize_command").(string),
		AutoSummarizeCronSchedule:          d.Get("auto_summarize_cron_schedule").(string),
		AutoSummarizeDispatchEarliestTime:  d.Get("auto_summarize_dispatch_earliest_time").(string),
		AutoSummarizeDispatchLatestTime:    d.Get("auto_summarize_dispatch_latest_time").(string),
		AutoSummarizeDispatchTimeFormat:    d.Get("auto_summarize_dispatch_time_format").(string),
		AutoSummarizeDispatchTTL:           d.Get("auto_summarize_dispatch_ttl").(string),
		AutoSummarizeMaxDisabledBuckets:    d.Get("auto_summarize_max_disabled_buckets").(int),
		AutoSummarizeMaxSummaryRatio:       d.Get("auto_summarize_max_summary_ratio").(float64),
		AutoSummarizeMaxSummarySize:        d.Get("auto_summarize_max_summary_size").(int),
		AutoSummarizeMaxTime:               d.Get("auto_summarize_max_time").(int),
		AutoSummarizeSuspendPeriod:         d.Get("auto_summarize_suspend_period").(string),
		AutoSummarizeTimespan:              d.Get("auto_summarize_timespan").(string),
		CronSchedule:                       d.Get("cron_schedule").(string),
		Description:                        d.Get("description").(string),
		Disabled:                           d.Get("disabled").(bool),
		DispatchBuckets:                    d.Get("dispatch_buckets").(int),
		DispatchEarliestTime:               d.Get("dispatch_earliest_time").(string),
		DispatchIndexedRealtime:            d.Get("dispatch_indexed_realtime").(bool),
		DispatchLatestTime:                 d.Get("dispatch_latest_time").(string),
		DispatchLookups:                    d.Get("dispatch_lookups").(bool),
		DispatchMaxCount:                   d.Get("dispatch_max_count").(int),
		DispatchMaxTime:                    d.Get("dispatch_max_time").(int),
		DispatchReduceFreq:                 d.Get("dispatch_reduce_freq").(int),
		DispatchRtBackfill:                 d.Get("dispatch_rt_backfill").(bool),
		DispatchSpawnProcess:               d.Get("dispatch_spawn_process").(bool),
		DispatchTimeFormat:                 d.Get("dispatch_time_format").(string),
		DispatchTTL:                        d.Get("dispatch_ttl").(string),
		DisplayView:                        d.Get("displayview").(string),
		IsScheduled:                        d.Get("is_scheduled").(bool),
		IsVisible:                          d.Get("is_visible").(bool),
		MaxConcurrent:                      d.Get("max_concurrent").(int),
		RealtimeSchedule:                   d.Get("realtime_schedule").(bool),
		RequestUIDispatchApp:               d.Get("request_ui_dispatch_app").(string),
		RequestUIDispatchView:              d.Get("request_ui_dispatch_view").(string),
		RestartOnSearchPeerAdd:             d.Get("restart_on_searchpeer_add").(bool),
		RunOnStartup:                       d.Get("run_on_startup").(bool),
		VSID:                               d.Get("vsid").(string),
	}

	return savedSearch
}
