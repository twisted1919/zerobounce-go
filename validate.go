package zerobounce

type ValidateResponseSuccess struct {
	// The email address you are validating.
	Address string
	// [valid, invalid, catch-all, unknown, spamtrap, abuse, do_not_mail]
	Status string
	//nolint:lll    // [antispam_system, greylisted, mail_server_temporary_error, forcible_disconnect, mail_server_did_not_respond, timeout_exceeded, failed_smtp_connection, mailbox_quota_exceeded, exception_occurred, possible_trap, role_based, global_suppression, mailbox_not_found, no_dns_entries, failed_syntax_check, possible_typo, unroutable_ip_address, leading_period_removed, does_not_accept_mail, alias_address, role_based_catch_all, disposable, toxic]
	SubStatus string `json:"sub_status"` //nolint:tagliatelle
	// [true/false] If the email comes from a free provider.
	FreeEmail bool `json:"free_email"` //nolint:tagliatelle
	// Suggestive Fix for an email typo
	DidYouMean *string `json:"did_you_mean"` //nolint:tagliatelle
	// The portion of the email address before the "@" symbol or null.
	Account *string `json:"account"`
	// The portion of the email address after the "@" symbol or null.
	Domain *string `json:"domain"`
	// Age of the email domain in days or [null].
	DomainAgeDays *string `json:"domain_age_days"` //nolint:tagliatelle
	// The SMTP Provider of the email or [null] [BETA].
	SMTPProvider *string `json:"smtp_provider"` //nolint:tagliatelle
	// The preferred MX record of the domain
	MxRecord *string `json:"mx_record"` //nolint:tagliatelle
	// [true/false] Does the domain have an MX record. [they return "bool" not bool, which makes it a string...]
	MxFound *string `json:"mx_found"` //nolint:tagliatelle
	// The first name of the owner of the email when available or [null].
	Firstname *string `json:"firstname"`
	// The last name of the owner of the email when available or [null].
	Lastname *string `json:"lastname"`
	// The gender of the owner of the email when available or [null].
	Gender *string `json:"gender"`
	// The country of the IP passed in or [null]
	Country *string `json:"country"`
	// The region/state of the IP passed in or [null]
	Region *string `json:"region"`
	// The city of the IP passed in or [null]
	City *string `json:"city"`
	// The zipcode of the IP passed in or [null]
	Zipcode *string `json:"zipcode"`
	// The UTC time the email was validated.
	ProcessedAt string `json:"processed_at"` //nolint:tagliatelle
}

func (r *ValidateResponseSuccess) IsValid() bool {
	return r.Status == "valid"
}

func (r *ValidateResponseSuccess) IsInvalid() bool {
	return r.Status == "invalid"
}

func (r *ValidateResponseSuccess) IsCatchAll() bool {
	return r.Status == "catch-all"
}

func (r *ValidateResponseSuccess) IsUnknown() bool {
	return r.Status == "unknown"
}

func (r *ValidateResponseSuccess) IsSpamtrap() bool {
	return r.Status == "spamtrap"
}

func (r *ValidateResponseSuccess) IsAbuse() bool {
	return r.Status == "abuse"
}

func (r *ValidateResponseSuccess) IsDoNotMail() bool {
	return r.Status == "do_not_mail"
}
