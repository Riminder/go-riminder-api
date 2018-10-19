# go-riminder-api
🐹 Riminder API Golang Wrapper

-------------------------------------

A go api client for Riminder api.

## Instalation with go gems

```shell
go get -u github.com/Xalrandion/go-riminder-api
```

## Authentification

To authenticate against the api, get your API SECRET KEY from your riminder dashboard: 
![findApiSecret](./secretLocation.png)

Then create a new `Riminder` object with this key:

```go
    import "github.com/Xalrandion/go-riminder-api/riminder"

    // Authentification to api
    client := riminder.New("your shinny key")

    // Finally you can use the api!!
```

## Api overview

```go
    import (
        "github.com/Xalrandion/go-riminder-api/riminder"
        "fmt")


    client := riminder.New("your shinny key")

    // Let's retrieve a profile.
    args := map[string]interface{}{
        "source_id":  "a source id",
        "profile_id": "id of the profile",
    }
    profile := client.Profile().Get(args)
    // And print his name !
    fmt.Printf("This profile name is %s", profile.Name)
```

## Api

The mentionned team is the team linked to your secret key.

When both `*_id` and `*_reference` arguments are requested only one is requiered.
For example `client.Filter().Get()` can take a

* `filter_id` (`client.Filter().Get(map[string]interface{}{"filter_id":filter_id})`)
* `filter_reference` (`client.Filter().Get(map[string]interface{}{"filter_reference":filter_reference})`)

and work as well.

Only the `data` field will be returned by methods.

