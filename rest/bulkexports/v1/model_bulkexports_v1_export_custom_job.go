/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Bulkexports
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// BulkexportsV1ExportCustomJob struct for BulkexportsV1ExportCustomJob
type BulkexportsV1ExportCustomJob struct {
	// The friendly name specified when creating the job
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
	// The start day for the custom export specified as a string in the format of yyyy-MM-dd
	StartDay *string `json:"start_day,omitempty"`
	// The end day for the custom export specified as a string in the format of yyyy-MM-dd. This will be the last day exported. For instance, to export a single day, choose the same day for start and end day. To export the first 4 days of July, you would set the start date to 2020-07-01 and the end date to 2020-07-04. The end date must be the UTC day before yesterday.
	EndDay *string `json:"end_day,omitempty"`
	// The optional webhook url called on completion
	WebhookUrl *string `json:"webhook_url,omitempty"`
	// This is the method used to call the webhook
	WebhookMethod *string `json:"webhook_method,omitempty"`
	// The optional email to send the completion notification to
	Email *string `json:"email,omitempty"`
	// The unique job_sid returned when the custom export was created. This can be used to look up the status of the job.
	JobSid *string `json:"job_sid,omitempty"`
	// The details of a job state which is an object that contains a `status` string, a day count integer, and list of days in the job
	Details *interface{} `json:"details,omitempty"`
	// This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease
	JobQueuePosition *string `json:"job_queue_position,omitempty"`
	// this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position
	EstimatedCompletionTime *string `json:"estimated_completion_time,omitempty"`
}
