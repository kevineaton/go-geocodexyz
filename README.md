# Go Geocode.xyz Client

This is a simple client for the Geocode.xyz API. This is in the very, very early testing stages, so should not be used by anyone. We welcome contributors, but we strongly recommend you reach out first to ensure that your work and efforts will be in-line with our design, scope, and goals.

_THE API CAN, AND WILL, CHANGE!_ Do not use until this message is removed. This is in a prototyping stage and may never be completed.

## Use Case

Using this library, you can make calls to the Geocode.xyz API. This library was developed for a proof of concept explring options for reverse geocoding APIs for an employer. It is not battlehardened and is designed to test integrating with the service. It does not intend to be a full-coverage SDK and may not be production-ready if another alternative is used.

### US-focus

As the primary audience of the client has customers solely in the US, the focus of the Reverse Geocode Lookup is to find a street, city, state, postal, and country. As such, returns that do not have those fields, and cannot by marshalled into a full result object, will return an error instead.

## Testing

Note that without an auth key, your testing may not be complete. For example, if you are rate limited, the "success" path in the tests will not pass, since the error handling would be called.

## Contributing

Pull Requests are welcome! See our `CONTRIBUTING.md` file for more information.
