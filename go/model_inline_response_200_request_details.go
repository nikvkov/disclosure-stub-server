/*
 * Shareholder Disclosure Hub API
 *
 * # About  Broadridge developed the Shareholder Disclosure Hub (SDH) to assist banks and other intermediaries with their compliance obligations with respect to the shareholder identification regulations made law in the 2017 amendment to the Shareholder Rights Directive (SRD II). SRD II allows issuers or issuer agents to submit a disclosure request to the chain of intermediaries at any time throughout the year at their chosen frequency. On behalf of Hub-participating intermediaries the Hub validates/authenticates these requests, provides the necessary request data to the Hub participant’s system, receives and validates the required response data, forwards the request to the Hub participant’s underlying intermediary clients and files the full response with the issuer or issuer agent.  The SDH API is built on the `Open API 3.0.3` specification. **[[LINK](https://swagger.io/specification/)]**  ## API Versioning --- The API is versioned using the Semantic versioning 2.0.0 (semver) specification. **[[LINK](https://semver.org/spec/v2.0.0.html)]**  The versioning format is `major.minor.path`. The `major.minor` portion of the semver (e.g. 2.0.0, 2.1.0) shall designate a change in the feature. `\".patch\"` versions addresses errors if any in the document/SDH API.  ## API Technical Specifications --- The API provides the following services:  - **API Authentication:** The SDH API requires Bearer HTTP authentication. Every request **must** contain an `Authorization` header using the username and password associated with a user account existing in SDH.    Format of the authorization header:    ```   Authorization: Bearer <token>    Example:   Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzT24iOjE1ODM0NTAxMDA1NTksImVtYWlsIjoidXNlcjFAc2MxLmNvbSIsIm1ldGEiOiJEcUE5Q3Rwdjg1RnVZa3B5TlM0RjAyMjdDVStRTUJpM25lekJwL0FtdVBtNGM3M1NFOWlQNC93bGFrZ2taSG1pWHg2cVgyNkk2RHczdTdmWDVRVGJIT1JtTVBtcUtZaGJRUGd5UHBxT2pzNHE0dEVyM1d2b1lxS0xEd3IwdkhTYmc1a3VkUHRXRno4ZTVWaWNkSTh6L08ydHFzYk9JTDc3NnFib2NzdnVneGN4ckg4Njg4akd2azVtaVhwbFY5UldIZlFlZEFmS3BxOWx5aEZuaWZOd3NQNU15aFIxRExKUGl2ZEtCRkJzY0pYWXRoUERsRGxvNExZOXRtdDVNTUVrckZIQUlFanhCUlBLbWhmYk9UNnowa3NyUnJnVGxRPT0iLCJ0b2tlbiI6IkF6bllnRG1zZjRVaFd3ZWFOZlRXZWRoTXlXSWpzSkd3VGlkd05hZmc2RUU9In0=.smRPmkC6TeZmUnwlrEjexEKHCehzS8qBC8mbIJOOJe0=   ```   The token is valid for `14400 seconds` i.e. `4 hours` from the date and time of its generation.  - **Disclosure requests:** This API endpoint enables the fetching of shareholder disclosure requests which are in `OPEN` and `PENDING` state.   This request must be accompanied by the JWT generated by the `API Authentication` endpoint, as mentioned above.  - **Disclosure response:** This section provides a set of endpoints for uploading a disclosure response.    - **Upload token:** A request for file upload must be accompanied by a CSRF token in the header of the request.       To generate a CSRF token for a file containing the disclosure response, an MD5 checksum of the file (in hex format), along with the name of the file       is to be provided in the request. This request must be accompanied by the JWT generated by the `API Authentication` endpoint, as mentioned above.    - **File upload:** Upload a file containing response to a disclosure request.     - The file should be of `JSON` format and should have the following naming convention **disclosureResponseIdentification_YYYYMMDDHHMMSS.json**.       - `disclosureResponseIdentification`: value of the disclosureResponseIdentification field in the response.       - `YYYYMMDDHHMMSS`: The date and time in UTC timezone in the said format.       For e.g. 4c582ed0eb3c01cb671000014d93738f_20200403103044.json     - The structure of the response is defined under `DisclosureResponse` schema.     - The file upload request must be accompanied by the CSRF and the authentication tokens.    - **File Processing status:** This API endpoint enables the fetching the statuses of disclosure responses uploaded.    **Note:** The SDH API currently doesn't support processing of bulk responses. Therefore only one response per file upload is allowed for processing.
 *
 * API version: 1.0.0-oas3
 * Contact: SRD-Hub-Dev@broadridge.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse200RequestDetails struct {
	// Official and unique identification assigned to a shareholders identification disclosure request process by the issuer or third party nominated by it.
	IssuerDisclosureRequestIdentification string `json:"issuerDisclosureRequestIdentification"`
	// Specifies the type of disclosure request.  Valid request types:   1) \"NEWM\": New - New Disclosure Request   2) \"REPL\": Replacement - Disclosure request replacing a previously sent request.
	DisclosureRequestType string `json:"disclosureRequestType"`
	// Indicates whether the request is to be forwarded to and responded by the other intermediaries down the chain of intermediaries or not.  If not to be forwarded, the indicator may not be present. 1) Yes: When True 2) No: When False
	ForwardRequestIndicator string `json:"forwardRequestIndicator,omitempty"`
	// Indicates whether the request was initiated by the first intermediary in the custody chain in accordance with SRD II. 1) Yes: When True 2) No: When False
	ShareholderRightsDirectiveIndicator string `json:"shareholderRightsDirectiveIndicator,omitempty"`

	FinancialInstrumentationIdentification *InlineResponse200RequestDetailsFinancialInstrumentationIdentification `json:"financialInstrumentationIdentification"`
	// Date set by the issuer on which shareholders identity is determined based on the settled positions struck in the books of the Issuer CSD or any other first intermediary at the close of business day.
	ShareholdersDisclosureRecordDate *OneOfinlineResponse200RequestDetailsShareholdersDisclosureRecordDate `json:"shareholdersDisclosureRecordDate"`
	// Minimum number of shares held by a shareholder above which the identification must be disclosed. Format: Decimal Number - Total Digits: 18, Fraction Digits: 17
	SharesQuantityThreshold float64 `json:"sharesQuantityThreshold,omitempty"`
	// Indicates whether the date from which the shares have been held must be communicated into the disclosure response and according to which method theses dates have to be communicated.
	RequestShareHeldDate *OneOfinlineResponse200RequestDetailsRequestShareHeldDate `json:"requestShareHeldDate,omitempty"`

	DisclosureResponseRecipient *InlineResponse200RequestDetailsDisclosureResponseRecipient `json:"disclosureResponseRecipient"`
	// Latest date/time set by the issuer or a third party appointed by the issuer at which a response to the request to disclose shareholder identity shall be provided by each intermediary to the recipient to whom the response must be sent.
	IssuerDisclosureDeadline *OneOfinlineResponse200RequestDetailsIssuerDisclosureDeadline `json:"issuerDisclosureDeadline"`
	// Indicates whether the shareholder identification disclosure response is to be sent back down the chain of intermediaries or directly to the identified response recipient.  Note: This is currently not supported by Broadridge Shareholder's Disclosure Hub.
	ResponseThroughChainIndicator string `json:"responseThroughChainIndicator,omitempty"`
	// Latest date/time set by an intermediary at which a response to the request to disclose shareholder identity shall be provided when sending the response though the chain of intermediaries.  Note: This field is set only when \"responseThroughChainIndicator\" is set to \"Yes\".
	DisclosureResponseDeadline *OneOfinlineResponse200RequestDetailsDisclosureResponseDeadline `json:"disclosureResponseDeadline,omitempty"`
	// Issuer of the financial instrument.
	Issuer *OneOfinlineResponse200RequestDetailsIssuer `json:"issuer,omitempty"`
	// Country code of the market in which the Issuer is registered. The code will indicate the market of the request and the respective market specific rule will be applied on the response data.
	IsinMarketCode *AllOfinlineResponse200RequestDetailsIsinMarketCode `json:"isinMarketCode"`

	SupplementaryData []SupplementaryDataInner `json:"supplementaryData,omitempty"`
}
