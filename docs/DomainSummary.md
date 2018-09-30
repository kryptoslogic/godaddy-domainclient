# DomainSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** | Authorization code for transferring the Domain | [optional] [default to null]
**ContactAdmin** | [***Contact**](Contact.md) | Administrative contact for the domain registration | [optional] [default to null]
**ContactBilling** | [***Contact**](Contact.md) | Billing contact for the domain registration | [optional] [default to null]
**ContactRegistrant** | [***Contact**](Contact.md) | Registration contact for the domain | [default to null]
**ContactTech** | [***Contact**](Contact.md) | Technical contact for the domain registration | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Date and time when this domain was created | [default to null]
**DeletedAt** | [**time.Time**](time.Time.md) | Date and time when this domain was deleted | [optional] [default to null]
**TransferAwayEligibleAt** | [**time.Time**](time.Time.md) | Date and time when this domain is eligible to transfer | [optional] [default to null]
**Domain** | **string** | Name of the domain | [default to null]
**DomainId** | **float64** | Unique identifier for this Domain | [default to null]
**ExpirationProtected** | **bool** | Whether or not the domain is protected from expiration | [default to null]
**Expires** | [**time.Time**](time.Time.md) | Date and time when this domain will expire | [optional] [default to null]
**HoldRegistrar** | **bool** | Whether or not the domain is on-hold by the registrar | [default to null]
**Locked** | **bool** | Whether or not the domain is locked to prevent transfers | [default to null]
**NameServers** | **[]string** | Fully-qualified domain names for DNS servers | [optional] [default to null]
**Privacy** | **bool** | Whether or not the domain has privacy protection | [default to null]
**RenewAuto** | **bool** | Whether or not the domain is configured to automatically renew | [default to null]
**RenewDeadline** | [**time.Time**](time.Time.md) | Date the domain must renew on | [default to null]
**Renewable** | **bool** | Whether or not the domain is eligble for renewal based on status | [optional] [default to null]
**Status** | **string** | Processing status of the domain&lt;br/&gt;&lt;ul&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;ACTIVE&lt;/strong&gt; - All is well&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;AWAITING*&lt;/strong&gt; - System is waiting for the end-user to complete an action&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;CANCELLED*&lt;/strong&gt; - Domain has been cancelled, and may or may not be reclaimable&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;CONFISCATED&lt;/strong&gt; - Domain has been confiscated, usually for abuse, chargeback, or fraud&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;DISABLED*&lt;/strong&gt; - Domain has been disabled&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;EXCLUDED*&lt;/strong&gt; - Domain has been excluded from Firehose registration&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;EXPIRED*&lt;/strong&gt; - Domain has expired&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;FAILED*&lt;/strong&gt; - Domain has failed a required action, and the system is no longer retrying&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;HELD*&lt;/strong&gt; - Domain has been placed on hold, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;LOCKED*&lt;/strong&gt; - Domain has been locked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;PARKED*&lt;/strong&gt; - Domain has been parked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;PENDING*&lt;/strong&gt; - Domain is working its way through an automated workflow&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;RESERVED*&lt;/strong&gt; - Domain is reserved, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REVERTED&lt;/strong&gt; - Domain has been reverted, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;SUSPENDED*&lt;/strong&gt; - Domain has been suspended, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;TRANSFERRED*&lt;/strong&gt; - Domain has been transferred out&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNKNOWN&lt;/strong&gt; - Domain is in an unknown state&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNLOCKED*&lt;/strong&gt; - Domain has been unlocked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNPARKED*&lt;/strong&gt; - Domain has been unparked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UPDATED*&lt;/strong&gt; - Domain ownership has been transferred to another account&lt;/li&gt; &lt;/ul&gt; | [default to null]
**TransferProtected** | **bool** | Whether or not the domain is protected from transfer | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