For details and examples see [our documentation](https://developers.Riminder.net/v1.0/reference).

### Filter

* Get all filters from the team.
  * returns: Array of Hash containing the filters, and an error if one occurs (nil if no error).

```go
 resp, err := client.Filter().List()
```

* Get a specific filter.
  * Arguments: Hash containing
    * `"filter_id"` id of the filter (*required*)
    * `"filter_reference"` reference of the filter (*required*)
  * Returns: Hash containing the filter, and an error if one occurs (nil if no error).

```go
 resp, err := client.Filter().Get(map[string]interface{}{"filter_id": "a filter id"})
```

More details about filters are available [here](https://developers.Riminder.net/v1.0/reference#jobs)

### Profile

* Retrieve the profiles information associated with specified source ids.
  * Arguments: Hash containing
    * `"source_ids"` Array of source ids (*requiered*)
    * `"seniority"` profile's seniority
    * `"filter_id"` filter id (to sort by filter stage and/or rating)
    * `"filter_reference"` filter_reference (to sort by filter stage and/or rating)
    * `"stage"` stage (to sort by filter stage and/or rating)
    * `"rating"` rating (to sort by filter stage and/or rating)
    * `"date_start"` profile's first date of reception (default 9/9/2012)
    * `"date_end"` profile's last date of reception (default: now)
    * `"page"` selected page (default: 1)
    * `"limit"` number of result by page
    * `"sort_by"` sort profile by (default: ranking)
    * `"order_by"` order profile by
  * Returns: Array of Hash containing the profiles, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().List(map[string]interface{}{ "source_ids": source_ids, "seniority": "all", "page": 2, "limit": 5})
```

* Add a new profile to a source on the platform.
  * Arguments: Hash containing
    * `"filepath"` path of the file to be uploaded. (*required*)
    * `"source_id"` id of the target source. (*required*)
    * `"timestamp_reception"` reception's timestamp
    * `"profile_reference"` reference for the new profile
    * `"training_metadata"` training metadatas to add.
  * Returns: Hash containing upload result, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().Add(map[string]interface{}{ "filepath" : "path/to/file", "source_id" : source_id, "profile_reference" : "reference0"})
```

* Get a specific profile.
  * Arguments: Hash containing
    * `"source_id"` id of the source containing the profile (*required*)
    * `"profile_id"` id of the profile (*required*)
    * `"profile_reference"` reference of the profile (*required*)
  * Returns: Hash containing the profile, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().Get(map[string]interface{}{"source_id": source_id, "profile_reference": profile_reference})
```

* Get attachements of a specific profile.
  * Arguments: Hash containing
    * `"source_id"` id of the source containing the profile (*required*)
    * `"profile_id"` id of the profile (*required*)
    * `"profile_reference"` reference of the profile (*required*)
  * Returns: Array of Hash containing the profile attachements, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().Document().List(map[string]interface{}{"source_id": source_id, "profile_id": profile_id})
```

* Get parsing result of a specific profile.
  * Arguments: Hash containing
    * `"source_id"` id of the source containing the profile (*required*)
    * `"profile_id"` id of the profile (*required*)
    * `"profile_reference"` reference of the profile (*required*)
  * Returns: Hash containing the profile parsing, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().Parsing().Get(map[string]interface{}{"source_id": source_id, "profile_id": profile_id})
```

* Get scoring result of a specific profile.
  * Arguments: Hash containing
    * `"source_id"` id of the source containing the profile (*required*)
    * `"profile_id"` id of the profile (*required*)
    * `"profile_reference"` reference of the profile (*required*)
  * Returns: Array of Hash containing the profile scoring, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().Scoring().List(map[string]interface{}{"source_id": source_id, "profile_reference": profile_reference})
```

* Set stage of a specific profile for a spcified filter.
  * Arguments: Hash containing
    * `"source_id"` id of the source containing the profile (*required*)
    * `"profile_id"` id of the profile (*required*)
    * `"profile_reference"` reference of the profile (*required*)
    * `"filter_id"` id of the target filter (*required*)
    * `"filter_reference"` reference of the target filter (*required*)
    * `"stage"` new stage
  * Returns: Hash containing operation result, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().Stage().Set(map[string]interface{}{"source_id": source_id, "profile_id": profile_id, "filter_reference": filter_reference, "stage": "YES"})
```

* Set rating of a specific profile for a spcified filter.
  * Arguments: Hash containing
    * `"source_id"` id of the source containing the profile (*required*)
    * `"profile_id"` id of the profile (*required*)
    * `"profile_reference"` reference of the profile (*required*)
    * `"filter_id"` id of the target filter (*required*)
    * `"filter_reference"` reference of the target filter (*required*)
    * `"stage"` new stage
  * Returns: Hash containing operation result, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().Rating().Set(map[string]interface{}{"source_id": source_id, "profile_id": profile_id, "filter_reference": filter_reference, "rating": 3})
```

* Check if a parsed profile is valid.
  * Arguments: Hash containing
    * `"profile_json"` profile data to check. (*required*)
    * `"training_metadata"` training metadatas to add.
  * Returns: Hash containing upload result, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().JSON().Check(map[string]interface{}{"profile_json": profile_datas})
```

* Add a parsed profile to a source on the platform.
  * Arguments: Hash containing
    * `"profile_json"` profile data to upload. (*required*)
    * `"source_id"` id of the target source. (*required*)
    * `"timestamp_reception"` reception's timestamp
    * `"profile_reference"` reference for the new profile
    * `"training_metadata"` training metadatas to add.
  * Returns: Hash containing upload result, and an error if one occurs (nil if no error).

```go
 resp, err := client.Profile().JSON().Add(map[string]interface{}{"profile_json": profile_datas, "source_id": "a source_id", "timestamp_reception": 1347209668})
```

More details about profiles are available [here](https://developers.Riminder.net/v1.0/reference#profile)

### Sources

* Get all source from the team.
  * returns: Array of Hash containing the sources, and an error if one occurs (nil if no error).

```go
 resp, err := client.Source().List()
```

* Get a specific source.
  * Arguments: 
    * source_id -> id of the source to retrieve.
  * Returns: Hash containing a specific source, and an error if one occurs (nil if no error).

```go
 resp, err := client.Source().Get("source_id": source_id)
```

More details about profiles are available [here](https://developers.Riminder.net/v1.0/reference#source)

### Webhooks

Webhooks methods permit you handle webhook events.
To handle event a webhook key is needed to be given, using `Riminder.SetWebhookKey(webhook_key)`

* Check if team's webhook integration is working.
  * Returns: Hash containing operation result, and an error if one occurs (nil if no error).

```go
 resp, err := client.Webhooks().check()
```

* Set an handler for a specified webhook event.
  * Arguments:
    * eventName -> name of the target event
    * callback -> callable called when target webhook is received.
      * takes 2 args:
        * eventName: name of the event (type field of webhooks)
        * payload: webhooks datas
  * returns: an error if one occurs (nil if no error)

```go
client.Webhooks().SetHandler("profile.parse.success", handler);
```

* Check if there is an handler for a specified event
  * Arguments:
    * eventName -> target event name
  * returns: an error if one occurs (nil if no error)

```go
client.Webhooks().isHandlerPresent(eventName);
```

* Remove the handler for an event
  * Arguments:
    * eventName -> target event name
  * returns: an error if one occurs (nil if no error)

```go
client.Webhooks().removeHandler(eventName);
```

* Start the selected handler depending of the event given.
  * Arguments:
    * receivedMessage -> received webhook request's headers as a Hash or the content of received webhook request's `"HTTP-RIMINDER-SIGNATURE"`header.
  * returns: an error if one occurs (nil if no error).

```go
client.Webhooks().handle(receivedHeader)
```

Example:

```go
client := riminder.New("an api key")
client.SetWebhookKey("your webhook key...")

cb  := func(event_type string, event_data interface{}) {
    // some treatment here
}
api.webhooks.SetHandler("profile.parse.success", cb)

encodedHeader := get_webhook_headers()
api.webhooks.handle encodedHeader
```

More details about webhooks are available [here](https://developers.Riminder.net/v1.0/reference#authentication-1)

## Tests

Some tests are available. To run them follow these steps:

* `git clone git@github.com:Riminder/go-riminder-api.git`
* `cd go-riminder-api`
* `go get ./...`
* `go test ./...`

## Help and documentation

If you need some more details about the api methods and routes see [Riminder API Docs](https://developers.Riminder.net/v1.0/reference).

If you need further explainations about how the api works see [Riminder API Overview](https://developers.riminder.net/v1.0/docs/website-api-overview) 